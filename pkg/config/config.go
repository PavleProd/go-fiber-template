package config

import (
	v "github.com/invopop/validation"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Application Application `mapstructure:"application"`
}

func (c Config) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.Application),
	)
}

const (
	Name      = "config"
	Extension = "yml"
	Path      = "."
)

func New() (*Config, error) {
	vpr := initViper()
	if err := vpr.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("failed to read config file")

		return nil, err
	}

	cfg := Config{}
	if err := vpr.Unmarshal(&cfg); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal config file into config structure")

		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		log.Error().Err(err).Msg("config validation returned with error")

		return nil, err
	}

	return &cfg, nil
}

func initViper() *viper.Viper {
	vpr := viper.New()

	vpr.SetConfigName(Name)
	vpr.SetConfigType(Extension)
	vpr.AddConfigPath(Path)

	return vpr
}
