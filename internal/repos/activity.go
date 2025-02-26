package repos

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lexantus/useless_project/configs"
	"github.com/lexantus/useless_project/internal/models"
	_ "github.com/lib/pq"
)

type ActivityRepo struct {
	db *sql.DB
}

func (r *ActivityRepo) GetActivity(ctx context.Context, timespanMs int) (models.Activity, error) {
	row := r.db.QueryRow("SELECT id, timespan_ms, description from activity WHERE timespan_ms = $1", timespanMs)
	var a models.Activity
	err := row.Scan(&a.ID, &a.TimespanMs, &a.Desc)
	return a, err
}

func NewActivityRepo(conf configs.DBConfig) (*ActivityRepo, error) {
	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", conf.User, conf.Password, conf.DBName, conf.Port, conf.SSLMode)
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return nil, err
	}
	return &ActivityRepo{db: db}, nil
}
