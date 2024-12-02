// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"voting-system/ent/generated/candidate"
	"voting-system/ent/generated/predicate"
	"voting-system/ent/generated/user"
	"voting-system/ent/generated/vote"

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

// SetIsActive sets the "is_active" field.
func (vu *VoteUpdate) SetIsActive(b bool) *VoteUpdate {
	vu.mutation.SetIsActive(b)
	return vu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (vu *VoteUpdate) SetNillableIsActive(b *bool) *VoteUpdate {
	if b != nil {
		vu.SetIsActive(*b)
	}
	return vu
}

// SetCandidateID sets the "candidate" edge to the Candidate entity by ID.
func (vu *VoteUpdate) SetCandidateID(id int) *VoteUpdate {
	vu.mutation.SetCandidateID(id)
	return vu
}

// SetNillableCandidateID sets the "candidate" edge to the Candidate entity by ID if the given value is not nil.
func (vu *VoteUpdate) SetNillableCandidateID(id *int) *VoteUpdate {
	if id != nil {
		vu = vu.SetCandidateID(*id)
	}
	return vu
}

// SetCandidate sets the "candidate" edge to the Candidate entity.
func (vu *VoteUpdate) SetCandidate(c *Candidate) *VoteUpdate {
	return vu.SetCandidateID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vu *VoteUpdate) SetUserID(id int) *VoteUpdate {
	vu.mutation.SetUserID(id)
	return vu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (vu *VoteUpdate) SetNillableUserID(id *int) *VoteUpdate {
	if id != nil {
		vu = vu.SetUserID(*id)
	}
	return vu
}

// SetUser sets the "user" edge to the User entity.
func (vu *VoteUpdate) SetUser(u *User) *VoteUpdate {
	return vu.SetUserID(u.ID)
}

// Mutation returns the VoteMutation object of the builder.
func (vu *VoteUpdate) Mutation() *VoteMutation {
	return vu.mutation
}

// ClearCandidate clears the "candidate" edge to the Candidate entity.
func (vu *VoteUpdate) ClearCandidate() *VoteUpdate {
	vu.mutation.ClearCandidate()
	return vu
}

// ClearUser clears the "user" edge to the User entity.
func (vu *VoteUpdate) ClearUser() *VoteUpdate {
	vu.mutation.ClearUser()
	return vu
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
	if value, ok := vu.mutation.IsActive(); ok {
		_spec.SetField(vote.FieldIsActive, field.TypeBool, value)
	}
	if vu.mutation.CandidateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.CandidateTable,
			Columns: []string{vote.CandidateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(candidate.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.CandidateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.CandidateTable,
			Columns: []string{vote.CandidateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(candidate.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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

// SetIsActive sets the "is_active" field.
func (vuo *VoteUpdateOne) SetIsActive(b bool) *VoteUpdateOne {
	vuo.mutation.SetIsActive(b)
	return vuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (vuo *VoteUpdateOne) SetNillableIsActive(b *bool) *VoteUpdateOne {
	if b != nil {
		vuo.SetIsActive(*b)
	}
	return vuo
}

// SetCandidateID sets the "candidate" edge to the Candidate entity by ID.
func (vuo *VoteUpdateOne) SetCandidateID(id int) *VoteUpdateOne {
	vuo.mutation.SetCandidateID(id)
	return vuo
}

// SetNillableCandidateID sets the "candidate" edge to the Candidate entity by ID if the given value is not nil.
func (vuo *VoteUpdateOne) SetNillableCandidateID(id *int) *VoteUpdateOne {
	if id != nil {
		vuo = vuo.SetCandidateID(*id)
	}
	return vuo
}

// SetCandidate sets the "candidate" edge to the Candidate entity.
func (vuo *VoteUpdateOne) SetCandidate(c *Candidate) *VoteUpdateOne {
	return vuo.SetCandidateID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vuo *VoteUpdateOne) SetUserID(id int) *VoteUpdateOne {
	vuo.mutation.SetUserID(id)
	return vuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (vuo *VoteUpdateOne) SetNillableUserID(id *int) *VoteUpdateOne {
	if id != nil {
		vuo = vuo.SetUserID(*id)
	}
	return vuo
}

// SetUser sets the "user" edge to the User entity.
func (vuo *VoteUpdateOne) SetUser(u *User) *VoteUpdateOne {
	return vuo.SetUserID(u.ID)
}

// Mutation returns the VoteMutation object of the builder.
func (vuo *VoteUpdateOne) Mutation() *VoteMutation {
	return vuo.mutation
}

// ClearCandidate clears the "candidate" edge to the Candidate entity.
func (vuo *VoteUpdateOne) ClearCandidate() *VoteUpdateOne {
	vuo.mutation.ClearCandidate()
	return vuo
}

// ClearUser clears the "user" edge to the User entity.
func (vuo *VoteUpdateOne) ClearUser() *VoteUpdateOne {
	vuo.mutation.ClearUser()
	return vuo
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
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Vote.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vote.FieldID)
		for _, f := range fields {
			if !vote.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
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
	if value, ok := vuo.mutation.IsActive(); ok {
		_spec.SetField(vote.FieldIsActive, field.TypeBool, value)
	}
	if vuo.mutation.CandidateCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.CandidateTable,
			Columns: []string{vote.CandidateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(candidate.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.CandidateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.CandidateTable,
			Columns: []string{vote.CandidateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(candidate.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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
