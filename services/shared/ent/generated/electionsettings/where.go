// Code generated by ent, DO NOT EDIT.

package electionsettings

import (
	"shared/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldCreateTime, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldIsActive, v))
}

// IsAnonymous applies equality check predicate on the "is_anonymous" field. It's identical to IsAnonymousEQ.
func IsAnonymous(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldIsAnonymous, v))
}

// AllowComments applies equality check predicate on the "allow_comments" field. It's identical to AllowCommentsEQ.
func AllowComments(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldAllowComments, v))
}

// MaxVotes applies equality check predicate on the "max_votes" field. It's identical to MaxVotesEQ.
func MaxVotes(v int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldMaxVotes, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldDuration, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldLTE(FieldCreateTime, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNEQ(FieldIsActive, v))
}

// IsAnonymousEQ applies the EQ predicate on the "is_anonymous" field.
func IsAnonymousEQ(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldIsAnonymous, v))
}

// IsAnonymousNEQ applies the NEQ predicate on the "is_anonymous" field.
func IsAnonymousNEQ(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNEQ(FieldIsAnonymous, v))
}

// AllowCommentsEQ applies the EQ predicate on the "allow_comments" field.
func AllowCommentsEQ(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldAllowComments, v))
}

// AllowCommentsNEQ applies the NEQ predicate on the "allow_comments" field.
func AllowCommentsNEQ(v bool) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNEQ(FieldAllowComments, v))
}

// MaxVotesEQ applies the EQ predicate on the "max_votes" field.
func MaxVotesEQ(v int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldMaxVotes, v))
}

// MaxVotesNEQ applies the NEQ predicate on the "max_votes" field.
func MaxVotesNEQ(v int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNEQ(FieldMaxVotes, v))
}

// MaxVotesIn applies the In predicate on the "max_votes" field.
func MaxVotesIn(vs ...int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldIn(FieldMaxVotes, vs...))
}

// MaxVotesNotIn applies the NotIn predicate on the "max_votes" field.
func MaxVotesNotIn(vs ...int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNotIn(FieldMaxVotes, vs...))
}

// MaxVotesGT applies the GT predicate on the "max_votes" field.
func MaxVotesGT(v int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldGT(FieldMaxVotes, v))
}

// MaxVotesGTE applies the GTE predicate on the "max_votes" field.
func MaxVotesGTE(v int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldGTE(FieldMaxVotes, v))
}

// MaxVotesLT applies the LT predicate on the "max_votes" field.
func MaxVotesLT(v int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldLT(FieldMaxVotes, v))
}

// MaxVotesLTE applies the LTE predicate on the "max_votes" field.
func MaxVotesLTE(v int) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldLTE(FieldMaxVotes, v))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v time.Time) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.FieldLTE(FieldDuration, v))
}

// HasElection applies the HasEdge predicate on the "election" edge.
func HasElection() predicate.ElectionSettings {
	return predicate.ElectionSettings(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ElectionTable, ElectionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasElectionWith applies the HasEdge predicate on the "election" edge with a given conditions (other predicates).
func HasElectionWith(preds ...predicate.Election) predicate.ElectionSettings {
	return predicate.ElectionSettings(func(s *sql.Selector) {
		step := newElectionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ElectionSettings) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ElectionSettings) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ElectionSettings) predicate.ElectionSettings {
	return predicate.ElectionSettings(sql.NotPredicates(p))
}
