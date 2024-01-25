package cart

const (
	fetchCategoriesQuery        = "SELECT id, name FROM category"
	fetchItemsByCategoryIDQuery = `SELECT i.id, i.name, i.description, i.price, c.id AS category_id, c.name as category_name
							  FROM item i
									 JOIN category c on i.category_id = c.id
							  WHERE c.id = ?
						      ORDER BY i.id DESC
							  LIMIT ?, ?`
	fetchCartByUserIDQuery      = "SELECT id, user_id FROM cart WHERE user_id = ?"
	fetchCartItemsByCartIDQuery = "SELECT ci.cart_id, ci.item_id, ci.quantity, i.name as item_name, i.price FROM cart_item ci JOIN item i ON ci.item_id = i.id WHERE cart_id = ?"
	createCartQuery             = "INSERT INTO cart (user_id) VALUES (?)"
	addItemToCartQuery          = "INSERT INTO cart_item (cart_id, item_id, quantity) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE quantity = quantity + VALUES(quantity)"
)
