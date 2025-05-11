package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s",
		greeter.LanguageName(), greeter.Greet(visitorName))
}

type Italian struct{}

func (_ Italian) LanguageName() string {
	return "Italian"
}

func (_ Italian) Greet(visitorName string) string {
	return "Ciao " + visitorName + "!"
}

type Portuguese struct{}

func (_ Portuguese) LanguageName() string {
	return "Portuguese"
}

func (_ Portuguese) Greet(visitorName string) string {
	return "Olá " + visitorName + "!"
}
