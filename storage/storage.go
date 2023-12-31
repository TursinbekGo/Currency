package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	Category() CategoryRepoI
	Product() ProductRepoI
	Market() MarketRepoI
	User() UserRepoI
	Sale() SaleRepoI
	SaleProduct() SaleProductRepoI
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(context.Context, *models.UpdateProduct) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.ProductPrimaryKey) error
}

type MarketRepoI interface {
	Create(context.Context, *models.CreateMarket) (string, error)
	GetByID(context.Context, *models.MarketPrimaryKey) (*models.Market, error)
	GetList(context.Context, *models.MarketGetListRequest) (*models.MarketGetListResponse, error)
	Update(context.Context, *models.UpdateMarket) (int64, error)
	Patch(context.Context, *models.PatchRequest) (int64, error)
	Delete(context.Context, *models.MarketPrimaryKey) error
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
}

type SaleRepoI interface {
	Create(context.Context, *models.CreateSale) (string, error)
	GetByID(context.Context, *models.SalePrimaryKey) (*models.Sale, error)
	GetList(context.Context, *models.SaleGetListRequest) (*models.SaleGetListResponse, error)
	Update(context.Context, *models.UpdateSale) (int64, error)
	Delete(context.Context, *models.SalePrimaryKey) error
}

type SaleProductRepoI interface {
	Create(context.Context, *models.CreateSaleProduct) (string, error)
	GetByID(context.Context, *models.SaleProductPrimaryKey) (*models.SaleProduct, error)
	GetList(context.Context, *models.SaleProductGetListRequest) (*models.SaleProductGetListResponse, error)
	Update(context.Context, *models.UpdateSaleProduct) (int64, error)
	Delete(context.Context, *models.SaleProductPrimaryKey) error
}
