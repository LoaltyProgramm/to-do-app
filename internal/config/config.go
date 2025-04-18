package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseFile string
}

func (c *Config) GetEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return fmt.Errorf("env file is not reading: %v", err)
	}

	c.Port = os.Getenv("PORT")
	c.Port = os.Getenv("TODO_DBFILE")

	return nil
}