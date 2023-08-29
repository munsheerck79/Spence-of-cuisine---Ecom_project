package repository

import (
	"context"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAddCoupon(t *testing.T) {
	ctx := context.Background()

	testCase := []struct {
		name          string
		input         domain.Coupon
		beforeTest    func(sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "addcoupon",
			input: domain.Coupon{
				Code:              "112233",
				MinOrderValue:     100.00,
				DiscountMaxAmount: 200.00,
				DiscountPercent:   10,
				Description:       "addcoupon of 100min and 200max",
				ValidTill:         10,
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO coupons (.+)$`
				mockSQL.ExpectExec(expectedQuery).WithArgs(
					"112233",
					"addcoupon of 100min and 200max",
					100.00,
					10,
					200.00,
					mock.AnythingOfType("int64"), // Use mock.AnythingOfType to match valid_till
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: fmt.Errorf("failed to add coupon "),
		},
		{
			name: "addcoupon2",
			input: domain.Coupon{
				Code:              "1",
				MinOrderValue:     100.00,
				DiscountMaxAmount: 200.00,
				DiscountPercent:   10,
				Description:       "",
				ValidTill:         10,
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				expectedQuery := `^INSERT INTO coupons (.+)$`
				mockSQL.ExpectExec(expectedQuery).WithArgs(
					"112223",
					"jhjfjfhkvjb",
					100.00,
					10,
					200.00,
					mock.AnythingOfType("int64"), // Use mock.AnythingOfType to match valid_till
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: fmt.Errorf("failed to add coupon "),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock database: %v", err)
			}
			defer mockDB.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			if err != nil {
				t.Fatalf("failed to create GORM database: %v", err)
			}

			tc.beforeTest(mockSQL)

			ud := NewCouponRepository(gormDB)

			err1 := ud.AddCoupon(ctx, tc.input)

			assert.Equal(t, tc.expectedError, err1)
		})
	}
}

func TestGetCouponByCode(t *testing.T) {
	ctx := context.Background()
	testCase := []struct {
		name        string
		code        string
		beforeTest  func(sqlmock.Sqlmock)
		want        domain.Coupon
		expectedErr error
	}{{
		name: "getbycode",
		code: "112233",
		beforeTest: func(mockSQL sqlmock.Sqlmock) {
			expectedQuery := `SELECT \* FROM coupons WHERE  code \= \$1`
			mockSQL.ExpectQuery(expectedQuery).WithArgs("112233").WillReturnRows(sqlmock.NewRows([]string{"id", "Code", "description", "min_order_value", "discount_percent", "discount_max_amount", "valid_till"}).AddRow(1, "112233", "addcoupon of 100min and 200max", 100.00, 10, 200.00, 10))
		},

		want: domain.Coupon{
			ID:                1,
			Code:              "112233",
			MinOrderValue:     100.00,
			DiscountMaxAmount: 200.00,
			DiscountPercent:   10,
			Description:       "addcoupon of 100min and 200max",
			ValidTill:         10,
		},
		expectedErr: nil,
	},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock database: %v", err)
			}
			defer mockDB.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			if err != nil {
				t.Fatalf("failed to create GORM database: %v", err)
			}

			tc.beforeTest(mockSQL)

			ud := NewCouponRepository(gormDB)

			coupon, err1 := ud.GetCouponByCode(ctx, tc.code)

			assert.Equal(t, tc.expectedErr, err1)
			assert.Equal(t, tc.want, coupon)
		})
	}
}

func TestGetCoupon(t *testing.T) {
	ctx := context.Background()

	testCase := []struct {
		name        string
		beforeTest  func(sqlmock.Sqlmock)
		want        []domain.Coupon
		expectedErr error
	}{{
		name: "get coupons",
		beforeTest: func(mockSQL sqlmock.Sqlmock) {
			expectedQuery := `SELECT \* FROM coupons`
			rows := sqlmock.NewRows([]string{"id", "Code", "description", "min_order_value", "discount_percent", "discount_max_amount", "valid_till"})
			rows.AddRow(1, "112233", "addcoupon of 100min and 200max", 100.00, 10, 200.00, 10)
			rows.AddRow(2, "112234", "addcoupon of 100min and 200max", 150.00, 15, 250.00, 20)
			// Add more rows as needed

			mockSQL.ExpectQuery(expectedQuery).WillReturnRows(rows)
		},

		want: []domain.Coupon{
			{
				ID:                1,
				Code:              "112233",
				MinOrderValue:     100.00,
				DiscountMaxAmount: 200.00,
				DiscountPercent:   10,
				Description:       "addcoupon of 100min and 200max",
				ValidTill:         10,
			},
			{
				ID:                2,
				Code:              "112234",
				MinOrderValue:     150.00,
				DiscountMaxAmount: 250.00,
				DiscountPercent:   15,
				Description:       "addcoupon of 100min and 200max",
				ValidTill:         20,
			},
		},
		expectedErr: nil,
	},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockDB, mockSQL, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock database: %v", err)
			}
			defer mockDB.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			if err != nil {
				t.Fatalf("failed to create GORM database: %v", err)
			}

			tc.beforeTest(mockSQL)

			ud := NewCouponRepository(gormDB)

			coupon, err1 := ud.GetCoupon(ctx)

			assert.Equal(t, tc.expectedErr, err1)
			assert.Equal(t, tc.want, coupon)
		})
	}

}
