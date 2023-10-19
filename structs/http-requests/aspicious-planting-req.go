package httpReq

import (
    "github.com/Sanzar-Group/std-model/enums/crops"
)

type AspiciousPlantingRequest struct {
    PlantingDate string `json:"planting_date"`
    Crop crops.Crops  `json:"crop"`
    Location string `json:"location"`
}


