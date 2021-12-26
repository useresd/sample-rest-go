package currency

type Repository interface {
	FindByID(id int) (*Currency, error)
	GetAll() []*Currency
	Save(*Currency) (*Currency, error)
	Delete(*Currency) error
}
