package main

import "fmt"

func Celsius2Fahrenheit(c float64) float64 {
	return (c*9/5 + 32)
}

func Celsius2Kelvin(c float64) float64 {
	return (c + 273.15)
}

func Kelvin2Celsius(k float64) float64 {
	return (k - 273.15)
}

func Fahrenheit2Celsius(f float64) float64 {
	return ((f - 32) * 5 / 9)
}

func convertKelvin(temperatureInput float64) (float64, float64) {
	// presumeing that the temperatureInput is Kelvin...
	resultCelsius := Kelvin2Celsius(temperatureInput)
	resultFahrenheit := Celsius2Fahrenheit(resultCelsius)

	//Do not remove this line
	return resultFahrenheit, resultCelsius
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	// presumeing that the temperatureInput is Celsius...
	resultFahrenheit := Celsius2Fahrenheit(temperatureInput)
	resultKelvin := Celsius2Kelvin(temperatureInput)
	//Do not remove this line
	return resultFahrenheit, resultKelvin
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	// presumeing that the temperatureInput is Fahrenheit...
	resultCelsius := Fahrenheit2Celsius(temperatureInput)
	resultKelvin := Celsius2Kelvin(resultCelsius)

	//Do not remove this line
	return resultKelvin, resultCelsius
}

func main() {
	var temperatureChoice int
	var temperatureInput float64
	fmt.Println("Enter your temperature format (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit: ")
	fmt.Scanln(&temperatureChoice)
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)

	if temperatureChoice == 1 {
		//Insert Code here
		resultFahrenheit, resultCelsius := convertKelvin(temperatureInput)
		//DO not remove this
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Celsius: ", resultCelsius)
	} else if temperatureChoice == 2 {
		//Insert Code here
		resultFahrenheit, resultKelvin := convertCelsius(temperatureInput)
		//DO not remove this
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Kelvin: ", resultKelvin)
	} else if temperatureChoice == 3 {
		//Insert Code here
		resultCelsius, resultKelvin := convertFahrenheit(temperatureInput)
		//DO not remove this
		fmt.Println("Kelvin: ", resultKelvin, " Celsius: ", resultCelsius)
	} else {
		fmt.Println("No Conversion")
	}

}
