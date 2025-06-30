package config

import (
	"os"

	"github.com/server/config/env"
	"github.com/server/internal/adapters/api"
	"github.com/server/internal/adapters/bootstrap"
	"github.com/server/internal/adapters/metrics"
	"github.com/server/internal/adapters/store/blob"
	"github.com/server/internal/adapters/store/postgres"
	"github.com/server/internal/adapters/store/redis"
	"github.com/server/internal/adapters/store/spicedb"
	"github.com/server/internal/core/auth"
	"github.com/server/pkg/logger"
	"github.com/server/pkg/mailer"
	"github.com/server/pkg/utils"
	"gopkg.in/yaml.v3"
)

type API struct {
	Server api.Config `yaml:"server"`
	Metrics metrics.Config `yaml:"metrics"`
	Mailer mailer.Config  `yaml:"mailer"`
	Auth auth.Config `yaml:"authentication"`
}

type Config struct {
	Enviroment env.ENVIROMENT `yaml:"enviroment" default:"0"`
	Api API `yaml:"api"`
	Bootstrap bootstrap.Config `yaml:"bootstrap"`
	Logger logger.Config `yaml:"logger"`	 
	DB postgres.Config `yaml:"db"`
	DBCache redis.Config `yaml:"db_cache"`
	SpiceDB spicedb.Config `yaml:"spice_db"`
	BlobFS blob.Config `yaml:"blob_storage"`
}

func LoadConfig(yamlFilePath string) (*Config, error) {

	cfg := &Config{}

	yamlFile, err := os.ReadFile(yamlFilePath)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "os.ReadFile")
	}

	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		return nil, utils.WrapErrorf(err, utils.ErrorCodeUnknown, "yaml.Unmarshal")
	}

	return cfg, err
}
