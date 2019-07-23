package emailvalidator

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/publicsuffix"
)

// NewEmailAddress will return a new email address
func NewEmailAddress(emailAddress string) (ep *EmailAddress, err error) {
	var e EmailAddress
	spl := strings.Split(emailAddress, "@")
	if len(spl) != 2 {
		err = fmt.Errorf("invalid number of \"@\", expected 1 and found %d", len(spl))
		return
	}

	e.LocalPart = spl[0]
	e.Domain = spl[1]
	ep = &e
	return
}

// EmailAddress represents an email address
type EmailAddress struct {
	LocalPart string `json:"localPart"`
	Domain    string `json:"domain"`
}

func (e *EmailAddress) validateLocalPart() (err error) {
	if len(e.LocalPart) == 0 {
		return ErrEmptyLocalPart
	}

	// Iterate through local part characters to ensure validity
	for i, char := range e.LocalPart {
		if !isValidLocalPartChar(char) {
			// Character is not valid, return error containin character and index
			err = fmt.Errorf("invalid character \"%s\" at index %d", string(char), i)
			return
		}
	}

	return e.validatePeriods()
}

func (e *EmailAddress) validatePeriods() (err error) {
	var wasPeriod bool
	lastRune := utf8.RuneCountInString(e.LocalPart) - 1
	// Iterate through local part characters to ensure validity
	for i, char := range e.LocalPart {
		if char != '.' {
			wasPeriod = false
			continue
		}

		switch {
		case i == 0:
			return ErrLocalPartLeadingDot
		case i == lastRune:
			return ErrLocalPartTrailingDot
		case wasPeriod:
			return ErrLocalPartMultiplePeriods
		}

		wasPeriod = true
	}

	return
}

func (e *EmailAddress) validateDomain() (err error) {
	if len(e.Domain) == 0 {
		return ErrEmptyDomain
	}

	var u *url.URL
	if u, err = url.Parse("https://" + e.Domain); err != nil {
		return
	}

	if u.Host == "" {
		err = fmt.Errorf("\"%s\" is not a valid domain", e.Domain)
	}

	var iCann bool
	if _, iCann = publicsuffix.PublicSuffix(e.Domain); !iCann {
		err = fmt.Errorf("\"%s\" does not have a valid TLD", e.Domain)
		return
	}

	return
}

// Validate will validate an email address
func (e *EmailAddress) Validate() (err error) {
	if err = e.validateLocalPart(); err != nil {
		return
	}

	if err = e.validateDomain(); err != nil {
		return
	}

	return nil
}

func (e *EmailAddress) String() string {
	return e.LocalPart + "@" + e.Domain
}
