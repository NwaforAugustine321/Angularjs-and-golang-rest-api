package config

import (
	"database/sql"
	"log"
	"os"
)

type configs struct {
	Port string
	env  string
}

type Application struct {
	configs
	Logger *log.Logger
}

func GetAppconfig() *Application {
	cfg := configs{":4000", ""}
	logger := log.New(os.Stdout, "", log.Ltime)
	return &Application{cfg, logger}
}

type Config struct {
}

var appConfigure = Config{}

type AppContext struct {
	DB *sql.DB
	Config
}

func NewAppConfigurationContex(db *sql.DB) *AppContext {
	return &AppContext{Config : appConfigure,DB : db}
}
