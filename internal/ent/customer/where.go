// Code generated by ent, DO NOT EDIT.

package customer

import (
	"entgo.io/ent/dialect/sql"
	"github.com/naelcodes/ab-backend/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldID, id))
}

// CustomerName applies equality check predicate on the "customer_name" field. It's identical to CustomerNameEQ.
func CustomerName(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCustomerName, v))
}

// AccountNumber applies equality check predicate on the "account_number" field. It's identical to AccountNumberEQ.
func AccountNumber(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAccountNumber, v))
}

// IDCurrency applies equality check predicate on the "id_currency" field. It's identical to IDCurrencyEQ.
func IDCurrency(v int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldIDCurrency, v))
}

// IDCountry applies equality check predicate on the "id_country" field. It's identical to IDCountryEQ.
func IDCountry(v int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldIDCountry, v))
}

// Alias applies equality check predicate on the "alias" field. It's identical to AliasEQ.
func Alias(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAlias, v))
}

// AbKey applies equality check predicate on the "ab_key" field. It's identical to AbKeyEQ.
func AbKey(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAbKey, v))
}

// TmcClientNumber applies equality check predicate on the "tmc_client_number" field. It's identical to TmcClientNumberEQ.
func TmcClientNumber(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldTmcClientNumber, v))
}

// CustomerNameEQ applies the EQ predicate on the "customer_name" field.
func CustomerNameEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCustomerName, v))
}

// CustomerNameNEQ applies the NEQ predicate on the "customer_name" field.
func CustomerNameNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCustomerName, v))
}

// CustomerNameIn applies the In predicate on the "customer_name" field.
func CustomerNameIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCustomerName, vs...))
}

// CustomerNameNotIn applies the NotIn predicate on the "customer_name" field.
func CustomerNameNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCustomerName, vs...))
}

// CustomerNameGT applies the GT predicate on the "customer_name" field.
func CustomerNameGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCustomerName, v))
}

// CustomerNameGTE applies the GTE predicate on the "customer_name" field.
func CustomerNameGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCustomerName, v))
}

// CustomerNameLT applies the LT predicate on the "customer_name" field.
func CustomerNameLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCustomerName, v))
}

// CustomerNameLTE applies the LTE predicate on the "customer_name" field.
func CustomerNameLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCustomerName, v))
}

// CustomerNameContains applies the Contains predicate on the "customer_name" field.
func CustomerNameContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldCustomerName, v))
}

// CustomerNameHasPrefix applies the HasPrefix predicate on the "customer_name" field.
func CustomerNameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldCustomerName, v))
}

// CustomerNameHasSuffix applies the HasSuffix predicate on the "customer_name" field.
func CustomerNameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldCustomerName, v))
}

// CustomerNameEqualFold applies the EqualFold predicate on the "customer_name" field.
func CustomerNameEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldCustomerName, v))
}

// CustomerNameContainsFold applies the ContainsFold predicate on the "customer_name" field.
func CustomerNameContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldCustomerName, v))
}

// AccountNumberEQ applies the EQ predicate on the "account_number" field.
func AccountNumberEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAccountNumber, v))
}

// AccountNumberNEQ applies the NEQ predicate on the "account_number" field.
func AccountNumberNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAccountNumber, v))
}

// AccountNumberIn applies the In predicate on the "account_number" field.
func AccountNumberIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldAccountNumber, vs...))
}

// AccountNumberNotIn applies the NotIn predicate on the "account_number" field.
func AccountNumberNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldAccountNumber, vs...))
}

// AccountNumberGT applies the GT predicate on the "account_number" field.
func AccountNumberGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldAccountNumber, v))
}

// AccountNumberGTE applies the GTE predicate on the "account_number" field.
func AccountNumberGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldAccountNumber, v))
}

// AccountNumberLT applies the LT predicate on the "account_number" field.
func AccountNumberLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldAccountNumber, v))
}

// AccountNumberLTE applies the LTE predicate on the "account_number" field.
func AccountNumberLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldAccountNumber, v))
}

// AccountNumberContains applies the Contains predicate on the "account_number" field.
func AccountNumberContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldAccountNumber, v))
}

// AccountNumberHasPrefix applies the HasPrefix predicate on the "account_number" field.
func AccountNumberHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldAccountNumber, v))
}

// AccountNumberHasSuffix applies the HasSuffix predicate on the "account_number" field.
func AccountNumberHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldAccountNumber, v))
}

// AccountNumberEqualFold applies the EqualFold predicate on the "account_number" field.
func AccountNumberEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldAccountNumber, v))
}

// AccountNumberContainsFold applies the ContainsFold predicate on the "account_number" field.
func AccountNumberContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldAccountNumber, v))
}

// IDCurrencyEQ applies the EQ predicate on the "id_currency" field.
func IDCurrencyEQ(v int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldIDCurrency, v))
}

// IDCurrencyNEQ applies the NEQ predicate on the "id_currency" field.
func IDCurrencyNEQ(v int) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldIDCurrency, v))
}

// IDCurrencyIn applies the In predicate on the "id_currency" field.
func IDCurrencyIn(vs ...int) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldIDCurrency, vs...))
}

// IDCurrencyNotIn applies the NotIn predicate on the "id_currency" field.
func IDCurrencyNotIn(vs ...int) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldIDCurrency, vs...))
}

// IDCurrencyGT applies the GT predicate on the "id_currency" field.
func IDCurrencyGT(v int) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldIDCurrency, v))
}

// IDCurrencyGTE applies the GTE predicate on the "id_currency" field.
func IDCurrencyGTE(v int) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldIDCurrency, v))
}

// IDCurrencyLT applies the LT predicate on the "id_currency" field.
func IDCurrencyLT(v int) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldIDCurrency, v))
}

// IDCurrencyLTE applies the LTE predicate on the "id_currency" field.
func IDCurrencyLTE(v int) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldIDCurrency, v))
}

// IDCountryEQ applies the EQ predicate on the "id_country" field.
func IDCountryEQ(v int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldIDCountry, v))
}

// IDCountryNEQ applies the NEQ predicate on the "id_country" field.
func IDCountryNEQ(v int) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldIDCountry, v))
}

// IDCountryIn applies the In predicate on the "id_country" field.
func IDCountryIn(vs ...int) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldIDCountry, vs...))
}

// IDCountryNotIn applies the NotIn predicate on the "id_country" field.
func IDCountryNotIn(vs ...int) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldIDCountry, vs...))
}

// IDCountryGT applies the GT predicate on the "id_country" field.
func IDCountryGT(v int) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldIDCountry, v))
}

// IDCountryGTE applies the GTE predicate on the "id_country" field.
func IDCountryGTE(v int) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldIDCountry, v))
}

// IDCountryLT applies the LT predicate on the "id_country" field.
func IDCountryLT(v int) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldIDCountry, v))
}

// IDCountryLTE applies the LTE predicate on the "id_country" field.
func IDCountryLTE(v int) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldIDCountry, v))
}

// AliasEQ applies the EQ predicate on the "alias" field.
func AliasEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAlias, v))
}

// AliasNEQ applies the NEQ predicate on the "alias" field.
func AliasNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAlias, v))
}

// AliasIn applies the In predicate on the "alias" field.
func AliasIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldAlias, vs...))
}

// AliasNotIn applies the NotIn predicate on the "alias" field.
func AliasNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldAlias, vs...))
}

// AliasGT applies the GT predicate on the "alias" field.
func AliasGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldAlias, v))
}

// AliasGTE applies the GTE predicate on the "alias" field.
func AliasGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldAlias, v))
}

// AliasLT applies the LT predicate on the "alias" field.
func AliasLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldAlias, v))
}

// AliasLTE applies the LTE predicate on the "alias" field.
func AliasLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldAlias, v))
}

// AliasContains applies the Contains predicate on the "alias" field.
func AliasContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldAlias, v))
}

// AliasHasPrefix applies the HasPrefix predicate on the "alias" field.
func AliasHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldAlias, v))
}

// AliasHasSuffix applies the HasSuffix predicate on the "alias" field.
func AliasHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldAlias, v))
}

// AliasEqualFold applies the EqualFold predicate on the "alias" field.
func AliasEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldAlias, v))
}

// AliasContainsFold applies the ContainsFold predicate on the "alias" field.
func AliasContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldAlias, v))
}

// AbKeyEQ applies the EQ predicate on the "ab_key" field.
func AbKeyEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAbKey, v))
}

// AbKeyNEQ applies the NEQ predicate on the "ab_key" field.
func AbKeyNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAbKey, v))
}

// AbKeyIn applies the In predicate on the "ab_key" field.
func AbKeyIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldAbKey, vs...))
}

// AbKeyNotIn applies the NotIn predicate on the "ab_key" field.
func AbKeyNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldAbKey, vs...))
}

// AbKeyGT applies the GT predicate on the "ab_key" field.
func AbKeyGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldAbKey, v))
}

// AbKeyGTE applies the GTE predicate on the "ab_key" field.
func AbKeyGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldAbKey, v))
}

// AbKeyLT applies the LT predicate on the "ab_key" field.
func AbKeyLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldAbKey, v))
}

// AbKeyLTE applies the LTE predicate on the "ab_key" field.
func AbKeyLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldAbKey, v))
}

// AbKeyContains applies the Contains predicate on the "ab_key" field.
func AbKeyContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldAbKey, v))
}

// AbKeyHasPrefix applies the HasPrefix predicate on the "ab_key" field.
func AbKeyHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldAbKey, v))
}

// AbKeyHasSuffix applies the HasSuffix predicate on the "ab_key" field.
func AbKeyHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldAbKey, v))
}

// AbKeyEqualFold applies the EqualFold predicate on the "ab_key" field.
func AbKeyEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldAbKey, v))
}

// AbKeyContainsFold applies the ContainsFold predicate on the "ab_key" field.
func AbKeyContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldAbKey, v))
}

// TmcClientNumberEQ applies the EQ predicate on the "tmc_client_number" field.
func TmcClientNumberEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldTmcClientNumber, v))
}

// TmcClientNumberNEQ applies the NEQ predicate on the "tmc_client_number" field.
func TmcClientNumberNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldTmcClientNumber, v))
}

// TmcClientNumberIn applies the In predicate on the "tmc_client_number" field.
func TmcClientNumberIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldTmcClientNumber, vs...))
}

// TmcClientNumberNotIn applies the NotIn predicate on the "tmc_client_number" field.
func TmcClientNumberNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldTmcClientNumber, vs...))
}

// TmcClientNumberGT applies the GT predicate on the "tmc_client_number" field.
func TmcClientNumberGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldTmcClientNumber, v))
}

// TmcClientNumberGTE applies the GTE predicate on the "tmc_client_number" field.
func TmcClientNumberGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldTmcClientNumber, v))
}

// TmcClientNumberLT applies the LT predicate on the "tmc_client_number" field.
func TmcClientNumberLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldTmcClientNumber, v))
}

// TmcClientNumberLTE applies the LTE predicate on the "tmc_client_number" field.
func TmcClientNumberLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldTmcClientNumber, v))
}

// TmcClientNumberContains applies the Contains predicate on the "tmc_client_number" field.
func TmcClientNumberContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldTmcClientNumber, v))
}

// TmcClientNumberHasPrefix applies the HasPrefix predicate on the "tmc_client_number" field.
func TmcClientNumberHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldTmcClientNumber, v))
}

// TmcClientNumberHasSuffix applies the HasSuffix predicate on the "tmc_client_number" field.
func TmcClientNumberHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldTmcClientNumber, v))
}

// TmcClientNumberEqualFold applies the EqualFold predicate on the "tmc_client_number" field.
func TmcClientNumberEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldTmcClientNumber, v))
}

// TmcClientNumberContainsFold applies the ContainsFold predicate on the "tmc_client_number" field.
func TmcClientNumberContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldTmcClientNumber, v))
}

// TagEQ applies the EQ predicate on the "Tag" field.
func TagEQ(v Tag) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldTag, v))
}

// TagNEQ applies the NEQ predicate on the "Tag" field.
func TagNEQ(v Tag) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldTag, v))
}

// TagIn applies the In predicate on the "Tag" field.
func TagIn(vs ...Tag) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldTag, vs...))
}

// TagNotIn applies the NotIn predicate on the "Tag" field.
func TagNotIn(vs ...Tag) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldTag, vs...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.NotPredicates(p))
}
