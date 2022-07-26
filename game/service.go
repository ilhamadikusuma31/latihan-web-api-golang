package game

type Service interface {
	S_FindAll() ([]Game, error)
	S_FindById(S_ID int) (Game, error)
	S_Create(S_gimR GameRequest) (Game, error)
	S_Update(id int, S_gimR GameRequest) (Game, error)
	S_Delete(id int) (Game, error)
}

type service struct {
	Rep Repository
}

func (s *service) S_FindAll() ([]Game, error) {
	gims, error := s.Rep.FindAll()
	return gims, error
}

func (s *service) S_FindById(param int) (Game, error) {
	gim, error := s.Rep.FindById(param)
	return gim, error
}

func (s *service) S_Create(param GameRequest) (Game, error) {
	penampung := Game{
		Judul:      param.Judul,
		Harga:      param.Harga,
		Genre:      param.Genre,
		TahunRilis: param.TahunRilis,
	}
	gim, err := s.Rep.Create(penampung)
	return gim, err
}
func (s *service) S_Update(param1 int, param2 GameRequest) (Game, error) {
	penampung, _ := s.Rep.FindById(param1)
	penampung.Judul = param2.Judul
	penampung.Harga = param2.Harga
	penampung.Genre = param2.Genre
	penampung.TahunRilis = param2.TahunRilis
	gim, err := s.Rep.Update(penampung)
	return gim, err
}

func (s *service) S_Delete(param1 int) (Game, error) {
	penampung, err := s.Rep.FindById(param1)
	gim, err := s.Rep.Delete(penampung)
	return gim, err
}

func NewService(param Repository) *service {
	return &service{param}
}

//Nb:
//ini param => ini parameter
