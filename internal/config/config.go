package configs

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	PostgresURL  string        `env:"POSTGRES_URL"`
	ServerPort   string        `env:"SERVER_PORT"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT"`
}

func GetConfig() *Config {
	var c Config
	var ok = false
	dbUrl, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		log.Fatalln("POSTGRES_URL not set in the ENV")
	}

	c.PostgresURL = dbUrl

	port, ok := os.LookupEnv("SERVER_PORT")

	if !ok {
		port = "8000"
	}

	c.ServerPort = port

	rTimeout, ok := os.LookupEnv("READ_TIMEOUT")

	if !ok {
		c.ReadTimeout = 10 * time.Second
	} else {
		num, err := strconv.ParseInt(rTimeout, 10, strconv.IntSize)
		if err != nil {
			log.Fatalln("Error reading READ_TIMEOUT from ENV:", err)
		} else {
			c.ReadTimeout = time.Duration(num) * time.Second
		}
	}

	wTimeout, ok := os.LookupEnv("WRITE_TIMEOUT")

	if !ok {
		c.WriteTimeout = 10 * time.Second
	} else {
		num, err := strconv.ParseInt(wTimeout, 10, strconv.IntSize)
		if err != nil {
			log.Fatalln("Error reading WRITE_TIMEOUT from ENV:", err)
		} else {
			c.WriteTimeout = time.Duration(num) * time.Second
		}
	}

	return &c
}
