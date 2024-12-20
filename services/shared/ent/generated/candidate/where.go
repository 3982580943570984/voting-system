// Code generated by ent, DO NOT EDIT.

package candidate

import (
	"shared/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Candidate {
	return predicate.Candidate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Candidate {
	return predicate.Candidate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Candidate {
	return predicate.Candidate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Candidate {
	return predicate.Candidate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Candidate {
	return predicate.Candidate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Candidate {
	return predicate.Candidate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Candidate {
	return predicate.Candidate(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldDescription, v))
}

// PhotoURL applies equality check predicate on the "photo_url" field. It's identical to PhotoURLEQ.
func PhotoURL(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldPhotoURL, v))
}

// VotesCount applies equality check predicate on the "votes_count" field. It's identical to VotesCountEQ.
func VotesCount(v int) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldVotesCount, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Candidate {
	return predicate.Candidate(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Candidate {
	return predicate.Candidate(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Candidate {
	return predicate.Candidate(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Candidate {
	return predicate.Candidate(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Candidate {
	return predicate.Candidate(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldContainsFold(FieldDescription, v))
}

// PhotoURLEQ applies the EQ predicate on the "photo_url" field.
func PhotoURLEQ(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldPhotoURL, v))
}

// PhotoURLNEQ applies the NEQ predicate on the "photo_url" field.
func PhotoURLNEQ(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldNEQ(FieldPhotoURL, v))
}

// PhotoURLIn applies the In predicate on the "photo_url" field.
func PhotoURLIn(vs ...string) predicate.Candidate {
	return predicate.Candidate(sql.FieldIn(FieldPhotoURL, vs...))
}

// PhotoURLNotIn applies the NotIn predicate on the "photo_url" field.
func PhotoURLNotIn(vs ...string) predicate.Candidate {
	return predicate.Candidate(sql.FieldNotIn(FieldPhotoURL, vs...))
}

// PhotoURLGT applies the GT predicate on the "photo_url" field.
func PhotoURLGT(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldGT(FieldPhotoURL, v))
}

// PhotoURLGTE applies the GTE predicate on the "photo_url" field.
func PhotoURLGTE(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldGTE(FieldPhotoURL, v))
}

// PhotoURLLT applies the LT predicate on the "photo_url" field.
func PhotoURLLT(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldLT(FieldPhotoURL, v))
}

// PhotoURLLTE applies the LTE predicate on the "photo_url" field.
func PhotoURLLTE(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldLTE(FieldPhotoURL, v))
}

// PhotoURLContains applies the Contains predicate on the "photo_url" field.
func PhotoURLContains(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldContains(FieldPhotoURL, v))
}

// PhotoURLHasPrefix applies the HasPrefix predicate on the "photo_url" field.
func PhotoURLHasPrefix(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldHasPrefix(FieldPhotoURL, v))
}

// PhotoURLHasSuffix applies the HasSuffix predicate on the "photo_url" field.
func PhotoURLHasSuffix(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldHasSuffix(FieldPhotoURL, v))
}

// PhotoURLIsNil applies the IsNil predicate on the "photo_url" field.
func PhotoURLIsNil() predicate.Candidate {
	return predicate.Candidate(sql.FieldIsNull(FieldPhotoURL))
}

// PhotoURLNotNil applies the NotNil predicate on the "photo_url" field.
func PhotoURLNotNil() predicate.Candidate {
	return predicate.Candidate(sql.FieldNotNull(FieldPhotoURL))
}

// PhotoURLEqualFold applies the EqualFold predicate on the "photo_url" field.
func PhotoURLEqualFold(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldEqualFold(FieldPhotoURL, v))
}

// PhotoURLContainsFold applies the ContainsFold predicate on the "photo_url" field.
func PhotoURLContainsFold(v string) predicate.Candidate {
	return predicate.Candidate(sql.FieldContainsFold(FieldPhotoURL, v))
}

// VotesCountEQ applies the EQ predicate on the "votes_count" field.
func VotesCountEQ(v int) predicate.Candidate {
	return predicate.Candidate(sql.FieldEQ(FieldVotesCount, v))
}

// VotesCountNEQ applies the NEQ predicate on the "votes_count" field.
func VotesCountNEQ(v int) predicate.Candidate {
	return predicate.Candidate(sql.FieldNEQ(FieldVotesCount, v))
}

// VotesCountIn applies the In predicate on the "votes_count" field.
func VotesCountIn(vs ...int) predicate.Candidate {
	return predicate.Candidate(sql.FieldIn(FieldVotesCount, vs...))
}

// VotesCountNotIn applies the NotIn predicate on the "votes_count" field.
func VotesCountNotIn(vs ...int) predicate.Candidate {
	return predicate.Candidate(sql.FieldNotIn(FieldVotesCount, vs...))
}

// VotesCountGT applies the GT predicate on the "votes_count" field.
func VotesCountGT(v int) predicate.Candidate {
	return predicate.Candidate(sql.FieldGT(FieldVotesCount, v))
}

// VotesCountGTE applies the GTE predicate on the "votes_count" field.
func VotesCountGTE(v int) predicate.Candidate {
	return predicate.Candidate(sql.FieldGTE(FieldVotesCount, v))
}

// VotesCountLT applies the LT predicate on the "votes_count" field.
func VotesCountLT(v int) predicate.Candidate {
	return predicate.Candidate(sql.FieldLT(FieldVotesCount, v))
}

// VotesCountLTE applies the LTE predicate on the "votes_count" field.
func VotesCountLTE(v int) predicate.Candidate {
	return predicate.Candidate(sql.FieldLTE(FieldVotesCount, v))
}

// HasElection applies the HasEdge predicate on the "election" edge.
func HasElection() predicate.Candidate {
	return predicate.Candidate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ElectionTable, ElectionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasElectionWith applies the HasEdge predicate on the "election" edge with a given conditions (other predicates).
func HasElectionWith(preds ...predicate.Election) predicate.Candidate {
	return predicate.Candidate(func(s *sql.Selector) {
		step := newElectionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVotes applies the HasEdge predicate on the "votes" edge.
func HasVotes() predicate.Candidate {
	return predicate.Candidate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VotesTable, VotesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVotesWith applies the HasEdge predicate on the "votes" edge with a given conditions (other predicates).
func HasVotesWith(preds ...predicate.Vote) predicate.Candidate {
	return predicate.Candidate(func(s *sql.Selector) {
		step := newVotesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Candidate) predicate.Candidate {
	return predicate.Candidate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Candidate) predicate.Candidate {
	return predicate.Candidate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Candidate) predicate.Candidate {
	return predicate.Candidate(sql.NotPredicates(p))
}
