package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/cheeeasy2501/go-form-builder/cmd/app"
	"github.com/cheeeasy2501/go-form-builder/config"
	"github.com/cheeeasy2501/go-form-builder/pkg/database"
	"github.com/joho/godotenv"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		cancel()
	}()

	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer l.Sync() // flushes buffer, if any
	sugar := l.Sugar()

	if err := godotenv.Load(".env"); err != nil {
		l.Fatal("No .env file found")
	}

	sugar.Infoln(".env variables is loaded")

	// инициализируем базу данных и конфиг
	config := config.NewConfig()
	db := database.NewDB(config.Database)

	conn, err := db.OpenConnection()
	if err != nil {
		sugar.Fatal("Connection isn't opened!", err)
	}
	defer db.CloseConnection()
	sugar.Infoln("Database connection is opened")

	app.Run(ctx, sugar, config, conn)
	sugar.Infoln("App is started")
	<-ctx.Done()
	sugar.Info("App is stopped")
}
