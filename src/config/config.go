package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURL  string
	ServerPort   string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	return Config{
		PostgresURL:  getEnvStringRequired("POSTGRES_URL"),
		ServerPort:   getEnvStringOptional("SERVER_PORT", "8000"),
		ReadTimeout:  time.Duration(getEnvIntOptional("READ_TIMEOUT", 10)) * time.Second,
		WriteTimeout: time.Duration(getEnvIntOptional("WRITE_TIMEOUT", 10)) * time.Second,
	}
}

func getEnvStringRequired(key string) string {
	var env string
	var ok bool
	if env, ok = os.LookupEnv(key); !ok {
		log.Fatalln(key + " missing in ENV")
	}

	return env
}

func getEnvStringOptional(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}

func getEnvIntRequired(key string) int {
	var num int
	if val, ok := os.LookupEnv(key); ok {
		parsed, err := strconv.ParseInt(val, 10, strconv.IntSize)
		if err != nil {
			log.Fatalln("Error parsing "+key+" from ENV", err)
		} else {
			num = int(parsed)
		}
	} else {
		log.Fatalln(key + " missing in ENV")
	}
	return num
}

func getEnvIntOptional(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		num, err := strconv.ParseInt(val, 10, strconv.IntSize)
		if err == nil {
			return int(num)
		} else {
			log.Println("Error parsing "+key+" from ENV", err)
		}
	}
	return fallback
}
