package game

type GameResponse struct {
	ID         int    `json:"id"`
	Judul      string `json:"judul" binding:"required"`
	TahunRilis int    `json:"tahun_rilis" binding:"required,number"`
	Harga      int    `json:"harga" binding:"required,number"`
	Genre      string `json:"genre" binding:"required"`
}
