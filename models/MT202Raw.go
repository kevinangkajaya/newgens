package models

import (
	"errors"
	"fmt"
	"newgens/src"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type MT202Raw struct {
	BasicHeader       string             `json:"basic_header"`       // 1
	ApplicationHeader string             `json:"application_header"` // 2
	UserHeader        MT202RawUserHeader `json:"user_header"`        // 3
	Body              MT202RawBody       `json:"body"`               // 4
	RawData           string             `json:"raw_data"`
}

func (r *MT202Raw) Validate() error {
	// --------------------------------- basic header ---------------------------------
	if len(r.BasicHeader) != 25 {
		return errors.New("basic header should be 25 characters")
	} else {
		basicHeaderRune := []rune(r.BasicHeader)
		if !unicode.IsLetter(basicHeaderRune[0]) {
			return errors.New("basic header at 0 should be a letter")
		}
		shouldBeDigit := []int{1, 2, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24} // index where basic header should be a digit
		for _, v := range shouldBeDigit {
			if !unicode.IsDigit(basicHeaderRune[v]) {
				return fmt.Errorf("basic header at %d should be a digit", v)
			}
		}

	}

	// --------------------------------- application header ---------------------------------
	if len(r.ApplicationHeader) < 1 {
		return errors.New("application header is empty")
	} else {
		applicationHeaderRune := []rune(r.ApplicationHeader)
		if applicationHeaderRune[0] == 'I' {
			if len(r.ApplicationHeader) < 16 || len(r.ApplicationHeader) > 21 {
				return errors.New("application header (input) should be 16-21 characters")
			}
		} else if applicationHeaderRune[0] == 'O' {
			if len(r.ApplicationHeader) < 46 || len(r.ApplicationHeader) > 47 {
				return errors.New("application header (output) should be 46-47 characters")
			}
		} else {
			return errors.New("wrong message type, should be I or O")
		}

		// --------------------------------- message type ---------------------------------
		messageType := r.ApplicationHeader[1:4]
		for _, v := range messageType {
			if !unicode.IsDigit(v) {
				return fmt.Errorf("message type (index 1-3) should be a digit")
			}
		}

		if applicationHeaderRune[0] == 'I' && len(r.ApplicationHeader) > 16 {
			// --------------------------------- priority ---------------------------------
			priority := applicationHeaderRune[16]
			if priority != 'S' && priority != 'U' && priority != 'N' {
				return fmt.Errorf("priority (index 16) should be S|U|N")
			}

			// --------------------------------- delivery monitoring ---------------------------------
			if len(r.ApplicationHeader) > 17 {
				deliveryMonitoring := applicationHeaderRune[17]
				if deliveryMonitoring != '1' && deliveryMonitoring != '2' && deliveryMonitoring != '3' {
					return fmt.Errorf("delivery monitoring (index 17) should be 1|2|3")
				}

				if priority == 'U' && deliveryMonitoring != '1' && deliveryMonitoring != '3' {
					return fmt.Errorf("when priority (index 16) is U, delivery monitoring (index 17) should be 1 or 3")
				} else if priority == 'N' && deliveryMonitoring != '2' {
					return fmt.Errorf("when priority (index 16) is N, delivery monitoring (index 17) should be 2 or empty")
				}
			}

			// --------------------------------- obsolescence period ---------------------------------
			if len(r.ApplicationHeader) > 20 {
				obsolescencePeriod := r.ApplicationHeader[18:21]
				for _, v := range obsolescencePeriod {
					if !unicode.IsDigit(v) {
						return fmt.Errorf("obsolescence period (index 18-20) should be a digit")
					}
				}

				if priority == 'U' && obsolescencePeriod != "003" {
					return fmt.Errorf("when priority (index 16) is U, obsolescence period (index 18-20) should be 003")
				} else if priority == 'N' && obsolescencePeriod != "020" {
					return fmt.Errorf("when priority (index 16) is N, obsolescence period (index 18-20) should be 020")
				}
			}
		} else if applicationHeaderRune[0] == 'O' && len(r.ApplicationHeader) > 46 {
			// --------------------------------- priority ---------------------------------
			priority := applicationHeaderRune[46]
			if priority != 'S' && priority != 'U' && priority != 'N' {
				return fmt.Errorf("priority (index 46) should be S|U|N")
			}
		}
	}

	// --------------------------------- user header ---------------------------------
	if err := r.UserHeader.Validate(); err != nil {
		return err
	}

	// --------------------------------- body ---------------------------------
	if err := r.Body.Validate(); err != nil {
		return err
	}
	return nil
}

// example file is on files/202MEP_52A_57A_58A.txt
func NewMt202RawFromFile(value string) (*MT202Raw, error) {
	// --------------------------------- basic header ---------------------------------
	basicHeaderRegex := regexp.MustCompile(`{1:([A-z0-9\/.:\-+, \n]+)}`)
	basicHeaderValue := basicHeaderRegex.FindStringSubmatch(value)
	if len(basicHeaderValue) != 2 {
		return nil, errors.New("basic header format is wrong")
	}

	// --------------------------------- application header ---------------------------------
	applicationHeaderRegex := regexp.MustCompile(`{2:([A-z0-9\/.:\-+, \n]+)}`)
	applicationHeaderValue := applicationHeaderRegex.FindStringSubmatch(value)
	if len(applicationHeaderValue) != 2 {
		return nil, errors.New("application header format is wrong")
	}

	// --------------------------------- user header ---------------------------------
	userHeaderRegex := regexp.MustCompile(`{3:({[{}A-z0-9\/.:\-+, \n]+})}`)
	userHeaderValue := userHeaderRegex.FindStringSubmatch(value)
	if len(applicationHeaderValue) != 2 {
		return nil, errors.New("user header format is wrong")
	}

	userHeaderChildRegex := regexp.MustCompile(`{([0-9]{3}):([A-z0-9\/.:\-+, \n]+)}`)
	userHeaderChildValue := userHeaderChildRegex.FindAllStringSubmatch(userHeaderValue[1], -1)

	var serviceIdentifier string
	var uniqueEndToEndId string
	for _, v := range userHeaderChildValue {
		if len(v) != 3 {
			return nil, errors.New("user header children format is wrong")
		}
		keyTag, err := strconv.Atoi(v[1])
		valueTag := src.TrimUnusedCharacters(v[2])
		if err != nil {
			return nil, err
		}
		switch keyTag {
		case 103:
			serviceIdentifier = valueTag
		case 121:
			uniqueEndToEndId = valueTag
		}
	}

	// --------------------------------- body ---------------------------------
	bodyRegex := regexp.MustCompile(`{4:([{}A-z0-9\/.:\-+, \n]+)}`)
	bodyValue := bodyRegex.FindStringSubmatch(value)
	if len(bodyValue) != 2 {
		return nil, errors.New("body format is wrong")
	}

	bodyChildRegex := regexp.MustCompile(`:([A-z0-9]{2,3}):([A-z0-9\/.\-+, \n]+)`)
	bodyChildValue := bodyChildRegex.FindAllStringSubmatch(bodyValue[1], -1)

	body := MT202RawBody{}
	for _, v := range bodyChildValue {
		if len(v) != 3 {
			return nil, errors.New("body children format is wrong")
		}

		keyTag := v[1]
		valueTag := src.TrimUnusedCharacters(v[2])
		switch {
		case keyTag == "20":
			body.TransactionReferenceNumber = valueTag
		case keyTag == "21":
			body.RelatedReference = valueTag
		case keyTag == "32A":
			body.ValueDateCurrencyAmount = valueTag
		case strings.HasPrefix(keyTag, "52"):
			if len(keyTag) != 3 {
				continue
			}
			body.OrderingInstitutionOption = []rune(keyTag)[2]
			body.OrderingInstitution = src.GetArrayOfStringSeparatedByEnter(valueTag)
		case strings.HasPrefix(keyTag, "57"):
			if len(keyTag) != 3 {
				continue
			}
			body.AccountWithInstitutionOption = []rune(keyTag)[2]
			body.AccountWithInstitution = src.GetArrayOfStringSeparatedByEnter(valueTag)
		case strings.HasPrefix(keyTag, "58"):
			if len(keyTag) != 3 {
				continue
			}
			body.BeneficiaryInstitutionOption = []rune(keyTag)[2]
			body.BeneficiaryInstitution = src.GetArrayOfStringSeparatedByEnter(valueTag)
		}
	}

	// --------------------------------- result ---------------------------------
	return &MT202Raw{
		BasicHeader:       basicHeaderValue[1],
		ApplicationHeader: applicationHeaderValue[1],
		UserHeader: MT202RawUserHeader{
			ServiceIdentifier: serviceIdentifier,
			UniqueEndToEndID:  uniqueEndToEndId,
		},
		Body:    body,
		RawData: value,
	}, nil
}
