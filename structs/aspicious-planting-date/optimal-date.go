package structs

type OptimalDate struct {
    Temperature []ReferencedDate `json:"temperature"`
    Humidity []ReferencedDate `json:"humidity"`
}
