package verifier

import (
	"errors"
	"fmt"
	"github.com/Verifier/emailVerifier/utils"
	"net"
	"net/smtp"
	"regexp"
	"strings"
)

type SmtpError struct {
	Err error
}

func (e SmtpError) Error() string {
	return e.Err.Error()
}

func (e SmtpError) Code() string {
	return e.Err.Error()[0:3]
}

func NewSmtpError(err error) SmtpError {
	return SmtpError{
		Err: err,
	}
}

var (
	ErrBadFormat        = errors.New("invalid format")
	ErrUnresolvableHost = errors.New("unresolvable host")
	emailRegexp         = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func ValidateHost(email string) error {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return ErrUnresolvableHost
	}

	fmt.Println("Valid Domain")

	client, err := smtp.Dial(fmt.Sprintf("%s:%d", mx[1].Host, 25))
	if err != nil {
		fmt.Println(err)
		return NewSmtpError(err)
	}
	fmt.Println("Dial done")
	defer client.Close()
	err = client.Hello("checkmail.me")
	if err != nil {
		fmt.Println(err)
		return NewSmtpError(err)
	}
	fmt.Println("Hello DOne")
	err = client.Mail("test@gmail.com")
	if err != nil {
		fmt.Println(err)
		return NewSmtpError(err)
	}

	err = client.Rcpt(email)
	if err != nil {
		fmt.Println(err)
		return NewSmtpError(err)
	}
	return nil
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}

func ValidateEmail(email string) (err error) {

	err = ValidateHost(email)
	if err != nil {
		utils.LogInfo("Invalid email found" + email)
		return err
	}

	return err

}
