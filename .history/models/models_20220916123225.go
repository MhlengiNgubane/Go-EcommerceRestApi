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
	P
}