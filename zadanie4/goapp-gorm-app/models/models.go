package models

import "gorm.io/gorm"

type Category struct {
    gorm.Model
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Products    []Product `json:"products"`
}

type Product struct {
    gorm.Model
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Price       float64  `json:"price"`
    CategoryID  uint     `json:"category_id"`
    Category    Category `json:"category"`
}

type Customer struct {
    gorm.Model
    Name   string  `json:"name"`
    Email  string  `json:"email"`
    Orders []Order `json:"orders"`
}

type Order struct {
    gorm.Model
    CustomerID uint        `json:"customer_id"`
    Customer   Customer    `json:"customer"`
    OrderItems []OrderItem `json:"order_items"`
    Total      float64     `json:"total"`
}

type OrderItem struct {
    gorm.Model
    OrderID   uint    `json:"order_id"`
    Order     Order   `json:"order"`
    ProductID uint    `json:"product_id"`
    Product   Product `json:"product"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}