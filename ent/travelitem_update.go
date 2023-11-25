// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/naelcodes/ab-backend/internal/ent/invoice"
	"github.com/naelcodes/ab-backend/internal/ent/predicate"
	"github.com/naelcodes/ab-backend/internal/ent/travelitem"
)

// TravelItemUpdate is the builder for updating TravelItem entities.
type TravelItemUpdate struct {
	config
	hooks    []Hook
	mutation *TravelItemMutation
}

// Where appends a list predicates to the TravelItemUpdate builder.
func (tiu *TravelItemUpdate) Where(ps ...predicate.TravelItem) *TravelItemUpdate {
	tiu.mutation.Where(ps...)
	return tiu
}

// SetTotalPrice sets the "total_price" field.
func (tiu *TravelItemUpdate) SetTotalPrice(f float64) *TravelItemUpdate {
	tiu.mutation.ResetTotalPrice()
	tiu.mutation.SetTotalPrice(f)
	return tiu
}

// SetNillableTotalPrice sets the "total_price" field if the given value is not nil.
func (tiu *TravelItemUpdate) SetNillableTotalPrice(f *float64) *TravelItemUpdate {
	if f != nil {
		tiu.SetTotalPrice(*f)
	}
	return tiu
}

// AddTotalPrice adds f to the "total_price" field.
func (tiu *TravelItemUpdate) AddTotalPrice(f float64) *TravelItemUpdate {
	tiu.mutation.AddTotalPrice(f)
	return tiu
}

// SetItinerary sets the "itinerary" field.
func (tiu *TravelItemUpdate) SetItinerary(s string) *TravelItemUpdate {
	tiu.mutation.SetItinerary(s)
	return tiu
}

// SetNillableItinerary sets the "itinerary" field if the given value is not nil.
func (tiu *TravelItemUpdate) SetNillableItinerary(s *string) *TravelItemUpdate {
	if s != nil {
		tiu.SetItinerary(*s)
	}
	return tiu
}

// SetTravelerName sets the "traveler_name" field.
func (tiu *TravelItemUpdate) SetTravelerName(s string) *TravelItemUpdate {
	tiu.mutation.SetTravelerName(s)
	return tiu
}

// SetNillableTravelerName sets the "traveler_name" field if the given value is not nil.
func (tiu *TravelItemUpdate) SetNillableTravelerName(s *string) *TravelItemUpdate {
	if s != nil {
		tiu.SetTravelerName(*s)
	}
	return tiu
}

// SetTicketNumber sets the "ticket_number" field.
func (tiu *TravelItemUpdate) SetTicketNumber(s string) *TravelItemUpdate {
	tiu.mutation.SetTicketNumber(s)
	return tiu
}

// SetNillableTicketNumber sets the "ticket_number" field if the given value is not nil.
func (tiu *TravelItemUpdate) SetNillableTicketNumber(s *string) *TravelItemUpdate {
	if s != nil {
		tiu.SetTicketNumber(*s)
	}
	return tiu
}

// SetConjunctionNumber sets the "conjunction_number" field.
func (tiu *TravelItemUpdate) SetConjunctionNumber(i int) *TravelItemUpdate {
	tiu.mutation.ResetConjunctionNumber()
	tiu.mutation.SetConjunctionNumber(i)
	return tiu
}

// SetNillableConjunctionNumber sets the "conjunction_number" field if the given value is not nil.
func (tiu *TravelItemUpdate) SetNillableConjunctionNumber(i *int) *TravelItemUpdate {
	if i != nil {
		tiu.SetConjunctionNumber(*i)
	}
	return tiu
}

// AddConjunctionNumber adds i to the "conjunction_number" field.
func (tiu *TravelItemUpdate) AddConjunctionNumber(i int) *TravelItemUpdate {
	tiu.mutation.AddConjunctionNumber(i)
	return tiu
}

// SetStatus sets the "status" field.
func (tiu *TravelItemUpdate) SetStatus(t travelitem.Status) *TravelItemUpdate {
	tiu.mutation.SetStatus(t)
	return tiu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tiu *TravelItemUpdate) SetNillableStatus(t *travelitem.Status) *TravelItemUpdate {
	if t != nil {
		tiu.SetStatus(*t)
	}
	return tiu
}

// SetInvoiceID sets the "invoice" edge to the Invoice entity by ID.
func (tiu *TravelItemUpdate) SetInvoiceID(id int) *TravelItemUpdate {
	tiu.mutation.SetInvoiceID(id)
	return tiu
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (tiu *TravelItemUpdate) SetInvoice(i *Invoice) *TravelItemUpdate {
	return tiu.SetInvoiceID(i.ID)
}

// Mutation returns the TravelItemMutation object of the builder.
func (tiu *TravelItemUpdate) Mutation() *TravelItemMutation {
	return tiu.mutation
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (tiu *TravelItemUpdate) ClearInvoice() *TravelItemUpdate {
	tiu.mutation.ClearInvoice()
	return tiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tiu *TravelItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tiu.sqlSave, tiu.mutation, tiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tiu *TravelItemUpdate) SaveX(ctx context.Context) int {
	affected, err := tiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tiu *TravelItemUpdate) Exec(ctx context.Context) error {
	_, err := tiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tiu *TravelItemUpdate) ExecX(ctx context.Context) {
	if err := tiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tiu *TravelItemUpdate) check() error {
	if v, ok := tiu.mutation.TotalPrice(); ok {
		if err := travelitem.TotalPriceValidator(v); err != nil {
			return &ValidationError{Name: "total_price", err: fmt.Errorf(`ent: validator failed for field "TravelItem.total_price": %w`, err)}
		}
	}
	if v, ok := tiu.mutation.Status(); ok {
		if err := travelitem.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TravelItem.status": %w`, err)}
		}
	}
	if _, ok := tiu.mutation.InvoiceID(); tiu.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TravelItem.invoice"`)
	}
	return nil
}

func (tiu *TravelItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(travelitem.Table, travelitem.Columns, sqlgraph.NewFieldSpec(travelitem.FieldID, field.TypeInt))
	if ps := tiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tiu.mutation.TotalPrice(); ok {
		_spec.SetField(travelitem.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := tiu.mutation.AddedTotalPrice(); ok {
		_spec.AddField(travelitem.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := tiu.mutation.Itinerary(); ok {
		_spec.SetField(travelitem.FieldItinerary, field.TypeString, value)
	}
	if value, ok := tiu.mutation.TravelerName(); ok {
		_spec.SetField(travelitem.FieldTravelerName, field.TypeString, value)
	}
	if value, ok := tiu.mutation.TicketNumber(); ok {
		_spec.SetField(travelitem.FieldTicketNumber, field.TypeString, value)
	}
	if value, ok := tiu.mutation.ConjunctionNumber(); ok {
		_spec.SetField(travelitem.FieldConjunctionNumber, field.TypeInt, value)
	}
	if value, ok := tiu.mutation.AddedConjunctionNumber(); ok {
		_spec.AddField(travelitem.FieldConjunctionNumber, field.TypeInt, value)
	}
	if value, ok := tiu.mutation.Status(); ok {
		_spec.SetField(travelitem.FieldStatus, field.TypeEnum, value)
	}
	if tiu.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   travelitem.InvoiceTable,
			Columns: []string{travelitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tiu.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   travelitem.InvoiceTable,
			Columns: []string{travelitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{travelitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tiu.mutation.done = true
	return n, nil
}

// TravelItemUpdateOne is the builder for updating a single TravelItem entity.
type TravelItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TravelItemMutation
}

// SetTotalPrice sets the "total_price" field.
func (tiuo *TravelItemUpdateOne) SetTotalPrice(f float64) *TravelItemUpdateOne {
	tiuo.mutation.ResetTotalPrice()
	tiuo.mutation.SetTotalPrice(f)
	return tiuo
}

// SetNillableTotalPrice sets the "total_price" field if the given value is not nil.
func (tiuo *TravelItemUpdateOne) SetNillableTotalPrice(f *float64) *TravelItemUpdateOne {
	if f != nil {
		tiuo.SetTotalPrice(*f)
	}
	return tiuo
}

// AddTotalPrice adds f to the "total_price" field.
func (tiuo *TravelItemUpdateOne) AddTotalPrice(f float64) *TravelItemUpdateOne {
	tiuo.mutation.AddTotalPrice(f)
	return tiuo
}

// SetItinerary sets the "itinerary" field.
func (tiuo *TravelItemUpdateOne) SetItinerary(s string) *TravelItemUpdateOne {
	tiuo.mutation.SetItinerary(s)
	return tiuo
}

// SetNillableItinerary sets the "itinerary" field if the given value is not nil.
func (tiuo *TravelItemUpdateOne) SetNillableItinerary(s *string) *TravelItemUpdateOne {
	if s != nil {
		tiuo.SetItinerary(*s)
	}
	return tiuo
}

// SetTravelerName sets the "traveler_name" field.
func (tiuo *TravelItemUpdateOne) SetTravelerName(s string) *TravelItemUpdateOne {
	tiuo.mutation.SetTravelerName(s)
	return tiuo
}

// SetNillableTravelerName sets the "traveler_name" field if the given value is not nil.
func (tiuo *TravelItemUpdateOne) SetNillableTravelerName(s *string) *TravelItemUpdateOne {
	if s != nil {
		tiuo.SetTravelerName(*s)
	}
	return tiuo
}

// SetTicketNumber sets the "ticket_number" field.
func (tiuo *TravelItemUpdateOne) SetTicketNumber(s string) *TravelItemUpdateOne {
	tiuo.mutation.SetTicketNumber(s)
	return tiuo
}

// SetNillableTicketNumber sets the "ticket_number" field if the given value is not nil.
func (tiuo *TravelItemUpdateOne) SetNillableTicketNumber(s *string) *TravelItemUpdateOne {
	if s != nil {
		tiuo.SetTicketNumber(*s)
	}
	return tiuo
}

// SetConjunctionNumber sets the "conjunction_number" field.
func (tiuo *TravelItemUpdateOne) SetConjunctionNumber(i int) *TravelItemUpdateOne {
	tiuo.mutation.ResetConjunctionNumber()
	tiuo.mutation.SetConjunctionNumber(i)
	return tiuo
}

// SetNillableConjunctionNumber sets the "conjunction_number" field if the given value is not nil.
func (tiuo *TravelItemUpdateOne) SetNillableConjunctionNumber(i *int) *TravelItemUpdateOne {
	if i != nil {
		tiuo.SetConjunctionNumber(*i)
	}
	return tiuo
}

// AddConjunctionNumber adds i to the "conjunction_number" field.
func (tiuo *TravelItemUpdateOne) AddConjunctionNumber(i int) *TravelItemUpdateOne {
	tiuo.mutation.AddConjunctionNumber(i)
	return tiuo
}

// SetStatus sets the "status" field.
func (tiuo *TravelItemUpdateOne) SetStatus(t travelitem.Status) *TravelItemUpdateOne {
	tiuo.mutation.SetStatus(t)
	return tiuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tiuo *TravelItemUpdateOne) SetNillableStatus(t *travelitem.Status) *TravelItemUpdateOne {
	if t != nil {
		tiuo.SetStatus(*t)
	}
	return tiuo
}

// SetInvoiceID sets the "invoice" edge to the Invoice entity by ID.
func (tiuo *TravelItemUpdateOne) SetInvoiceID(id int) *TravelItemUpdateOne {
	tiuo.mutation.SetInvoiceID(id)
	return tiuo
}

// SetInvoice sets the "invoice" edge to the Invoice entity.
func (tiuo *TravelItemUpdateOne) SetInvoice(i *Invoice) *TravelItemUpdateOne {
	return tiuo.SetInvoiceID(i.ID)
}

// Mutation returns the TravelItemMutation object of the builder.
func (tiuo *TravelItemUpdateOne) Mutation() *TravelItemMutation {
	return tiuo.mutation
}

// ClearInvoice clears the "invoice" edge to the Invoice entity.
func (tiuo *TravelItemUpdateOne) ClearInvoice() *TravelItemUpdateOne {
	tiuo.mutation.ClearInvoice()
	return tiuo
}

// Where appends a list predicates to the TravelItemUpdate builder.
func (tiuo *TravelItemUpdateOne) Where(ps ...predicate.TravelItem) *TravelItemUpdateOne {
	tiuo.mutation.Where(ps...)
	return tiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tiuo *TravelItemUpdateOne) Select(field string, fields ...string) *TravelItemUpdateOne {
	tiuo.fields = append([]string{field}, fields...)
	return tiuo
}

// Save executes the query and returns the updated TravelItem entity.
func (tiuo *TravelItemUpdateOne) Save(ctx context.Context) (*TravelItem, error) {
	return withHooks(ctx, tiuo.sqlSave, tiuo.mutation, tiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tiuo *TravelItemUpdateOne) SaveX(ctx context.Context) *TravelItem {
	node, err := tiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tiuo *TravelItemUpdateOne) Exec(ctx context.Context) error {
	_, err := tiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tiuo *TravelItemUpdateOne) ExecX(ctx context.Context) {
	if err := tiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tiuo *TravelItemUpdateOne) check() error {
	if v, ok := tiuo.mutation.TotalPrice(); ok {
		if err := travelitem.TotalPriceValidator(v); err != nil {
			return &ValidationError{Name: "total_price", err: fmt.Errorf(`ent: validator failed for field "TravelItem.total_price": %w`, err)}
		}
	}
	if v, ok := tiuo.mutation.Status(); ok {
		if err := travelitem.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TravelItem.status": %w`, err)}
		}
	}
	if _, ok := tiuo.mutation.InvoiceID(); tiuo.mutation.InvoiceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TravelItem.invoice"`)
	}
	return nil
}

func (tiuo *TravelItemUpdateOne) sqlSave(ctx context.Context) (_node *TravelItem, err error) {
	if err := tiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(travelitem.Table, travelitem.Columns, sqlgraph.NewFieldSpec(travelitem.FieldID, field.TypeInt))
	id, ok := tiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TravelItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, travelitem.FieldID)
		for _, f := range fields {
			if !travelitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != travelitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tiuo.mutation.TotalPrice(); ok {
		_spec.SetField(travelitem.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := tiuo.mutation.AddedTotalPrice(); ok {
		_spec.AddField(travelitem.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := tiuo.mutation.Itinerary(); ok {
		_spec.SetField(travelitem.FieldItinerary, field.TypeString, value)
	}
	if value, ok := tiuo.mutation.TravelerName(); ok {
		_spec.SetField(travelitem.FieldTravelerName, field.TypeString, value)
	}
	if value, ok := tiuo.mutation.TicketNumber(); ok {
		_spec.SetField(travelitem.FieldTicketNumber, field.TypeString, value)
	}
	if value, ok := tiuo.mutation.ConjunctionNumber(); ok {
		_spec.SetField(travelitem.FieldConjunctionNumber, field.TypeInt, value)
	}
	if value, ok := tiuo.mutation.AddedConjunctionNumber(); ok {
		_spec.AddField(travelitem.FieldConjunctionNumber, field.TypeInt, value)
	}
	if value, ok := tiuo.mutation.Status(); ok {
		_spec.SetField(travelitem.FieldStatus, field.TypeEnum, value)
	}
	if tiuo.mutation.InvoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   travelitem.InvoiceTable,
			Columns: []string{travelitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tiuo.mutation.InvoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   travelitem.InvoiceTable,
			Columns: []string{travelitem.InvoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TravelItem{config: tiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{travelitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tiuo.mutation.done = true
	return _node, nil
}