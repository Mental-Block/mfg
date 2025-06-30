package api

import (
	"net/http"
	"strings"
	"time"

	loggr "github.com/server/pkg/logger"
)

func LoggerMiddleWare(logger loggr.Logger) func(h http.Handler) http.Handler  {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info(r.Method,
				loggr.ZapTime("time", time.Now()),
				loggr.ZapString("url", r.URL.String()),
			)

			h.ServeHTTP(w, r)
		})
	}
}

type CorsConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins" mapstructure:"allowed_origins" default:"[http://localhost:8080]"`
	AllowedMethods []string `yaml:"allowed_methods" mapstructure:"allowed_methods" default:"[GET,POST,OPTIONS,PUT,DELETE,PATCH]"`
	AllowedHeaders []string `yaml:"allowed_headers" mapstructure:"allowed_headers" default:"[Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,accept,origin,Cache-Control,X-Requested-With]"`
	ExposedHeaders []string `yaml:"exposed_headers" mapstructure:"exposed_headers" default:"[Content-Type]"`
}

func CORSMiddleWare(cors CorsConfig) func(h http.Handler) http.Handler {
	AllowedOrigins := strings.Join(cors.AllowedOrigins, ",")
	AllowedMethods := strings.Join(cors.AllowedMethods, ",")
	AllowedHeaders := strings.Join(cors.AllowedHeaders, ",")
	ExposedHeaders := strings.Join(cors.ExposedHeaders, ",")

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 
			w.Header().Set(AccessControlAlowOrgin, AllowedMethods)
			w.Header().Set(AccessControlAllowMethods, AllowedOrigins)
			w.Header().Set(AccessControlAllowCredentials, "true")
			
			if r.Method == http.MethodOptions && r.Header.Get("Origin") != "" && r.Header.Get("Access-Control-Request-Method") != "" {
				w.Header().Set(AccessControlAllowHeaders, AllowedHeaders)
				w.Header().Set(AccessControlMaxAge, "5")
				w.WriteHeader(204)
				return
			}

			w.Header().Set(AccessControlAllowHeaders, AllowedHeaders);
			w.Header().Set(AccessControlExposeHeaders, ExposedHeaders);
		
			w.Header().Add("Vary", "Origin")
			h.ServeHTTP(w, r)
		})
	}
}