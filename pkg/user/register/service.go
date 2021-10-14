package register

type UserRepository interface {
}

type RegisterService interface {
}

type service struct {
	r UserRepository
}
