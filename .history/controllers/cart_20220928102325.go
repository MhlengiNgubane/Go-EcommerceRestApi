package controllers

import( 
	"time"
	"context"
	"errors"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

	type Application struct {
		prodCollection *mongo.Collection
		userCollection * mongo.Collection
	}

	func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
		return &Application{
			prodCollection
		}
	}

func AddToCart() gin.HandlerFunc{

}

func RemoveItem() gin.HandlerFunc{

}

func GetItemFromCart() gin.HandlerFunc{

}

func BuyFromCart() gin.HandlerFunc{

}

func InstantBuy() gin.HandlerFunc{
	
}