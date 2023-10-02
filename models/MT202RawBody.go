package models

import (
	"fmt"
	"newgens/src"
	"strings"
	"time"
	"unicode"

	"github.com/shopspring/decimal"
)

type MT202RawBody struct {
	TransactionReferenceNumber   string   `json:"transaction_reference_number"`    //20
	RelatedReference             string   `json:"related_reference"`               //21
	ValueDateCurrencyAmount      string   `json:"value_date_currency_amount"`      //32A
	OrderingInstitutionOption    rune     `json:"ordering_institution_option"`     //52a
	OrderingInstitution          []string `json:"ordering_institution"`            //52a
	AccountWithInstitutionOption rune     `json:"account_with_institution_option"` //57a
	AccountWithInstitution       []string `json:"account_with_institution"`        //57a
	BeneficiaryInstitutionOption rune     `json:"beneficiary_institution_option"`  //58a
	BeneficiaryInstitution       []string `json:"beneficiary_institution"`         //58a
}

func (r *MT202RawBody) Validate() error {
	// --------------------------------- transaction reference number ---------------------------------
	if len(r.TransactionReferenceNumber) < 1 || len(r.TransactionReferenceNumber) > 16 {
		return fmt.Errorf("tag 20 length should be 1-16 characters")
	}

	// --------------------------------- related reference ---------------------------------
	if len(r.RelatedReference) < 1 || len(r.RelatedReference) > 16 {
		return fmt.Errorf("tag 21 length should be 1-16 characters")
	}

	// --------------------------------- [value date, currency code, amount] ---------------------------------
	if len(r.ValueDateCurrencyAmount) < 10 || len(r.ValueDateCurrencyAmount) > 24 {
		return fmt.Errorf("tag 32A length should be 10-24 characters")
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

	// --------------------------------- ordering institution ---------------------------------
	if r.OrderingInstitutionOption == 'A' {
		if len(r.OrderingInstitution) < 1 || len(r.OrderingInstitution) > 2 {
			return fmt.Errorf("tag 52A should have 1-2 lines")
		}
		if err := r.checkPartyIdentifier(r.OrderingInstitution[0]); err != nil {
			return fmt.Errorf("tag 52A %s", err)
		}
		if err := r.checkIdentifierCode(r.OrderingInstitution[1]); err != nil {
			return fmt.Errorf("tag 52A %s", err)
		}
	} else if r.OrderingInstitutionOption == 'D' {
		if len(r.OrderingInstitution) < 1 || len(r.OrderingInstitution) > 5 {
			return fmt.Errorf("tag 52D should have 1-5 lines")
		}
		if err := r.checkPartyIdentifier(r.OrderingInstitution[0]); err != nil {
			return fmt.Errorf("tag 52D %s", err)
		}
		for i := 1; i < len(r.OrderingInstitution); i++ {
			if err := r.checkNameAndAddress(r.OrderingInstitution[i]); err != nil {
				return fmt.Errorf("tag 52D %s", err)
			}
		}
	} else {
		return fmt.Errorf("tag 52 option should be A or D")
	}

	// --------------------------------- acccount with institution ---------------------------------
	if r.AccountWithInstitutionOption == 'A' {
		if len(r.AccountWithInstitution) < 1 || len(r.AccountWithInstitution) > 2 {
			return fmt.Errorf("tag 57A should have 1-2 lines")
		}
		if err := r.checkPartyIdentifier(r.AccountWithInstitution[0]); err != nil {
			return fmt.Errorf("tag 57A %s", err)
		}
		if err := r.checkIdentifierCode(r.AccountWithInstitution[1]); err != nil {
			return fmt.Errorf("tag 57A %s", err)
		}
	} else if r.AccountWithInstitutionOption == 'B' {
		if len(r.AccountWithInstitution) < 1 || len(r.AccountWithInstitution) > 2 {
			return fmt.Errorf("tag 57B should have 1-2 lines")
		}
		if err := r.checkPartyIdentifier(r.AccountWithInstitution[0]); err != nil {
			return fmt.Errorf("tag 57B %s", err)
		}
		if err := r.checkNameAndAddress(r.AccountWithInstitution[1]); err != nil {
			return fmt.Errorf("tag 57B %s", err)
		}
	} else if r.AccountWithInstitutionOption == 'D' {
		if len(r.AccountWithInstitution) < 1 || len(r.AccountWithInstitution) > 5 {
			return fmt.Errorf("tag 57D should have 1-5 lines")
		}
		if err := r.checkPartyIdentifier(r.AccountWithInstitution[0]); err != nil {
			return fmt.Errorf("tag 57D %s", err)
		}
		for i := 1; i < len(r.AccountWithInstitution); i++ {
			if err := r.checkNameAndAddress(r.AccountWithInstitution[i]); err != nil {
				return fmt.Errorf("tag 57D %s", err)
			}
		}
	} else {
		return fmt.Errorf("tag 57 option should be A or B or D")
	}

	// --------------------------------- benificiary institution ---------------------------------
	if r.BeneficiaryInstitutionOption == 'A' {
		if len(r.BeneficiaryInstitution) < 1 || len(r.BeneficiaryInstitution) > 2 {
			return fmt.Errorf("tag 58A should have 1-2 lines")
		}
		if err := r.checkPartyIdentifier(r.BeneficiaryInstitution[0]); err != nil {
			return fmt.Errorf("tag 58A %s", err)
		}
		if err := r.checkIdentifierCode(r.BeneficiaryInstitution[1]); err != nil {
			return fmt.Errorf("tag 58A %s", err)
		}
	} else if r.BeneficiaryInstitutionOption == 'D' {
		if len(r.BeneficiaryInstitution) < 1 || len(r.BeneficiaryInstitution) > 5 {
			return fmt.Errorf("tag 58D should have 1-5 lines")
		}
		if err := r.checkPartyIdentifier(r.BeneficiaryInstitution[0]); err != nil {
			return fmt.Errorf("tag 58D %s", err)
		}
		for i := 1; i < len(r.BeneficiaryInstitution); i++ {
			if err := r.checkNameAndAddress(r.BeneficiaryInstitution[i]); err != nil {
				return fmt.Errorf("tag 58D %s", err)
			}
		}
	} else {
		return fmt.Errorf("tag 58 option should be A or D")
	}

	return nil
}
func (r *MT202RawBody) checkPartyIdentifier(value string) error {
	if len(value) < 2 || len(value) > 37 {
		return fmt.Errorf("party identifier should have 2-37 characters")
	}
	if value[0] != '/' {
		return fmt.Errorf("party identifier should start with '/'")
	}

	splitted := strings.Split(value, "/")
	if len(splitted) > 3 {
		return fmt.Errorf("party identifier should only have 2 parts")
	}

	var firstpart string
	var secondpart string
	if len(splitted) == 3 {
		firstpart = splitted[1]
		secondpart = splitted[2]
	} else {
		secondpart = splitted[1]
	}

	if firstpart != "" {
		if len(firstpart) != 1 {
			return fmt.Errorf("party identifier first part should have 1 alphabet")
		}
		if !unicode.IsLetter([]rune(firstpart)[0]) {
			return fmt.Errorf("party identifier first part should only be alphabetic")
		}
	}
	if len(secondpart) < 1 || len(secondpart) > 34 {
		return fmt.Errorf("party identifier second part should have 1-34 characters")
	}

	return nil
}

func (r *MT202RawBody) checkIdentifierCode(value string) error {
	if len(value) < 8 || len(value) > 11 {
		return fmt.Errorf("identifier code should have 8-11 characters")
	}
	for i := 0; i < 6; i++ {
		if !unicode.IsLetter([]rune(value)[i]) {
			return fmt.Errorf("identifier code - index %d should be alphabetic", i)
		}
	}
	return nil
}
func (r *MT202RawBody) checkNameAndAddress(value string) error {
	if len(value) < 1 || len(value) > 35 {
		return fmt.Errorf("name and address should have 1-35 characters")
	}
	return nil
}
