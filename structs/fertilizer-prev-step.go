package structs

type FertilizerPrevStep struct {
    Npk NPK `json:"npk"`
    HumidityNpk NPK `json:"humidity_npk"`
    TemperatureNpk NPK `json:"temperature_npk"`
    NdviNpk NPK `json:"ndvi_npk"`
}
