// Code generated by ent, DO NOT EDIT.

package imputation

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the imputation type in the database.
	Label = "imputation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmountApply holds the string denoting the amount_apply field in the database.
	FieldAmountApply = "amount_apply"
	// FieldPaymentAmount holds the string denoting the payment_amount field in the database.
	FieldPaymentAmount = "payment_amount"
	// FieldInvoiceAmount holds the string denoting the invoice_amount field in the database.
	FieldInvoiceAmount = "invoice_amount"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// EdgeInvoice holds the string denoting the invoice edge name in mutations.
	EdgeInvoice = "invoice"
	// EdgePayment holds the string denoting the payment edge name in mutations.
	EdgePayment = "payment"
	// Table holds the table name of the imputation in the database.
	Table = "invoice_payment_received"
	// InvoiceTable is the table that holds the invoice relation/edge.
	InvoiceTable = "invoice_payment_received"
	// InvoiceInverseTable is the table name for the Invoice entity.
	// It exists in this package in order to avoid circular dependency with the "invoice" package.
	InvoiceInverseTable = "invoice"
	// InvoiceColumn is the table column denoting the invoice relation/edge.
	InvoiceColumn = "id_invoice"
	// PaymentTable is the table that holds the payment relation/edge.
	PaymentTable = "invoice_payment_received"
	// PaymentInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	PaymentInverseTable = "payment_received"
	// PaymentColumn is the table column denoting the payment relation/edge.
	PaymentColumn = "id_payment_received"
)

// Columns holds all SQL columns for imputation fields.
var Columns = []string{
	FieldID,
	FieldAmountApply,
	FieldPaymentAmount,
	FieldInvoiceAmount,
	FieldTag,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "invoice_payment_received"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"id_invoice",
	"id_payment_received",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAmountApply holds the default value on creation for the "amount_apply" field.
	DefaultAmountApply float64
	// AmountApplyValidator is a validator for the "amount_apply" field. It is called by the builders before save.
	AmountApplyValidator func(float64) error
	// DefaultPaymentAmount holds the default value on creation for the "payment_amount" field.
	DefaultPaymentAmount float64
	// PaymentAmountValidator is a validator for the "payment_amount" field. It is called by the builders before save.
	PaymentAmountValidator func(float64) error
	// DefaultInvoiceAmount holds the default value on creation for the "invoice_amount" field.
	DefaultInvoiceAmount float64
	// InvoiceAmountValidator is a validator for the "invoice_amount" field. It is called by the builders before save.
	InvoiceAmountValidator func(float64) error
)

// Tag defines the type for the "tag" enum field.
type Tag string

// Tag3 is the default value of the Tag enum.
const DefaultTag = Tag3

// Tag values.
const (
	Tag1 Tag = "1"
	Tag2 Tag = "2"
	Tag3 Tag = "3"
)

func (t Tag) String() string {
	return string(t)
}

// TagValidator is a validator for the "tag" field enum values. It is called by the builders before save.
func TagValidator(t Tag) error {
	switch t {
	case Tag1, Tag2, Tag3:
		return nil
	default:
		return fmt.Errorf("imputation: invalid enum value for tag field: %q", t)
	}
}

// OrderOption defines the ordering options for the Imputation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAmountApply orders the results by the amount_apply field.
func ByAmountApply(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountApply, opts...).ToFunc()
}

// ByPaymentAmount orders the results by the payment_amount field.
func ByPaymentAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentAmount, opts...).ToFunc()
}

// ByInvoiceAmount orders the results by the invoice_amount field.
func ByInvoiceAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceAmount, opts...).ToFunc()
}

// ByTag orders the results by the tag field.
func ByTag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTag, opts...).ToFunc()
}

// ByInvoiceField orders the results by invoice field.
func ByInvoiceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvoiceStep(), sql.OrderByField(field, opts...))
	}
}

// ByPaymentField orders the results by payment field.
func ByPaymentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPaymentStep(), sql.OrderByField(field, opts...))
	}
}
func newInvoiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvoiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InvoiceTable, InvoiceColumn),
	)
}
func newPaymentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PaymentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PaymentTable, PaymentColumn),
	)
}