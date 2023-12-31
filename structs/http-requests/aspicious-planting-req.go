package httpReq

import (
    "github.com/Sanzar-Group/std-model/enums/crops"
)

type AspiciousPlantingRequest struct {
    PlantingDate string `json:"planting_date"`
    Crop crops.Crops  `json:"crop"`
    Location struct{
        Latitude string `json:"latitude"`
        Longitude string `json:"longitude"`
    }`json:"location"`
}


