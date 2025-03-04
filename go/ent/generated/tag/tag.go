// Code generated by ent, DO NOT EDIT.

package tag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeElections holds the string denoting the elections edge name in mutations.
	EdgeElections = "elections"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// ElectionsTable is the table that holds the elections relation/edge. The primary key declared below.
	ElectionsTable = "tag_elections"
	// ElectionsInverseTable is the table name for the Election entity.
	// It exists in this package in order to avoid circular dependency with the "election" package.
	ElectionsInverseTable = "elections"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ElectionsPrimaryKey and ElectionsColumn2 are the table columns denoting the
	// primary key for the elections relation (M2M).
	ElectionsPrimaryKey = []string{"tag_id", "election_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByElectionsCount orders the results by elections count.
func ByElectionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newElectionsStep(), opts...)
	}
}

// ByElections orders the results by elections terms.
func ByElections(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newElectionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newElectionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ElectionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ElectionsTable, ElectionsPrimaryKey...),
	)
}
