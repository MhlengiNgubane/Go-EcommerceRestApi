package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{

	ID
	First_Name
	Password
	Email
	Phone
	Token
	Refresh_Token
	Created_At
	Updated_At
	User_ID
	UserCart
	Address_Details
	Order_Status
}

type Product struct{
	Product_ID
	Product_Name
	Price
	Rating
	Image
}

type ProductUser struct{
	Product_ID
	Product_Name
	Price
	Rating
	Image
}

type Address struct{
	Address_id
	House
	Street
	City
	PostalAddress
}

type Order struct{
	
}

type Payment struct{

}