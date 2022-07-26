package game

type Service interface {
	S_FindAll() ([]Game, error)
	S_FindById(S_ID int) (Game, error)
	S_Create(S_gimR GameRequest) (Game, error)
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

func NewService(param Repository) *service {
	return &service{param}
}

//Nb:
//ini param => ini parameter
