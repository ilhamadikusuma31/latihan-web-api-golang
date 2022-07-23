package game

import "time"

type Game struct {
	ID         int
	Judul      string
	TahunRilis int
	Harga      int
	Genre      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
