package httpRes

import "github.com/Sanzar-Group/std-model/enums/alkalinity"

type Whitewash_Res struct {
    Factor int `json:"factor"`
    Alkalinity alkalinity.Alkalinity `json:"alkalinity"`
}
