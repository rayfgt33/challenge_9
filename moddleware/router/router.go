package router

import (
	"moddleware/controllers"
	"moddleware/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		// panggil middlware
		productRouter.Use(middleware.Authentications())

		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/:productId", middleware.ProductAuth(), controllers.ReadProduct)
		productRouter.PUT("/:productId", middleware.ProductAuthAdm(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middleware.ProductAuthAdm(), controllers.DeleteProduct)
	}

	return r

}
