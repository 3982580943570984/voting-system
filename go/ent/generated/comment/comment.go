// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldContents holds the string denoting the contents field in the database.
	FieldContents = "contents"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeElection holds the string denoting the election edge name in mutations.
	EdgeElection = "election"
	// Table holds the table name of the comment in the database.
	Table = "comments"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "comments"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "comment_children"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "comments"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "comment_children"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "comments"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_comments"
	// ElectionTable is the table that holds the election relation/edge.
	ElectionTable = "comments"
	// ElectionInverseTable is the table name for the Election entity.
	// It exists in this package in order to avoid circular dependency with the "election" package.
	ElectionInverseTable = "elections"
	// ElectionColumn is the table column denoting the election relation/edge.
	ElectionColumn = "election_comments"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldContents,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "comments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"comment_children",
	"election_comments",
	"user_comments",
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// ContentsValidator is a validator for the "contents" field. It is called by the builders before save.
	ContentsValidator func(string) error
)

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByContents orders the results by the contents field.
func ByContents(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContents, opts...).ToFunc()
}

// ByParentField orders the results by parent field.
func ByParentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParentStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByElectionField orders the results by election field.
func ByElectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newElectionStep(), sql.OrderByField(field, opts...))
	}
}
func newParentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newElectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ElectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ElectionTable, ElectionColumn),
	)
}
