package ports

type DBPort interface {
	CloseConnection() error
	AddToHistory(value int32, operation string) error
}
