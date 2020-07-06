package checkmail

import (
	"errors"
	"testing"
)

var (
	samples = []struct {
		mail    string
		format  bool
		account bool // host+user should be in the test string
	}{
		{mail: "cmelgarejodev@gmail.com", format: true, account: true},
		{mail: "cmelgarejo.dev@gmail.com", format: true, account: true},
		{mail: "unsupported@gmail2.com", format: true, account: false},
		{mail: "cmelgarejo@loopcontext.com", format: true, account: false},
		{mail: " info@loopcontext.com", format: false, account: false},
		{mail: "info@loopcontext.com ", format: false, account: false},
		{mail: "test@thisdomainshouldnotexistever.com", format: true, account: false},
		{mail: "loopcontext1234567890@loopcontext.com", format: true, account: false},
		{mail: "@loopcontext.com", format: false, account: false},
		{mail: " test@loopcontext.com", format: false, account: false},
		{mail: "test test@loopcontext.com", format: false, account: false},
		{mail: "ááá&é&ààà@loopcontext.com", format: false, account: false},
		{mail: "test@mail@loopcontext.com", format: false, account: false},
		{mail: "test@wong domain wabbit.com", format: false, account: false},
		{mail: "admin@googlemaps.com", format: true, account: false},
		{mail: "a@loopcontext.semperfi", format: true, account: false},
		{mail: "foobar@", format: false, account: false},
	}
)

func TestErrors(t *testing.T) {
	err := NewSMTPError(errors.New("500 *buzzer sounds*"))
	if err.Error() != "500 *buzzer sounds*" {
		t.Errorf("failed: %v", NewSMTPError(errors.New("500 *buzzer sounds*")).Error())
	}
	err = NewSMTPError(errors.New("200 OK"))
	if err.Code() != "200" {
		t.Errorf("failed: %v", NewSMTPError(errors.New("200 OK")).Error())
	}
}

func TestValidateHost(t *testing.T) {
	for _, s := range samples {
		if !s.format {
			continue
		}
		err := ValidateHost(s.mail)
		if err != nil && s.account == true {
			t.Errorf(`"%s" => unexpected error: %q`, s.mail, err)
		}
		if err == nil && s.account == false {
			t.Errorf(`"%s" => expected error`, s.mail)
		}
	}
}

func TestValidateFormat(t *testing.T) {
	for _, s := range samples {
		err := ValidateFormat(s.mail)
		if err != nil && s.format == true {
			t.Errorf(`"%s" => unexpected error: %q`, s.mail, err)
		}
		if err == nil && s.format == false {
			t.Errorf(`"%s" => expected error`, s.mail)
		}
	}
}
