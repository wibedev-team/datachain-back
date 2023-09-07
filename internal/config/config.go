package config

import (
	"github.com/wibedev-team/datachain-back/pkg/db/postgresql"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgresql struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	} `yaml:"postgresql"`
}

func New(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	log.Println(cfg)
	return &cfg
}

func Init() *postgresql.PgConfig {
	args := os.Args
	if len(args) != 2 {
		if os.Getenv("POSTGRES_DB") == "" {
			log.Fatalf("provide path to congig file")
		}
	}

	var cfg *Config
	var pgCfg *postgresql.PgConfig

	if os.Getenv("POSTGRES_HOST") == "" {
		cfg = New(args[1])

		pgCfg = postgresql.NewConfig(
			cfg.Postgresql.Username,
			cfg.Postgresql.Password,
			cfg.Postgresql.Host,
			cfg.Postgresql.Port,
			cfg.Postgresql.Database,
		)
	} else {
		pgCfg = postgresql.NewConfig(
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		)
	}

	return pgCfg
}
