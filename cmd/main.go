package main

import (
	"errors"
	"flag"
	"fmt"
	"go-starter-clean/pkg/entity"
	"go-starter-clean/pkg/logger"
	"go-starter-clean/pkg/repository"

	"go-starter-clean/pkg/config"
	"go-starter-clean/pkg/http"
	"go-starter-clean/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	var cfgApp *config.Config
	cfgApp, _ = config.NewConfig()
	cfgFlag := config.Config{}
	flag.StringVar(&cfgFlag.AppName, "name", "", "-name	Name of application to set default application name")
	flag.StringVar(&cfgFlag.Port, "port", "", "-port 	Port of application for set application set to default port initial, default: 9090")
	flag.StringVar(&cfgFlag.Logger.LogLevel, "log-level", "info", "-log-level	Set the logging level ('debug'|'info'|'warn'|'error'|'fatal') (default 'info')")
	flag.StringVar(&cfgFlag.Logger.Environment, "log-env", "prod", "-log-env	Set the logging environment ('prod'|'dev') (default 'prod')")
	flag.StringVar(&cfgFlag.DB.Use, "db-driver", "postgres", "-db-driver	Set the driver database for connector application, default: postgres")
	flag.StringVar(&cfgFlag.DB.Host, "db-host", "localhost", "-db-host	Set the host of database, default: localhost")
	flag.StringVar(&cfgFlag.DB.Port, "db-port", "5432", "-db-port	Set the port of database, default: 5432")
	flag.StringVar(&cfgFlag.DB.UserName, "db-username", "root", "-db-username	Set the username for crendentials database, default: root")
	flag.StringVar(&cfgFlag.DB.Password, "db-password", "root", "-db-password	Set the password for crendentials database, default: root")
	flag.StringVar(&cfgFlag.DB.Database, "db-name", "default", "-db-name	Set the database name, default: default")
	flag.Parse()
	if isFlagPassed("name") {
		cfgApp.AppName = cfgFlag.AppName
	}
	if isFlagPassed("port") {
		cfgApp.Port = cfgFlag.Port
	}
	if isFlagPassed("log-level") {
		cfgApp.Logger.LogLevel = cfgFlag.Logger.LogLevel
	}
	if isFlagPassed("log-env") {
		cfgApp.Logger.Environment = cfgFlag.Logger.Environment
	}
	if isFlagPassed("db-driver") {
		cfgApp.DB.Use = cfgFlag.DB.Use
	}
	if isFlagPassed("db-host") {
		cfgApp.DB.Host = cfgFlag.DB.Host
	}
	if isFlagPassed("db-port") {
		cfgApp.DB.Port = cfgFlag.DB.Port
	}
	if isFlagPassed("db-username") {
		cfgApp.DB.UserName = cfgFlag.DB.UserName
	}
	if isFlagPassed("db-password") {
		cfgApp.DB.Password = cfgFlag.DB.Password
	}
	if isFlagPassed("db-name") {
		cfgApp.DB.Database = cfgFlag.DB.Database
	}
	// Load config application with configor
	errCfgApp := configor.Load(cfgApp)
	if errCfgApp != nil {
		panic(errors.New(errCfgApp.Error()))
	}
	// initial logger with parameter config app
	l, errLog := logger.NewLogger(cfgApp)
	if errLog != nil {
		panic(errors.New(errLog.Error()))
	}

	// Check Driver Database for connecting to Gorm
	var connection string
	if cfgApp.DB.Use == "mysql" {
		connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf-8&parseTime=True&loc=Local",
			cfgApp.DB.UserName,
			cfgApp.DB.Password,
			cfgApp.DB.Host,
			cfgApp.DB.Port,
			cfgApp.DB.Database)
	} else if cfgApp.DB.Use == "postgres" {
		connection = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfgApp.DB.Host,
			cfgApp.DB.Port,
			cfgApp.DB.UserName,
			cfgApp.DB.Password,
			cfgApp.DB.Database)
	} else {
		panic(errors.New(fmt.Sprintf("Application not supported for driver : %s", cfgApp.DB.Use)))
	}

	DB, err := gorm.Open(cfgApp.DB.Use, connection)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(entity.User{})

	server := gin.Default()
	usr := http.NewUserHandler(l, service.New(repository.NewUserRepository(DB, l)))
	server.GET("/users", usr.GetAll)
	server.POST("/users", usr.Store)

	server.Run(":9090")
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
