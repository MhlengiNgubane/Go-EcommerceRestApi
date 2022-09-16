package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{

	ID				primitive.ObjectID
	First_Name		*string
	Last_Name		*string
	Password		*string
	Email			*string
	Phone			*string
	Token			*string
	Refresh_Token	*string
	Created_At		time.Time
	Updated_At		time.Time
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
	Address_ID
	House
	Street
	City
	Postal_Address
}

type Order struct{
	Order_ID
	Order_Cart
	Orderd_At
	Price
	Discount
	Payment_Method
}

type Payment struct{
	Digital bool
	COD     bool
}