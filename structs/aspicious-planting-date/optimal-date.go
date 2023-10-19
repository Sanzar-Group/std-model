package aspiciousPlanting 

type OptimalDate struct {
    Temperature []ReferencedDate `json:"temperature"`
    Humidity []ReferencedDate `json:"humidity"`
}
