package datasource

import (
	"fmt"
	"log/slog"
	"t_astrum/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(cfg config.Database, log *slog.Logger) *sqlx.DB {
	fmt.Println("func NewDatabase(cfg config.Database, log *slog.Logger) *sqlx.DB")

	db, errOpen := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if errOpen != nil {
		log.Error(errOpen.Error())
		return nil
	} else {
		log.Info("Database opened")
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Error(errPing.Error())
		return nil
	} else {
		log.Error("DB pinged")
	}

	return db
}
