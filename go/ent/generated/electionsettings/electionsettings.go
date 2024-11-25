// Code generated by ent, DO NOT EDIT.

package electionsettings

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the electionsettings type in the database.
	Label = "election_settings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldIsAnonymous holds the string denoting the is_anonymous field in the database.
	FieldIsAnonymous = "is_anonymous"
	// FieldAllowComments holds the string denoting the allow_comments field in the database.
	FieldAllowComments = "allow_comments"
	// FieldMaxVotes holds the string denoting the max_votes field in the database.
	FieldMaxVotes = "max_votes"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// EdgeElection holds the string denoting the election edge name in mutations.
	EdgeElection = "election"
	// Table holds the table name of the electionsettings in the database.
	Table = "election_settings"
	// ElectionTable is the table that holds the election relation/edge.
	ElectionTable = "election_settings"
	// ElectionInverseTable is the table name for the Election entity.
	// It exists in this package in order to avoid circular dependency with the "election" package.
	ElectionInverseTable = "elections"
	// ElectionColumn is the table column denoting the election relation/edge.
	ElectionColumn = "election_settings"
)

// Columns holds all SQL columns for electionsettings fields.
var Columns = []string{
	FieldID,
	FieldIsActive,
	FieldIsAnonymous,
	FieldAllowComments,
	FieldMaxVotes,
	FieldStartDate,
	FieldEndDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "election_settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"election_settings",
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
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultIsAnonymous holds the default value on creation for the "is_anonymous" field.
	DefaultIsAnonymous bool
	// DefaultAllowComments holds the default value on creation for the "allow_comments" field.
	DefaultAllowComments bool
	// DefaultMaxVotes holds the default value on creation for the "max_votes" field.
	DefaultMaxVotes int
	// MaxVotesValidator is a validator for the "max_votes" field. It is called by the builders before save.
	MaxVotesValidator func(int) error
	// DefaultStartDate holds the default value on creation for the "start_date" field.
	DefaultStartDate func() time.Time
	// DefaultEndDate holds the default value on creation for the "end_date" field.
	DefaultEndDate time.Time
)

// OrderOption defines the ordering options for the ElectionSettings queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsActive orders the results by the is_active field.
func ByIsActive(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsActive, opts...).ToFunc()
}

// ByIsAnonymous orders the results by the is_anonymous field.
func ByIsAnonymous(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAnonymous, opts...).ToFunc()
}

// ByAllowComments orders the results by the allow_comments field.
func ByAllowComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAllowComments, opts...).ToFunc()
}

// ByMaxVotes orders the results by the max_votes field.
func ByMaxVotes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxVotes, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByElectionField orders the results by election field.
func ByElectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newElectionStep(), sql.OrderByField(field, opts...))
	}
}
func newElectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ElectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ElectionTable, ElectionColumn),
	)
}
