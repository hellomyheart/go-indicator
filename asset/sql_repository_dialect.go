package asset

// SQLRepositoryDialect defines the SQL dialect for the SQL repository.
type SQLRepositoryDialect interface {
	// CreateTable returns the SQL statement to create the repository table.
	CreateTable() string

	// DropTable returns the SQL statement to drop the repository table.
	DropTable() string

	// Assets returns the SQL statement to get the names of all assets in the respository.
	Assets() string

	// GetSince returns the SQL statement to query snapshots for the asset with the given name since the given date.
	GetSince() string

	// LastDate returns the SQL statement to query for the last date for the asset with the given name.
	LastDate() string

	// Appends returns the SQL statement to add the given snapshots to the asset with the given name.
	Append() string
}
