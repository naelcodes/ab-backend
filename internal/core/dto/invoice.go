package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type TravelItemDTO struct {
	Id           uint    `json:"id"`
	Itinerary    *string `json:"itinerary,omitempty"`
	TravelerName *string `json:"travelerName,omitempty"`
	TicketNumber *string `json:"ticketNumber,omitempty"`
	TotalPrice   float64 `json:"totalPrice"`
}

func (t TravelItemDTO) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Id, validation.Required),
		validation.Field(&t.TotalPrice, validation.Required),
	)
}

type GetInvoiceDTO struct {
	Id            int             `json:"id"`
	InvoiceNumber string          `json:"invoiceNumber"`
	CreationDate  string          `json:"creationDate"`
	DueDate       string          `json:"dueDate"`
	Amount        float64         `json:"amount"`
	Balance       float64         `json:"balance"`
	Credit_apply  float64         `json:"credit_apply"`
	Status        string          `json:"status"`
	IdCustomer    *int            `json:"idCustomer,omitempty"`
	Customer      *GetCustomerDTO `json:"customer,omitempty"`
	TravelItems   []TravelItemDTO `json:"travelItems"`
}

type CreateInvoiceDTO struct {
	IdCustomer   int             `json:"idCustomer"`
	CreationDate string          `json:"creationDate"`
	DueDate      string          `json:"dueDate"`
	TravelItems  []TravelItemDTO `json:"travelItems"`
}

func (c CreateInvoiceDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.IdCustomer, validation.Required),
		validation.Field(&c.CreationDate, validation.Required),
		validation.Field(&c.DueDate, validation.Required),
		validation.Field(&c.TravelItems, validation.Required),
	)
}

type UpdateInvoiceDTO struct {
	Id           int              `json:"id"`
	IdCustomer   *int             `json:"idCustomer,omitempty"`
	CreationDate *string          `json:"creationDate,omitempty"`
	DueDate      *string          `json:"dueDate,omitempty"`
	TravelItems  *[]TravelItemDTO `json:"travelItems,omitempty"`
}

func (u UpdateInvoiceDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.IdCustomer, validation.NilOrNotEmpty),
		validation.Field(&u.CreationDate, validation.NilOrNotEmpty),
		validation.Field(&u.DueDate, validation.NilOrNotEmpty),
		validation.Field(&u.TravelItems, validation.NilOrNotEmpty),
	)
}

type GetAllInvoiceDTO types.GetAllDTO[[]*GetInvoiceDTO]
type GetAllTravelItemDTO types.GetAllDTO[[]*TravelItemDTO]
