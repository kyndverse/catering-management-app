package model

type ProductResponse struct {
	ID         int
	Nama       string
	Harga      int
	Stok       int
	Keterangan string
}

type CreateProductRequest struct {
	Nama       string `form:"nama" validate:"required"`
	Harga      int    `form:"harga" validate:"required,min=1"`
	Stok       int    `form:"stok" validate:"required,min=1"`
	Keterangan string `form:"keterangan"`
}
