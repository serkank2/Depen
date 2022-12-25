package services

type UserService struct {
	UserRepository string
}

func (u *UserService) Login() string {
	return "Login"
}

func (u *UserService) Register() string {
	return "Register"
}

func (u *UserService) Update() string {
	return "Update"
}

func (u *UserService) Delete() string {
	return "Delete"
}

func NewUserService(userRepo string) *UserService {
	return &UserService{UserRepository: userRepo}
}
