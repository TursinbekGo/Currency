package models

type SaleProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateSaleProduct struct {
	SaleID            string  `json:"sale_id"`
	ProductID         string  `json:"product_id"`
	Discount          int64   `json:"discount"`
	DiscountType      string  `json:"discount_type"`
	ProductName       string  `json:"product_name"`
	ProductPrice      float64 `json:"product_price"`
	PriceWithDiscount float64 `json:"price_with_discount"`
	DiscountPrice     float64 `json:"discount_price"`
	Count             int64   `json:"count"`
}

type SaleProduct struct {
	Id                string  `json:"id"`
	SaleID            string  `json:"sale_id"`
	ProductID         string  `json:"product_id"`
	Discount          int64   `json:"discount"`
	DiscountType      string  `json:"discount_type"`
	ProductName       string  `json:"product_name"`
	ProductPrice      float64 `json:"product_price"`
	PriceWithDiscount float64 `json:"price_with_discount"`
	DiscountPrice     float64 `json:"discount_price"`
	Count             int64   `json:"count"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type UpdateSaleProduct struct {
	Id                string  `json:"id"`
	SaleID            string  `json:"sale_id"`
	ProductID         string  `json:"product_id"`
	Discount          int64   `json:"discount"`
	DiscountType      string  `json:"discount_type"`
	ProductName       string  `json:"product_name"`
	ProductPrice      float64 `json:"product_price"`
	PriceWithDiscount float64 `json:"price_with_discount"`
	DiscountPrice     float64 `json:"discount_price"`
	Count             int64   `json:"count"`
}

type SaleProductGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type SaleProductGetListResponse struct {
	Count        int            `json:"count"`
	SaleProducts []*SaleProduct `json:"sale_products"`
}
