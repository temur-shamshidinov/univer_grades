package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/univer_grades/config"
)

func ConnectToDb(pgCfg config.Config) (*pgx.Conn, error) {

	var ctx = context.Background()

	url := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		pgCfg.PgConfig.Username,
		pgCfg.PgConfig.Password,
		pgCfg.PgConfig.Host,
		pgCfg.PgConfig.Port,
		pgCfg.PgConfig.DatabaseName,
	)

	conn, err := pgx.Connect(ctx,url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v \n", err)
		return nil, err
	}

	return conn, nil
}