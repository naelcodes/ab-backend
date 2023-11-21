// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/naelcodes/ab-backend/internal/ent/country"
	"github.com/naelcodes/ab-backend/internal/ent/customer"
	"github.com/naelcodes/ab-backend/internal/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCountry  = "Country"
	TypeCustomer = "Customer"
)

// CountryMutation represents an operation that mutates the Country nodes in the graph.
type CountryMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	code          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Country, error)
	predicates    []predicate.Country
}

var _ ent.Mutation = (*CountryMutation)(nil)

// countryOption allows management of the mutation configuration using functional options.
type countryOption func(*CountryMutation)

// newCountryMutation creates new mutation for the Country entity.
func newCountryMutation(c config, op Op, opts ...countryOption) *CountryMutation {
	m := &CountryMutation{
		config:        c,
		op:            op,
		typ:           TypeCountry,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCountryID sets the ID field of the mutation.
func withCountryID(id int) countryOption {
	return func(m *CountryMutation) {
		var (
			err   error
			once  sync.Once
			value *Country
		)
		m.oldValue = func(ctx context.Context) (*Country, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Country.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCountry sets the old Country of the mutation.
func withCountry(node *Country) countryOption {
	return func(m *CountryMutation) {
		m.oldValue = func(context.Context) (*Country, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CountryMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CountryMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CountryMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CountryMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Country.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *CountryMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CountryMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Country entity.
// If the Country object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CountryMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *CountryMutation) ResetName() {
	m.name = nil
}

// SetCode sets the "code" field.
func (m *CountryMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *CountryMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the Country entity.
// If the Country object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CountryMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *CountryMutation) ResetCode() {
	m.code = nil
}

// Where appends a list predicates to the CountryMutation builder.
func (m *CountryMutation) Where(ps ...predicate.Country) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CountryMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CountryMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Country, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CountryMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CountryMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Country).
func (m *CountryMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CountryMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, country.FieldName)
	}
	if m.code != nil {
		fields = append(fields, country.FieldCode)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CountryMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case country.FieldName:
		return m.Name()
	case country.FieldCode:
		return m.Code()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CountryMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case country.FieldName:
		return m.OldName(ctx)
	case country.FieldCode:
		return m.OldCode(ctx)
	}
	return nil, fmt.Errorf("unknown Country field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CountryMutation) SetField(name string, value ent.Value) error {
	switch name {
	case country.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case country.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	}
	return fmt.Errorf("unknown Country field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CountryMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CountryMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CountryMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Country numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CountryMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CountryMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CountryMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Country nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CountryMutation) ResetField(name string) error {
	switch name {
	case country.FieldName:
		m.ResetName()
		return nil
	case country.FieldCode:
		m.ResetCode()
		return nil
	}
	return fmt.Errorf("unknown Country field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CountryMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CountryMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CountryMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CountryMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CountryMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CountryMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CountryMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Country unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CountryMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Country edge %s", name)
}

// CustomerMutation represents an operation that mutates the Customer nodes in the graph.
type CustomerMutation struct {
	config
	op                Op
	typ               string
	id                *int
	customer_name     *string
	account_number    *string
	id_currency       *int
	addid_currency    *int
	id_country        *int
	addid_country     *int
	alias             *string
	ab_key            *string
	tmc_client_number *string
	_Tag              *customer.Tag
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*Customer, error)
	predicates        []predicate.Customer
}

var _ ent.Mutation = (*CustomerMutation)(nil)

// customerOption allows management of the mutation configuration using functional options.
type customerOption func(*CustomerMutation)

// newCustomerMutation creates new mutation for the Customer entity.
func newCustomerMutation(c config, op Op, opts ...customerOption) *CustomerMutation {
	m := &CustomerMutation{
		config:        c,
		op:            op,
		typ:           TypeCustomer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCustomerID sets the ID field of the mutation.
func withCustomerID(id int) customerOption {
	return func(m *CustomerMutation) {
		var (
			err   error
			once  sync.Once
			value *Customer
		)
		m.oldValue = func(ctx context.Context) (*Customer, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Customer.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCustomer sets the old Customer of the mutation.
func withCustomer(node *Customer) customerOption {
	return func(m *CustomerMutation) {
		m.oldValue = func(context.Context) (*Customer, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CustomerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CustomerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CustomerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CustomerMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Customer.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCustomerName sets the "customer_name" field.
func (m *CustomerMutation) SetCustomerName(s string) {
	m.customer_name = &s
}

// CustomerName returns the value of the "customer_name" field in the mutation.
func (m *CustomerMutation) CustomerName() (r string, exists bool) {
	v := m.customer_name
	if v == nil {
		return
	}
	return *v, true
}

// OldCustomerName returns the old "customer_name" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldCustomerName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCustomerName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCustomerName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCustomerName: %w", err)
	}
	return oldValue.CustomerName, nil
}

// ResetCustomerName resets all changes to the "customer_name" field.
func (m *CustomerMutation) ResetCustomerName() {
	m.customer_name = nil
}

// SetAccountNumber sets the "account_number" field.
func (m *CustomerMutation) SetAccountNumber(s string) {
	m.account_number = &s
}

// AccountNumber returns the value of the "account_number" field in the mutation.
func (m *CustomerMutation) AccountNumber() (r string, exists bool) {
	v := m.account_number
	if v == nil {
		return
	}
	return *v, true
}

// OldAccountNumber returns the old "account_number" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldAccountNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAccountNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAccountNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccountNumber: %w", err)
	}
	return oldValue.AccountNumber, nil
}

// ResetAccountNumber resets all changes to the "account_number" field.
func (m *CustomerMutation) ResetAccountNumber() {
	m.account_number = nil
}

// SetIDCurrency sets the "id_currency" field.
func (m *CustomerMutation) SetIDCurrency(i int) {
	m.id_currency = &i
	m.addid_currency = nil
}

// IDCurrency returns the value of the "id_currency" field in the mutation.
func (m *CustomerMutation) IDCurrency() (r int, exists bool) {
	v := m.id_currency
	if v == nil {
		return
	}
	return *v, true
}

// OldIDCurrency returns the old "id_currency" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldIDCurrency(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIDCurrency is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIDCurrency requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIDCurrency: %w", err)
	}
	return oldValue.IDCurrency, nil
}

// AddIDCurrency adds i to the "id_currency" field.
func (m *CustomerMutation) AddIDCurrency(i int) {
	if m.addid_currency != nil {
		*m.addid_currency += i
	} else {
		m.addid_currency = &i
	}
}

// AddedIDCurrency returns the value that was added to the "id_currency" field in this mutation.
func (m *CustomerMutation) AddedIDCurrency() (r int, exists bool) {
	v := m.addid_currency
	if v == nil {
		return
	}
	return *v, true
}

// ResetIDCurrency resets all changes to the "id_currency" field.
func (m *CustomerMutation) ResetIDCurrency() {
	m.id_currency = nil
	m.addid_currency = nil
}

// SetIDCountry sets the "id_country" field.
func (m *CustomerMutation) SetIDCountry(i int) {
	m.id_country = &i
	m.addid_country = nil
}

// IDCountry returns the value of the "id_country" field in the mutation.
func (m *CustomerMutation) IDCountry() (r int, exists bool) {
	v := m.id_country
	if v == nil {
		return
	}
	return *v, true
}

// OldIDCountry returns the old "id_country" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldIDCountry(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIDCountry is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIDCountry requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIDCountry: %w", err)
	}
	return oldValue.IDCountry, nil
}

// AddIDCountry adds i to the "id_country" field.
func (m *CustomerMutation) AddIDCountry(i int) {
	if m.addid_country != nil {
		*m.addid_country += i
	} else {
		m.addid_country = &i
	}
}

// AddedIDCountry returns the value that was added to the "id_country" field in this mutation.
func (m *CustomerMutation) AddedIDCountry() (r int, exists bool) {
	v := m.addid_country
	if v == nil {
		return
	}
	return *v, true
}

// ResetIDCountry resets all changes to the "id_country" field.
func (m *CustomerMutation) ResetIDCountry() {
	m.id_country = nil
	m.addid_country = nil
}

// SetAlias sets the "alias" field.
func (m *CustomerMutation) SetAlias(s string) {
	m.alias = &s
}

// Alias returns the value of the "alias" field in the mutation.
func (m *CustomerMutation) Alias() (r string, exists bool) {
	v := m.alias
	if v == nil {
		return
	}
	return *v, true
}

// OldAlias returns the old "alias" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldAlias(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAlias is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAlias requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAlias: %w", err)
	}
	return oldValue.Alias, nil
}

// ResetAlias resets all changes to the "alias" field.
func (m *CustomerMutation) ResetAlias() {
	m.alias = nil
}

// SetAbKey sets the "ab_key" field.
func (m *CustomerMutation) SetAbKey(s string) {
	m.ab_key = &s
}

// AbKey returns the value of the "ab_key" field in the mutation.
func (m *CustomerMutation) AbKey() (r string, exists bool) {
	v := m.ab_key
	if v == nil {
		return
	}
	return *v, true
}

// OldAbKey returns the old "ab_key" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldAbKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAbKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAbKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAbKey: %w", err)
	}
	return oldValue.AbKey, nil
}

// ResetAbKey resets all changes to the "ab_key" field.
func (m *CustomerMutation) ResetAbKey() {
	m.ab_key = nil
}

// SetTmcClientNumber sets the "tmc_client_number" field.
func (m *CustomerMutation) SetTmcClientNumber(s string) {
	m.tmc_client_number = &s
}

// TmcClientNumber returns the value of the "tmc_client_number" field in the mutation.
func (m *CustomerMutation) TmcClientNumber() (r string, exists bool) {
	v := m.tmc_client_number
	if v == nil {
		return
	}
	return *v, true
}

// OldTmcClientNumber returns the old "tmc_client_number" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldTmcClientNumber(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTmcClientNumber is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTmcClientNumber requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTmcClientNumber: %w", err)
	}
	return oldValue.TmcClientNumber, nil
}

// ResetTmcClientNumber resets all changes to the "tmc_client_number" field.
func (m *CustomerMutation) ResetTmcClientNumber() {
	m.tmc_client_number = nil
}

// SetTag sets the "Tag" field.
func (m *CustomerMutation) SetTag(c customer.Tag) {
	m._Tag = &c
}

// Tag returns the value of the "Tag" field in the mutation.
func (m *CustomerMutation) Tag() (r customer.Tag, exists bool) {
	v := m._Tag
	if v == nil {
		return
	}
	return *v, true
}

// OldTag returns the old "Tag" field's value of the Customer entity.
// If the Customer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CustomerMutation) OldTag(ctx context.Context) (v customer.Tag, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTag is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTag requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTag: %w", err)
	}
	return oldValue.Tag, nil
}

// ResetTag resets all changes to the "Tag" field.
func (m *CustomerMutation) ResetTag() {
	m._Tag = nil
}

// Where appends a list predicates to the CustomerMutation builder.
func (m *CustomerMutation) Where(ps ...predicate.Customer) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CustomerMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CustomerMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Customer, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CustomerMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CustomerMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Customer).
func (m *CustomerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CustomerMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.customer_name != nil {
		fields = append(fields, customer.FieldCustomerName)
	}
	if m.account_number != nil {
		fields = append(fields, customer.FieldAccountNumber)
	}
	if m.id_currency != nil {
		fields = append(fields, customer.FieldIDCurrency)
	}
	if m.id_country != nil {
		fields = append(fields, customer.FieldIDCountry)
	}
	if m.alias != nil {
		fields = append(fields, customer.FieldAlias)
	}
	if m.ab_key != nil {
		fields = append(fields, customer.FieldAbKey)
	}
	if m.tmc_client_number != nil {
		fields = append(fields, customer.FieldTmcClientNumber)
	}
	if m._Tag != nil {
		fields = append(fields, customer.FieldTag)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CustomerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case customer.FieldCustomerName:
		return m.CustomerName()
	case customer.FieldAccountNumber:
		return m.AccountNumber()
	case customer.FieldIDCurrency:
		return m.IDCurrency()
	case customer.FieldIDCountry:
		return m.IDCountry()
	case customer.FieldAlias:
		return m.Alias()
	case customer.FieldAbKey:
		return m.AbKey()
	case customer.FieldTmcClientNumber:
		return m.TmcClientNumber()
	case customer.FieldTag:
		return m.Tag()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CustomerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case customer.FieldCustomerName:
		return m.OldCustomerName(ctx)
	case customer.FieldAccountNumber:
		return m.OldAccountNumber(ctx)
	case customer.FieldIDCurrency:
		return m.OldIDCurrency(ctx)
	case customer.FieldIDCountry:
		return m.OldIDCountry(ctx)
	case customer.FieldAlias:
		return m.OldAlias(ctx)
	case customer.FieldAbKey:
		return m.OldAbKey(ctx)
	case customer.FieldTmcClientNumber:
		return m.OldTmcClientNumber(ctx)
	case customer.FieldTag:
		return m.OldTag(ctx)
	}
	return nil, fmt.Errorf("unknown Customer field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CustomerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case customer.FieldCustomerName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCustomerName(v)
		return nil
	case customer.FieldAccountNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccountNumber(v)
		return nil
	case customer.FieldIDCurrency:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIDCurrency(v)
		return nil
	case customer.FieldIDCountry:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIDCountry(v)
		return nil
	case customer.FieldAlias:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAlias(v)
		return nil
	case customer.FieldAbKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAbKey(v)
		return nil
	case customer.FieldTmcClientNumber:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTmcClientNumber(v)
		return nil
	case customer.FieldTag:
		v, ok := value.(customer.Tag)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTag(v)
		return nil
	}
	return fmt.Errorf("unknown Customer field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CustomerMutation) AddedFields() []string {
	var fields []string
	if m.addid_currency != nil {
		fields = append(fields, customer.FieldIDCurrency)
	}
	if m.addid_country != nil {
		fields = append(fields, customer.FieldIDCountry)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CustomerMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case customer.FieldIDCurrency:
		return m.AddedIDCurrency()
	case customer.FieldIDCountry:
		return m.AddedIDCountry()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CustomerMutation) AddField(name string, value ent.Value) error {
	switch name {
	case customer.FieldIDCurrency:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIDCurrency(v)
		return nil
	case customer.FieldIDCountry:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIDCountry(v)
		return nil
	}
	return fmt.Errorf("unknown Customer numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CustomerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CustomerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CustomerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Customer nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CustomerMutation) ResetField(name string) error {
	switch name {
	case customer.FieldCustomerName:
		m.ResetCustomerName()
		return nil
	case customer.FieldAccountNumber:
		m.ResetAccountNumber()
		return nil
	case customer.FieldIDCurrency:
		m.ResetIDCurrency()
		return nil
	case customer.FieldIDCountry:
		m.ResetIDCountry()
		return nil
	case customer.FieldAlias:
		m.ResetAlias()
		return nil
	case customer.FieldAbKey:
		m.ResetAbKey()
		return nil
	case customer.FieldTmcClientNumber:
		m.ResetTmcClientNumber()
		return nil
	case customer.FieldTag:
		m.ResetTag()
		return nil
	}
	return fmt.Errorf("unknown Customer field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CustomerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CustomerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CustomerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CustomerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CustomerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CustomerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CustomerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Customer unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CustomerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Customer edge %s", name)
}
