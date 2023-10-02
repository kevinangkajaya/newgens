package models

import (
	"fmt"
	"regexp"
	"unicode"
)

type MT202RawUserHeader struct {
	ServiceIdentifier string `json:"service_identifier"` //103
	// MessageUserReference string `json:"message_user_reference"` //108
	UniqueEndToEndID string `json:"unique_end_to_end_id"` //121

}

func (r *MT202RawUserHeader) Validate() error {
	// --------------------------------- service identifier ---------------------------------
	if r.ServiceIdentifier != "" && len(r.ServiceIdentifier) != 3 {
		return fmt.Errorf("service identifier length should be 0 or 3")
	}
	for _, v := range r.ServiceIdentifier {
		if !unicode.IsLetter(v) {
			return fmt.Errorf("service identifier should be alphabetic")
		}
	}

	// --------------------------------- service identifier ---------------------------------
	// if r.MessageUserReference != "" && len(r.MessageUserReference) != 16 {
	// 	return fmt.Errorf("message user reference length should be 0 or 16")
	// }

	// --------------------------------- service identifier ---------------------------------
	if len(r.UniqueEndToEndID) != 36 {
		return fmt.Errorf("unique end to end transaction reference length should be 36")
	}

	regex, err := regexp.Compile(`[a-z0-9]{8}-[a-z0-9]{4}-4[a-z0-9]{3}-[89ab][a-z0-9]{3}-[a-z0-9]{12}`) //xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
	if err != nil {
		return err
	}
	isMatch := regex.MatchString(r.UniqueEndToEndID)
	if !isMatch {
		return fmt.Errorf("unique end to end transaction reference format is wrong")
	}

	return nil
}
