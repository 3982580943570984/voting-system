// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"voting-system/ent/generated"
)

// The CandidateFunc type is an adapter to allow the use of ordinary
// function as Candidate mutator.
type CandidateFunc func(context.Context, *generated.CandidateMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.CandidateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.CandidateMutation", m)
}

// The CommentFunc type is an adapter to allow the use of ordinary
// function as Comment mutator.
type CommentFunc func(context.Context, *generated.CommentMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f CommentFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.CommentMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.CommentMutation", m)
}

// The ElectionFunc type is an adapter to allow the use of ordinary
// function as Election mutator.
type ElectionFunc func(context.Context, *generated.ElectionMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ElectionFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ElectionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ElectionMutation", m)
}

// The ElectionFiltersFunc type is an adapter to allow the use of ordinary
// function as ElectionFilters mutator.
type ElectionFiltersFunc func(context.Context, *generated.ElectionFiltersMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ElectionFiltersFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ElectionFiltersMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ElectionFiltersMutation", m)
}

// The ElectionSettingsFunc type is an adapter to allow the use of ordinary
// function as ElectionSettings mutator.
type ElectionSettingsFunc func(context.Context, *generated.ElectionSettingsMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ElectionSettingsFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ElectionSettingsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ElectionSettingsMutation", m)
}

// The ProfileFunc type is an adapter to allow the use of ordinary
// function as Profile mutator.
type ProfileFunc func(context.Context, *generated.ProfileMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ProfileFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ProfileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ProfileMutation", m)
}

// The TagFunc type is an adapter to allow the use of ordinary
// function as Tag mutator.
type TagFunc func(context.Context, *generated.TagMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f TagFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.TagMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.TagMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *generated.UserMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserMutation", m)
}

// The VoteFunc type is an adapter to allow the use of ordinary
// function as Vote mutator.
type VoteFunc func(context.Context, *generated.VoteMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f VoteFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.VoteMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.VoteMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, generated.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op generated.Op) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk generated.Hook, cond Condition) generated.Hook {
	return func(next generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(ctx context.Context, m generated.Mutation) (generated.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, generated.Delete|generated.Create)
func On(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, generated.Update|generated.UpdateOne)
func Unless(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) generated.Hook {
	return func(generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(context.Context, generated.Mutation) (generated.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []generated.Hook {
//		return []generated.Hook{
//			Reject(generated.Delete|generated.Update),
//		}
//	}
func Reject(op generated.Op) generated.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []generated.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...generated.Hook) Chain {
	return Chain{append([]generated.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() generated.Hook {
	return func(mutator generated.Mutator) generated.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...generated.Hook) Chain {
	newHooks := make([]generated.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
