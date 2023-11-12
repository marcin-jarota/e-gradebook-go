package app

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Dsn          string
	Port         string
	RedisAddr    string
	Secret       string
	BaseUrl      string
	BaseUrlFront string
}

func NewConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	dsn := flag.String("dsn", os.Getenv("DSN"), "Connection DSN")
	port := flag.String("port", os.Getenv("PORT"), "Port to listen on")
	secret := flag.String("secret", os.Getenv("APP_SECRET"), "App secret to encode tokens ect.")
	redisAddr := flag.String("redisAddr", os.Getenv("REDIS_ADDR"), "Addres of redis")
	baseUrl := flag.String("baseUrl", os.Getenv("BASE_URL"), "Addres of backend application")
	baseUrlFront := flag.String("baseUrlFront", os.Getenv("BASE_URL_FRONT"), "Addres of frontend application")

	flag.Parse()

	// time.Sleep(time.Second * 5)

	return &Config{
		Dsn:          *dsn,
		Port:         *port,
		Secret:       *secret,
		RedisAddr:    *redisAddr,
		BaseUrl:      *baseUrl,
		BaseUrlFront: *baseUrlFront,
	}
}
