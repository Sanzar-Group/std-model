package ApiResponses

type WheaterForecast struct {
    Hourly struct {
        Time []string `json:"time"`
        Temperature []float64 `json:"temperature_2m"`
        Humidity []int8`json:"relativehumidity_2m"`
    } `json:"hourly"`
}
