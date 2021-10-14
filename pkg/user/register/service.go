package register

type UserRepository interface {
	create(User) error
}
type Service interface {
	Register(User) error
}

type service struct {
	r UserRepository
}

func NewService(r UserRepository) Service {
	return service{r}
}

func (s service) Register(u User) error {
	return s.r.create(u)
}
