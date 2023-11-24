// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/naelcodes/ab-backend/internal/ent/predicate"
	"github.com/naelcodes/ab-backend/internal/ent/travelitem"
)

// TravelItemDelete is the builder for deleting a TravelItem entity.
type TravelItemDelete struct {
	config
	hooks    []Hook
	mutation *TravelItemMutation
}

// Where appends a list predicates to the TravelItemDelete builder.
func (tid *TravelItemDelete) Where(ps ...predicate.TravelItem) *TravelItemDelete {
	tid.mutation.Where(ps...)
	return tid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tid *TravelItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tid.sqlExec, tid.mutation, tid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tid *TravelItemDelete) ExecX(ctx context.Context) int {
	n, err := tid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tid *TravelItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(travelitem.Table, sqlgraph.NewFieldSpec(travelitem.FieldID, field.TypeInt))
	if ps := tid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tid.mutation.done = true
	return affected, err
}

// TravelItemDeleteOne is the builder for deleting a single TravelItem entity.
type TravelItemDeleteOne struct {
	tid *TravelItemDelete
}

// Where appends a list predicates to the TravelItemDelete builder.
func (tido *TravelItemDeleteOne) Where(ps ...predicate.TravelItem) *TravelItemDeleteOne {
	tido.tid.mutation.Where(ps...)
	return tido
}

// Exec executes the deletion query.
func (tido *TravelItemDeleteOne) Exec(ctx context.Context) error {
	n, err := tido.tid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{travelitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tido *TravelItemDeleteOne) ExecX(ctx context.Context) {
	if err := tido.Exec(ctx); err != nil {
		panic(err)
	}
}
