package database

// DBHandler defines the interface for a generic database handler
type DBHandler interface {
	Execute(statement string, args ...interface{}) error
	Query(statement string, args ...interface{}) (Row, error)
	Close() error
}

// Row defines the interface for a generic database row
type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}
