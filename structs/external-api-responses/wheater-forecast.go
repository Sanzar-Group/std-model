package ApiResponses

type WheaterForecast struct {
    Hourly []struct {
        Time []string `json:"time"`
    } `json:"hourly"`
    Temperature []float32 `json:"temperature_2m"`
    Humidity []int8`json:"relativehumidity_2m"`
}
