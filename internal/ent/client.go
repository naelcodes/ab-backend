// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/naelcodes/ab-backend/internal/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/naelcodes/ab-backend/internal/ent/customer"
	"github.com/naelcodes/ab-backend/internal/ent/imputation"
	"github.com/naelcodes/ab-backend/internal/ent/invoice"
	"github.com/naelcodes/ab-backend/internal/ent/payment"
	"github.com/naelcodes/ab-backend/internal/ent/travelitem"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Imputation is the client for interacting with the Imputation builders.
	Imputation *ImputationClient
	// Invoice is the client for interacting with the Invoice builders.
	Invoice *InvoiceClient
	// Payment is the client for interacting with the Payment builders.
	Payment *PaymentClient
	// TravelItem is the client for interacting with the TravelItem builders.
	TravelItem *TravelItemClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Customer = NewCustomerClient(c.config)
	c.Imputation = NewImputationClient(c.config)
	c.Invoice = NewInvoiceClient(c.config)
	c.Payment = NewPaymentClient(c.config)
	c.TravelItem = NewTravelItemClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Customer:   NewCustomerClient(cfg),
		Imputation: NewImputationClient(cfg),
		Invoice:    NewInvoiceClient(cfg),
		Payment:    NewPaymentClient(cfg),
		TravelItem: NewTravelItemClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Customer:   NewCustomerClient(cfg),
		Imputation: NewImputationClient(cfg),
		Invoice:    NewInvoiceClient(cfg),
		Payment:    NewPaymentClient(cfg),
		TravelItem: NewTravelItemClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Customer.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Customer.Use(hooks...)
	c.Imputation.Use(hooks...)
	c.Invoice.Use(hooks...)
	c.Payment.Use(hooks...)
	c.TravelItem.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Customer.Intercept(interceptors...)
	c.Imputation.Intercept(interceptors...)
	c.Invoice.Intercept(interceptors...)
	c.Payment.Intercept(interceptors...)
	c.TravelItem.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CustomerMutation:
		return c.Customer.mutate(ctx, m)
	case *ImputationMutation:
		return c.Imputation.mutate(ctx, m)
	case *InvoiceMutation:
		return c.Invoice.mutate(ctx, m)
	case *PaymentMutation:
		return c.Payment.mutate(ctx, m)
	case *TravelItemMutation:
		return c.TravelItem.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `customer.Intercept(f(g(h())))`.
func (c *CustomerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Customer = append(c.inters.Customer, interceptors...)
}

// Create returns a builder for creating a Customer entity.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Customer entities.
func (c *CustomerClient) CreateBulk(builders ...*CustomerCreate) *CustomerCreateBulk {
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CustomerClient) MapCreateBulk(slice any, setFunc func(*CustomerCreate, int)) *CustomerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CustomerCreateBulk{err: fmt.Errorf("calling to CustomerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CustomerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id int) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerClient) DeleteOneID(id int) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Query returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCustomer},
		inters: c.Interceptors(),
	}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id int) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id int) *Customer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInvoices queries the invoices edge of a Customer.
func (c *CustomerClient) QueryInvoices(cu *Customer) *InvoiceQuery {
	query := (&InvoiceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.InvoicesTable, customer.InvoicesColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPayments queries the payments edge of a Customer.
func (c *CustomerClient) QueryPayments(cu *Customer) *PaymentQuery {
	query := (&PaymentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(payment.Table, payment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.PaymentsTable, customer.PaymentsColumn),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// Interceptors returns the client interceptors.
func (c *CustomerClient) Interceptors() []Interceptor {
	return c.inters.Customer
}

func (c *CustomerClient) mutate(ctx context.Context, m *CustomerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Customer mutation op: %q", m.Op())
	}
}

// ImputationClient is a client for the Imputation schema.
type ImputationClient struct {
	config
}

// NewImputationClient returns a client for the Imputation from the given config.
func NewImputationClient(c config) *ImputationClient {
	return &ImputationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `imputation.Hooks(f(g(h())))`.
func (c *ImputationClient) Use(hooks ...Hook) {
	c.hooks.Imputation = append(c.hooks.Imputation, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `imputation.Intercept(f(g(h())))`.
func (c *ImputationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Imputation = append(c.inters.Imputation, interceptors...)
}

// Create returns a builder for creating a Imputation entity.
func (c *ImputationClient) Create() *ImputationCreate {
	mutation := newImputationMutation(c.config, OpCreate)
	return &ImputationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Imputation entities.
func (c *ImputationClient) CreateBulk(builders ...*ImputationCreate) *ImputationCreateBulk {
	return &ImputationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ImputationClient) MapCreateBulk(slice any, setFunc func(*ImputationCreate, int)) *ImputationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ImputationCreateBulk{err: fmt.Errorf("calling to ImputationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ImputationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ImputationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Imputation.
func (c *ImputationClient) Update() *ImputationUpdate {
	mutation := newImputationMutation(c.config, OpUpdate)
	return &ImputationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImputationClient) UpdateOne(i *Imputation) *ImputationUpdateOne {
	mutation := newImputationMutation(c.config, OpUpdateOne, withImputation(i))
	return &ImputationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImputationClient) UpdateOneID(id int) *ImputationUpdateOne {
	mutation := newImputationMutation(c.config, OpUpdateOne, withImputationID(id))
	return &ImputationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Imputation.
func (c *ImputationClient) Delete() *ImputationDelete {
	mutation := newImputationMutation(c.config, OpDelete)
	return &ImputationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ImputationClient) DeleteOne(i *Imputation) *ImputationDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ImputationClient) DeleteOneID(id int) *ImputationDeleteOne {
	builder := c.Delete().Where(imputation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImputationDeleteOne{builder}
}

// Query returns a query builder for Imputation.
func (c *ImputationClient) Query() *ImputationQuery {
	return &ImputationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeImputation},
		inters: c.Interceptors(),
	}
}

// Get returns a Imputation entity by its id.
func (c *ImputationClient) Get(ctx context.Context, id int) (*Imputation, error) {
	return c.Query().Where(imputation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImputationClient) GetX(ctx context.Context, id int) *Imputation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInvoice queries the invoice edge of a Imputation.
func (c *ImputationClient) QueryInvoice(i *Imputation) *InvoiceQuery {
	query := (&InvoiceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(imputation.Table, imputation.FieldID, id),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, imputation.InvoiceTable, imputation.InvoiceColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPayment queries the payment edge of a Imputation.
func (c *ImputationClient) QueryPayment(i *Imputation) *PaymentQuery {
	query := (&PaymentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(imputation.Table, imputation.FieldID, id),
			sqlgraph.To(payment.Table, payment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, imputation.PaymentTable, imputation.PaymentColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImputationClient) Hooks() []Hook {
	return c.hooks.Imputation
}

// Interceptors returns the client interceptors.
func (c *ImputationClient) Interceptors() []Interceptor {
	return c.inters.Imputation
}

func (c *ImputationClient) mutate(ctx context.Context, m *ImputationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ImputationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ImputationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ImputationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ImputationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Imputation mutation op: %q", m.Op())
	}
}

// InvoiceClient is a client for the Invoice schema.
type InvoiceClient struct {
	config
}

// NewInvoiceClient returns a client for the Invoice from the given config.
func NewInvoiceClient(c config) *InvoiceClient {
	return &InvoiceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `invoice.Hooks(f(g(h())))`.
func (c *InvoiceClient) Use(hooks ...Hook) {
	c.hooks.Invoice = append(c.hooks.Invoice, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `invoice.Intercept(f(g(h())))`.
func (c *InvoiceClient) Intercept(interceptors ...Interceptor) {
	c.inters.Invoice = append(c.inters.Invoice, interceptors...)
}

// Create returns a builder for creating a Invoice entity.
func (c *InvoiceClient) Create() *InvoiceCreate {
	mutation := newInvoiceMutation(c.config, OpCreate)
	return &InvoiceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Invoice entities.
func (c *InvoiceClient) CreateBulk(builders ...*InvoiceCreate) *InvoiceCreateBulk {
	return &InvoiceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *InvoiceClient) MapCreateBulk(slice any, setFunc func(*InvoiceCreate, int)) *InvoiceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &InvoiceCreateBulk{err: fmt.Errorf("calling to InvoiceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*InvoiceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &InvoiceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Invoice.
func (c *InvoiceClient) Update() *InvoiceUpdate {
	mutation := newInvoiceMutation(c.config, OpUpdate)
	return &InvoiceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InvoiceClient) UpdateOne(i *Invoice) *InvoiceUpdateOne {
	mutation := newInvoiceMutation(c.config, OpUpdateOne, withInvoice(i))
	return &InvoiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InvoiceClient) UpdateOneID(id int) *InvoiceUpdateOne {
	mutation := newInvoiceMutation(c.config, OpUpdateOne, withInvoiceID(id))
	return &InvoiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Invoice.
func (c *InvoiceClient) Delete() *InvoiceDelete {
	mutation := newInvoiceMutation(c.config, OpDelete)
	return &InvoiceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *InvoiceClient) DeleteOne(i *Invoice) *InvoiceDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *InvoiceClient) DeleteOneID(id int) *InvoiceDeleteOne {
	builder := c.Delete().Where(invoice.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InvoiceDeleteOne{builder}
}

// Query returns a query builder for Invoice.
func (c *InvoiceClient) Query() *InvoiceQuery {
	return &InvoiceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeInvoice},
		inters: c.Interceptors(),
	}
}

// Get returns a Invoice entity by its id.
func (c *InvoiceClient) Get(ctx context.Context, id int) (*Invoice, error) {
	return c.Query().Where(invoice.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InvoiceClient) GetX(ctx context.Context, id int) *Invoice {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCustomer queries the customer edge of a Invoice.
func (c *InvoiceClient) QueryCustomer(i *Invoice) *CustomerQuery {
	query := (&CustomerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, invoice.CustomerTable, invoice.CustomerColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryImputations queries the imputations edge of a Invoice.
func (c *InvoiceClient) QueryImputations(i *Invoice) *ImputationQuery {
	query := (&ImputationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, id),
			sqlgraph.To(imputation.Table, imputation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, invoice.ImputationsTable, invoice.ImputationsColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTravelItems queries the travel_items edge of a Invoice.
func (c *InvoiceClient) QueryTravelItems(i *Invoice) *TravelItemQuery {
	query := (&TravelItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, id),
			sqlgraph.To(travelitem.Table, travelitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, invoice.TravelItemsTable, invoice.TravelItemsColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InvoiceClient) Hooks() []Hook {
	return c.hooks.Invoice
}

// Interceptors returns the client interceptors.
func (c *InvoiceClient) Interceptors() []Interceptor {
	return c.inters.Invoice
}

func (c *InvoiceClient) mutate(ctx context.Context, m *InvoiceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&InvoiceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&InvoiceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&InvoiceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&InvoiceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Invoice mutation op: %q", m.Op())
	}
}

// PaymentClient is a client for the Payment schema.
type PaymentClient struct {
	config
}

// NewPaymentClient returns a client for the Payment from the given config.
func NewPaymentClient(c config) *PaymentClient {
	return &PaymentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `payment.Hooks(f(g(h())))`.
func (c *PaymentClient) Use(hooks ...Hook) {
	c.hooks.Payment = append(c.hooks.Payment, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `payment.Intercept(f(g(h())))`.
func (c *PaymentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Payment = append(c.inters.Payment, interceptors...)
}

// Create returns a builder for creating a Payment entity.
func (c *PaymentClient) Create() *PaymentCreate {
	mutation := newPaymentMutation(c.config, OpCreate)
	return &PaymentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Payment entities.
func (c *PaymentClient) CreateBulk(builders ...*PaymentCreate) *PaymentCreateBulk {
	return &PaymentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PaymentClient) MapCreateBulk(slice any, setFunc func(*PaymentCreate, int)) *PaymentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PaymentCreateBulk{err: fmt.Errorf("calling to PaymentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PaymentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PaymentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Payment.
func (c *PaymentClient) Update() *PaymentUpdate {
	mutation := newPaymentMutation(c.config, OpUpdate)
	return &PaymentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaymentClient) UpdateOne(pa *Payment) *PaymentUpdateOne {
	mutation := newPaymentMutation(c.config, OpUpdateOne, withPayment(pa))
	return &PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PaymentClient) UpdateOneID(id int) *PaymentUpdateOne {
	mutation := newPaymentMutation(c.config, OpUpdateOne, withPaymentID(id))
	return &PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Payment.
func (c *PaymentClient) Delete() *PaymentDelete {
	mutation := newPaymentMutation(c.config, OpDelete)
	return &PaymentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PaymentClient) DeleteOne(pa *Payment) *PaymentDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PaymentClient) DeleteOneID(id int) *PaymentDeleteOne {
	builder := c.Delete().Where(payment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PaymentDeleteOne{builder}
}

// Query returns a query builder for Payment.
func (c *PaymentClient) Query() *PaymentQuery {
	return &PaymentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePayment},
		inters: c.Interceptors(),
	}
}

// Get returns a Payment entity by its id.
func (c *PaymentClient) Get(ctx context.Context, id int) (*Payment, error) {
	return c.Query().Where(payment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaymentClient) GetX(ctx context.Context, id int) *Payment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCustomer queries the customer edge of a Payment.
func (c *PaymentClient) QueryCustomer(pa *Payment) *CustomerQuery {
	query := (&CustomerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(payment.Table, payment.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payment.CustomerTable, payment.CustomerColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryImputations queries the imputations edge of a Payment.
func (c *PaymentClient) QueryImputations(pa *Payment) *ImputationQuery {
	query := (&ImputationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(payment.Table, payment.FieldID, id),
			sqlgraph.To(imputation.Table, imputation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, payment.ImputationsTable, payment.ImputationsColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PaymentClient) Hooks() []Hook {
	return c.hooks.Payment
}

// Interceptors returns the client interceptors.
func (c *PaymentClient) Interceptors() []Interceptor {
	return c.inters.Payment
}

func (c *PaymentClient) mutate(ctx context.Context, m *PaymentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PaymentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PaymentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PaymentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Payment mutation op: %q", m.Op())
	}
}

// TravelItemClient is a client for the TravelItem schema.
type TravelItemClient struct {
	config
}

// NewTravelItemClient returns a client for the TravelItem from the given config.
func NewTravelItemClient(c config) *TravelItemClient {
	return &TravelItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `travelitem.Hooks(f(g(h())))`.
func (c *TravelItemClient) Use(hooks ...Hook) {
	c.hooks.TravelItem = append(c.hooks.TravelItem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `travelitem.Intercept(f(g(h())))`.
func (c *TravelItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.TravelItem = append(c.inters.TravelItem, interceptors...)
}

// Create returns a builder for creating a TravelItem entity.
func (c *TravelItemClient) Create() *TravelItemCreate {
	mutation := newTravelItemMutation(c.config, OpCreate)
	return &TravelItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TravelItem entities.
func (c *TravelItemClient) CreateBulk(builders ...*TravelItemCreate) *TravelItemCreateBulk {
	return &TravelItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TravelItemClient) MapCreateBulk(slice any, setFunc func(*TravelItemCreate, int)) *TravelItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TravelItemCreateBulk{err: fmt.Errorf("calling to TravelItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TravelItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TravelItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TravelItem.
func (c *TravelItemClient) Update() *TravelItemUpdate {
	mutation := newTravelItemMutation(c.config, OpUpdate)
	return &TravelItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TravelItemClient) UpdateOne(ti *TravelItem) *TravelItemUpdateOne {
	mutation := newTravelItemMutation(c.config, OpUpdateOne, withTravelItem(ti))
	return &TravelItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TravelItemClient) UpdateOneID(id int) *TravelItemUpdateOne {
	mutation := newTravelItemMutation(c.config, OpUpdateOne, withTravelItemID(id))
	return &TravelItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TravelItem.
func (c *TravelItemClient) Delete() *TravelItemDelete {
	mutation := newTravelItemMutation(c.config, OpDelete)
	return &TravelItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TravelItemClient) DeleteOne(ti *TravelItem) *TravelItemDeleteOne {
	return c.DeleteOneID(ti.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TravelItemClient) DeleteOneID(id int) *TravelItemDeleteOne {
	builder := c.Delete().Where(travelitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TravelItemDeleteOne{builder}
}

// Query returns a query builder for TravelItem.
func (c *TravelItemClient) Query() *TravelItemQuery {
	return &TravelItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTravelItem},
		inters: c.Interceptors(),
	}
}

// Get returns a TravelItem entity by its id.
func (c *TravelItemClient) Get(ctx context.Context, id int) (*TravelItem, error) {
	return c.Query().Where(travelitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TravelItemClient) GetX(ctx context.Context, id int) *TravelItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInvoice queries the invoice edge of a TravelItem.
func (c *TravelItemClient) QueryInvoice(ti *TravelItem) *InvoiceQuery {
	query := (&InvoiceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ti.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(travelitem.Table, travelitem.FieldID, id),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, travelitem.InvoiceTable, travelitem.InvoiceColumn),
		)
		fromV = sqlgraph.Neighbors(ti.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TravelItemClient) Hooks() []Hook {
	return c.hooks.TravelItem
}

// Interceptors returns the client interceptors.
func (c *TravelItemClient) Interceptors() []Interceptor {
	return c.inters.TravelItem
}

func (c *TravelItemClient) mutate(ctx context.Context, m *TravelItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TravelItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TravelItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TravelItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TravelItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TravelItem mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Customer, Imputation, Invoice, Payment, TravelItem []ent.Hook
	}
	inters struct {
		Customer, Imputation, Invoice, Payment, TravelItem []ent.Interceptor
	}
)
