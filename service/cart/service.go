package cart

type Service struct {
	repository *repository
}

func New() *Service {
	return &Service{
		repository: &repository{},
	}
}
func (s *Service) FetchCategories() ([]Category, error) {
	return s.repository.fetchCategories()
}

func (s *Service) FetchCategoryItems(categoryID int, page int) ([]Item, error) {
	return s.repository.fetchCategoryItems(categoryID, page)
}

func (s *Service) FetchCart(userID int) (*Cart, error) {
	return s.repository.fetchCart(userID)
}

func (s *Service) CreateCart(userID int) (*Cart, error) {
	return s.repository.createCart(userID)
}

func (s *Service) AddItemToCart(cartID, itemID, quantity int) error {
	return s.repository.addItemToCart(cartID, itemID, quantity)
}

func (s *Service) RemoveItemFromCart(cartID, itemID int) error {
	return s.repository.removeItemFromCart(cartID, itemID)
}

func (s *Service) ClearCart(cartID int) error {
	return s.repository.clearCart(cartID)
}
