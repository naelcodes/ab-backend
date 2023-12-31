// Code generated by ent, DO NOT EDIT.

package invoice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/naelcodes/ab-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldID, id))
}

// CreationDate applies equality check predicate on the "creation_date" field. It's identical to CreationDateEQ.
func CreationDate(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldCreationDate, v))
}

// InvoiceNumber applies equality check predicate on the "invoice_number" field. It's identical to InvoiceNumberEQ.
func InvoiceNumber(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldInvoiceNumber, v))
}

// DueDate applies equality check predicate on the "due_date" field. It's identical to DueDateEQ.
func DueDate(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldDueDate, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Invoice {
	vc, err := ValueScanner.Amount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldAmount, vc), err)
}

// NetAmount applies equality check predicate on the "net_amount" field. It's identical to NetAmountEQ.
func NetAmount(v float64) predicate.Invoice {
	vc, err := ValueScanner.NetAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldNetAmount, vc), err)
}

// BaseAmount applies equality check predicate on the "base_amount" field. It's identical to BaseAmountEQ.
func BaseAmount(v float64) predicate.Invoice {
	vc, err := ValueScanner.BaseAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldBaseAmount, vc), err)
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v float64) predicate.Invoice {
	vc, err := ValueScanner.Balance.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldBalance, vc), err)
}

// CreditApply applies equality check predicate on the "credit_apply" field. It's identical to CreditApplyEQ.
func CreditApply(v float64) predicate.Invoice {
	vc, err := ValueScanner.CreditApply.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldCreditApply, vc), err)
}

// CreationDateEQ applies the EQ predicate on the "creation_date" field.
func CreationDateEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldCreationDate, v))
}

// CreationDateNEQ applies the NEQ predicate on the "creation_date" field.
func CreationDateNEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldCreationDate, v))
}

// CreationDateIn applies the In predicate on the "creation_date" field.
func CreationDateIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldCreationDate, vs...))
}

// CreationDateNotIn applies the NotIn predicate on the "creation_date" field.
func CreationDateNotIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldCreationDate, vs...))
}

// CreationDateGT applies the GT predicate on the "creation_date" field.
func CreationDateGT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldCreationDate, v))
}

// CreationDateGTE applies the GTE predicate on the "creation_date" field.
func CreationDateGTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldCreationDate, v))
}

// CreationDateLT applies the LT predicate on the "creation_date" field.
func CreationDateLT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldCreationDate, v))
}

// CreationDateLTE applies the LTE predicate on the "creation_date" field.
func CreationDateLTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldCreationDate, v))
}

// CreationDateContains applies the Contains predicate on the "creation_date" field.
func CreationDateContains(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContains(FieldCreationDate, v))
}

// CreationDateHasPrefix applies the HasPrefix predicate on the "creation_date" field.
func CreationDateHasPrefix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasPrefix(FieldCreationDate, v))
}

// CreationDateHasSuffix applies the HasSuffix predicate on the "creation_date" field.
func CreationDateHasSuffix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasSuffix(FieldCreationDate, v))
}

// CreationDateEqualFold applies the EqualFold predicate on the "creation_date" field.
func CreationDateEqualFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEqualFold(FieldCreationDate, v))
}

// CreationDateContainsFold applies the ContainsFold predicate on the "creation_date" field.
func CreationDateContainsFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContainsFold(FieldCreationDate, v))
}

// InvoiceNumberEQ applies the EQ predicate on the "invoice_number" field.
func InvoiceNumberEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldInvoiceNumber, v))
}

// InvoiceNumberNEQ applies the NEQ predicate on the "invoice_number" field.
func InvoiceNumberNEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldInvoiceNumber, v))
}

// InvoiceNumberIn applies the In predicate on the "invoice_number" field.
func InvoiceNumberIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldInvoiceNumber, vs...))
}

// InvoiceNumberNotIn applies the NotIn predicate on the "invoice_number" field.
func InvoiceNumberNotIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldInvoiceNumber, vs...))
}

// InvoiceNumberGT applies the GT predicate on the "invoice_number" field.
func InvoiceNumberGT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldInvoiceNumber, v))
}

// InvoiceNumberGTE applies the GTE predicate on the "invoice_number" field.
func InvoiceNumberGTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldInvoiceNumber, v))
}

// InvoiceNumberLT applies the LT predicate on the "invoice_number" field.
func InvoiceNumberLT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldInvoiceNumber, v))
}

// InvoiceNumberLTE applies the LTE predicate on the "invoice_number" field.
func InvoiceNumberLTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldInvoiceNumber, v))
}

// InvoiceNumberContains applies the Contains predicate on the "invoice_number" field.
func InvoiceNumberContains(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContains(FieldInvoiceNumber, v))
}

// InvoiceNumberHasPrefix applies the HasPrefix predicate on the "invoice_number" field.
func InvoiceNumberHasPrefix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasPrefix(FieldInvoiceNumber, v))
}

// InvoiceNumberHasSuffix applies the HasSuffix predicate on the "invoice_number" field.
func InvoiceNumberHasSuffix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasSuffix(FieldInvoiceNumber, v))
}

// InvoiceNumberEqualFold applies the EqualFold predicate on the "invoice_number" field.
func InvoiceNumberEqualFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEqualFold(FieldInvoiceNumber, v))
}

// InvoiceNumberContainsFold applies the ContainsFold predicate on the "invoice_number" field.
func InvoiceNumberContainsFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContainsFold(FieldInvoiceNumber, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldStatus, vs...))
}

// DueDateEQ applies the EQ predicate on the "due_date" field.
func DueDateEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldDueDate, v))
}

// DueDateNEQ applies the NEQ predicate on the "due_date" field.
func DueDateNEQ(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldDueDate, v))
}

// DueDateIn applies the In predicate on the "due_date" field.
func DueDateIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldDueDate, vs...))
}

// DueDateNotIn applies the NotIn predicate on the "due_date" field.
func DueDateNotIn(vs ...string) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldDueDate, vs...))
}

// DueDateGT applies the GT predicate on the "due_date" field.
func DueDateGT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGT(FieldDueDate, v))
}

// DueDateGTE applies the GTE predicate on the "due_date" field.
func DueDateGTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldGTE(FieldDueDate, v))
}

// DueDateLT applies the LT predicate on the "due_date" field.
func DueDateLT(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLT(FieldDueDate, v))
}

// DueDateLTE applies the LTE predicate on the "due_date" field.
func DueDateLTE(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldLTE(FieldDueDate, v))
}

// DueDateContains applies the Contains predicate on the "due_date" field.
func DueDateContains(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContains(FieldDueDate, v))
}

// DueDateHasPrefix applies the HasPrefix predicate on the "due_date" field.
func DueDateHasPrefix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasPrefix(FieldDueDate, v))
}

// DueDateHasSuffix applies the HasSuffix predicate on the "due_date" field.
func DueDateHasSuffix(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldHasSuffix(FieldDueDate, v))
}

// DueDateEqualFold applies the EqualFold predicate on the "due_date" field.
func DueDateEqualFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldEqualFold(FieldDueDate, v))
}

// DueDateContainsFold applies the ContainsFold predicate on the "due_date" field.
func DueDateContainsFold(v string) predicate.Invoice {
	return predicate.Invoice(sql.FieldContainsFold(FieldDueDate, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.Amount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldAmount, vc), err)
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.Amount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldNEQ(FieldAmount, vc), err)
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.Amount.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldIn(FieldAmount, v...), err)
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.Amount.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldNotIn(FieldAmount, v...), err)
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Invoice {
	vc, err := ValueScanner.Amount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGT(FieldAmount, vc), err)
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.Amount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGTE(FieldAmount, vc), err)
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Invoice {
	vc, err := ValueScanner.Amount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLT(FieldAmount, vc), err)
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.Amount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLTE(FieldAmount, vc), err)
}

// NetAmountEQ applies the EQ predicate on the "net_amount" field.
func NetAmountEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.NetAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldNetAmount, vc), err)
}

// NetAmountNEQ applies the NEQ predicate on the "net_amount" field.
func NetAmountNEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.NetAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldNEQ(FieldNetAmount, vc), err)
}

// NetAmountIn applies the In predicate on the "net_amount" field.
func NetAmountIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.NetAmount.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldIn(FieldNetAmount, v...), err)
}

// NetAmountNotIn applies the NotIn predicate on the "net_amount" field.
func NetAmountNotIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.NetAmount.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldNotIn(FieldNetAmount, v...), err)
}

// NetAmountGT applies the GT predicate on the "net_amount" field.
func NetAmountGT(v float64) predicate.Invoice {
	vc, err := ValueScanner.NetAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGT(FieldNetAmount, vc), err)
}

// NetAmountGTE applies the GTE predicate on the "net_amount" field.
func NetAmountGTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.NetAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGTE(FieldNetAmount, vc), err)
}

// NetAmountLT applies the LT predicate on the "net_amount" field.
func NetAmountLT(v float64) predicate.Invoice {
	vc, err := ValueScanner.NetAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLT(FieldNetAmount, vc), err)
}

// NetAmountLTE applies the LTE predicate on the "net_amount" field.
func NetAmountLTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.NetAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLTE(FieldNetAmount, vc), err)
}

// BaseAmountEQ applies the EQ predicate on the "base_amount" field.
func BaseAmountEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.BaseAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldBaseAmount, vc), err)
}

// BaseAmountNEQ applies the NEQ predicate on the "base_amount" field.
func BaseAmountNEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.BaseAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldNEQ(FieldBaseAmount, vc), err)
}

// BaseAmountIn applies the In predicate on the "base_amount" field.
func BaseAmountIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.BaseAmount.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldIn(FieldBaseAmount, v...), err)
}

// BaseAmountNotIn applies the NotIn predicate on the "base_amount" field.
func BaseAmountNotIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.BaseAmount.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldNotIn(FieldBaseAmount, v...), err)
}

// BaseAmountGT applies the GT predicate on the "base_amount" field.
func BaseAmountGT(v float64) predicate.Invoice {
	vc, err := ValueScanner.BaseAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGT(FieldBaseAmount, vc), err)
}

// BaseAmountGTE applies the GTE predicate on the "base_amount" field.
func BaseAmountGTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.BaseAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGTE(FieldBaseAmount, vc), err)
}

// BaseAmountLT applies the LT predicate on the "base_amount" field.
func BaseAmountLT(v float64) predicate.Invoice {
	vc, err := ValueScanner.BaseAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLT(FieldBaseAmount, vc), err)
}

// BaseAmountLTE applies the LTE predicate on the "base_amount" field.
func BaseAmountLTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.BaseAmount.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLTE(FieldBaseAmount, vc), err)
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.Balance.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldBalance, vc), err)
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.Balance.Value(v)
	return predicate.InvoiceOrErr(sql.FieldNEQ(FieldBalance, vc), err)
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.Balance.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldIn(FieldBalance, v...), err)
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.Balance.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldNotIn(FieldBalance, v...), err)
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v float64) predicate.Invoice {
	vc, err := ValueScanner.Balance.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGT(FieldBalance, vc), err)
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.Balance.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGTE(FieldBalance, vc), err)
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v float64) predicate.Invoice {
	vc, err := ValueScanner.Balance.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLT(FieldBalance, vc), err)
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.Balance.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLTE(FieldBalance, vc), err)
}

// CreditApplyEQ applies the EQ predicate on the "credit_apply" field.
func CreditApplyEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.CreditApply.Value(v)
	return predicate.InvoiceOrErr(sql.FieldEQ(FieldCreditApply, vc), err)
}

// CreditApplyNEQ applies the NEQ predicate on the "credit_apply" field.
func CreditApplyNEQ(v float64) predicate.Invoice {
	vc, err := ValueScanner.CreditApply.Value(v)
	return predicate.InvoiceOrErr(sql.FieldNEQ(FieldCreditApply, vc), err)
}

// CreditApplyIn applies the In predicate on the "credit_apply" field.
func CreditApplyIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.CreditApply.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldIn(FieldCreditApply, v...), err)
}

// CreditApplyNotIn applies the NotIn predicate on the "credit_apply" field.
func CreditApplyNotIn(vs ...float64) predicate.Invoice {
	var (
		err error
		v   = make([]any, len(vs))
	)
	for i := range v {
		if v[i], err = ValueScanner.CreditApply.Value(vs[i]); err != nil {
			break
		}
	}
	return predicate.InvoiceOrErr(sql.FieldNotIn(FieldCreditApply, v...), err)
}

// CreditApplyGT applies the GT predicate on the "credit_apply" field.
func CreditApplyGT(v float64) predicate.Invoice {
	vc, err := ValueScanner.CreditApply.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGT(FieldCreditApply, vc), err)
}

// CreditApplyGTE applies the GTE predicate on the "credit_apply" field.
func CreditApplyGTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.CreditApply.Value(v)
	return predicate.InvoiceOrErr(sql.FieldGTE(FieldCreditApply, vc), err)
}

// CreditApplyLT applies the LT predicate on the "credit_apply" field.
func CreditApplyLT(v float64) predicate.Invoice {
	vc, err := ValueScanner.CreditApply.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLT(FieldCreditApply, vc), err)
}

// CreditApplyLTE applies the LTE predicate on the "credit_apply" field.
func CreditApplyLTE(v float64) predicate.Invoice {
	vc, err := ValueScanner.CreditApply.Value(v)
	return predicate.InvoiceOrErr(sql.FieldLTE(FieldCreditApply, vc), err)
}

// TagEQ applies the EQ predicate on the "tag" field.
func TagEQ(v Tag) predicate.Invoice {
	return predicate.Invoice(sql.FieldEQ(FieldTag, v))
}

// TagNEQ applies the NEQ predicate on the "tag" field.
func TagNEQ(v Tag) predicate.Invoice {
	return predicate.Invoice(sql.FieldNEQ(FieldTag, v))
}

// TagIn applies the In predicate on the "tag" field.
func TagIn(vs ...Tag) predicate.Invoice {
	return predicate.Invoice(sql.FieldIn(FieldTag, vs...))
}

// TagNotIn applies the NotIn predicate on the "tag" field.
func TagNotIn(vs ...Tag) predicate.Invoice {
	return predicate.Invoice(sql.FieldNotIn(FieldTag, vs...))
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := newCustomerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasImputations applies the HasEdge predicate on the "imputations" edge.
func HasImputations() predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImputationsTable, ImputationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImputationsWith applies the HasEdge predicate on the "imputations" edge with a given conditions (other predicates).
func HasImputationsWith(preds ...predicate.Imputation) predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := newImputationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTravelItems applies the HasEdge predicate on the "travel_items" edge.
func HasTravelItems() predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TravelItemsTable, TravelItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTravelItemsWith applies the HasEdge predicate on the "travel_items" edge with a given conditions (other predicates).
func HasTravelItemsWith(preds ...predicate.TravelItem) predicate.Invoice {
	return predicate.Invoice(func(s *sql.Selector) {
		step := newTravelItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Invoice) predicate.Invoice {
	return predicate.Invoice(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Invoice) predicate.Invoice {
	return predicate.Invoice(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Invoice) predicate.Invoice {
	return predicate.Invoice(sql.NotPredicates(p))
}
