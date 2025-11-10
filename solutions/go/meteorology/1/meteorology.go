package meteorology

import "strconv"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return strconv.Itoa(t.degree) + " " + t.unit.String()
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[s]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
	return strconv.Itoa(speed.magnitude) + " " + speed.unit.String()
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string {
	return m.location + ": " + m.temperature.String() + ", " + "Wind " + m.windDirection + " at " + m.windSpeed.String() + ", " + strconv.Itoa(m.humidity) + "% Humidity"
}
