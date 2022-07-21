package game

type GameInput struct {
	Judul string `json:"judul" binding:"required"`
	Harga int    `json:"harga" binding:"required"`
}
