package account

// Repository Interface
type Repository interface {
	FindByID(int) (*Account, error)
	GetAll() ([]*Account, error)
	Save(*Account) (*Account, error)
	Delete(*Account) error
}
