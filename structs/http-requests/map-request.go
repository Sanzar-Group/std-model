package httpReq

type Date struct {
    Start string `json:"start" binding:"required"`
    End string `json:"end" binding:"required"`
}

type MapRequest struct {
    Geojson string `json:"geojson" binding:"required"`
    Date Date `json:"date" binding:"required"`
    Index string `json:"index" binding:"required"`
}
