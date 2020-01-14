package models

import (
    "time"
)

// Model Struct
type Ration struct {
    Id   int
    PacketType string `orm:"size(11)"`
    PacketContent string `orm:"size(100)"`
    Calories int `orm:"size(5)"`
    ExpiryDate time.Time `orm:"type(date)"`
    Quantity int `orm:"size(5)"`
    CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}