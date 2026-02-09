package airportrobot

import "fmt"

type Portuguese struct{}

func (i Portuguese) LanguageName() (languageName string) {
	return "Portuguese"
}

func (i Portuguese) Greet(visitorName string) (greeting string) {
	return fmt.Sprintf("I can speak Portuguese: Olá %s!", visitorName)
}
