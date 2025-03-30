package db

import (
	"context"
	"fmt"

	"github.com/huuloc2026/go-social/internal/config"
	"github.com/huuloc2026/go-social/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

type PostgresDB struct {
	conn *pgx.Conn
}

func NewPostgresDB(cfg config.DBConfig) (*PostgresDB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{conn: conn}, nil
}

func (db *PostgresDB) Save(user *domain.User) error {
	query := "INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id"
	err := db.conn.QueryRow(context.Background(), query, user.Username, user.Email).Scan(&user.ID)
	return err
}

func (db *PostgresDB) FindByID(id int) (*domain.User, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	user := &domain.User{}
	err := db.conn.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
