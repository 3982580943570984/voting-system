// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"
	"voting-system/ent/generated/comment"
	"voting-system/ent/generated/election"
	"voting-system/ent/generated/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cc *CommentCreate) SetCreateTime(t time.Time) *CommentCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreateTime(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CommentCreate) SetUpdateTime(t time.Time) *CommentCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUpdateTime(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetContents sets the "contents" field.
func (cc *CommentCreate) SetContents(s string) *CommentCreate {
	cc.mutation.SetContents(s)
	return cc
}

// SetParentID sets the "parent" edge to the Comment entity by ID.
func (cc *CommentCreate) SetParentID(id int) *CommentCreate {
	cc.mutation.SetParentID(id)
	return cc
}

// SetNillableParentID sets the "parent" edge to the Comment entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableParentID(id *int) *CommentCreate {
	if id != nil {
		cc = cc.SetParentID(*id)
	}
	return cc
}

// SetParent sets the "parent" edge to the Comment entity.
func (cc *CommentCreate) SetParent(c *Comment) *CommentCreate {
	return cc.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Comment entity by IDs.
func (cc *CommentCreate) AddChildIDs(ids ...int) *CommentCreate {
	cc.mutation.AddChildIDs(ids...)
	return cc
}

// AddChildren adds the "children" edges to the Comment entity.
func (cc *CommentCreate) AddChildren(c ...*Comment) *CommentCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddChildIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cc *CommentCreate) SetUserID(id int) *CommentCreate {
	cc.mutation.SetUserID(id)
	return cc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableUserID(id *int) *CommentCreate {
	if id != nil {
		cc = cc.SetUserID(*id)
	}
	return cc
}

// SetUser sets the "user" edge to the User entity.
func (cc *CommentCreate) SetUser(u *User) *CommentCreate {
	return cc.SetUserID(u.ID)
}

// SetElectionID sets the "election" edge to the Election entity by ID.
func (cc *CommentCreate) SetElectionID(id int) *CommentCreate {
	cc.mutation.SetElectionID(id)
	return cc
}

// SetNillableElectionID sets the "election" edge to the Election entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableElectionID(id *int) *CommentCreate {
	if id != nil {
		cc = cc.SetElectionID(*id)
	}
	return cc
}

// SetElection sets the "election" edge to the Election entity.
func (cc *CommentCreate) SetElection(e *Election) *CommentCreate {
	return cc.SetElectionID(e.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommentCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := comment.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := comment.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`generated: missing required field "Comment.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`generated: missing required field "Comment.update_time"`)}
	}
	if _, ok := cc.mutation.Contents(); !ok {
		return &ValidationError{Name: "contents", err: errors.New(`generated: missing required field "Comment.contents"`)}
	}
	if v, ok := cc.mutation.Contents(); ok {
		if err := comment.ContentsValidator(v); err != nil {
			return &ValidationError{Name: "contents", err: fmt.Errorf(`generated: validator failed for field "Comment.contents": %w`, err)}
		}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(comment.Table, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.SetField(comment.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.SetField(comment.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Contents(); ok {
		_spec.SetField(comment.FieldContents, field.TypeString, value)
		_node.Contents = value
	}
	if nodes := cc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ParentTable,
			Columns: []string{comment.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.comment_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ElectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ElectionTable,
			Columns: []string{comment.ElectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(election.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.election_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
