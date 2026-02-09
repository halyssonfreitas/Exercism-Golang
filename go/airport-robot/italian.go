package airportrobot

import "fmt"

type Italian struct{}

func (i Italian) LanguageName() (languageName string) {
	return "Italian"
}

func (i Italian) Greet(visitorName string) (greeting string) {
	return fmt.Sprintf("I can speak Italian: Ciao %s!", visitorName)
}
