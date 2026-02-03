package main

import "time"

type User struct {
	Id       int
	Email    string
	Password string
	Role     string
}

type Product struct {
	Id          int
	Name        string
	Description string
	Price       int
}

type CartItem struct {
	Id        int
	UserId    int
	ProductId int
	UnitPrice int
	Quantity  int
}

type OrderItem struct {
	Id        int
	OrderId   int
	ProductId int
	UnitPrice int
	Quantity  int
}

type Order struct {
	Id         int
	UserId     int
	TotalPrice int
	CreatedAt  time.Time
	Products   []Product
}
