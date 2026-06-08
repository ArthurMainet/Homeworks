package order

import "Email-API/internal/products"

type OrderRequest struct {
	UserID   uint
	Products []products.Product
}
