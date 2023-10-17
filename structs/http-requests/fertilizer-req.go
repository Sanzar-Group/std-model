package httpReq

import (
    "github.com/Sanzar-Group/std-model/structs"
)

type Fertilizer_req struct {
    ExpectedProduction int `json:"expectedProduction"`
    SoilTest structs.SoilTest `json:"soilTest"`
}
