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
	User_ID			*string
	UserCart		[]ProductUser
	Address_Details	[]Address
	Order_Status	[]Order
}

type Product struct{
	Product_ID		primitive.ObjectID
	Product_Name	*string
	Price			*uint64
	Rating			*uint8
	Image			*string
}

type ProductUser struct{
	Product_ID		primitive.ObjectID
	Product_Name	*strin
	Price
	Rating
	Image
}

type Address struct{
	Address_ID		primitive.ObjectID
	House
	Street
	City
	Postal_Address
}

type Order struct{
	Order_ID			primitive.ObjectID
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