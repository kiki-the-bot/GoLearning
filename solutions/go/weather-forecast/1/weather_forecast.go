//Package weather takes two string parameters and returns them as a single string.
package weather


var (
//CurrentCondition represents the current location.
	CurrentCondition string    
//CurrentLocation represents the city.
	CurrentLocation  string
)

//Forecast returns a message
//displaying the weather condition in the given city. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
