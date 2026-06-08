package order

type OrderRequest struct {
	UserID     uint   `json:"userId"`
	ProductIDs []uint `json:"product_ids" validate:"required,min=1,dive,gt=0"`
}
