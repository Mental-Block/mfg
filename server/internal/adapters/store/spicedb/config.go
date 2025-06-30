package spicedb

import "strings"

type Config struct {
	Host         string `yaml:"host" default:"localhost"`
	Port         string `yaml:"port" default:"50051"`
	PreSharedKey string `yaml:"pre_shared_key" mapstructure:"pre_shared_key"`

	// FullyConsistent ensures APIs although slower than usual will result in responses always most consistent
	Consistent bool `yaml:"consistent" mapstructure:"consistent" default:"false"`

	// Consistency ensures Authz server consistency guarantees for various operations
	// Possible values are:
	// - "full": Guarantees that the data is always fresh
	// - "best_effort": Guarantees that the data is the best effort fresh
	// - "minimize_latency": Tries to prioritise minimal latency
	Consistency string `yaml:"consistency" mapstructure:"consistency" default:"best_effort"`

	// CheckTrace enables tracing in check api for spicedb, it adds considerable
	// latency to the check calls and shouldn't be enabled in production
	CheckTrace bool `yaml:"check_trace" mapstructure:"check_trace" default:"false"`
}

func (c *Config) ConsistencyPolicy() ConsistencyLevel {
	switch strings.ToLower(c.Consistency) {
	case "full":
		return  ConsistencyLevelFull
	case "best_effort":
		return ConsistencyLevelBestEffort
	case "minimize_latency":
		return ConsistencyLevelMinimizeLatency
	}
	return ConsistencyLevelFull
}

type ConsistencyLevel string

const (
	ConsistencyLevelFull            ConsistencyLevel = "full"
	ConsistencyLevelBestEffort      ConsistencyLevel = "best_effort"
	ConsistencyLevelMinimizeLatency ConsistencyLevel = "minimize_latency"
)
