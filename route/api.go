package route

import(
	"github.com/gin-gonic/gin"
	"rest_api_golang/controllers/ProductController"
	"rest_api_golang/controllers/UserController"

)

func RegisterRoutes(r *gin.Engine) {
	// Product routes
	products := r.Group("/api/products")
	{
		products.GET("/", ProductController.Index)
		products.GET("/:id", ProductController.Show)
		products.POST("/", ProductController.Create)
		products.PUT("/:id", ProductController.Update)
		products.DELETE("/", ProductController.Delete)
	}

	users := r.Group("/api/users")
	{
		users.POST("/", UserController.Register)
		users.GET("/detail/:id", UserController.GetUserWithProducts)
	}

}