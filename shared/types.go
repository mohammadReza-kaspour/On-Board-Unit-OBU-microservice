package shared

import (
	"math/rand"
	"time"
	"uuid"
)

type ObuData struct {
	Id   string  `json:"id"`
	Time int64   `json:"time"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func NewObuData() *ObuData {
	return &ObuData{
		Id:   uuid.NewV4().String(),
		Time: time.Now().UnixMilli(),
		Lat:  rand.Float64()*180 - 90,
		Long: rand.Float64()*360 - 180,
	}
}
