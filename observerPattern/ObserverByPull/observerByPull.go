package observerByPull

import "fmt"

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type Observer interface {
	Update(subject interface{})
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
		observer.Update(wd)
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

func (ccd *CurrentConditionsDisplay) Update(subject interface{}) {
	if wd, ok := subject.(*WeatherData); ok {
		ccd.temperature = wd.temperature
		ccd.humidity = wd.humidity
		ccd.Display()
	}
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

func (sd *statisticsDisplay) Update(subject interface{}) {
	if wd, ok := subject.(*WeatherData); ok {
		sd.temperatures = append(sd.temperatures, wd.temperature)
		sd.humidities = append(sd.humidities, wd.humidity)
		sd.pressures = append(sd.pressures, wd.pressure)

		sd.sum_temperature += wd.temperature
		sd.sum_humidity += wd.humidity
		sd.sum_pressure += wd.pressure

		sd.Display()
	}
}

func (sd *statisticsDisplay) Display() {
	fmt.Println("Average temperature: ", sd.sum_temperature/float64(len(sd.temperatures)))
	fmt.Println("Average humidity: ", sd.sum_humidity/float64(len(sd.humidities)))
	fmt.Println("Average pressure: ", sd.sum_pressure/float64(len(sd.pressures)))
}

func SimulateWeatherData() {
	weatherData := WeatherData{}

	currentConditionsDisplay := CurrentConditionsDisplay{}
	statisticsDisplay := statisticsDisplay{}

	weatherData.RegisterObserver(&currentConditionsDisplay)
	weatherData.RegisterObserver(&statisticsDisplay)

	weatherData.MeasurementsChanged(80, 65, 30.4)
	weatherData.MeasurementsChanged(82, 70, 29.2)
	weatherData.MeasurementsChanged(78, 90, 29.2)
}
