package structs

import (
    "github.com/Sanzar-Group/std-model/enums/crops"
)

type SoilTest struct {
    Crop crops.Crops `json:"crop"`
    N int `json:"n"`
    P int `json:"p"`
    K int `json:"k"`
    Clay int `json:"clay"`
    OrganicMatter int `json:"organicMatter"`
    Ph int `json:"ph"`
}
