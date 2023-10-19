package aspiciousPlanting

type ReferencedDate struct {
    Date string `json:"date"`
    Temperature float64 `json:"temperature"`
    Humidity float64 `json:"humidity"`
    Difference int `json:"difference"`
}
