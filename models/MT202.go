package models

import (
	"errors"
	"fmt"
	"time"
	"unicode"

	"github.com/shopspring/decimal"
)

type MT202SwiftAddress string
type MT202 struct {
	ID             int               `gorm:"primarykey" json:"id" db:"ID"`
	SenderBIC      MT202SwiftAddress `json:"sender_bic" db:"SenderBIC"`
	ReceiverBIC    MT202SwiftAddress `json:"receiver_bic" db:"ReceiverBIC"`
	Direction      string            `json:"direction" db:"Direction"`
	MTType         string            `json:"mt_type" db:"MTType"`
	UETR           string            `json:"uetr" db:"UETR"`
	F20            string            `json:"f20" db:"F20"`
	F21            string            `json:"f21" db:"F21"`
	F32A_ValueDate time.Time         `json:"f32a_value_date" db:"F32A_ValueDate"`
	F32A_Currency  string            `json:"f32a_currency" db:"F32A_Currency"`
	F32A_Amount    decimal.Decimal   `json:"f32a_amount" db:"F32A_Amount"`
	F52a           string            `json:"f52a" db:"F52a"`
	F57a           string            `json:"f57a" db:"F57a"`
	F58a           string            `json:"f58a" db:"F58a"`
	RawData        string            `json:"raw_data" db:"RawData"`
	CreatedAt      time.Time         `json:"created_at" db:"created_at"`
}

func (r *MT202) Validate() error {
	if err := r.SenderBIC.Validate("SenderBIC"); err != nil {
		return err
	}
	if err := r.ReceiverBIC.Validate("ReceiverBIC"); err != nil {
		return err
	}
	if r.Direction != "O" && r.Direction != "I" {
		return errors.New("direction should be 'O' or 'I'")
	}
	if len(r.MTType) != 3 {
		return errors.New("MTType should have 3 characters")
	}
	if len(r.UETR) != 36 {
		return errors.New("UETR should have 36 characters")
	}

	if len(r.F20) < 1 || len(r.F20) > 16 {
		return errors.New("F20 should have 1-16 characters")
	}
	if len(r.F21) < 1 || len(r.F21) > 16 {
		return errors.New("F20 should have 1-16 characters")
	}
	if r.F32A_ValueDate.IsZero() {
		return errors.New("wrong F32_ValueDate")
	}
	if len(r.F32A_Currency) != 3 {
		return errors.New("F32A_Currency should be 3 characters")
	}
	if len(r.F58a) < 1 {
		return errors.New("F58a should have more than 0 characters")
	}
	if len(r.RawData) < 1 {
		return errors.New("raw Data should have more than 0 characters")
	}
	return nil
}

func (address MT202SwiftAddress) Validate(name string) error {
	if len(address) != 12 {
		return fmt.Errorf("%s should have 12 characters", name)
	} else if !unicode.IsLetter([]rune(address)[8]) {
		return fmt.Errorf("8th character of %s can only be alphabetic", name)
	}

	return nil
}
