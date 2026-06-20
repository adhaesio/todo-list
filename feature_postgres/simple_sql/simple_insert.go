package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
INSERT INTO tasks (title, description, completed, created_at)
VALUES ('Завтрак','сделать домашку по арабскому до 20.10.26', FALSE, '2026-06-20 18:00:05')

	`
	_, err := conn.Exec(ctx, sqlQuery)

	return err

}
