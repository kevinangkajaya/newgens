package models

import (
	"fmt"
	"newgens/src"
	"time"
	"unicode"

	"github.com/shopspring/decimal"
)

type MT202RawBody struct {
	TransactionReferenceNumber string   `json:"transaction_reference_number"` //20
	RelatedReference           string   `json:"related_reference"`            //21
	ValueDateCurrencyAmount    string   `json:"value_date_currency_amount"`   //32A
	OrderingInstitution        []string `json:"ordering_institution"`         //52a
	AccountWithInstitution     []string `json:"account_with_institution"`     //57a
	BeneficiaryInstitution     []string `json:"beneficiary_institution"`      //58a
}

func (r *MT202RawBody) Validate() error {
	// --------------------------------- transaction reference number ---------------------------------
	if len(r.TransactionReferenceNumber) < 1 || len(r.TransactionReferenceNumber) > 16 {
		return fmt.Errorf("transaction reference number length should be 1-16 characters")
	}

	// --------------------------------- related reference ---------------------------------
	if len(r.RelatedReference) < 1 || len(r.RelatedReference) > 16 {
		return fmt.Errorf("related reference length should be 1-16 characters")
	}

	// --------------------------------- [value date, currency code, amount] ---------------------------------
	if len(r.ValueDateCurrencyAmount) < 10 {
		return fmt.Errorf("[value date, currency code, amount] length should be at least 10 characters")
	}
	_, err := time.Parse("060102", r.ValueDateCurrencyAmount[0:6])
	if err != nil {
		return err
	}
	for _, v := range r.ValueDateCurrencyAmount[6:9] {
		if !unicode.IsLetter(v) {
			return fmt.Errorf("currency code should be alphabetic")
		}
	}
	_, err = decimal.NewFromString(src.ReplaceCommaWithDot(r.ValueDateCurrencyAmount[9:]))
	if err != nil {
		return err
	}

	return nil
}
