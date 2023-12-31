// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Customer is the predicate function for customer builders.
type Customer func(*sql.Selector)

// Imputation is the predicate function for imputation builders.
type Imputation func(*sql.Selector)

// ImputationOrErr calls the predicate only if the error is not nit.
func ImputationOrErr(p Imputation, err error) Imputation {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// Invoice is the predicate function for invoice builders.
type Invoice func(*sql.Selector)

// InvoiceOrErr calls the predicate only if the error is not nit.
func InvoiceOrErr(p Invoice, err error) Invoice {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// Payment is the predicate function for payment builders.
type Payment func(*sql.Selector)

// PaymentOrErr calls the predicate only if the error is not nit.
func PaymentOrErr(p Payment, err error) Payment {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// TravelItem is the predicate function for travelitem builders.
type TravelItem func(*sql.Selector)

// TravelItemOrErr calls the predicate only if the error is not nit.
func TravelItemOrErr(p TravelItem, err error) TravelItem {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}
