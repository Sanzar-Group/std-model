package structs

import "sanzar/stdModel/enums"

type whitewash_res struct {
    Factor int `json:"factor"`
    Status enums.LoamStatus `json:"status"`
}
