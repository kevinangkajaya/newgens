package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type MT202 struct {
	ID             int             `gorm:"primarykey" json:"id"`
	SenderBIC      string          `json:"sender_bic"`
	ReceiverBIC    string          `json:"receiver_bic"`
	Direction      string          `json:"direction"`
	MTType         string          `json:"mt_type"`
	UETR           string          `json:"uetr"`
	F20            string          `json:"f20"`
	F21            string          `json:"f21"`
	F32A_ValueDate time.Time       `json:"f32a_value_date"`
	F32A_Currency  string          `json:"f32a_currency"`
	F32A_Amount    decimal.Decimal `json:"f32a_amount"`
	F52a           string          `json:"f52a"`
	F57a           string          `json:"f57a"`
	F58a           string          `json:"f58a"`
	RawData        string          `json:"raw_data"`
	CreatedAt      time.Time       `json:"created_at"`
}

func (r *MT202) Validate() error {
	return nil
}
