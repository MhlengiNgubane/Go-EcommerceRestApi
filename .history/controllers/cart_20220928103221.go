package controllers

import( 
	"time"
	"context"
	"errors"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

	type Application struct {
		prodCollection *mongo.Collection
		userCollection * mongo.Collection
	}

	func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
		return &Application{
			prodCollection: prodCollection,
			userCollection: userProduction
		}
	}

func (app *Application) AddToCart() gin.HandlerFunc{
	return 
}

func RemoveItem() gin.HandlerFunc{

}

func GetItemFromCart() gin.HandlerFunc{

}

func BuyFromCart() gin.HandlerFunc{

}

func InstantBuy() gin.HandlerFunc{
	
}