package models

type Car struct {
	ID       uint   `gorm:"primaryKey"`
	Merk     string `gorm:"not null; type:varchar(191)"`
	TypeCars string `gorm:"not null; type:varchar(191)"`
	Harga    int    `gorm:"not null; type:int"`
}
