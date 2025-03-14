// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"voting-system/ent/generated/election"
	"voting-system/ent/generated/electionfilters"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ElectionFiltersCreate is the builder for creating a ElectionFilters entity.
type ElectionFiltersCreate struct {
	config
	mutation *ElectionFiltersMutation
	hooks    []Hook
}

// SetHasFirstName sets the "has_first_name" field.
func (efc *ElectionFiltersCreate) SetHasFirstName(b bool) *ElectionFiltersCreate {
	efc.mutation.SetHasFirstName(b)
	return efc
}

// SetNillableHasFirstName sets the "has_first_name" field if the given value is not nil.
func (efc *ElectionFiltersCreate) SetNillableHasFirstName(b *bool) *ElectionFiltersCreate {
	if b != nil {
		efc.SetHasFirstName(*b)
	}
	return efc
}

// SetHasLastName sets the "has_last_name" field.
func (efc *ElectionFiltersCreate) SetHasLastName(b bool) *ElectionFiltersCreate {
	efc.mutation.SetHasLastName(b)
	return efc
}

// SetNillableHasLastName sets the "has_last_name" field if the given value is not nil.
func (efc *ElectionFiltersCreate) SetNillableHasLastName(b *bool) *ElectionFiltersCreate {
	if b != nil {
		efc.SetHasLastName(*b)
	}
	return efc
}

// SetHasBirthdate sets the "has_birthdate" field.
func (efc *ElectionFiltersCreate) SetHasBirthdate(b bool) *ElectionFiltersCreate {
	efc.mutation.SetHasBirthdate(b)
	return efc
}

// SetNillableHasBirthdate sets the "has_birthdate" field if the given value is not nil.
func (efc *ElectionFiltersCreate) SetNillableHasBirthdate(b *bool) *ElectionFiltersCreate {
	if b != nil {
		efc.SetHasBirthdate(*b)
	}
	return efc
}

// SetHasPhoneNumber sets the "has_phone_number" field.
func (efc *ElectionFiltersCreate) SetHasPhoneNumber(b bool) *ElectionFiltersCreate {
	efc.mutation.SetHasPhoneNumber(b)
	return efc
}

// SetNillableHasPhoneNumber sets the "has_phone_number" field if the given value is not nil.
func (efc *ElectionFiltersCreate) SetNillableHasPhoneNumber(b *bool) *ElectionFiltersCreate {
	if b != nil {
		efc.SetHasPhoneNumber(*b)
	}
	return efc
}

// SetHasBio sets the "has_bio" field.
func (efc *ElectionFiltersCreate) SetHasBio(b bool) *ElectionFiltersCreate {
	efc.mutation.SetHasBio(b)
	return efc
}

// SetNillableHasBio sets the "has_bio" field if the given value is not nil.
func (efc *ElectionFiltersCreate) SetNillableHasBio(b *bool) *ElectionFiltersCreate {
	if b != nil {
		efc.SetHasBio(*b)
	}
	return efc
}

// SetHasAddress sets the "has_address" field.
func (efc *ElectionFiltersCreate) SetHasAddress(b bool) *ElectionFiltersCreate {
	efc.mutation.SetHasAddress(b)
	return efc
}

// SetNillableHasAddress sets the "has_address" field if the given value is not nil.
func (efc *ElectionFiltersCreate) SetNillableHasAddress(b *bool) *ElectionFiltersCreate {
	if b != nil {
		efc.SetHasAddress(*b)
	}
	return efc
}

// SetHasPhotoURL sets the "has_photo_url" field.
func (efc *ElectionFiltersCreate) SetHasPhotoURL(b bool) *ElectionFiltersCreate {
	efc.mutation.SetHasPhotoURL(b)
	return efc
}

// SetNillableHasPhotoURL sets the "has_photo_url" field if the given value is not nil.
func (efc *ElectionFiltersCreate) SetNillableHasPhotoURL(b *bool) *ElectionFiltersCreate {
	if b != nil {
		efc.SetHasPhotoURL(*b)
	}
	return efc
}

// SetElectionID sets the "election" edge to the Election entity by ID.
func (efc *ElectionFiltersCreate) SetElectionID(id int) *ElectionFiltersCreate {
	efc.mutation.SetElectionID(id)
	return efc
}

// SetElection sets the "election" edge to the Election entity.
func (efc *ElectionFiltersCreate) SetElection(e *Election) *ElectionFiltersCreate {
	return efc.SetElectionID(e.ID)
}

// Mutation returns the ElectionFiltersMutation object of the builder.
func (efc *ElectionFiltersCreate) Mutation() *ElectionFiltersMutation {
	return efc.mutation
}

// Save creates the ElectionFilters in the database.
func (efc *ElectionFiltersCreate) Save(ctx context.Context) (*ElectionFilters, error) {
	efc.defaults()
	return withHooks(ctx, efc.sqlSave, efc.mutation, efc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (efc *ElectionFiltersCreate) SaveX(ctx context.Context) *ElectionFilters {
	v, err := efc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (efc *ElectionFiltersCreate) Exec(ctx context.Context) error {
	_, err := efc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efc *ElectionFiltersCreate) ExecX(ctx context.Context) {
	if err := efc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (efc *ElectionFiltersCreate) defaults() {
	if _, ok := efc.mutation.HasFirstName(); !ok {
		v := electionfilters.DefaultHasFirstName
		efc.mutation.SetHasFirstName(v)
	}
	if _, ok := efc.mutation.HasLastName(); !ok {
		v := electionfilters.DefaultHasLastName
		efc.mutation.SetHasLastName(v)
	}
	if _, ok := efc.mutation.HasBirthdate(); !ok {
		v := electionfilters.DefaultHasBirthdate
		efc.mutation.SetHasBirthdate(v)
	}
	if _, ok := efc.mutation.HasPhoneNumber(); !ok {
		v := electionfilters.DefaultHasPhoneNumber
		efc.mutation.SetHasPhoneNumber(v)
	}
	if _, ok := efc.mutation.HasBio(); !ok {
		v := electionfilters.DefaultHasBio
		efc.mutation.SetHasBio(v)
	}
	if _, ok := efc.mutation.HasAddress(); !ok {
		v := electionfilters.DefaultHasAddress
		efc.mutation.SetHasAddress(v)
	}
	if _, ok := efc.mutation.HasPhotoURL(); !ok {
		v := electionfilters.DefaultHasPhotoURL
		efc.mutation.SetHasPhotoURL(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (efc *ElectionFiltersCreate) check() error {
	if _, ok := efc.mutation.HasFirstName(); !ok {
		return &ValidationError{Name: "has_first_name", err: errors.New(`generated: missing required field "ElectionFilters.has_first_name"`)}
	}
	if _, ok := efc.mutation.HasLastName(); !ok {
		return &ValidationError{Name: "has_last_name", err: errors.New(`generated: missing required field "ElectionFilters.has_last_name"`)}
	}
	if _, ok := efc.mutation.HasBirthdate(); !ok {
		return &ValidationError{Name: "has_birthdate", err: errors.New(`generated: missing required field "ElectionFilters.has_birthdate"`)}
	}
	if _, ok := efc.mutation.HasPhoneNumber(); !ok {
		return &ValidationError{Name: "has_phone_number", err: errors.New(`generated: missing required field "ElectionFilters.has_phone_number"`)}
	}
	if _, ok := efc.mutation.HasBio(); !ok {
		return &ValidationError{Name: "has_bio", err: errors.New(`generated: missing required field "ElectionFilters.has_bio"`)}
	}
	if _, ok := efc.mutation.HasAddress(); !ok {
		return &ValidationError{Name: "has_address", err: errors.New(`generated: missing required field "ElectionFilters.has_address"`)}
	}
	if _, ok := efc.mutation.HasPhotoURL(); !ok {
		return &ValidationError{Name: "has_photo_url", err: errors.New(`generated: missing required field "ElectionFilters.has_photo_url"`)}
	}
	if len(efc.mutation.ElectionIDs()) == 0 {
		return &ValidationError{Name: "election", err: errors.New(`generated: missing required edge "ElectionFilters.election"`)}
	}
	return nil
}

func (efc *ElectionFiltersCreate) sqlSave(ctx context.Context) (*ElectionFilters, error) {
	if err := efc.check(); err != nil {
		return nil, err
	}
	_node, _spec := efc.createSpec()
	if err := sqlgraph.CreateNode(ctx, efc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	efc.mutation.id = &_node.ID
	efc.mutation.done = true
	return _node, nil
}

func (efc *ElectionFiltersCreate) createSpec() (*ElectionFilters, *sqlgraph.CreateSpec) {
	var (
		_node = &ElectionFilters{config: efc.config}
		_spec = sqlgraph.NewCreateSpec(electionfilters.Table, sqlgraph.NewFieldSpec(electionfilters.FieldID, field.TypeInt))
	)
	if value, ok := efc.mutation.HasFirstName(); ok {
		_spec.SetField(electionfilters.FieldHasFirstName, field.TypeBool, value)
		_node.HasFirstName = value
	}
	if value, ok := efc.mutation.HasLastName(); ok {
		_spec.SetField(electionfilters.FieldHasLastName, field.TypeBool, value)
		_node.HasLastName = value
	}
	if value, ok := efc.mutation.HasBirthdate(); ok {
		_spec.SetField(electionfilters.FieldHasBirthdate, field.TypeBool, value)
		_node.HasBirthdate = value
	}
	if value, ok := efc.mutation.HasPhoneNumber(); ok {
		_spec.SetField(electionfilters.FieldHasPhoneNumber, field.TypeBool, value)
		_node.HasPhoneNumber = value
	}
	if value, ok := efc.mutation.HasBio(); ok {
		_spec.SetField(electionfilters.FieldHasBio, field.TypeBool, value)
		_node.HasBio = value
	}
	if value, ok := efc.mutation.HasAddress(); ok {
		_spec.SetField(electionfilters.FieldHasAddress, field.TypeBool, value)
		_node.HasAddress = value
	}
	if value, ok := efc.mutation.HasPhotoURL(); ok {
		_spec.SetField(electionfilters.FieldHasPhotoURL, field.TypeBool, value)
		_node.HasPhotoURL = value
	}
	if nodes := efc.mutation.ElectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   electionfilters.ElectionTable,
			Columns: []string{electionfilters.ElectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(election.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.election_filters = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ElectionFiltersCreateBulk is the builder for creating many ElectionFilters entities in bulk.
type ElectionFiltersCreateBulk struct {
	config
	err      error
	builders []*ElectionFiltersCreate
}

// Save creates the ElectionFilters entities in the database.
func (efcb *ElectionFiltersCreateBulk) Save(ctx context.Context) ([]*ElectionFilters, error) {
	if efcb.err != nil {
		return nil, efcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(efcb.builders))
	nodes := make([]*ElectionFilters, len(efcb.builders))
	mutators := make([]Mutator, len(efcb.builders))
	for i := range efcb.builders {
		func(i int, root context.Context) {
			builder := efcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ElectionFiltersMutation)
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
					_, err = mutators[i+1].Mutate(root, efcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, efcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, efcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (efcb *ElectionFiltersCreateBulk) SaveX(ctx context.Context) []*ElectionFilters {
	v, err := efcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (efcb *ElectionFiltersCreateBulk) Exec(ctx context.Context) error {
	_, err := efcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efcb *ElectionFiltersCreateBulk) ExecX(ctx context.Context) {
	if err := efcb.Exec(ctx); err != nil {
		panic(err)
	}
}
