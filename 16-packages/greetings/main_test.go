package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Batman"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Batman")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello (Batman) = %q, %v, quiere coincidencia para %#q,nil`, msg, err, want)

	}

}
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello ("") = %q,%v quiere ""`, msg, err)
	}

}
