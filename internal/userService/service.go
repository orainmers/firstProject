package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}
func (s *UserService) CreateUser(user User) (User, error) {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return User{}, err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}
func (s *UserService) UpdateUserById(id uint, updatedUser User) (User, error) {
	if updatedUser.Password != "" {
		hashedPassword, err := HashPassword(updatedUser.Password)
		if err != nil {
			return User{}, err
		}
		updatedUser.Password = hashedPassword
	}
	return s.repo.UpdateUserById(id, updatedUser)
}
func (s *UserService) DeleteUserById(id uint) error {
	return s.repo.DeleteUserById(id)
}
