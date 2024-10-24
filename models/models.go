package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"column:user_name;not null"`
	Email       string `gorm:"column:email;not null"`
	Password    string `gorm:"column:password;not null" json:"-"`
	PhoneNumber string `gorm:"column:phonenumber;not null"`
	Status      string `gorm:"check(status IN('Active', 'Inactive', 'Blocked'))"`
}

type Address struct {
	AddressID    int `gorm:"primaryKey;autoIncrement"`
	UserID       int `gorm:"not null;index;constraint:OnDelete:CASCADE;foreignKey:UserID;references:UserID"`
	AddressLine1 string
	AddressLine2 string
	Country      string
	City         string
	PostalCode   string
	Landmark     string
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Admin struct {
	AdminID   int `gorm:"primaryKey;autoIncrement"`
	AdminName string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Category struct {
	CategoryID   uint   `gorm:"primaryKey" json:"category_id"`
	CategoryName string `json:"name"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	ProductID   int     `gorm:"primaryKey"`
	ProductName string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  uint    `gorm:"not null;index;constraint:OnDelete:CASCADE;foreignKey:CategoryID;references:CategoryID" json:"category_id"`
	ImgURL      string  `json:"img_url"`
	Status      string  `gorm:"check(status IN('Available', 'Unavailable', 'Out of stock'))"`
	Quantity    int    `json:"quantity"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Wishlist struct {
	WishlistID int `gorm:"primaryKey;autoIncrement"`
	ProductID  int `gorm:"not null;foreignKey:ProductID;references:ProductID"`
	UserID     int `gorm:"not null;index;foreignKey:UserID;references:UserID"`
	CreatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}

type Cart struct {
	CartID    int `gorm:"primaryKey;autoIncrement"`
	UserID    int `gorm:"not null;index"`
	ProductID int `gorm:"not null"`
	Total     int
	Quantity  int
	User      User    `gorm:"foreignKey:UserID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

type Order struct {
	OrderID       int `gorm:"primaryKey;autoIncrement"`
	UserID        int `gorm:"not null;index;foreignKey:UserID;references:UserID"`
	ProductID     int `gorm:"not null;index;foreignKey:ProductID;references:ProductID"`
	PaymentID     int
	OrderDate     time.Time
	Total         int
	CouponID      int `gorm:"foreignKey:CouponID;references:CouponID"`
	Discount      int
	Quantity      int
	Status        string `gorm:"check(status IN('Pending', 'Processing', 'Delivered', 'Canceled'));"`
	Amount        float64
	Method        string `gorm:"check(method IN('Credit Card', 'PayPal', 'Bank Transfer'));"`
	PaymentStatus string `gorm:"check(status IN('Processing', 'Success', 'Failed'));"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	OrderItems    []OrderItem `gorm:"foreignKey:OrderID;references:OrderID"`
}

type OrderItem struct {
	OrderItemsID int `gorm:"primaryKey;autoIncrement"`
	OrderID      int `gorm:"not null;index;foreignKey:OrderID;references:OrderID"`
	UserID       int `gorm:"not null;index;foreignKey:UserID;references:UserID"`
	ProductID    int `gorm:"not null;index;foreignKey:ProductID;references:ProductID"`
	Quantity     int
	Price        float64
	Discount     int
}

type Coupon struct {
	CouponID             int    `gorm:"primaryKey;autoIncrement"`
	CouponCode           string `gorm:"unique"`
	CouponDiscountAmount int
	Description          string
	StartDate            time.Time
	Period               int
	MinPurchaseAmount    int
	MaxPurchaseAmount    int
	IsActive             string `gorm:"check(is_active IN('Active', 'Inactive'))"`
}

type ReviewRating struct {
	ReviewRatingID int `gorm:"primaryKey;autoIncrement"`
	UserID         int `gorm:"not null;foreignKey:UserID;references:UserID"`
	ProductID      int `gorm:"not null;foreignKey:ProductID;references:ProductID"`
	Review         string
	Rating         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type TempUser struct {
	UserName    string `json:"username"`
	Address     string
	Email       string `json:"email"`
	Password    string
	PhoneNumber string
}

type UserLoginMethod struct {
	UserLoginMethodEmail string
	LoginMethod          string
}

type OTP struct {
	Email  string
	Code   string
	Expiry time.Time
}
