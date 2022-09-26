package controllers

import "github.com/mhlengi/Go-EcommerceRestApi/models"

func HashPassword(password string) string{

}

func VerifyPassword(userPassword string, givenPassword) (bool, string) {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeOut(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		
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