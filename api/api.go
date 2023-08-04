package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	handler := handler.NewHandler(cfg, storage, logger)

	// r.Use(customCORSMiddleware())

	// v1 := r.Group("/v1")
	r.POST("/register", handler.Register)

	r.POST("/login", handler.LOGIN)

	r.POST("/category", handler.AuthMiddleware(), handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.AuthMiddleware(), handler.UpdateCategory)
	r.DELETE("/category/:id", handler.AuthMiddleware(), handler.DeleteCategory)

	r.POST("/sale", handler.AuthMiddleware(), handler.CreateSale)
	r.GET("/sale/:id", handler.GetByIdSale)
	r.GET("/sale", handler.GetListSale)
	r.PUT("/sale/:id", handler.AuthMiddleware(), handler.UpdateSale)
	r.DELETE("/sale/:id", handler.AuthMiddleware(), handler.DeleteSale)

	r.POST("/sale_product", handler.AuthMiddleware(), handler.CreateSaleProduct)
	r.GET("/sale_product/:id", handler.GetByIdSaleProduct)
	r.GET("/sale_product", handler.GetListSaleProduct)
	r.PUT("/sale_product/:id", handler.AuthMiddleware(), handler.UpdateSaleProduct)
	r.DELETE("/sale_product/:id", handler.AuthMiddleware(), handler.DeleteSaleProduct)

	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.AuthMiddleware(), handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.POST("/product", handler.AuthMiddleware(), handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.AuthMiddleware(), handler.UpdateProduct)
	r.PATCH("/product/:id", handler.PatchProduct)
	r.DELETE("/product/:id", handler.AuthMiddleware(), handler.DeleteProduct)

	// v1.Use(handler.AuthMiddleware())
	r.POST("/market", handler.AuthMiddleware(), handler.CreateMarket)
	r.GET("/market/:id", handler.GetByIdMarket)
	r.GET("/market", handler.GetListMarket)
	r.PUT("/market/:id", handler.AuthMiddleware(), handler.UpdateMarket)
	r.PATCH("/market/:id", handler.PatchMarket)
	r.DELETE("/market/:id", handler.AuthMiddleware(), handler.DeleteMarket)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// func customCORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
// 		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
