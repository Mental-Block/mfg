package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/server/internal"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/core/domain"
)

type ResourcePermissionStore struct {
	db *postgres.Store
}

func NewResourcePermissionStore(db *postgres.Store) *ResourcePermissionStore {
	return &ResourcePermissionStore{
		db,
	}
}

func (pg *ResourcePermissionStore) Insert(ctx context.Context, resourceId domain.Id, permissionId domain.Id) (*domain.ResourcePermission, error) {
	var (
		resourcePermission *domain.ResourcePermission
	)

	conn, err := pg.db.Acquire(context.Background())

	if (err != nil) {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
	}

	defer conn.Release()

	err = postgres.Transaction(ctx, conn.Conn(), func(tx pgx.Tx) error {
		rpq := resourcePermisisonQueries{ conn: tx }

		resourcePermissionId, err := rpq.insertQ(ctx, resourceId, permissionId)

		if (err != nil) {
			return internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
		}
		
		rp, err := rpq.selectQ(ctx, *resourcePermissionId)

		if err != nil {
			return internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error())
		}

		resourcePermission = rp

		return nil
	})

	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeInvalidArgument, err.Error()) 
	}

	return resourcePermission, nil
}

func (pg *ResourcePermissionStore) Delete(ctx context.Context, id domain.Id) (*domain.Id, error) {
	query := fmt.Sprintf(`DELETE FROM %v.resource_permission WHERE resource_permission_id = @id`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := pg.db.Exec(ctx, query, args)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NewErrorf(internal.ErrorCodeNotFound, domain.ErrResourceNotFound.Error()) 
		}

		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to delete row: %s", err.Error())) 
	}
	
	return &id, nil
}

type resourcePermisisonQueries struct {
	conn postgres.StoreTX
}

func (rpq resourcePermisisonQueries) insertQ(ctx context.Context, resourceId domain.Id, permissionId domain.Id) (*domain.Id, error)  {
	query := fmt.Sprintf(`
		INSERT INTO %v.resource_permission
		(
			permission_id
			,resource_id
		) VALUES(
			@permissionId
			,@resourceId
		)
		RETURNING resource_permission_id as id; 
	`, postgres.PublicSchema)

	args := pgx.NamedArgs{
		"permissionId": permissionId,
		"resourceId": resourceId,
	}

	var id *domain.Id
	err := rpq.conn.QueryRow(ctx, query, args).Scan(&id)

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert row: %s", err.Error())) 
	}

	return id, nil
}

var selectResourcePermissionQuery string = fmt.Sprintf(`
	SELECT 
		res_perm.resource_permission_id AS id 
		res.resource_id as res_id
		res.name as res_name
		perm.permission_id as perm_id
		perm.name as perm_name
	FROM %v.resource_permission AS res_perm
		INNER JOIN %v.resource AS res ON res.resource_id = res_perm.resource_id
		INNER JOIN %v.permission AS perm perm.resource_id = res_perm.permission_id
	WHERE
		res_perm.resource_permission_id = @id
`, postgres.PublicSchema, postgres.PublicSchema, postgres.PublicSchema)

func (rpq resourcePermisisonQueries) selectQ(ctx context.Context, resourcePermissionId domain.Id) (*domain.ResourcePermission, error) {
	args := pgx.NamedArgs{
		"id": resourcePermissionId,
	}

	var resourcePermission = &domain.ResourcePermission{}
	var resource = domain.Resource{}
	var permission = domain.Permission{}

	err := rpq.conn.QueryRow(ctx, selectResourcePermissionQuery, args).Scan(
		&resourcePermission.Id,
		&resource.Id,
		&resource.Name,
		&permission.Id,
		&permission.Name,
	)

	resourcePermission.Permission = permission
	resourcePermission.Resource = resource

	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, fmt.Sprintf("unable to insert row: %s", err.Error())) 
	}

	return resourcePermission, nil
} 