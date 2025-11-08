// Package weather 用于预测天气.
package weather

var (
	// CurrentCondition 表示当前天气状况.
	CurrentCondition string
	// CurrentLocation 表示当前地理位置.
	CurrentLocation string
)

// Forecast 用于预测给定城市的天气状况.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
