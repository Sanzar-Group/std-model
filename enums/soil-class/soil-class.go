package soilClass

type SoilClass int

const (
    VeryLow SoilClass = iota
    Low
    Medium
    High
    // Organic Matter
    OM_Low
    OM_Medium
    OM_High
)
