package cart

import (
	"database/sql"
	"errors"
	"github.com/katerji/ecommerce/db"
)

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
		if errors.Is(err, sql.ErrNoRows) {
			return r.createCart(userID)
		}
		return nil, err
	}

	cartItems, err := db.Fetch[dbCartItem, CartItem](fetchCartItemsByCartIDQuery, cart.ID)
	if err != nil && !errors.Is(err, db.ErrNoRows) {
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

func (r *repository) removeItemFromCart(cartID, itemID int) error {
	cartItem, err := db.FetchOne[dbCartItem, CartItem](fetchCartItemQuantityQuery, cartID, itemID)
	if err != nil {
		if errors.Is(err, db.ErrNoRows) {
			return nil
		}
		return err
	}
	if cartItem.InCartQuantity == 0 {
		return db.Delete(deleteItemFromCartQuery, cartID, itemID)
	}

	return db.Update(removeItemFromCartQuery, cartID, itemID)
}

func (r *repository) clearCart(cartID int) error {
	return db.Update(clearCartQuery, cartID)
}
