// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"voting-system/ent/generated/election"
	"voting-system/ent/generated/electionfilters"
	"voting-system/ent/generated/electionsettings"
	"voting-system/ent/generated/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Election is the model entity for the Election schema.
type Election struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Completed holds the value of the "completed" field.
	Completed bool `json:"completed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ElectionQuery when eager-loading is set.
	Edges          ElectionEdges `json:"edges"`
	user_elections *int
	selectValues   sql.SelectValues
}

// ElectionEdges holds the relations/edges for other nodes in the graph.
type ElectionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Candidates holds the value of the candidates edge.
	Candidates []*Candidate `json:"candidates,omitempty"`
	// Settings holds the value of the settings edge.
	Settings *ElectionSettings `json:"settings,omitempty"`
	// Filters holds the value of the filters edge.
	Filters *ElectionFilters `json:"filters,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ElectionEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ElectionEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[1] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e ElectionEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[2] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// CandidatesOrErr returns the Candidates value or an error if the edge
// was not loaded in eager-loading.
func (e ElectionEdges) CandidatesOrErr() ([]*Candidate, error) {
	if e.loadedTypes[3] {
		return e.Candidates, nil
	}
	return nil, &NotLoadedError{edge: "candidates"}
}

// SettingsOrErr returns the Settings value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ElectionEdges) SettingsOrErr() (*ElectionSettings, error) {
	if e.Settings != nil {
		return e.Settings, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: electionsettings.Label}
	}
	return nil, &NotLoadedError{edge: "settings"}
}

// FiltersOrErr returns the Filters value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ElectionEdges) FiltersOrErr() (*ElectionFilters, error) {
	if e.Filters != nil {
		return e.Filters, nil
	} else if e.loadedTypes[5] {
		return nil, &NotFoundError{label: electionfilters.Label}
	}
	return nil, &NotLoadedError{edge: "filters"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Election) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case election.FieldCompleted:
			values[i] = new(sql.NullBool)
		case election.FieldID:
			values[i] = new(sql.NullInt64)
		case election.FieldTitle, election.FieldDescription:
			values[i] = new(sql.NullString)
		case election.ForeignKeys[0]: // user_elections
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Election fields.
func (e *Election) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case election.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case election.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				e.Title = value.String
			}
		case election.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		case election.FieldCompleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field completed", values[i])
			} else if value.Valid {
				e.Completed = value.Bool
			}
		case election.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_elections", value)
			} else if value.Valid {
				e.user_elections = new(int)
				*e.user_elections = int(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Election.
// This includes values selected through modifiers, order, etc.
func (e *Election) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Election entity.
func (e *Election) QueryUser() *UserQuery {
	return NewElectionClient(e.config).QueryUser(e)
}

// QueryTags queries the "tags" edge of the Election entity.
func (e *Election) QueryTags() *TagQuery {
	return NewElectionClient(e.config).QueryTags(e)
}

// QueryComments queries the "comments" edge of the Election entity.
func (e *Election) QueryComments() *CommentQuery {
	return NewElectionClient(e.config).QueryComments(e)
}

// QueryCandidates queries the "candidates" edge of the Election entity.
func (e *Election) QueryCandidates() *CandidateQuery {
	return NewElectionClient(e.config).QueryCandidates(e)
}

// QuerySettings queries the "settings" edge of the Election entity.
func (e *Election) QuerySettings() *ElectionSettingsQuery {
	return NewElectionClient(e.config).QuerySettings(e)
}

// QueryFilters queries the "filters" edge of the Election entity.
func (e *Election) QueryFilters() *ElectionFiltersQuery {
	return NewElectionClient(e.config).QueryFilters(e)
}

// Update returns a builder for updating this Election.
// Note that you need to call Election.Unwrap() before calling this method if this Election
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Election) Update() *ElectionUpdateOne {
	return NewElectionClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Election entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Election) Unwrap() *Election {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("generated: Election is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Election) String() string {
	var builder strings.Builder
	builder.WriteString("Election(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("title=")
	builder.WriteString(e.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(e.Description)
	builder.WriteString(", ")
	builder.WriteString("completed=")
	builder.WriteString(fmt.Sprintf("%v", e.Completed))
	builder.WriteByte(')')
	return builder.String()
}

// Elections is a parsable slice of Election.
type Elections []*Election
