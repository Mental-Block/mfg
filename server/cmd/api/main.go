package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/server/config"

	"github.com/server/pkg/logger"
	"github.com/server/pkg/mailer"
	"github.com/server/pkg/utils"

	"github.com/server/internal/adapters/api"
	"github.com/server/internal/adapters/bootstrap"
	"github.com/server/internal/adapters/metrics"
	"github.com/server/internal/adapters/store/blob"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/adapters/store/postgres/tx"
	"github.com/server/internal/adapters/store/redis"
	"github.com/server/internal/adapters/store/spicedb"

	auth_store "github.com/server/internal/adapters/store/postgres/auth"
	namespace_store "github.com/server/internal/adapters/store/postgres/namespace"
	permission_store "github.com/server/internal/adapters/store/postgres/permission"
	relation_store "github.com/server/internal/adapters/store/postgres/relation"
	role_store "github.com/server/internal/adapters/store/postgres/role"
	serviceuser_store "github.com/server/internal/adapters/store/postgres/serviceuser"
	user_store "github.com/server/internal/adapters/store/postgres/user"
	session_store "github.com/server/internal/adapters/store/redis"

	"github.com/server/internal/core/auth"
	"github.com/server/internal/core/auth/session"
	"github.com/server/internal/core/auth/token"
	"github.com/server/internal/core/namespace"
	"github.com/server/internal/core/permission"
	"github.com/server/internal/core/relation"
	"github.com/server/internal/core/role"
	"github.com/server/internal/core/serviceuser"
	"github.com/server/internal/core/user"
)

func main() {
	var conf string

	flag.StringVar(&conf, "config", "../infra/api/config.yaml", "filename")
	flag.Parse()

	cfg, err := config.LoadConfig(conf)

	if (err != nil) {
		log.Fatal(err)
	}

	errC, err := run(cfg)

	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(cfg *config.Config) (<-chan error, error) {
	logr := logger.Set(cfg.Logger)

	errC := make(chan error, 1)
	
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	// RDBMS database
	postgres, err := postgres.NewStore(ctx, cfg.DB)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "postgres.NewStore")
	}

	// In-memory database
	redis, err := redis.NewStore(ctx, cfg.DBCache)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "redis.NewStore")
	}

	// blob/file storage
	blobFS, err := blob.NewStore(ctx, cfg.BlobFS)
	
	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "blob.NewStore")
	}
	
	// schema, authorization DSL file location for migration
	resourceStorage := blob.NewResourceStore(logr, blobFS)
	
	if err := resourceStorage.InitCache(ctx, time.Minute * 2); err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "resourceStorage.InitCache")
	}

	// metrics collection for Spicedb/authzed
	stdLibpostgres := postgres.PgxToStdLib()
	_, promMetrics := metrics.Initalize(stdLibpostgres, "metrics")

	// Spicedb/authzed GRPC client
	spicedb, err := spicedb.NewClient(cfg.SpiceDB, promMetrics)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "spicedb.NewClient")
	}
	
	// build app dependanicies getting ready to serve
	services := buildDependancies(
		cfg,
		logr,
		postgres,
		redis,
		spicedb,
		resourceStorage,
	)

	// migrate schemas
	if err = services.BootstrapService.MigrateSchema(ctx); err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "services.BootstrapService.MigrateSchema")
	}

	//  migrate roles
	if err = services.BootstrapService.MigrateRoles(ctx); err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "services.BootstrapService.MigrateRoles")
	}

	//  promote users
	if err = services.BootstrapService.MakeSuperUsers(ctx); err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "services.BootstrapService.MakeSuperUsers")
	}
	
	// start API
	api := api.Serve(
		*services.ApiServices,
		api.WithMiddlewares([]func(next http.Handler) http.Handler{
			api.LoggerMiddleWare(logr), 
			api.CORSMiddleWare(cfg.Api.Server.CorsConfig),
		}),
		api.WithDocUI(cfg.Api.Server.DocUI),
		api.WithDocPath(cfg.Api.Server.DocsPath),
		api.WithHost(cfg.Api.Server.Host),
		api.WithPort(strconv.Itoa(cfg.Api.Server.Port)),
		api.WithVersion(cfg.Api.Server.Version),
		api.WithName(cfg.Api.Server.Name),
		api.WithEnivorment(cfg.Enviroment),
	)	

	// if any service critically fails... flush logger and gracefully shutdown services
	go func() {
		<-ctx.Done()

		logr.Info("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

		defer func() {
			err = logr.Sync()

			postgres.Close()
			_ = redis.Close()
			_ = resourceStorage.Close()
			_ = spicedb.Close()
			_ = api.Server.Close()
			
			stop()
			cancel()
			close(errC)
		}()

		api.Server.SetKeepAlivesEnabled(false)

		if err := api.Server.Shutdown(ctxTimeout); err != nil { 
			errC <- err
		}

		logr.Info("Shutdown completed")
	}()

	go func() {
		uri := net.JoinHostPort(cfg.Api.Server.Host, strconv.Itoa(cfg.Api.Server.Port)) + cfg.Api.Server.DocsPath

		logr.Info(fmt.Sprintf("Listening and serving: %v", uri))

		if err := api.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}


type Dependancies struct {
	ApiServices  		*api.Services
	BootstrapService 	*bootstrap.Service    
}

func buildDependancies(
	cfg   		*config.Config,
	logger      logger.Logger,
	dbCon 		*postgres.Store,
	dbcacheCon	*redis.Store,
	spicedbCon  *spicedb.SpiceDB,
	resourceBlob *blob.ResourceStore,
) (*Dependancies) {	
	/* 
		At runtime this will inject the needed stores for transactions.  
		Prefer to keep transactions within the store repository if posible if the context makes sense. 
		This pattern should be used sparingly as it is complex and very easy to introduce bugs.

		Article explaining GO anti patterns: https://threedots.tech/post/database-transactions-in-go/
		Github repo of transaction strategies: https://github.com/ThreeDotsLabs/go-web-app-antipatterns
		Prefered strategy when using pgx driver: https://github.com/MarioCarrion/videos/tree/5dbfad345ff8e39c6539ce4f51aea5db582ed59a/2023/transaction-in-context/internal/postgresql
	*/
	txProvider := tx.NewTXProvider(dbCon)

	//mailer
	mailer := mailer.NewDialer(cfg.Api.Mailer)

	// authorization relations & schemas 
	authzRelationStore := spicedb.NewRelationStore(
		spicedbCon, 
		cfg.SpiceDB.Consistency, 
		cfg.SpiceDB.CheckTrace,
		logger,
	)

	// base relation - app specific and global
	relationStore := relation_store.NewRelationStore(dbCon)
	relationService := relation.NewRelationService(
		relationStore, 
		authzRelationStore,
	)

	nameSpaceStore := namespace_store.NewNamespaceStore(dbCon)
	nameSpaceService := namespace.NewService(nameSpaceStore)

	permissionStore := permission_store.NewPermissionStore(dbCon)
	permissionService := permission.NewService(permissionStore)

	roleStore := role_store.NewRoleStore(dbCon)
	roleService := role.NewRoleService(roleStore, relationService, permissionService)

	userStore := user_store.NewUserStore(dbCon)
	userService := user.NewUserService(userStore, relationService)

	//authentication
	authTxConsumer := tx.NewAuthTxConsumer(txProvider)
	authStore := auth_store.NewAuthStore(dbCon)

	sessionStore := session_store.NewSessionStore(dbcacheCon)
	sessionService := session.NewSessionService(cfg.Api.Auth.Session, authStore, sessionStore)

	tokenService := token.NewTokenService(cfg.Api.Auth.Token)

	flowStore := redis.NewFlowStore(dbcacheCon)

	serviceUserStore := serviceuser_store.NewServiceUserStore(dbCon)
	serviceUserCredsStore := serviceuser_store.NewServiceUserCredentialStore(dbCon)
	serviceUserService := serviceuser.NewServiceUserService(
		cfg.Api.Auth.Password.Params, 
		serviceUserStore, 
		serviceUserCredsStore, 
		relationService,
	)
	
	authService := auth.NewAuthService(
		cfg.Api.Auth,
		authStore,
		flowStore,
		authTxConsumer,
		mailer,
		tokenService,
		serviceUserService,
		sessionService,
	)
	
	// schema relation - contains relations files for migrations
	resourceRelationSchemastore := blob.NewSchemaStore(resourceBlob.Bucket)
	
	// current schemas - in spicedb
	authzSchemaStore := spicedb.NewSchemaStore(spicedbCon)

	// bootstrap 
	// - migrate authz DSL files and app specific relations, role, permissions and schemas
	// - create any super users if specified
	bootstrapService := bootstrap.NewBootstrapService(
		cfg.Bootstrap,
		resourceRelationSchemastore,
		nameSpaceService,
		roleService,
		permissionService,
		userService,
		authzSchemaStore,
	)
	
	return &Dependancies{
		ApiServices: &api.Services{
			AuthService:  authService,
			UserService:  userService,
		},
		BootstrapService: bootstrapService,
	}
}
