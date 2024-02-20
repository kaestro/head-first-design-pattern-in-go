package badObserver

import "fmt"

type WeatherData struct {
	temperature float64
	humidity    float64
	pressure    float64

	currentConditionsDisplay CurrentConditionsDisplay
	statisticsDisplay        statisticsDisplay
}

func (wd *WeatherData) MeasurementsChanged(currentTemperature, currentHumidity, currentPressure float64) float64 {
	wd.temperature = currentTemperature
	wd.humidity = currentHumidity
	wd.pressure = currentPressure

	wd.currentConditionsDisplay.Update(wd.temperature, wd.humidity, wd.pressure)
	wd.statisticsDisplay.Update(wd.temperature, wd.humidity, wd.pressure)

	return wd.temperature
}

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
}

func (ccd *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	ccd.temperature = temperature
	ccd.humidity = humidity

	ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Println("Current conditions: ", ccd.temperature, "F degrees and ", ccd.humidity, "% humidity")
}

type statisticsDisplay struct {
	temperatures []float64
	humidities   []float64
	pressures    []float64

	sum_temperature float64
	sum_humidity    float64
	sum_pressure    float64
}

func (sd *statisticsDisplay) Update(temperature, humidity, pressure float64) {
	sd.temperatures = append(sd.temperatures, temperature)
	sd.humidities = append(sd.humidities, humidity)
	sd.pressures = append(sd.pressures, pressure)

	sd.sum_temperature += temperature
	sd.sum_humidity += humidity
	sd.sum_pressure += pressure

	sd.Display()
}

func (sd *statisticsDisplay) Display() {
	fmt.Println("Average temperature: ", sd.sum_temperature/float64(len(sd.temperatures)))
	fmt.Println("Average humidity: ", sd.sum_humidity/float64(len(sd.humidities)))
	fmt.Println("Average pressure: ", sd.sum_pressure/float64(len(sd.pressures)))
}

func SimulateWeatherData() {
	weatherData := WeatherData{
		temperature: 80.0,
		humidity:    65.0,
		pressure:    30.4,
	}

	weatherData.currentConditionsDisplay = CurrentConditionsDisplay{}
	weatherData.statisticsDisplay = statisticsDisplay{}

	weatherData.MeasurementsChanged(82, 70, 29.2)
}
