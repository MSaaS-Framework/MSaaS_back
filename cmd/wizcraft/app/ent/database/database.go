// Code generated by ent, DO NOT EDIT.

package database

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the database type in the database.
	Label = "database"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldConnectionPath holds the string denoting the connection_path field in the database.
	FieldConnectionPath = "connection_path"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldDbType holds the string denoting the db_type field in the database.
	FieldDbType = "db_type"
	// EdgeService holds the string denoting the service edge name in mutations.
	EdgeService = "service"
	// EdgeGeneralspec holds the string denoting the generalspec edge name in mutations.
	EdgeGeneralspec = "generalspec"
	// Table holds the table name of the database in the database.
	Table = "databases"
	// ServiceTable is the table that holds the service relation/edge.
	ServiceTable = "databases"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "service_databases"
	// GeneralspecTable is the table that holds the generalspec relation/edge.
	GeneralspecTable = "databases"
	// GeneralspecInverseTable is the table name for the GeneralSpec entity.
	// It exists in this package in order to avoid circular dependency with the "generalspec" package.
	GeneralspecInverseTable = "general_specs"
	// GeneralspecColumn is the table column denoting the generalspec relation/edge.
	GeneralspecColumn = "general_spec_database"
)

// Columns holds all SQL columns for database fields.
var Columns = []string{
	FieldID,
	FieldConnectionPath,
	FieldPassword,
	FieldDbType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "databases"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"general_spec_database",
	"service_databases",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ConnectionPathValidator is a validator for the "connection_path" field. It is called by the builders before save.
	ConnectionPathValidator func(string) error
	// DbTypeValidator is a validator for the "db_type" field. It is called by the builders before save.
	DbTypeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Database queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByConnectionPath orders the results by the connection_path field.
func ByConnectionPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConnectionPath, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByDbType orders the results by the db_type field.
func ByDbType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDbType, opts...).ToFunc()
}

// ByServiceField orders the results by service field.
func ByServiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newServiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByGeneralspecField orders the results by generalspec field.
func ByGeneralspecField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGeneralspecStep(), sql.OrderByField(field, opts...))
	}
}
func newServiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ServiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ServiceTable, ServiceColumn),
	)
}
func newGeneralspecStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GeneralspecInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, GeneralspecTable, GeneralspecColumn),
	)
}
