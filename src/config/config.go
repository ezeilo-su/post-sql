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
	if dbUrl, ok := os.LookupEnv("POSTGRES_URL"); !ok || dbUrl == "" {
		log.Fatalln("POSTGRES_URL not set in the ENV")
	} else {
		c.PostgresURL = dbUrl
	}

	if port, ok := os.LookupEnv("SERVER_PORT"); !ok || port == "" {
		log.Println("SERVER_PORT not set in the ENV " + port)
		c.ServerPort = "8000"
	} else {
		log.Println("PORT: " + port)
		c.ServerPort = port
	}

	if rTimeout, ok := os.LookupEnv("READ_TIMEOUT"); !ok || rTimeout == "" {
		c.ReadTimeout = 10 * time.Second
	} else {
		num, err := strconv.ParseInt(rTimeout, 10, strconv.IntSize)
		if err != nil {
			log.Fatalln("Error reading READ_TIMEOUT from ENV:", err)
		} else {
			c.ReadTimeout = time.Duration(num) * time.Second
		}
	}

	if wTimeout, ok := os.LookupEnv("WRITE_TIMEOUT"); !ok || wTimeout == "" {
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
