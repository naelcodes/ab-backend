// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/naelcodes/ab-backend/ent/imputation"
	"github.com/naelcodes/ab-backend/ent/invoice"
	"github.com/naelcodes/ab-backend/ent/payment"
	"github.com/naelcodes/ab-backend/ent/predicate"
)

// ImputationQuery is the builder for querying Imputation entities.
type ImputationQuery struct {
	config
	ctx         *QueryContext
	order       []imputation.OrderOption
	inters      []Interceptor
	predicates  []predicate.Imputation
	withInvoice *InvoiceQuery
	withPayment *PaymentQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImputationQuery builder.
func (iq *ImputationQuery) Where(ps ...predicate.Imputation) *ImputationQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *ImputationQuery) Limit(limit int) *ImputationQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *ImputationQuery) Offset(offset int) *ImputationQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *ImputationQuery) Unique(unique bool) *ImputationQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *ImputationQuery) Order(o ...imputation.OrderOption) *ImputationQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryInvoice chains the current query on the "invoice" edge.
func (iq *ImputationQuery) QueryInvoice() *InvoiceQuery {
	query := (&InvoiceClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(imputation.Table, imputation.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, imputation.InvoiceTable, imputation.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPayment chains the current query on the "payment" edge.
func (iq *ImputationQuery) QueryPayment() *PaymentQuery {
	query := (&PaymentClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(imputation.Table, imputation.FieldID, selector),
			sqlgraph.To(payment.Table, payment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, imputation.PaymentTable, imputation.PaymentColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Imputation entity from the query.
// Returns a *NotFoundError when no Imputation was found.
func (iq *ImputationQuery) First(ctx context.Context) (*Imputation, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{imputation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *ImputationQuery) FirstX(ctx context.Context) *Imputation {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Imputation ID from the query.
// Returns a *NotFoundError when no Imputation ID was found.
func (iq *ImputationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{imputation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *ImputationQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Imputation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Imputation entity is found.
// Returns a *NotFoundError when no Imputation entities are found.
func (iq *ImputationQuery) Only(ctx context.Context) (*Imputation, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{imputation.Label}
	default:
		return nil, &NotSingularError{imputation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *ImputationQuery) OnlyX(ctx context.Context) *Imputation {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Imputation ID in the query.
// Returns a *NotSingularError when more than one Imputation ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *ImputationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{imputation.Label}
	default:
		err = &NotSingularError{imputation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *ImputationQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Imputations.
func (iq *ImputationQuery) All(ctx context.Context) ([]*Imputation, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Imputation, *ImputationQuery]()
	return withInterceptors[[]*Imputation](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *ImputationQuery) AllX(ctx context.Context) []*Imputation {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Imputation IDs.
func (iq *ImputationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(imputation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *ImputationQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *ImputationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*ImputationQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *ImputationQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *ImputationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *ImputationQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImputationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *ImputationQuery) Clone() *ImputationQuery {
	if iq == nil {
		return nil
	}
	return &ImputationQuery{
		config:      iq.config,
		ctx:         iq.ctx.Clone(),
		order:       append([]imputation.OrderOption{}, iq.order...),
		inters:      append([]Interceptor{}, iq.inters...),
		predicates:  append([]predicate.Imputation{}, iq.predicates...),
		withInvoice: iq.withInvoice.Clone(),
		withPayment: iq.withPayment.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ImputationQuery) WithInvoice(opts ...func(*InvoiceQuery)) *ImputationQuery {
	query := (&InvoiceClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withInvoice = query
	return iq
}

// WithPayment tells the query-builder to eager-load the nodes that are connected to
// the "payment" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *ImputationQuery) WithPayment(opts ...func(*PaymentQuery)) *ImputationQuery {
	query := (&PaymentClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withPayment = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AmountApply float64 `json:"amount_apply,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Imputation.Query().
//		GroupBy(imputation.FieldAmountApply).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *ImputationQuery) GroupBy(field string, fields ...string) *ImputationGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ImputationGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = imputation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AmountApply float64 `json:"amount_apply,omitempty"`
//	}
//
//	client.Imputation.Query().
//		Select(imputation.FieldAmountApply).
//		Scan(ctx, &v)
func (iq *ImputationQuery) Select(fields ...string) *ImputationSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &ImputationSelect{ImputationQuery: iq}
	sbuild.label = imputation.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ImputationSelect configured with the given aggregations.
func (iq *ImputationQuery) Aggregate(fns ...AggregateFunc) *ImputationSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *ImputationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !imputation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *ImputationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Imputation, error) {
	var (
		nodes       = []*Imputation{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [2]bool{
			iq.withInvoice != nil,
			iq.withPayment != nil,
		}
	)
	if iq.withInvoice != nil || iq.withPayment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, imputation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Imputation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Imputation{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withInvoice; query != nil {
		if err := iq.loadInvoice(ctx, query, nodes, nil,
			func(n *Imputation, e *Invoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withPayment; query != nil {
		if err := iq.loadPayment(ctx, query, nodes, nil,
			func(n *Imputation, e *Payment) { n.Edges.Payment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *ImputationQuery) loadInvoice(ctx context.Context, query *InvoiceQuery, nodes []*Imputation, init func(*Imputation), assign func(*Imputation, *Invoice)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Imputation)
	for i := range nodes {
		if nodes[i].id_invoice == nil {
			continue
		}
		fk := *nodes[i].id_invoice
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(invoice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "id_invoice" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *ImputationQuery) loadPayment(ctx context.Context, query *PaymentQuery, nodes []*Imputation, init func(*Imputation), assign func(*Imputation, *Payment)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Imputation)
	for i := range nodes {
		if nodes[i].id_payment_received == nil {
			continue
		}
		fk := *nodes[i].id_payment_received
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(payment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "id_payment_received" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iq *ImputationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *ImputationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(imputation.Table, imputation.Columns, sqlgraph.NewFieldSpec(imputation.FieldID, field.TypeInt))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, imputation.FieldID)
		for i := range fields {
			if fields[i] != imputation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *ImputationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(imputation.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = imputation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImputationGroupBy is the group-by builder for Imputation entities.
type ImputationGroupBy struct {
	selector
	build *ImputationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *ImputationGroupBy) Aggregate(fns ...AggregateFunc) *ImputationGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *ImputationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImputationQuery, *ImputationGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *ImputationGroupBy) sqlScan(ctx context.Context, root *ImputationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ImputationSelect is the builder for selecting fields of Imputation entities.
type ImputationSelect struct {
	*ImputationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *ImputationSelect) Aggregate(fns ...AggregateFunc) *ImputationSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *ImputationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ImputationQuery, *ImputationSelect](ctx, is.ImputationQuery, is, is.inters, v)
}

func (is *ImputationSelect) sqlScan(ctx context.Context, root *ImputationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
