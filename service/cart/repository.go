package cart

import "github.com/katerji/ecommerce/db"

type repository struct{}

func (r *repository) fetchCategories() ([]Category, error) {
	return db.Fetch[dbCategory, Category](fetchCategoriesQuery)
}

func (r *repository) fetchCategoryItems(categoryID int, page int) ([]Item, error) {
	if page < 1 {
		page = 1
	}
	const itemsPerPage = 20
	offset := (page - 1) * itemsPerPage

	return db.Fetch[dbItem, Item](fetchItemsByCategoryIDQuery, categoryID, offset, itemsPerPage)
}

func (r *repository) fetchCart(userID int) (*Cart, error) {
	cart, err := db.FetchOne[dbCart, Cart](fetchCartByUserIDQuery, userID)
	if err != nil {
		return nil, err
	}

	cartItems, err := db.Fetch[dbCartItem, CartItem](fetchCartItemsByCartIDQuery, cart.ID)
	if err != nil {
		return nil, err
	}
	cart.CartItems = cartItems

	return &cart, nil
}

func (r *repository) createCart(userID int) (*Cart, error) {
	cartID, err := db.Insert(createCartQuery, userID)
	if err != nil {
		return nil, err
	}

	return &Cart{
		ID:        cartID,
		UserID:    userID,
		CartItems: nil,
	}, nil
}

func (r *repository) addItemToCart(cartID, itemID, quantity int) error {
	_, err := db.Insert(addItemToCartQuery, cartID, itemID, quantity)

	return err
}
