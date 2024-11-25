// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"
	"voting-system/ent/generated/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// Contents applies equality check predicate on the "contents" field. It's identical to ContentsEQ.
func Contents(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContents, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldTimestamp, v))
}

// ContentsEQ applies the EQ predicate on the "contents" field.
func ContentsEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContents, v))
}

// ContentsNEQ applies the NEQ predicate on the "contents" field.
func ContentsNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldContents, v))
}

// ContentsIn applies the In predicate on the "contents" field.
func ContentsIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldContents, vs...))
}

// ContentsNotIn applies the NotIn predicate on the "contents" field.
func ContentsNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldContents, vs...))
}

// ContentsGT applies the GT predicate on the "contents" field.
func ContentsGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldContents, v))
}

// ContentsGTE applies the GTE predicate on the "contents" field.
func ContentsGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldContents, v))
}

// ContentsLT applies the LT predicate on the "contents" field.
func ContentsLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldContents, v))
}

// ContentsLTE applies the LTE predicate on the "contents" field.
func ContentsLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldContents, v))
}

// ContentsContains applies the Contains predicate on the "contents" field.
func ContentsContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldContents, v))
}

// ContentsHasPrefix applies the HasPrefix predicate on the "contents" field.
func ContentsHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldContents, v))
}

// ContentsHasSuffix applies the HasSuffix predicate on the "contents" field.
func ContentsHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldContents, v))
}

// ContentsEqualFold applies the EqualFold predicate on the "contents" field.
func ContentsEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldContents, v))
}

// ContentsContainsFold applies the ContainsFold predicate on the "contents" field.
func ContentsContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldContents, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldTimestamp, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasElection applies the HasEdge predicate on the "election" edge.
func HasElection() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ElectionTable, ElectionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasElectionWith applies the HasEdge predicate on the "election" edge with a given conditions (other predicates).
func HasElectionWith(preds ...predicate.Election) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newElectionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.NotPredicates(p))
}
