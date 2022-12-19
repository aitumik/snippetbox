package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Addr        string
	StaticDir   string
	SecretKey   string
	DatabaseURI string
}

func (c *Config) getVarFromEnv(envName string) (env string, err error) {
	env, exists := os.LookupEnv(envName)
	if !exists {
		err = fmt.Errorf("config: env variable %s is not set", envName)
	}
	return
}

func (c *Config) loadEnvVariables() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	c.StaticDir, err = c.getVarFromEnv("STATIC_DIR")
	if err != nil {
		log.Error(err)
		return err
	}
	c.SecretKey, err = c.getVarFromEnv("SECRET_KEY")
	if err != nil {
		log.Error(err)
		return err
	}
	c.Addr, err = c.getVarFromEnv("APP_ADDR")
	if err != nil {
		log.Error(err)
		return err
	}
	c.DatabaseURI, err = c.getVarFromEnv("DATABASE_URI")
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func NewConfig() (config *Config, err error) {
	config = &Config{}
	err = config.loadEnvVariables()
	return
}
