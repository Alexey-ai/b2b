package main

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/Alexey-ai/b2b"
	"github.com/Alexey-ai/b2b/pkg/handler"
	"github.com/Alexey-ai/b2b/pkg/repository"
	"github.com/Alexey-ai/b2b/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	//----logger
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cwd, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("Failed to determine working directory: %s", err)
	}
	runID := time.Now().Format("run-2022-10-20-15-04-05")
	logLocation := filepath.Join(cwd, runID+".log")
	logFile, err := os.OpenFile(logLocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatalf("Failed to open log file %s for output: %s", logLocation, err)
	}
	//Logger.SetOutput(io.MultiWriter(os.Stderr, logFile))
	logrus.SetOutput(io.MultiWriter(os.Stderr, logFile))
	logrus.WithFields(logrus.Fields{"at": "start", "log-location": logLocation}).Info()
	defer logFile.Close()
	defer logrus.Exit(0)

	//----logger

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	logrus.Infof("success start")
	srv := new(b2b.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
