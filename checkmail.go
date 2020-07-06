package checkmail

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"regexp"
	"strings"
	"time"
)

// STMPError describes an error from the SMTP server
type STMPError struct {
	Err error
}

func (e STMPError) Error() string {
	return e.Err.Error()
}

// Code Status code
func (e STMPError) Code() string {
	return e.Err.Error()[0:3]
}

// NewSMTPError creates a new STMPError instance
func NewSMTPError(err error) STMPError {
	return STMPError{
		Err: err,
	}
}

const forceDisconnectAfter = time.Second * 5

// Error vars
var (
	ErrBadFormat        = errors.New("invalid format")
	ErrUnresolvableHost = errors.New("unresolvable host")
	// email Regexp
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// ValidateFormat validates email format
func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

// ValidateHost validates host
func ValidateHost(email string) error {
	_, host := splitAccountHost(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return ErrUnresolvableHost
	}
	client, err := dialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), forceDisconnectAfter)
	if err != nil {
		return NewSMTPError(err)
	}
	defer client.Close()
	err = client.Hello("loopcontext.checkemail")
	if err != nil {
		return NewSMTPError(err)
	}
	err = client.Mail("smtp@loopcontext.com")
	if err != nil {
		return NewSMTPError(err)
	}
	err = client.Rcpt(email)
	if err != nil {
		return NewSMTPError(err)
	}
	return nil
}

// dialTimeout returns a new Client connected to an SMTP server at addr(ess)
// - addr(ess) must include a port, i.e: "mail.example.com:smtp".
// - timeout is, by default, 5 seconds
func dialTimeout(addr string, timeout time.Duration) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}
	t := time.AfterFunc(timeout, func() { conn.Close() })
	defer t.Stop()
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func splitAccountHost(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}
