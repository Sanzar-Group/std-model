package structs

import "github.com/Sanzar-Group/std-model/enums/alkalinity"

type Whitewash_res struct {
    Factor int `json:"factor"`
    Alkalinity alkalinity.Alkalinity `json:"alkalinity"`
}
