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
	Adress_Details
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

type Adress struct{

}

type Order struct{

}

type Payment struct{

}