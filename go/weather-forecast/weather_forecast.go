// Package weather is able to help to forecast weather.
package weather

var (
	// CurrentCondition represents the currrent condition.
	CurrentCondition string
	// CurrentLocation represents the current location.
	CurrentLocation string
)

// Forecast returna a message with forecast formated.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
