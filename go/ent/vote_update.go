// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"voting-system/ent/predicate"
	"voting-system/ent/vote"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoteUpdate is the builder for updating Vote entities.
type VoteUpdate struct {
	config
	hooks    []Hook
	mutation *VoteMutation
}

// Where appends a list predicates to the VoteUpdate builder.
func (vu *VoteUpdate) Where(ps ...predicate.Vote) *VoteUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// Mutation returns the VoteMutation object of the builder.
func (vu *VoteUpdate) Mutation() *VoteMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VoteUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VoteUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VoteUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VoteUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vu *VoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(vote.Table, vote.Columns, sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VoteUpdateOne is the builder for updating a single Vote entity.
type VoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VoteMutation
}

// Mutation returns the VoteMutation object of the builder.
func (vuo *VoteUpdateOne) Mutation() *VoteMutation {
	return vuo.mutation
}

// Where appends a list predicates to the VoteUpdate builder.
func (vuo *VoteUpdateOne) Where(ps ...predicate.Vote) *VoteUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VoteUpdateOne) Select(field string, fields ...string) *VoteUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vote entity.
func (vuo *VoteUpdateOne) Save(ctx context.Context) (*Vote, error) {
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VoteUpdateOne) SaveX(ctx context.Context) *Vote {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VoteUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VoteUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vuo *VoteUpdateOne) sqlSave(ctx context.Context) (_node *Vote, err error) {
	_spec := sqlgraph.NewUpdateSpec(vote.Table, vote.Columns, sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Vote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vote.FieldID)
		for _, f := range fields {
			if !vote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Vote{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vote.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
