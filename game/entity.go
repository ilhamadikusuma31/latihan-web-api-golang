package game

import "time"

type Game struct {
	Id          string
	Judul       string
	Tahun_rilis int8
	Harga       int
	Genre       string
	Created_at  time.Time
	Update_at   time.Time
}
