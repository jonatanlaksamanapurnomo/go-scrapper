package postgres

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"toped-scrapper/pkg/database"
)

// PostgresHandler struct for PostgreSQL that implements DBHandler
type PostgresHandler struct {
	Conn *sql.DB
}

// NewPostgresHandler creates a new handler for PostgreSQL operations
func NewPostgresHandler(dataSourceName string) (*PostgresHandler, error) {
	conn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &PostgresHandler{Conn: conn}, nil
}

// Execute executes a query without returning any rows
func (handler *PostgresHandler) Execute(statement string, args ...interface{}) error {
	_, err := handler.Conn.Exec(statement, args...)
	return err
}

// Query executes a query that returns rows, typically a SELECT
func (handler *PostgresHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	return handler.Conn.Query(statement, args...)
}

// Close closes the database connection
func (handler *PostgresHandler) Close() error {
	return handler.Conn.Close()
}

func (handler *PostgresHandler) CheckAndCreateTable() error {
	tableExistsQuery := `SELECT to_regclass('public.products')`
	var tableName string
	err := handler.Conn.QueryRow(tableExistsQuery).Scan(&tableName)
	if err != nil || tableName == "" {
		createTableSQL := `
			CREATE TABLE products (
				id SERIAL PRIMARY KEY,
				name TEXT NOT NULL,
				description TEXT,
				image_link TEXT,
				rating TEXT,
				price TEXT,
				store_name TEXT
			);`
		if _, err := handler.Conn.Exec(createTableSQL); err != nil {
			log.Fatalf("Failed to create table: %v", err)
			return err
		}
	}
	return nil
}
