package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{

	ID				primitive.ObjectID		`json:"_id" bson:"_id"`
	First_Name		*string					`json:"las_name`
	Last_Name		*string					`json:`
	Password		*string					`json:`
	Email			*string					`json:`
	Phone			*string					`json:`
	Token			*string					`json:`
	Refresh_Token	*string					`json:`
	Created_At		time.Time				`json:`
	Updated_At		time.Time				`json:`
	User_ID			*string					`json:`
	UserCart		[]ProductUser			`json:`
	Address_Details	[]Address				`json:`
	Order_Status	[]Order					`json:`
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
	Product_Name	*string
	Price			int
	Rating			*uint
	Image			*string
}

type Address struct{
	Address_ID		primitive.ObjectID
	House			*string
	Street			*string
	City			*string
	Postal_Address	*string
}

type Order struct{
	Order_ID			primitive.ObjectID
	Order_Cart			[]ProductUser
	Orderd_At			time.Time
	Price				int
	Discount			*int
	Payment_Method		Payment
}

type Payment struct{
	Digital bool
	COD     bool
}