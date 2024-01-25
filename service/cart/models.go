package cart

type Cart struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	CartItems []CartItem `json:"cart_items"`
}

type dbCart struct {
	ID     int `db:"id"`
	UserID int `db:"user_id"`
}

func (c dbCart) ToModel() any {
	return Cart{
		ID:     c.ID,
		UserID: c.UserID,
	}
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type dbCategory struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func (c dbCategory) ToModel() any {
	return Category(c)
}

type Item struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	Category    Category `json:"category"`
}

type dbItem struct {
	ID           int     `db:"id"`
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	Price        float64 `db:"price"`
	Quantity     int     `db:"quantity"`
	CategoryID   int     `db:"category_id"`
	CategoryName string  `db:"category_name"`
}

func (i dbItem) ToModel() any {
	return Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		Quantity:    i.Quantity,
		Category: Category{
			ID:   i.CategoryID,
			Name: i.CategoryName,
		},
	}
}

type ItemLight struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CartItem struct {
	CartID         int       `json:"cart_id"`
	Item           ItemLight `json:"item"`
	InCartQuantity int       `json:"in_cart_quantity"`
}

type dbCartItem struct {
	CartID         int     `db:"cart_i	d"`
	ItemID         int     `db:"item_id"`
	ItemName       string  `db:"item_name"`
	ItemPrice      float64 `db:"item_price"`
	InCartQuantity int     `db:"quantity"`
}

func (ci dbCartItem) ToModel() any {
	return CartItem{
		CartID: ci.CartID,
		Item: ItemLight{
			ID:    ci.ItemID,
			Name:  ci.ItemName,
			Price: ci.ItemPrice,
		},
		InCartQuantity: ci.InCartQuantity,
	}
}
