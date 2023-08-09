package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/munsheerck79/Ecom_project.git/pkg/api/auth"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) interfacess.UserService {
	return &UserUseCase{userRepository: repo}
}

func (u *UserUseCase) SignUp(ctx context.Context, user domain.Users) error {
	// Check if user already exist
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return err
	}
	//
	if DBUser.ID == 0 {
		// Hash user password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			fmt.Println("Hashing failed")
			return err
		}
		user.Password = string(hashedPass)

		// Save user if not exist
		err = u.userRepository.SaveUser(ctx, user)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("%v user already exists", DBUser.UserName)
	}

	return nil
}

func (u *UserUseCase) Login(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		fmt.Println("error3333")

		return user, err
	}
	if DBUser.ID == 0 {
		fmt.Println("name", DBUser.FirstName)
		return user, errors.New("user not exist")
	}
	// Check if the user blocked by admin
	if DBUser.BlockStatus {
		fmt.Println("status=======", DBUser.BlockStatus)
		return user, errors.New("user blocked by admin")
	}
	fmt.Println("phone= ", DBUser.Phone)
	if _, err := auth.TwilioSendOTP("+91" + DBUser.Phone); err != nil {

		return user, fmt.Errorf("failed to send otp %v", err)
	}
	// check password with hashed pass
	if bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password)) != nil {
		fmt.Println("password incorrect")
		return user, errors.New("password incorrect")
	}
	fmt.Println("sucsess and return")
	return DBUser, nil

}
func (u *UserUseCase) OTPLogin(ctx context.Context, user domain.Users) (domain.Users, error) {
	// Find user in db
	DBUser, err := u.userRepository.FindUser(ctx, user)
	if err != nil {
		return user, err
	} else if DBUser.ID == 0 {
		return user, errors.New("user not exist")
	}
	return DBUser, nil
}

func (u *UserUseCase) GetAddress(ctx context.Context, userId uint) (domain.Address, error) {
	var address domain.Address
	address.UsersID = userId
	DBAddress, err := u.userRepository.FindAddress(ctx, userId)
	if err != nil {
		return DBAddress, err
	}
	return DBAddress, nil
}
func (u *UserUseCase) ListCart(ctx context.Context, userId uint) ([]response.Cart, float32, float32, error) {
	var TotalPrice float32
	var TotalDiscPrice float32
	cartList, err := u.userRepository.CartList(ctx, userId)
	for i := 0; i < len(cartList); i++ {
		TotalPrice += cartList[i].ActualPrice * float32(cartList[i].Quantity)
		TotalDiscPrice += cartList[i].DiscountPrice * float32(cartList[i].Quantity)
	}

	if err != nil {

		return cartList, TotalPrice, TotalDiscPrice, err
	}
	return cartList, TotalPrice, TotalDiscPrice, nil

}

func (u *UserUseCase) AddAddress(ctx context.Context, address domain.Address) error {
	DBAddress, err := u.userRepository.FindAddress(ctx, address.UsersID)
	if err != nil {
		return err
	}
	if DBAddress.ID == 0 {
		// Save user if not exist
		err = u.userRepository.SaveAddress(ctx, address)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("%v user already exists", DBAddress.UsersID)
	}
	return nil
}

func (u *UserUseCase) EditAddress(ctx context.Context, address domain.Address) error {
	DBAddress, err := u.userRepository.FindAddress(ctx, address.UsersID)
	if err != nil {
		return err
	}
	if DBAddress.ID == 0 {
		return fmt.Errorf("%v user not exists", DBAddress.UsersID)
	} else {

		if address.Address == "" {
			address.Address = DBAddress.Address
			fmt.Println("1")
		}
		if address.District == "" {
			address.District = DBAddress.District
			fmt.Println("2")
		}
		if address.LandMark == "" {
			fmt.Println("3")
			address.LandMark = DBAddress.LandMark
		}
		if address.State == "" {
			fmt.Println("4")
			address.State = DBAddress.State
		}
		if address.Muncipality == "" {
			fmt.Println("5")
			address.Muncipality = DBAddress.Muncipality
		}
		if address.PhoneNumber == "" {
			fmt.Println("6")
			address.PhoneNumber = DBAddress.PhoneNumber
		}
		if address.PinCode == "" {
			fmt.Println("7")
			address.PinCode = DBAddress.PinCode
		}
		// Save user if not exist
		err = u.userRepository.UpdateAddress(ctx, address)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UserUseCase) AddToCart(ctx context.Context, item request.Cart) error {

	err := u.userRepository.AddToCart(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) DeleteFromCart(c context.Context, cartId uint, userId uint) error {

	err := u.userRepository.DeleteFromCart(c, cartId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) AddToWishList(ctx context.Context, item domain.WishList) error {

	err := u.userRepository.AddToWishList(ctx, item)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUseCase) DeleteFromWishLIst(c context.Context, Id uint, userId uint) error {
	fmt.Println("enter usecase")
	err := u.userRepository.DeleteFromWishLIst(c, Id, userId)

	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) ListWishList(ctx context.Context, userId uint) ([]response.ProductDetails, error) {

	wishList, err := u.userRepository.ListWishList(ctx, userId)
	if err != nil {
		return wishList, err
	}

	return wishList, nil

}

func (u *UserUseCase) OrderHistory(ctx context.Context, userId uint, page request.ReqPagination) ([]response.Order, error) {

	orderHistory, err := u.userRepository.OrderHistory(ctx, userId, page)
	if err != nil {
		return orderHistory, err
	}

	return orderHistory, nil

}

// func (u *UserUseCase) GetWallet(ctx context.Context, userId uint) (domain.Wallet, error) {

// 	wallet, err := u.userRepository.GetWallet(ctx, userId)
// 	if err != nil {
// 		return wallet, err
// 	}

// 	return wallet, nil

// }

func (u *UserUseCase) EditCartProduct(ctx context.Context, item domain.Cart) error {
	err := u.userRepository.EditCartProduct(ctx, item)
	if err != nil {
		return err
	}
	return nil
}
