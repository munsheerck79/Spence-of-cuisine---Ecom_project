package usecase

import (
	"context"
	"fmt"

	"github.com/munsheerck79/Ecom_project.git/pkg/domain"
	"github.com/munsheerck79/Ecom_project.git/pkg/repository/interfaces"
	"github.com/munsheerck79/Ecom_project.git/pkg/usecase/interfacess"
	"github.com/munsheerck79/Ecom_project.git/util/request"
	"github.com/munsheerck79/Ecom_project.git/util/response"
)

type ProductUsecase struct {
	productRepository interfaces.ProductRepository
}

func NewProductService(repo interfaces.ProductRepository) interfacess.ProductService {
	return &ProductUsecase{productRepository: repo}
}

func (p *ProductUsecase) GetCategory(ctx context.Context) ([]domain.Category, error) {

	categoryList, err := p.productRepository.GetCategory(ctx)
	if err != nil {
		return categoryList, err
	}

	return categoryList, nil
}

func (p *ProductUsecase) Getvariations(ctx context.Context) ([]domain.Variation, error) {

	variationList, err := p.productRepository.Getvariations(ctx)
	if err != nil {
		return variationList, err
	}

	return variationList, nil
}

func (p *ProductUsecase) AddCategory(ctx context.Context, category domain.Category) error {

	DBProduct, err := p.productRepository.FindCategory(ctx, category)
	if err != nil {
		fmt.Println("qwerttyui")
		return err
	}

	if DBProduct.ID == 0 {

		err = p.productRepository.SaveCategory(ctx, category)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("%v catogory already exists", DBProduct.CategoryName)
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (p *ProductUsecase) AddVarient(ctx context.Context, varient domain.Variation) error {

	DBVarient, err := p.productRepository.FindVarient(ctx, varient)
	if err != nil {
		return err
	}
	if DBVarient.ID == 0 {

		err = p.productRepository.SaveVarient(ctx, varient)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("%v varient already exists", DBVarient.Name)
	}

	return nil

}

/////////////////////////////////////////////////////////////////////////////////

func (p *ProductUsecase) AddProduct(ctx context.Context, product domain.Product) error {

	DBProduct, err := p.productRepository.FindProduct(ctx, product)
	if err != nil {
		return err
	}
	if DBProduct.ID == 0 {

		err = p.productRepository.SaveProduct(ctx, product)
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("%v product already exists", DBProduct.ID)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *ProductUsecase) EditProduct(ctx context.Context, product domain.Product) error {

	DBProduct, err := p.productRepository.FindProduct(ctx, product)
	if err != nil {
		return err
	}
	if DBProduct.ID == 0 {
		return fmt.Errorf("%v product is not exists", DBProduct.Code)

	} else {

		if product.Name == "" {
			product.Name = DBProduct.Name
		}
		if product.Description == "" {
			product.Description = DBProduct.Description
		}
		if product.QtyInStock == 0 {
			product.QtyInStock = DBProduct.QtyInStock
		}
		if product.Image == "" {
			product.Image = DBProduct.Image
		}
		if product.CategoryID == 0 {
			product.CategoryID = DBProduct.CategoryID
		}

		err = p.productRepository.EditProduct(ctx, product)
		if err != nil {
			return err
		}

		return nil

	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *ProductUsecase) AddPrice(ctx context.Context, price domain.Price) error {

	DBProductPrice, err := p.productRepository.AddPrice(ctx, price)
	if err != nil {
		return err
	}
	if DBProductPrice.ID == 0 {

		err = p.productRepository.SaveProductPrice(ctx, price)
		if err != nil {
			return err
		}

	} else {
		// return fmt.Errorf("%v product already exists", price.ProductID)
		err = p.productRepository.UpdateProductPrice(ctx, price)
		if err != nil {
			return err

		}
	}

	return nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (p *ProductUsecase) GetProductList(ctx context.Context, page request.ReqPagination) ([]response.ProductRes, []response.VariationR, error) {

	DBProduct, variations, err := p.productRepository.GetProductList(ctx,page)
	if err != nil {
		return DBProduct, variations, err
	}
	return DBProduct, variations, nil

}

func (p *ProductUsecase) GetProductsByCategoryName(ctc context.Context, CID uint) ([]response.ProductDetails, error) {

	DBProduct, err := p.productRepository.GetProductsByCategoryName(ctc, CID)
	if err != nil {
		return DBProduct, err
	}
	return DBProduct, nil

}

func (p *ProductUsecase) GetOrderStatus(ctx context.Context) ([]domain.OrderStatus, error) {

	statusList, err := p.productRepository.GetOrderStatus(ctx)
	if err != nil {
		return statusList, err
	}

	return statusList, nil
}

func (p *ProductUsecase) AddOrderStatus(ctx context.Context, body domain.OrderStatus) error {

	orderstatus, err := p.productRepository.FindOrderStatus(ctx, body)
	if err != nil {
		return err
	}

	if orderstatus.Id == 0 {

		err = p.productRepository.SaveOrderStatus(ctx, body)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("%v status already exists", orderstatus.Status)

}

func (p *ProductUsecase) GetProduct(ctx context.Context, Id uint) (response.ProductDetails, error) {

	product, err := p.productRepository.GetProduct(ctx, Id)
	if err != nil {
		return product, err
	}

	return product, nil
}
