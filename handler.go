package main

import "database/sql"

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

// create user


// read user by email


// create product


// read all products


// create cart item


// read cart items by user id


// delete cart items by user id


// create order by user id


// read orders by user id


// create order items

