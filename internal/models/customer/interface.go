package customer

// CustomerI  customer interface
type CustomerI interface {
	GetTableName() string
	CustomerID() string
	CustomerName() string
	CustomerPhone() string
}
