package goodObserver

import "fmt"

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type Observer interface {
	Update(temperature, humidity, pressure float64)
}

type DisplayElement interface {
	Display()
}

type WeatherData struct {
	temperature float64
	humidity    float64
	pressure    float64
	observers   []Observer
}

func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) RemoveObserver(o Observer) {
	for i, observer := range wd.observers {
		if observer == o {
			wd.observers = append(wd.observers[:i], wd.observers[i+1:]...)
		}
	}
}

func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observers {
		observer.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) MeasurementsChanged(currentTemperature, currentHumidity, currentPressure float64) {
	wd.temperature = currentTemperature
	wd.humidity = currentHumidity
	wd.pressure = currentPressure
	wd.NotifyObservers()
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

type StatisticsDisplay struct {
	temperatures []float64
	humidities   []float64
	pressures    []float64

	sum_temperature float64
	sum_humidity    float64
	sum_pressure    float64
}

func (sd *StatisticsDisplay) Update(temperature, humidity, pressure float64) {
	sd.temperatures = append(sd.temperatures, temperature)
	sd.humidities = append(sd.humidities, humidity)
	sd.pressures = append(sd.pressures, pressure)

	sd.sum_temperature += temperature
	sd.sum_humidity += humidity
	sd.sum_pressure += pressure

	sd.Display()
}

func (sd *StatisticsDisplay) Display() {
	fmt.Println("Average temperature: ", sd.sum_temperature/float64(len(sd.temperatures)))
	fmt.Println("Average humidity: ", sd.sum_humidity/float64(len(sd.humidities)))
	fmt.Println("Average pressure: ", sd.sum_pressure/float64(len(sd.pressures)))
}

func SimulateWeatherData() {
	weatherData := WeatherData{}

	currentConditionsDisplay := CurrentConditionsDisplay{}
	statisticsDisplay := StatisticsDisplay{}

	weatherData.RegisterObserver(&currentConditionsDisplay)
	weatherData.RegisterObserver(&statisticsDisplay)

	weatherData.MeasurementsChanged(80, 65, 30.4)
	weatherData.MeasurementsChanged(82, 70, 29.2)
	weatherData.MeasurementsChanged(78, 90, 29.2)
}
