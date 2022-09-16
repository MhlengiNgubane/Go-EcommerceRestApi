package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{

	ID				primitive.ObjectID		`json:"_id" bson:"_id"`
	First_Name		*string					`json:"first_name"`
	Last_Name		*string					`json:"last_name"`
	Password		*string					`json:"password"`
	Email			*string					`json:"email"`
	Phone			*string					`json:"phone"`
	Token			*string					`json:"token"`
	Refresh_Token	*string					`json:"refresh_token"`
	Created_At		time.Time				`json:"created_at"`
	Updated_At		time.Time				`json:"updated_at"`
	User_ID			*string					`json:"user_id"`
	UserCart		[]ProductUser			`json:"usercart" bson:"usercart"`
	Address_Details	[]Address				`json:"address" bson:"address"`
	Order_Status	[]Order					`json:"orders" bson:"orders"`
}

type Product struct{
	Product_ID		primitive.ObjectID		`bson:"_id"`
	Product_Name	*string					`json:"product_name"`
	Price			*uint64					`json:"price"`
	Rating			*uint8					`json:"rating"`
	Image			*string					`json:"image"`
}

type ProductUser struct{
	Product_ID		primitive.ObjectID		`bson:"_id"`
	Product_Name	*string					`json:"product_name" bson:"product_name"`
	Price			int						`json:"price" bson:"price"`
	Rating			*uint					`json:"rating" bson:"rating"`
	Image			*string					`json:"image" bson:"image"`
}

type Address struct{
	Address_ID		primitive.ObjectID		`bson:"_id"`
	House			*string					`json:"ouse_name" bson:"house_name"`
	Street			*string					`json:"street_name" bson:"street_name"`
	City			*string					`json:"city_name" bson:"city_name"`
	Postal_Address	*string					`json:"postal_address" bson:"postal_address"`
}

type Order struct{
	Order_ID		primitive.ObjectID		`bson:"_id"`
	
	Order_Cart		[]ProductUser
	Orderd_At		time.Time
	Price			int
	Discount		*int
	Payment_Method	Payment
}

type Payment struct{
	Digital bool
	COD     bool
}