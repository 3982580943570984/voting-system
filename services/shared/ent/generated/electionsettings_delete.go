// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"shared/ent/generated/electionsettings"
	"shared/ent/generated/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ElectionSettingsDelete is the builder for deleting a ElectionSettings entity.
type ElectionSettingsDelete struct {
	config
	hooks    []Hook
	mutation *ElectionSettingsMutation
}

// Where appends a list predicates to the ElectionSettingsDelete builder.
func (esd *ElectionSettingsDelete) Where(ps ...predicate.ElectionSettings) *ElectionSettingsDelete {
	esd.mutation.Where(ps...)
	return esd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (esd *ElectionSettingsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, esd.sqlExec, esd.mutation, esd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (esd *ElectionSettingsDelete) ExecX(ctx context.Context) int {
	n, err := esd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (esd *ElectionSettingsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(electionsettings.Table, sqlgraph.NewFieldSpec(electionsettings.FieldID, field.TypeInt))
	if ps := esd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, esd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	esd.mutation.done = true
	return affected, err
}

// ElectionSettingsDeleteOne is the builder for deleting a single ElectionSettings entity.
type ElectionSettingsDeleteOne struct {
	esd *ElectionSettingsDelete
}

// Where appends a list predicates to the ElectionSettingsDelete builder.
func (esdo *ElectionSettingsDeleteOne) Where(ps ...predicate.ElectionSettings) *ElectionSettingsDeleteOne {
	esdo.esd.mutation.Where(ps...)
	return esdo
}

// Exec executes the deletion query.
func (esdo *ElectionSettingsDeleteOne) Exec(ctx context.Context) error {
	n, err := esdo.esd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{electionsettings.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (esdo *ElectionSettingsDeleteOne) ExecX(ctx context.Context) {
	if err := esdo.Exec(ctx); err != nil {
		panic(err)
	}
}
