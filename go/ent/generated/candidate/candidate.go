// Code generated by ent, DO NOT EDIT.

package candidate

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the candidate type in the database.
	Label = "candidate"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPhotoURL holds the string denoting the photo_url field in the database.
	FieldPhotoURL = "photo_url"
	// FieldVotesCount holds the string denoting the votes_count field in the database.
	FieldVotesCount = "votes_count"
	// EdgeElection holds the string denoting the election edge name in mutations.
	EdgeElection = "election"
	// EdgeVotes holds the string denoting the votes edge name in mutations.
	EdgeVotes = "votes"
	// Table holds the table name of the candidate in the database.
	Table = "candidates"
	// ElectionTable is the table that holds the election relation/edge.
	ElectionTable = "candidates"
	// ElectionInverseTable is the table name for the Election entity.
	// It exists in this package in order to avoid circular dependency with the "election" package.
	ElectionInverseTable = "elections"
	// ElectionColumn is the table column denoting the election relation/edge.
	ElectionColumn = "election_candidates"
	// VotesTable is the table that holds the votes relation/edge.
	VotesTable = "votes"
	// VotesInverseTable is the table name for the Vote entity.
	// It exists in this package in order to avoid circular dependency with the "vote" package.
	VotesInverseTable = "votes"
	// VotesColumn is the table column denoting the votes relation/edge.
	VotesColumn = "candidate_votes"
)

// Columns holds all SQL columns for candidate fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldPhotoURL,
	FieldVotesCount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "candidates"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"election_candidates",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultVotesCount holds the default value on creation for the "votes_count" field.
	DefaultVotesCount int
	// VotesCountValidator is a validator for the "votes_count" field. It is called by the builders before save.
	VotesCountValidator func(int) error
)

// OrderOption defines the ordering options for the Candidate queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByPhotoURL orders the results by the photo_url field.
func ByPhotoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL, opts...).ToFunc()
}

// ByVotesCountField orders the results by the votes_count field.
func ByVotesCountField(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVotesCount, opts...).ToFunc()
}

// ByElectionField orders the results by election field.
func ByElectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newElectionStep(), sql.OrderByField(field, opts...))
	}
}

// ByVotesCount orders the results by votes count.
func ByVotesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVotesStep(), opts...)
	}
}

// ByVotes orders the results by votes terms.
func ByVotes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVotesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newElectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ElectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ElectionTable, ElectionColumn),
	)
}
func newVotesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VotesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VotesTable, VotesColumn),
	)
}
