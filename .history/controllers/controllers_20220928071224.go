package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/mhlengi/Go-EcommerceRestApi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HashPassword(password string) string{

}

func VerifyPassword(userPassword string, givenPassword) (bool, string) {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeOut(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user);
		err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"}) 
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone no. is already used"}) 
		}
		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format())
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format())
		user.ID = primitive.NewObjectID()
		user.
	}
}

func Login() gin.HandlerFunc {

}

func ProdutViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {
	
}