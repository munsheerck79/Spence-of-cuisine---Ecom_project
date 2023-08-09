package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	interfaces "github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: DB}
}

func (i *userDatabase) FindUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Check any of the user details matching with db user list
	fmt.Println(user.FirstName)
	query := `SELECT * FROM users WHERE id = ? OR email = ? OR user_name = ?`
	if err := i.DB.Raw(query, user.ID, user.Email, user.UserName).Scan(&user).Error; err != nil {
		return user, errors.New("failed to get user")
	}
	fmt.Println("get user")
	return user, nil
}

func (i *userDatabase) SaveUser(ctx context.Context, user domain.Users) error {
	query := `INSERT INTO users (user_name, first_name, last_name, age, email, phone, password,created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now()
	err := i.DB.Exec(query, user.UserName, user.FirstName, user.LastName, user.Age,
		user.Email, user.Phone, user.Password, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save user %s", user.UserName)
	}
	return nil
}

func (i *userDatabase) FindAddress(ctx context.Context, id uint) (domain.Address, error) {
	var address domain.Address
	// Check any of the user details matching with db user list

	query := `SELECT * FROM addresses WHERE users_id = ?`
	if err := i.DB.Raw(query, id).Scan(&address).Error; err != nil {
		return domain.Address{}, errors.New("failed to get address")
	}
	fmt.Println("get address", address)
	return address, nil

}

func (i *userDatabase) SaveAddress(ctx context.Context, address domain.Address) error {
	query := `INSERT INTO addresses (users_id, address, muncipality, land_mark, district, state, phone_number,pin_code,created_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	createdAt := time.Now()
	fmt.Println("add", address.Address)
	fmt.Println("usr", address.UsersID)
	err := i.DB.Exec(query, address.UsersID, address.Address, address.Muncipality, address.LandMark,
		address.District, address.State, address.PhoneNumber, address.PinCode, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to save address %d", address.ID)
	}
	return nil
}

func (i *userDatabase) UpdateAddress(ctx context.Context, address domain.Address) error {

	query := `UPDATE addresses SET address = $1, muncipality = $2, land_mark= $3,
			  district= $4, state = $5, phone_number = $6  ,pin_code = $7 , updated_at =$8 WHERE users_id=$9`

	updatedAt := time.Now()

	err := i.DB.Exec(query, address.Address, address.Muncipality, address.LandMark,
		address.District, address.State, address.PhoneNumber, address.PinCode, updatedAt, address.UsersID).Error
	if err != nil {
		return fmt.Errorf("failed to update address of %d", address.UsersID)
	}
	return nil
}
func (i *userDatabase) AddToCart(ctx context.Context, item request.Cart) error {
	// var ex request.Cart
	// query1 := `SELECT product_id,users_id, FROM carts WHERE users_id = ? AND  product_id = ? AND variation_id = ? ;`
	// if err := i.DB.Raw(query1, item.UsersID, item.ProductID, item.VariationID).Scan(&ex).Error; err != nil {
	// 	return errors.New("failed at in searching existing data")
	// }
	// if ex.UsersID != 0 && ex.ProductID != 0 {
	// 	return errors.New("product is allredy at in cart")
	// }

	query := `INSERT INTO carts (users_id, product_id, variation_id,quantity,created_at) 
	VALUES ($1, $2, $3, $4,$5)`
	createdAt := time.Now()
	err := i.DB.Exec(query, item.UsersID, item.ProductID, item.VariationID, item.Quantity, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to add item to cart %d", item.ProductID)
	}
	return nil
}

func (i *userDatabase) DeleteFromCart(c context.Context, cartId uint, userId uint) error {

	query := `DELETE FROM carts WHERE users_id = ? AND id = ?`
	err := i.DB.Exec(query, userId, cartId).Error
	if err != nil {
		return fmt.Errorf("failed to delete item from cart %d", cartId)
	}
	return nil
}

func (i *userDatabase) AddToWishList(ctx context.Context, item domain.WishList) error {
	query := `INSERT INTO wish_lists (users_id, product_id,created_at) 
	VALUES ($1, $2, $3)`
	createdAt := time.Now()
	err := i.DB.Exec(query, item.UsersID, item.ProductID, createdAt).Error
	if err != nil {
		return fmt.Errorf("failed to add item to wishlist %d", item.ProductID)
	}
	return nil
}
func (i *userDatabase) DeleteFromWishLIst(c context.Context, Id uint, userId uint) error {
	fmt.Println("enter repo")
	query := `DELETE FROM wish_lists WHERE users_id = ? AND id = ?`
	err := i.DB.Exec(query, userId, Id).Error
	if err != nil {
		return fmt.Errorf("failed to delete item from cart %d", Id)
	}
	return nil
}

func (i *userDatabase) CartList(ctx context.Context, userId uint) ([]response.Cart, error) {
	var Cartlist []response.Cart
	query := `SELECT carts.id,products.Id AS product_id,products.name AS product_name,variations.name AS variation_name,
	prices.actual_price,prices.discount_price,carts.quantity,products.qty_in_stock 
	FROM carts
	LEFT JOIN products ON carts.product_id = products.id
	LEFT JOIN variations ON carts.variation_id = variations.id
	LEFT JOIN prices ON carts.product_id = prices.product_id AND carts.variation_id = prices.variation_id 
	WHERE carts.users_id = ?;`

	if err := i.DB.Raw(query, userId).Scan(&Cartlist).Error; err != nil {
		return Cartlist, errors.New("failed to get user")
	}

	return Cartlist, nil

}
func (i *userDatabase) ListWishList(ctx context.Context, userId uint) ([]response.ProductDetails, error) {
	var W []response.ProductDetails
	// query := `SELECT wish_lists.id AS ID,products.code,products.name AS product_name,products.discription,categotyies.category_name,variations.name AS variation_name,
	// products.qty_in_stock,products.stock_status,prices.actual_price,prices.discount_price,products.image
	// FROM wish_lists
	// INNER JOIN products ON wish_lists.product_id = products.id
	// RIGHT JOIN prices ON variations.id = prices.variation_id
	// INNER JOIN prices ON variations.id = prices.variation_id

	// WHERE wish_lists.users_id = ?`
	// if err := i.DB.Raw(query, userId).Scan(&wishlist).Error; err != nil {
	// 	return wishlist, errors.New("failed to get user")
	// }
	// return wishlist, nil

	query := `SELECT products.id AS id,products.code,products.name AS product_name,products.description,products.qty_in_stock,products.image,category_Name
	FROM products 
	LEFT JOIN wish_lists ON wish_lists.product_id = products.id
	LEFT JOIN categories ON products.category_id =categories.id 
	WHERE wish_lists.users_id = ? `
	if err := i.DB.Raw(query, userId).Scan(&W).Error; err != nil {
		return W, errors.New("faild to show products")
	}

	for k := 0; k < len(W); k++ {
		query2 := `SELECT variations.name,prices.actual_price,prices.discount_price 
		FROM prices
		INNER JOIN variations ON variations.Id = prices.variation_id
		INNER JOIN products ON products.Id = prices.product_id
		WHERE  prices.product_id =? `
		if err := i.DB.Raw(query2, W[k].ID).Scan(&W[k].PriceList).Error; err != nil {
			return W, errors.New("faild to show products")
		}
	}
	return W, nil
}

func (i *userDatabase) OrderHistory(ctx context.Context, userId uint, page request.ReqPagination) ([]response.Order, error) {
	var orderHistory []response.Order
	limit := page.Count
	offset := (page.PageNumber - 1) * limit
	query := `SELECT orders.id ,users_id ,actual_price ,discount_price ,net_amount ,order_statuses.status AS order_status ,payment_method ,order_date ,payment_id
	FROM orders
	LEFT JOIN order_statuses ON orders.order_status_id = order_statuses.id
	WHERE users_id = ? ORDER BY order_date DESC LIMIT ? OFFSET ?`
	if err := i.DB.Raw(query, userId, limit, offset).Scan(&orderHistory).Error; err != nil {
		return orderHistory, errors.New("failed to get history")
	}
	return orderHistory, nil

}

func (i *userDatabase) GetWalletx(ctx context.Context, userId uint) (domain.Wallet, error) {
	var wallet domain.Wallet
	query := `SELECT * FROM wallets WHERE users_id = ?`
	if err := i.DB.Raw(query, userId).Scan(&wallet).Error; err != nil {
		return wallet, errors.New("failed to get wallet")
	}
	return wallet, nil

}

func (i *userDatabase) EditCartProduct(ctx context.Context, item domain.Cart) error {

	var product domain.Cart
	query := `SELECT * FROM carts WHERE users_id =? AND product_id = ?`
	if err := i.DB.Raw(query, item.UsersID, item.ProductID).Scan(&product).Error; err != nil {
		return errors.New("failed to get wallet")
	}
	if product.ID == 0 {
		return fmt.Errorf("%v item not exists", item.ProductID)
	} else {
		if item.VariationID == 0 {
			item.VariationID = product.VariationID
		}
		if item.Quantity == 0 {
			item.Quantity = product.Quantity
		}
		query := `UPDATE carts SET variation_id = $1, quantity = $2, updated_at =$3 WHERE id=$4`
		updatedAt := time.Now()
		err := i.DB.Exec(query, item.VariationID, item.Quantity, updatedAt, product.ID).Error
		if err != nil {
			return fmt.Errorf("failed to update cart of %d", product.ProductID)
		}
		return nil
	}
}
