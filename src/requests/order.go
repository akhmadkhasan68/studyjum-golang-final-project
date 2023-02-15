package requests

type ProductsOrderRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Quantity  uint64 `json:"quantity" binding:"required"`
}

type CreateOrderRequest struct {
	OutletID string                 `json:"outlet_id" binding:"required"`
	Products []ProductsOrderRequest `json:"products" binding:"dive,required"`
}
