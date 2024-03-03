package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:app`
		HTTP `yaml:http`
		SMS  `yaml:sms`
	}

	App struct {
		PathToData string `env-required:"true" yaml:"path_to_data"    env:"PATH_TO_DATA"`
	}

	HTTP struct {
		LocalPort string `env-required:"true" yaml:"local_port" env:"LOCAL_PORT"`
		DataPort  string `env-required:"true" yaml:"data_port" env:"DATA_PORT"`
		LocalHost string `env-required:"true" yaml:"local_host" env:"LOCAL_HOST"`
		DataHost  string `env-required:"true" yaml:"data_host" env:"DATA_HOST"`
	}

	SMS struct {
		SMSDataFile       string   `env-required:"true" yaml:"sms_data_file" env:"SMS_DATA_FILE"`
		SMSNumberOfFields int      `env-required:"true" yaml:"sms_number_of_fields" env:"SMS_NUMBER_OF_FIELDS"`
		SMSValidProviders []string `env-required:"true" yaml:"sms_valid_providers" env:"SMS_VALID_PROVIDERS"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
