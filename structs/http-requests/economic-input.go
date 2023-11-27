package httpReq

import (
    "github.com/Sanzar-Group/std-model/structs"
)

type production struct {
    Expected int `json:"expected"`
    Previous int `json:"previous"`
}

type EconomicInput struct {
    NPK structs.NPK `json:"npk"`
    Ha int `json:"ha"`
    Posology int `json:"posology"`
    Production production `json:"production"`
}
