package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"t_astrum/internal/config"
	v1 "t_astrum/internal/promo/handlers/http"
	"t_astrum/pkg/datasource"
	"time"
)

type App struct {
	db  *sqlx.DB
	gin *gin.Engine
	cfg *config.Config
	log *slog.Logger
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		db:  datasource.NewDatabase(cfg.Database, log),
		gin: v1.NewGinRouter(),
		cfg: cfg,
		log: log,
	}
}

func (app *App) Run() error {
	if err := app.InitService(); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	server := &http.Server{
		Addr:    app.cfg.HTTPServer.Address,
		Handler: app.gin,
	}

	go func() {
		<-quit
		app.log.Info("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			app.log.Error("Server Shutdown:", err)
		}

		// Close the database connection
		if err := app.db.Close(); err != nil {
			app.log.Error("Database connection close failed: %v", err)
		} else {
			app.log.Info("Database connection closed")
		}
	}()

	app.log.Info("Server is starting on port", slog.String("addr", server.Addr))

	if err := app.gin.Run(server.Addr); err != nil {
		app.log.Error("Error starting server: %v", err)
		return err
	}

	return nil
}
