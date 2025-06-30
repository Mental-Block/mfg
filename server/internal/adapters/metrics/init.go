package metrics

import (
	"database/sql"

	prometheusmiddleware "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

func Initalize(db *sql.DB, dbName string) (*prometheus.Registry, *prometheusmiddleware.ClientMetrics) {
	promRegistry := prometheus.NewRegistry()
	
	promMetrics := prometheusmiddleware.NewClientMetrics(
		prometheusmiddleware.WithClientHandlingTimeHistogram(),
	)
	
	promRegistry.MustRegister(promMetrics)

	dbPromCollector := collectors.NewDBStatsCollector(db, dbName)
	
	promRegistry.MustRegister(
		dbPromCollector,
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	prometheus.DefaultRegisterer = promRegistry
		
	initDB()
	initService()

	return promRegistry, promMetrics
}