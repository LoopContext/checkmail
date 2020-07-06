# checkmail

[![Godoc Reference](https://godoc.org/github.com/loopcontext/checkmail?status.svg)](http://godoc.org/github.com/loopcontext/checkmail)
[![Coverage](http://gocover.io/_badge/github.com/loopcontext/checkmail)](http://gocover.io/github.com/loopcontext/checkmail)
[![Go Report Card](https://goreportcard.com/badge/github.com/loopcontext/checkmail)](https://goreportcard.com/report/github.com/loopcontext/checkmail)

checkmail is a simple go package to check the validity of an email, it can check:

- [Format](https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address)
- Domain validity
- Mailbox existance, by checking the user and mailbox host

## Install

```bash
go get -u github.com/loopcontext/checkmail
```

## Example

### Check email format

```go
email := "email@email.com"
err := ValidateFormat(email)
if err != nil {
    fmt.Printf(`"%s" -> format error: %q`, mail, err)
}
// Send email ...
```

### Check the host and email existance

```go
email := "email@email.com"
err := ValidateHost(email)
if err != nil {
    fmt.Printf(`"%s" -> host error: %q`, mail, err)
}
// Send email ...
```
