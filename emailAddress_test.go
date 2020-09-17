package emailvalidator

import (
	"testing"

	"github.com/Hatch1fy/errors"
)

func TestNewEmailAddress(t *testing.T) {
	var err error
	if _, err = NewEmailAddress("engineering@hatchify.co"); err != nil {
		t.Fatal(err)
	}
}

func TestEmailAddress_Validate(t *testing.T) {
	var err error
	testCases := []testCase{
		{
			emailAddress: "engineering@hatchify.co",
			err:          nil,
		},
		{
			emailAddress: "John.Doe@example.com",
			err:          nil,
		},
		{
			emailAddress: "John_Doe@example.com",
			err:          nil,
		},
		{
			emailAddress: "John__Doe@example.com",
			err:          nil,
		},
		{
			emailAddress: "_John_Doe@example.com",
			err:          nil,
		},
		{
			emailAddress: "John_Doe_@example.com",
			err:          nil,
		},
		{
			emailAddress: "engineering@hatchifyco",
			err:          errors.Error("\"hatchifyco\" does not have a valid TLD"),
		},
		{
			emailAddress: "@",
			err:          ErrEmptyLocalPart,
		},
		{
			emailAddress: "helloworld@",
			err:          ErrEmptyDomain,
		},
		{
			emailAddress: "a\"b(c)d,e:f;g<h>i[j\\k]l@example.com",
			err:          errors.Error("invalid character \"\"\" at index 1"),
		},
		{
			emailAddress: "hello\\world@example.com",
			err:          errors.Error("invalid character \"\\\" at index 5"),
		},
		{
			emailAddress: "just\"not\"right@example.com",
			err:          errors.Error("invalid character \"\"\" at index 4"),
		},
		{
			emailAddress: "John..Doe@example.com",
			err:          ErrLocalPartMultiplePeriods,
		},
		{
			emailAddress: ".John.Doe@example.com",
			err:          ErrLocalPartLeadingDot,
		},
		{
			emailAddress: "John.Doe.@example.com",
			err:          ErrLocalPartTrailingDot,
		},
	}

	for _, testCase := range testCases {
		var e *EmailAddress
		if e, err = NewEmailAddress(testCase.emailAddress); err != nil {
			t.Fatal(err)
		}

		if err = e.Validate(); !testCase.isErrorMatch(err) {
			t.Fatalf("invalid error, expected %v and received %v", testCase.err, err)
		}
	}
}

type testCase struct {
	emailAddress string
	err          error
}

func (t *testCase) isErrorMatch(err error) (ok bool) {
	if t.err == nil {
		return err == nil
	}

	// At this point, t.err is NOT nil. We can safely use this assumption

	if err == nil {
		// t.err is NOT nil and err is nil, return false
		return false
	}

	return t.err.Error() == err.Error()
}
