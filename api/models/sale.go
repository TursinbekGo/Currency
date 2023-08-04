package models

type SalePrimaryKey struct {
	Id string `json:"id"`
}

type CreateSale struct {
	UserID     string  `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	TotalCount int64   `json:"total_count"`
}

type Sale struct {
	Id         string  `json:"id"`
	UserID     string  `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	TotalCount int64   `json:"total_count"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type UpdateSale struct {
	Id         string  `json:"id"`
	UserID     string  `json:"user_id"`
	TotalPrice float64 `json:"total_price"`
	TotalCount int64   `json:"total_count"`
}

type SaleGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type SaleGetListResponse struct {
	Count int     `json:"count"`
	Sales []*Sale `json:"sales"`
}
