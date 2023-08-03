package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	interfaces "github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/util/response"
	"gorm.io/gorm"
)

type paymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentDatabase{DB: DB}
}

func (p *paymentDatabase) AddMonyToWallet(ctx context.Context, userId uint, RazorPayPaymentId string, money float32) error {

	var User domain.Wallet
	query1 := `INSERT INTO wallet_histories (users_id, Payment_id, amount, date) VALUES (?, ?, ?, ?)`

	date := time.Now()
	if err := p.DB.Exec(query1, userId, RazorPayPaymentId, money, date).Error; err != nil {
		fmt.Println("33333")
		return errors.New("failed to add history")
	}
	fmt.Println("iiiiiii")
	query := `SELECT id,users_id,Balence FROM wallets WHERE users_id = ?`
	if err := p.DB.Raw(query, userId).Scan(&User).Error; err != nil {
		fmt.Println("1e1")
		return errors.New("failed to get key")
	}

	if User.ID == 0 {
		fmt.Println("iittt")
		query2 := `INSERT INTO wallets (users_id, balence,created_at) VALUES (?, ?, ?)`
		date := time.Now()
		if err := p.DB.Exec(query2, userId, money, date).Error; err != nil {
			fmt.Println("555i")
			return errors.New("failed to add history")
		}
		fmt.Println("user id not get")

	} else {
		fmt.Println("user id get")
		query := `UPDATE wallets SET balence = ?,updated_at =? WHERE users_id=?`
		updatedAt := time.Now()
		fmt.Println("old", User.Balence)
		fmt.Println(User)
		fmt.Println("nwe", money)
		netamount := User.Balence + money
		fmt.Println("net", netamount)
		err := p.DB.Exec(query, netamount, updatedAt, userId).Error
		if err != nil {
			return fmt.Errorf("failed to update address of %d", userId)
		}
		fmt.Println("2nd qry")

		return nil
	}
	return nil
}

func (p *paymentDatabase) FindTempData(ctx context.Context, RazorPayKey string) (response.Order, error) {
	var Data response.Order

	fmt.Println("jhvkh", RazorPayKey)

	query := `SELECT id,users_id,actual_price,discount_price,net_amount,status AS order_status, Razor_pay_order_id,order_date
	FROM orders_temps WHERE razor_pay_order_id = ?`
	if err := p.DB.Raw(query, RazorPayKey).Scan(&Data).Error; err != nil {
		fmt.Println("11111")
		return Data, errors.New("failed to get key")
	}
	fmt.Println("success", Data.UsersID)
	fmt.Println("sts", Data.OrderStatus)
	return Data, nil
}

func (p *paymentDatabase) GetWallet(c context.Context, userid uint) (response.WalletRes, error) {
	var wallet response.WalletRes
	query := `SELECT Balence FROM wallets WHERE users_id = ?`
	if err := p.DB.Raw(query, userid).Scan(&wallet.Balence).Error; err != nil {
		return wallet, errors.New("failed to get key")
	}
	query1 := `SELECT Payment_id, amount, date FROM wallet_histories WHERE users_id = ? ORDER BY date DESC`
	if err := p.DB.Raw(query1, userid).Scan(&wallet.History).Error; err != nil {
		return wallet, errors.New("failed to get key")
	}
	return wallet, nil
}
func (p *paymentDatabase) AddPaymentId(paymentID string, RazorPayOrderId string) error {

	query := `UPDATE orders_temps SET payment_id = ? WHERE razor_pay_order_id=?`

	err := p.DB.Exec(query, paymentID, RazorPayOrderId).Error
	if err != nil {
		return fmt.Errorf("failed to update razorpay")
	}
	return nil
}
