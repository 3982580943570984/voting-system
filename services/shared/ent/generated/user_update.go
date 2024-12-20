// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"shared/ent/generated/comment"
	"shared/ent/generated/election"
	"shared/ent/generated/predicate"
	"shared/ent/generated/profile"
	"shared/ent/generated/user"
	"shared/ent/generated/vote"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *UserUpdate) SetUpdateTime(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetLastLogin sets the "last_login" field.
func (uu *UserUpdate) SetLastLogin(t time.Time) *UserUpdate {
	uu.mutation.SetLastLogin(t)
	return uu
}

// SetIsActive sets the "is_active" field.
func (uu *UserUpdate) SetIsActive(b bool) *UserUpdate {
	uu.mutation.SetIsActive(b)
	return uu
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsActive(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsActive(*b)
	}
	return uu
}

// SetIsOrganizer sets the "is_organizer" field.
func (uu *UserUpdate) SetIsOrganizer(b bool) *UserUpdate {
	uu.mutation.SetIsOrganizer(b)
	return uu
}

// SetNillableIsOrganizer sets the "is_organizer" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsOrganizer(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsOrganizer(*b)
	}
	return uu
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (uu *UserUpdate) SetProfileID(id int) *UserUpdate {
	uu.mutation.SetProfileID(id)
	return uu
}

// SetNillableProfileID sets the "profile" edge to the Profile entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableProfileID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetProfileID(*id)
	}
	return uu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (uu *UserUpdate) SetProfile(p *Profile) *UserUpdate {
	return uu.SetProfileID(p.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uu *UserUpdate) AddCommentIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCommentIDs(ids...)
	return uu
}

// AddComments adds the "comments" edges to the Comment entity.
func (uu *UserUpdate) AddComments(c ...*Comment) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCommentIDs(ids...)
}

// AddElectionIDs adds the "elections" edge to the Election entity by IDs.
func (uu *UserUpdate) AddElectionIDs(ids ...int) *UserUpdate {
	uu.mutation.AddElectionIDs(ids...)
	return uu
}

// AddElections adds the "elections" edges to the Election entity.
func (uu *UserUpdate) AddElections(e ...*Election) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.AddElectionIDs(ids...)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (uu *UserUpdate) AddVoteIDs(ids ...int) *UserUpdate {
	uu.mutation.AddVoteIDs(ids...)
	return uu
}

// AddVotes adds the "votes" edges to the Vote entity.
func (uu *UserUpdate) AddVotes(v ...*Vote) *UserUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uu.AddVoteIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (uu *UserUpdate) ClearProfile() *UserUpdate {
	uu.mutation.ClearProfile()
	return uu
}

// ClearComments clears all "comments" edges to the Comment entity.
func (uu *UserUpdate) ClearComments() *UserUpdate {
	uu.mutation.ClearComments()
	return uu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (uu *UserUpdate) RemoveCommentIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCommentIDs(ids...)
	return uu
}

// RemoveComments removes "comments" edges to Comment entities.
func (uu *UserUpdate) RemoveComments(c ...*Comment) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCommentIDs(ids...)
}

// ClearElections clears all "elections" edges to the Election entity.
func (uu *UserUpdate) ClearElections() *UserUpdate {
	uu.mutation.ClearElections()
	return uu
}

// RemoveElectionIDs removes the "elections" edge to Election entities by IDs.
func (uu *UserUpdate) RemoveElectionIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveElectionIDs(ids...)
	return uu
}

// RemoveElections removes "elections" edges to Election entities.
func (uu *UserUpdate) RemoveElections(e ...*Election) *UserUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.RemoveElectionIDs(ids...)
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (uu *UserUpdate) ClearVotes() *UserUpdate {
	uu.mutation.ClearVotes()
	return uu
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (uu *UserUpdate) RemoveVoteIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveVoteIDs(ids...)
	return uu
}

// RemoveVotes removes "votes" edges to Vote entities.
func (uu *UserUpdate) RemoveVotes(v ...*Vote) *UserUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uu.RemoveVoteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		if user.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized user.UpdateDefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := user.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
	if _, ok := uu.mutation.LastLogin(); !ok {
		if user.UpdateDefaultLastLogin == nil {
			return fmt.Errorf("generated: uninitialized user.UpdateDefaultLastLogin (forgotten import generated/runtime?)")
		}
		v := user.UpdateDefaultLastLogin()
		uu.mutation.SetLastLogin(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`generated: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if value, ok := uu.mutation.IsActive(); ok {
		_spec.SetField(user.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := uu.mutation.IsOrganizer(); ok {
		_spec.SetField(user.FieldIsOrganizer, field.TypeBool, value)
	}
	if uu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ProfileTable,
			Columns: []string{user.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ProfileTable,
			Columns: []string{user.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !uu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ElectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ElectionsTable,
			Columns: []string{user.ElectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(election.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedElectionsIDs(); len(nodes) > 0 && !uu.mutation.ElectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ElectionsTable,
			Columns: []string{user.ElectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(election.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ElectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ElectionsTable,
			Columns: []string{user.ElectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(election.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedVotesIDs(); len(nodes) > 0 && !uu.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *UserUpdateOne) SetUpdateTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetLastLogin sets the "last_login" field.
func (uuo *UserUpdateOne) SetLastLogin(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastLogin(t)
	return uuo
}

// SetIsActive sets the "is_active" field.
func (uuo *UserUpdateOne) SetIsActive(b bool) *UserUpdateOne {
	uuo.mutation.SetIsActive(b)
	return uuo
}

// SetNillableIsActive sets the "is_active" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsActive(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsActive(*b)
	}
	return uuo
}

// SetIsOrganizer sets the "is_organizer" field.
func (uuo *UserUpdateOne) SetIsOrganizer(b bool) *UserUpdateOne {
	uuo.mutation.SetIsOrganizer(b)
	return uuo
}

// SetNillableIsOrganizer sets the "is_organizer" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsOrganizer(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsOrganizer(*b)
	}
	return uuo
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (uuo *UserUpdateOne) SetProfileID(id int) *UserUpdateOne {
	uuo.mutation.SetProfileID(id)
	return uuo
}

// SetNillableProfileID sets the "profile" edge to the Profile entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProfileID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetProfileID(*id)
	}
	return uuo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (uuo *UserUpdateOne) SetProfile(p *Profile) *UserUpdateOne {
	return uuo.SetProfileID(p.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (uuo *UserUpdateOne) AddCommentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCommentIDs(ids...)
	return uuo
}

// AddComments adds the "comments" edges to the Comment entity.
func (uuo *UserUpdateOne) AddComments(c ...*Comment) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCommentIDs(ids...)
}

// AddElectionIDs adds the "elections" edge to the Election entity by IDs.
func (uuo *UserUpdateOne) AddElectionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddElectionIDs(ids...)
	return uuo
}

// AddElections adds the "elections" edges to the Election entity.
func (uuo *UserUpdateOne) AddElections(e ...*Election) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.AddElectionIDs(ids...)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (uuo *UserUpdateOne) AddVoteIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddVoteIDs(ids...)
	return uuo
}

// AddVotes adds the "votes" edges to the Vote entity.
func (uuo *UserUpdateOne) AddVotes(v ...*Vote) *UserUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uuo.AddVoteIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (uuo *UserUpdateOne) ClearProfile() *UserUpdateOne {
	uuo.mutation.ClearProfile()
	return uuo
}

// ClearComments clears all "comments" edges to the Comment entity.
func (uuo *UserUpdateOne) ClearComments() *UserUpdateOne {
	uuo.mutation.ClearComments()
	return uuo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (uuo *UserUpdateOne) RemoveCommentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCommentIDs(ids...)
	return uuo
}

// RemoveComments removes "comments" edges to Comment entities.
func (uuo *UserUpdateOne) RemoveComments(c ...*Comment) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCommentIDs(ids...)
}

// ClearElections clears all "elections" edges to the Election entity.
func (uuo *UserUpdateOne) ClearElections() *UserUpdateOne {
	uuo.mutation.ClearElections()
	return uuo
}

// RemoveElectionIDs removes the "elections" edge to Election entities by IDs.
func (uuo *UserUpdateOne) RemoveElectionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveElectionIDs(ids...)
	return uuo
}

// RemoveElections removes "elections" edges to Election entities.
func (uuo *UserUpdateOne) RemoveElections(e ...*Election) *UserUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.RemoveElectionIDs(ids...)
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (uuo *UserUpdateOne) ClearVotes() *UserUpdateOne {
	uuo.mutation.ClearVotes()
	return uuo
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (uuo *UserUpdateOne) RemoveVoteIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveVoteIDs(ids...)
	return uuo
}

// RemoveVotes removes "votes" edges to Vote entities.
func (uuo *UserUpdateOne) RemoveVotes(v ...*Vote) *UserUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uuo.RemoveVoteIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		if user.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("generated: uninitialized user.UpdateDefaultUpdateTime (forgotten import generated/runtime?)")
		}
		v := user.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
	if _, ok := uuo.mutation.LastLogin(); !ok {
		if user.UpdateDefaultLastLogin == nil {
			return fmt.Errorf("generated: uninitialized user.UpdateDefaultLastLogin (forgotten import generated/runtime?)")
		}
		v := user.UpdateDefaultLastLogin()
		uuo.mutation.SetLastLogin(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`generated: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.IsActive(); ok {
		_spec.SetField(user.FieldIsActive, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.IsOrganizer(); ok {
		_spec.SetField(user.FieldIsOrganizer, field.TypeBool, value)
	}
	if uuo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ProfileTable,
			Columns: []string{user.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.ProfileTable,
			Columns: []string{user.ProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !uuo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CommentsTable,
			Columns: []string{user.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ElectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ElectionsTable,
			Columns: []string{user.ElectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(election.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedElectionsIDs(); len(nodes) > 0 && !uuo.mutation.ElectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ElectionsTable,
			Columns: []string{user.ElectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(election.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ElectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ElectionsTable,
			Columns: []string{user.ElectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(election.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedVotesIDs(); len(nodes) > 0 && !uuo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
