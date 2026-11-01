package entity

import "time"

type Product struct {
	ID         int
	Nama       string
	Harga      int
	Stok       int
	Keterangan string
	DeletedAt  time.Time
}
