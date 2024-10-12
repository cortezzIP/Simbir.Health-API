package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var hospitalDBPool *pgxpool.Pool

func ConnectToHospitalDB(cfg *config.DatabaseConfig) {
	var err error
	
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	hospitalDBPool, err = pgxpool.New(context.TODO(), connectionString)
	if err != nil {
		slog.Error("Unable to create hospital connection pool: " + err.Error())
		os.Exit(1001)
	}

	err = hospitalDBPool.Ping(context.TODO())
	if err != nil {
		slog.Error("Unable to connect to database: " + err.Error())
		os.Exit(1002)
	}

	slog.Info("Successful connection to the database")
}

func GetHospitalDB() *pgxpool.Pool {
	return hospitalDBPool
}

func CloseHospitalDB() {
	if hospitalDBPool != nil {
		hospitalDBPool.Close()
	}
}