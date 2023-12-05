package dicts

import (
    "github.com/Sanzar-Group/std-model/enums/soil-class"
	"github.com/Sanzar-Group/std-model/structs"
)

type Productivity map[int]map[soilClass.SoilClass]int

type NPKTable struct{
    N Productivity
    P Productivity
    P_clay_less400 Productivity
    K Productivity
    Limit_factor structs.NPK
    MaxProd [3]int
}
