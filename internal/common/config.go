package common

import (
	"flag"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Dsn    string
	Port   string
	Secret string
}

func NewConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	dsn := flag.String("dsn", os.Getenv("DSN"), "Connection DSN")
	port := flag.String("port", os.Getenv("PORT"), "Port to listen on")
	secret := flag.String("secret", os.Getenv("APP_SECRET"), "App secret to encode tokens ect.")

	flag.Parse()

	time.Sleep(time.Second * 5)

	return &Config{
		Dsn:    *dsn,
		Port:   *port,
		Secret: *secret,
	}
}
