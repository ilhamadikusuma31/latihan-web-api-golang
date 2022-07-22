package game

import "time"

type game struct {
	id          string
	judul       string
	tahun_rilis int8
	harga       int
	genre       string
	created_at  time.Time
	update_at   time.Time
}
