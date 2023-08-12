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
)

type AdminUsecase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminService(repo interfaces.AdminRepository) interfacess.AdminService {
	return &AdminUsecase{adminRepository: repo}
}

func (u *AdminUsecase) LoginAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	// Find user in db
	fmt.Println(admin.Email)

	DBAdmin, err := u.adminRepository.FindAdmin(ctx, admin)
	fmt.Println("dbadmin",)

	if err != nil {
		fmt.Println("err@find admin",err)
		return admin, err
	} else if DBAdmin.ID == 0 {
		return admin, errors.New("user not exist")
	}
	fmt.Println(DBAdmin.Phone)
	if _, err := auth.TwilioSendOTP("+91" + DBAdmin.Phone); err != nil {
		fmt.Println("error @send otp")
		return admin, fmt.Errorf("failed to send otp %v", err)
	}
	// check password with hashed pass
	if DBAdmin.Password != admin.Password {
		fmt.Println("error @check password")
		return admin, errors.New("password incorrect")
	}
	fmt.Println("return to handler")
	return DBAdmin, nil

}

func (u *AdminUsecase) OTPLogin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	// Find user in db
	DBAdmin, err := u.adminRepository.FindAdmin(ctx, admin)
	if err != nil {
		return admin, err
	} else if DBAdmin.ID == 0 {
		return admin, errors.New("user not exist")
	}
	return DBAdmin, nil
}

func (a *AdminUsecase) GetUserlist(ctx context.Context, page request.ReqPagination) (userList []domain.Users, err error) {

	GetUser, err := a.adminRepository.GetUserlist(ctx, page)
	if err != nil {
		return GetUser, err
	}

	return GetUser, nil
}
func (a *AdminUsecase) BlockUser(ctx context.Context, body request.BlockUser) (string, error) {
	msg, err := a.adminRepository.BlockUser(ctx, body)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
func (a *AdminUsecase) UserDetails(ctx context.Context, body request.UserDetails) (response.UserDetails, error) {

	userDetails, err := a.adminRepository.UserDetails(ctx, body)

	if err != nil {
		return userDetails, err
	}
	if userDetails.Age == 0 {
		return userDetails, errors.New("user not exist")
	}
	return userDetails, nil
}

func (a *AdminUsecase) GetOrderlist(ctx context.Context, page request.ReqPagination) (orderList []response.AdminOrderList, err error) {

	GetOrder, err := a.adminRepository.GetOrderlist(ctx, page)
	if err != nil {
		return GetOrder, err
	}

	return GetOrder, nil
}

func (a *AdminUsecase) SalesReport(c context.Context, daterange request.DateRange) (salesReport []response.SalesReport, err error) {

	salesReport, err = a.adminRepository.GenerateSalesReport(c, daterange)
	if err != nil {
		return salesReport, err
	}
	return salesReport, nil
}

func (a *AdminUsecase) CancelOrder(ctx context.Context, body request.CancelOrder) (string, error) {
	msg, err := a.adminRepository.CancelOrder(ctx, body)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
