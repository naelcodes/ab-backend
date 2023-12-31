// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/naelcodes/ab-backend/ent/customer"
	"github.com/naelcodes/ab-backend/ent/imputation"
	"github.com/naelcodes/ab-backend/ent/payment"
	"github.com/naelcodes/ab-backend/ent/predicate"
)

// PaymentUpdate is the builder for updating Payment entities.
type PaymentUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentMutation
}

// Where appends a list predicates to the PaymentUpdate builder.
func (pu *PaymentUpdate) Where(ps ...predicate.Payment) *PaymentUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetNumber sets the "number" field.
func (pu *PaymentUpdate) SetNumber(s string) *PaymentUpdate {
	pu.mutation.SetNumber(s)
	return pu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableNumber(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetNumber(*s)
	}
	return pu
}

// SetDate sets the "date" field.
func (pu *PaymentUpdate) SetDate(s string) *PaymentUpdate {
	pu.mutation.SetDate(s)
	return pu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableDate(s *string) *PaymentUpdate {
	if s != nil {
		pu.SetDate(*s)
	}
	return pu
}

// SetBalance sets the "balance" field.
func (pu *PaymentUpdate) SetBalance(f float64) *PaymentUpdate {
	pu.mutation.ResetBalance()
	pu.mutation.SetBalance(f)
	return pu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableBalance(f *float64) *PaymentUpdate {
	if f != nil {
		pu.SetBalance(*f)
	}
	return pu
}

// AddBalance adds f to the "balance" field.
func (pu *PaymentUpdate) AddBalance(f float64) *PaymentUpdate {
	pu.mutation.AddBalance(f)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PaymentUpdate) SetAmount(f float64) *PaymentUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(f)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableAmount(f *float64) *PaymentUpdate {
	if f != nil {
		pu.SetAmount(*f)
	}
	return pu
}

// AddAmount adds f to the "amount" field.
func (pu *PaymentUpdate) AddAmount(f float64) *PaymentUpdate {
	pu.mutation.AddAmount(f)
	return pu
}

// SetBaseAmount sets the "base_amount" field.
func (pu *PaymentUpdate) SetBaseAmount(f float64) *PaymentUpdate {
	pu.mutation.ResetBaseAmount()
	pu.mutation.SetBaseAmount(f)
	return pu
}

// SetNillableBaseAmount sets the "base_amount" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableBaseAmount(f *float64) *PaymentUpdate {
	if f != nil {
		pu.SetBaseAmount(*f)
	}
	return pu
}

// AddBaseAmount adds f to the "base_amount" field.
func (pu *PaymentUpdate) AddBaseAmount(f float64) *PaymentUpdate {
	pu.mutation.AddBaseAmount(f)
	return pu
}

// SetUsedAmount sets the "used_amount" field.
func (pu *PaymentUpdate) SetUsedAmount(f float64) *PaymentUpdate {
	pu.mutation.ResetUsedAmount()
	pu.mutation.SetUsedAmount(f)
	return pu
}

// SetNillableUsedAmount sets the "used_amount" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableUsedAmount(f *float64) *PaymentUpdate {
	if f != nil {
		pu.SetUsedAmount(*f)
	}
	return pu
}

// AddUsedAmount adds f to the "used_amount" field.
func (pu *PaymentUpdate) AddUsedAmount(f float64) *PaymentUpdate {
	pu.mutation.AddUsedAmount(f)
	return pu
}

// SetType sets the "type" field.
func (pu *PaymentUpdate) SetType(pa payment.Type) *PaymentUpdate {
	pu.mutation.SetType(pa)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableType(pa *payment.Type) *PaymentUpdate {
	if pa != nil {
		pu.SetType(*pa)
	}
	return pu
}

// SetFop sets the "fop" field.
func (pu *PaymentUpdate) SetFop(pa payment.Fop) *PaymentUpdate {
	pu.mutation.SetFop(pa)
	return pu
}

// SetNillableFop sets the "fop" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableFop(pa *payment.Fop) *PaymentUpdate {
	if pa != nil {
		pu.SetFop(*pa)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *PaymentUpdate) SetStatus(pa payment.Status) *PaymentUpdate {
	pu.mutation.SetStatus(pa)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableStatus(pa *payment.Status) *PaymentUpdate {
	if pa != nil {
		pu.SetStatus(*pa)
	}
	return pu
}

// SetIDChartOfAccounts sets the "id_chart_of_accounts" field.
func (pu *PaymentUpdate) SetIDChartOfAccounts(i int) *PaymentUpdate {
	pu.mutation.ResetIDChartOfAccounts()
	pu.mutation.SetIDChartOfAccounts(i)
	return pu
}

// SetNillableIDChartOfAccounts sets the "id_chart_of_accounts" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableIDChartOfAccounts(i *int) *PaymentUpdate {
	if i != nil {
		pu.SetIDChartOfAccounts(*i)
	}
	return pu
}

// AddIDChartOfAccounts adds i to the "id_chart_of_accounts" field.
func (pu *PaymentUpdate) AddIDChartOfAccounts(i int) *PaymentUpdate {
	pu.mutation.AddIDChartOfAccounts(i)
	return pu
}

// SetIDCurrency sets the "id_currency" field.
func (pu *PaymentUpdate) SetIDCurrency(i int) *PaymentUpdate {
	pu.mutation.ResetIDCurrency()
	pu.mutation.SetIDCurrency(i)
	return pu
}

// SetNillableIDCurrency sets the "id_currency" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableIDCurrency(i *int) *PaymentUpdate {
	if i != nil {
		pu.SetIDCurrency(*i)
	}
	return pu
}

// AddIDCurrency adds i to the "id_currency" field.
func (pu *PaymentUpdate) AddIDCurrency(i int) *PaymentUpdate {
	pu.mutation.AddIDCurrency(i)
	return pu
}

// SetTag sets the "Tag" field.
func (pu *PaymentUpdate) SetTag(pa payment.Tag) *PaymentUpdate {
	pu.mutation.SetTag(pa)
	return pu
}

// SetNillableTag sets the "Tag" field if the given value is not nil.
func (pu *PaymentUpdate) SetNillableTag(pa *payment.Tag) *PaymentUpdate {
	if pa != nil {
		pu.SetTag(*pa)
	}
	return pu
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (pu *PaymentUpdate) SetCustomerID(id int) *PaymentUpdate {
	pu.mutation.SetCustomerID(id)
	return pu
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (pu *PaymentUpdate) SetCustomer(c *Customer) *PaymentUpdate {
	return pu.SetCustomerID(c.ID)
}

// AddImputationIDs adds the "imputations" edge to the Imputation entity by IDs.
func (pu *PaymentUpdate) AddImputationIDs(ids ...int) *PaymentUpdate {
	pu.mutation.AddImputationIDs(ids...)
	return pu
}

// AddImputations adds the "imputations" edges to the Imputation entity.
func (pu *PaymentUpdate) AddImputations(i ...*Imputation) *PaymentUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.AddImputationIDs(ids...)
}

// Mutation returns the PaymentMutation object of the builder.
func (pu *PaymentUpdate) Mutation() *PaymentMutation {
	return pu.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (pu *PaymentUpdate) ClearCustomer() *PaymentUpdate {
	pu.mutation.ClearCustomer()
	return pu
}

// ClearImputations clears all "imputations" edges to the Imputation entity.
func (pu *PaymentUpdate) ClearImputations() *PaymentUpdate {
	pu.mutation.ClearImputations()
	return pu
}

// RemoveImputationIDs removes the "imputations" edge to Imputation entities by IDs.
func (pu *PaymentUpdate) RemoveImputationIDs(ids ...int) *PaymentUpdate {
	pu.mutation.RemoveImputationIDs(ids...)
	return pu
}

// RemoveImputations removes "imputations" edges to Imputation entities.
func (pu *PaymentUpdate) RemoveImputations(i ...*Imputation) *PaymentUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.RemoveImputationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PaymentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PaymentUpdate) check() error {
	if v, ok := pu.mutation.Number(); ok {
		if err := payment.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Payment.number": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Date(); ok {
		if err := payment.DateValidator(v); err != nil {
			return &ValidationError{Name: "date", err: fmt.Errorf(`ent: validator failed for field "Payment.date": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Balance(); ok {
		if err := payment.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`ent: validator failed for field "Payment.balance": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if v, ok := pu.mutation.BaseAmount(); ok {
		if err := payment.BaseAmountValidator(v); err != nil {
			return &ValidationError{Name: "base_amount", err: fmt.Errorf(`ent: validator failed for field "Payment.base_amount": %w`, err)}
		}
	}
	if v, ok := pu.mutation.UsedAmount(); ok {
		if err := payment.UsedAmountValidator(v); err != nil {
			return &ValidationError{Name: "used_amount", err: fmt.Errorf(`ent: validator failed for field "Payment.used_amount": %w`, err)}
		}
	}
	if v, ok := pu.mutation.GetType(); ok {
		if err := payment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Payment.type": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Fop(); ok {
		if err := payment.FopValidator(v); err != nil {
			return &ValidationError{Name: "fop", err: fmt.Errorf(`ent: validator failed for field "Payment.fop": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Tag(); ok {
		if err := payment.TagValidator(v); err != nil {
			return &ValidationError{Name: "Tag", err: fmt.Errorf(`ent: validator failed for field "Payment.Tag": %w`, err)}
		}
	}
	if _, ok := pu.mutation.CustomerID(); pu.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Payment.customer"`)
	}
	return nil
}

func (pu *PaymentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Number(); ok {
		_spec.SetField(payment.FieldNumber, field.TypeString, value)
	}
	if value, ok := pu.mutation.Date(); ok {
		_spec.SetField(payment.FieldDate, field.TypeString, value)
	}
	if value, ok := pu.mutation.Balance(); ok {
		vv, err := payment.ValueScanner.Balance.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(payment.FieldBalance, field.TypeFloat64, vv)
	}
	if value, ok := pu.mutation.AddedBalance(); ok {
		vv, err := payment.ValueScanner.Balance.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.AddField(payment.FieldBalance, field.TypeFloat64, vv)
	}
	if value, ok := pu.mutation.Amount(); ok {
		vv, err := payment.ValueScanner.Amount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, vv)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		vv, err := payment.ValueScanner.Amount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, vv)
	}
	if value, ok := pu.mutation.BaseAmount(); ok {
		vv, err := payment.ValueScanner.BaseAmount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(payment.FieldBaseAmount, field.TypeFloat64, vv)
	}
	if value, ok := pu.mutation.AddedBaseAmount(); ok {
		vv, err := payment.ValueScanner.BaseAmount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.AddField(payment.FieldBaseAmount, field.TypeFloat64, vv)
	}
	if value, ok := pu.mutation.UsedAmount(); ok {
		vv, err := payment.ValueScanner.UsedAmount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(payment.FieldUsedAmount, field.TypeFloat64, vv)
	}
	if value, ok := pu.mutation.AddedUsedAmount(); ok {
		vv, err := payment.ValueScanner.UsedAmount.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.AddField(payment.FieldUsedAmount, field.TypeFloat64, vv)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(payment.FieldType, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.Fop(); ok {
		_spec.SetField(payment.FieldFop, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.IDChartOfAccounts(); ok {
		_spec.SetField(payment.FieldIDChartOfAccounts, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedIDChartOfAccounts(); ok {
		_spec.AddField(payment.FieldIDChartOfAccounts, field.TypeInt, value)
	}
	if value, ok := pu.mutation.IDCurrency(); ok {
		_spec.SetField(payment.FieldIDCurrency, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedIDCurrency(); ok {
		_spec.AddField(payment.FieldIDCurrency, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Tag(); ok {
		_spec.SetField(payment.FieldTag, field.TypeEnum, value)
	}
	if pu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.CustomerTable,
			Columns: []string{payment.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.CustomerTable,
			Columns: []string{payment.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ImputationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.ImputationsTable,
			Columns: []string{payment.ImputationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedImputationsIDs(); len(nodes) > 0 && !pu.mutation.ImputationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.ImputationsTable,
			Columns: []string{payment.ImputationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ImputationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.ImputationsTable,
			Columns: []string{payment.ImputationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PaymentUpdateOne is the builder for updating a single Payment entity.
type PaymentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentMutation
}

// SetNumber sets the "number" field.
func (puo *PaymentUpdateOne) SetNumber(s string) *PaymentUpdateOne {
	puo.mutation.SetNumber(s)
	return puo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableNumber(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetNumber(*s)
	}
	return puo
}

// SetDate sets the "date" field.
func (puo *PaymentUpdateOne) SetDate(s string) *PaymentUpdateOne {
	puo.mutation.SetDate(s)
	return puo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableDate(s *string) *PaymentUpdateOne {
	if s != nil {
		puo.SetDate(*s)
	}
	return puo
}

// SetBalance sets the "balance" field.
func (puo *PaymentUpdateOne) SetBalance(f float64) *PaymentUpdateOne {
	puo.mutation.ResetBalance()
	puo.mutation.SetBalance(f)
	return puo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableBalance(f *float64) *PaymentUpdateOne {
	if f != nil {
		puo.SetBalance(*f)
	}
	return puo
}

// AddBalance adds f to the "balance" field.
func (puo *PaymentUpdateOne) AddBalance(f float64) *PaymentUpdateOne {
	puo.mutation.AddBalance(f)
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PaymentUpdateOne) SetAmount(f float64) *PaymentUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(f)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableAmount(f *float64) *PaymentUpdateOne {
	if f != nil {
		puo.SetAmount(*f)
	}
	return puo
}

// AddAmount adds f to the "amount" field.
func (puo *PaymentUpdateOne) AddAmount(f float64) *PaymentUpdateOne {
	puo.mutation.AddAmount(f)
	return puo
}

// SetBaseAmount sets the "base_amount" field.
func (puo *PaymentUpdateOne) SetBaseAmount(f float64) *PaymentUpdateOne {
	puo.mutation.ResetBaseAmount()
	puo.mutation.SetBaseAmount(f)
	return puo
}

// SetNillableBaseAmount sets the "base_amount" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableBaseAmount(f *float64) *PaymentUpdateOne {
	if f != nil {
		puo.SetBaseAmount(*f)
	}
	return puo
}

// AddBaseAmount adds f to the "base_amount" field.
func (puo *PaymentUpdateOne) AddBaseAmount(f float64) *PaymentUpdateOne {
	puo.mutation.AddBaseAmount(f)
	return puo
}

// SetUsedAmount sets the "used_amount" field.
func (puo *PaymentUpdateOne) SetUsedAmount(f float64) *PaymentUpdateOne {
	puo.mutation.ResetUsedAmount()
	puo.mutation.SetUsedAmount(f)
	return puo
}

// SetNillableUsedAmount sets the "used_amount" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableUsedAmount(f *float64) *PaymentUpdateOne {
	if f != nil {
		puo.SetUsedAmount(*f)
	}
	return puo
}

// AddUsedAmount adds f to the "used_amount" field.
func (puo *PaymentUpdateOne) AddUsedAmount(f float64) *PaymentUpdateOne {
	puo.mutation.AddUsedAmount(f)
	return puo
}

// SetType sets the "type" field.
func (puo *PaymentUpdateOne) SetType(pa payment.Type) *PaymentUpdateOne {
	puo.mutation.SetType(pa)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableType(pa *payment.Type) *PaymentUpdateOne {
	if pa != nil {
		puo.SetType(*pa)
	}
	return puo
}

// SetFop sets the "fop" field.
func (puo *PaymentUpdateOne) SetFop(pa payment.Fop) *PaymentUpdateOne {
	puo.mutation.SetFop(pa)
	return puo
}

// SetNillableFop sets the "fop" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableFop(pa *payment.Fop) *PaymentUpdateOne {
	if pa != nil {
		puo.SetFop(*pa)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *PaymentUpdateOne) SetStatus(pa payment.Status) *PaymentUpdateOne {
	puo.mutation.SetStatus(pa)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableStatus(pa *payment.Status) *PaymentUpdateOne {
	if pa != nil {
		puo.SetStatus(*pa)
	}
	return puo
}

// SetIDChartOfAccounts sets the "id_chart_of_accounts" field.
func (puo *PaymentUpdateOne) SetIDChartOfAccounts(i int) *PaymentUpdateOne {
	puo.mutation.ResetIDChartOfAccounts()
	puo.mutation.SetIDChartOfAccounts(i)
	return puo
}

// SetNillableIDChartOfAccounts sets the "id_chart_of_accounts" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableIDChartOfAccounts(i *int) *PaymentUpdateOne {
	if i != nil {
		puo.SetIDChartOfAccounts(*i)
	}
	return puo
}

// AddIDChartOfAccounts adds i to the "id_chart_of_accounts" field.
func (puo *PaymentUpdateOne) AddIDChartOfAccounts(i int) *PaymentUpdateOne {
	puo.mutation.AddIDChartOfAccounts(i)
	return puo
}

// SetIDCurrency sets the "id_currency" field.
func (puo *PaymentUpdateOne) SetIDCurrency(i int) *PaymentUpdateOne {
	puo.mutation.ResetIDCurrency()
	puo.mutation.SetIDCurrency(i)
	return puo
}

// SetNillableIDCurrency sets the "id_currency" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableIDCurrency(i *int) *PaymentUpdateOne {
	if i != nil {
		puo.SetIDCurrency(*i)
	}
	return puo
}

// AddIDCurrency adds i to the "id_currency" field.
func (puo *PaymentUpdateOne) AddIDCurrency(i int) *PaymentUpdateOne {
	puo.mutation.AddIDCurrency(i)
	return puo
}

// SetTag sets the "Tag" field.
func (puo *PaymentUpdateOne) SetTag(pa payment.Tag) *PaymentUpdateOne {
	puo.mutation.SetTag(pa)
	return puo
}

// SetNillableTag sets the "Tag" field if the given value is not nil.
func (puo *PaymentUpdateOne) SetNillableTag(pa *payment.Tag) *PaymentUpdateOne {
	if pa != nil {
		puo.SetTag(*pa)
	}
	return puo
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (puo *PaymentUpdateOne) SetCustomerID(id int) *PaymentUpdateOne {
	puo.mutation.SetCustomerID(id)
	return puo
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (puo *PaymentUpdateOne) SetCustomer(c *Customer) *PaymentUpdateOne {
	return puo.SetCustomerID(c.ID)
}

// AddImputationIDs adds the "imputations" edge to the Imputation entity by IDs.
func (puo *PaymentUpdateOne) AddImputationIDs(ids ...int) *PaymentUpdateOne {
	puo.mutation.AddImputationIDs(ids...)
	return puo
}

// AddImputations adds the "imputations" edges to the Imputation entity.
func (puo *PaymentUpdateOne) AddImputations(i ...*Imputation) *PaymentUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.AddImputationIDs(ids...)
}

// Mutation returns the PaymentMutation object of the builder.
func (puo *PaymentUpdateOne) Mutation() *PaymentMutation {
	return puo.mutation
}

// ClearCustomer clears the "customer" edge to the Customer entity.
func (puo *PaymentUpdateOne) ClearCustomer() *PaymentUpdateOne {
	puo.mutation.ClearCustomer()
	return puo
}

// ClearImputations clears all "imputations" edges to the Imputation entity.
func (puo *PaymentUpdateOne) ClearImputations() *PaymentUpdateOne {
	puo.mutation.ClearImputations()
	return puo
}

// RemoveImputationIDs removes the "imputations" edge to Imputation entities by IDs.
func (puo *PaymentUpdateOne) RemoveImputationIDs(ids ...int) *PaymentUpdateOne {
	puo.mutation.RemoveImputationIDs(ids...)
	return puo
}

// RemoveImputations removes "imputations" edges to Imputation entities.
func (puo *PaymentUpdateOne) RemoveImputations(i ...*Imputation) *PaymentUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.RemoveImputationIDs(ids...)
}

// Where appends a list predicates to the PaymentUpdate builder.
func (puo *PaymentUpdateOne) Where(ps ...predicate.Payment) *PaymentUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PaymentUpdateOne) Select(field string, fields ...string) *PaymentUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Payment entity.
func (puo *PaymentUpdateOne) Save(ctx context.Context) (*Payment, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentUpdateOne) SaveX(ctx context.Context) *Payment {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PaymentUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PaymentUpdateOne) check() error {
	if v, ok := puo.mutation.Number(); ok {
		if err := payment.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Payment.number": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Date(); ok {
		if err := payment.DateValidator(v); err != nil {
			return &ValidationError{Name: "date", err: fmt.Errorf(`ent: validator failed for field "Payment.date": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Balance(); ok {
		if err := payment.BalanceValidator(v); err != nil {
			return &ValidationError{Name: "balance", err: fmt.Errorf(`ent: validator failed for field "Payment.balance": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Amount(); ok {
		if err := payment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Payment.amount": %w`, err)}
		}
	}
	if v, ok := puo.mutation.BaseAmount(); ok {
		if err := payment.BaseAmountValidator(v); err != nil {
			return &ValidationError{Name: "base_amount", err: fmt.Errorf(`ent: validator failed for field "Payment.base_amount": %w`, err)}
		}
	}
	if v, ok := puo.mutation.UsedAmount(); ok {
		if err := payment.UsedAmountValidator(v); err != nil {
			return &ValidationError{Name: "used_amount", err: fmt.Errorf(`ent: validator failed for field "Payment.used_amount": %w`, err)}
		}
	}
	if v, ok := puo.mutation.GetType(); ok {
		if err := payment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Payment.type": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Fop(); ok {
		if err := payment.FopValidator(v); err != nil {
			return &ValidationError{Name: "fop", err: fmt.Errorf(`ent: validator failed for field "Payment.fop": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := payment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Payment.status": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Tag(); ok {
		if err := payment.TagValidator(v); err != nil {
			return &ValidationError{Name: "Tag", err: fmt.Errorf(`ent: validator failed for field "Payment.Tag": %w`, err)}
		}
	}
	if _, ok := puo.mutation.CustomerID(); puo.mutation.CustomerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Payment.customer"`)
	}
	return nil
}

func (puo *PaymentUpdateOne) sqlSave(ctx context.Context) (_node *Payment, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(payment.Table, payment.Columns, sqlgraph.NewFieldSpec(payment.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Payment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payment.FieldID)
		for _, f := range fields {
			if !payment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Number(); ok {
		_spec.SetField(payment.FieldNumber, field.TypeString, value)
	}
	if value, ok := puo.mutation.Date(); ok {
		_spec.SetField(payment.FieldDate, field.TypeString, value)
	}
	if value, ok := puo.mutation.Balance(); ok {
		vv, err := payment.ValueScanner.Balance.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(payment.FieldBalance, field.TypeFloat64, vv)
	}
	if value, ok := puo.mutation.AddedBalance(); ok {
		vv, err := payment.ValueScanner.Balance.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.AddField(payment.FieldBalance, field.TypeFloat64, vv)
	}
	if value, ok := puo.mutation.Amount(); ok {
		vv, err := payment.ValueScanner.Amount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(payment.FieldAmount, field.TypeFloat64, vv)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		vv, err := payment.ValueScanner.Amount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.AddField(payment.FieldAmount, field.TypeFloat64, vv)
	}
	if value, ok := puo.mutation.BaseAmount(); ok {
		vv, err := payment.ValueScanner.BaseAmount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(payment.FieldBaseAmount, field.TypeFloat64, vv)
	}
	if value, ok := puo.mutation.AddedBaseAmount(); ok {
		vv, err := payment.ValueScanner.BaseAmount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.AddField(payment.FieldBaseAmount, field.TypeFloat64, vv)
	}
	if value, ok := puo.mutation.UsedAmount(); ok {
		vv, err := payment.ValueScanner.UsedAmount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(payment.FieldUsedAmount, field.TypeFloat64, vv)
	}
	if value, ok := puo.mutation.AddedUsedAmount(); ok {
		vv, err := payment.ValueScanner.UsedAmount.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.AddField(payment.FieldUsedAmount, field.TypeFloat64, vv)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(payment.FieldType, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.Fop(); ok {
		_spec.SetField(payment.FieldFop, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(payment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.IDChartOfAccounts(); ok {
		_spec.SetField(payment.FieldIDChartOfAccounts, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedIDChartOfAccounts(); ok {
		_spec.AddField(payment.FieldIDChartOfAccounts, field.TypeInt, value)
	}
	if value, ok := puo.mutation.IDCurrency(); ok {
		_spec.SetField(payment.FieldIDCurrency, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedIDCurrency(); ok {
		_spec.AddField(payment.FieldIDCurrency, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Tag(); ok {
		_spec.SetField(payment.FieldTag, field.TypeEnum, value)
	}
	if puo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.CustomerTable,
			Columns: []string{payment.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payment.CustomerTable,
			Columns: []string{payment.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ImputationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.ImputationsTable,
			Columns: []string{payment.ImputationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedImputationsIDs(); len(nodes) > 0 && !puo.mutation.ImputationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.ImputationsTable,
			Columns: []string{payment.ImputationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ImputationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   payment.ImputationsTable,
			Columns: []string{payment.ImputationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Payment{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
