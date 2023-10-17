package httpReq

import (
    "github.com/Sanzar-Group/std-model/enums/loam-class"
    "github.com/Sanzar-Group/std-model/structs"
)

type Whitewashing_Req struct {
    SoilTest structs.SoilTest `json:"soil_test"`
    LoamType loamClass.LoamClass `json:"loam_type"`
}
