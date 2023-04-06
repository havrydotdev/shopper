package handler

import (
	"github.com/gin-gonic/gin"
	"shopper/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		items := api.Group("/items")
		{
			items.POST("/", h.addItem)
			items.GET("/", h.getAllItems)
			id := items.Group("/:id")
			{
				id.GET("/", h.getItemById)
				id.PUT("/", h.updateItem)
				id.DELETE("/", h.deleteItem)

				comments := id.Group("/comments")
				{
					comments.GET("/", h.getAllComments)
					comments.POST("/", h.addComment)
					comments.PUT("/:comment_id", h.updateComment)
					comments.DELETE("/:comment_id", h.deleteComment)
				}

				ratings := id.Group("/rating")
				{
					ratings.POST("/", h.addNewRate)
				}
			}
		}

		discounts := api.Group("/discounts")
		{
			discounts.GET("/", h.getAllDiscounts)
			discounts.POST("/", h.addNewDiscount)
		}

		users := api.Group("/user/:id")
		{
			users.DELETE("/", h.deleteUser)
			users.PUT("/", h.updateUser)
			users.POST("/balance", h.updateUserBalance)
			users.GET("/history", h.getUserHistory)
			users.GET("/return/:item_id", h.returnItem)
			users.GET("/notifications", h.getUserNotifications)
		}

		notifications := api.Group("/notifications")
		{
			notifications.POST("/", h.addNewNotification)
		}

		companies := api.Group("/companies")
		{
			companies.POST("/", h.addNewCompany)
			companies.GET("/:id", h.getCompanyById)
			companies.PUT("/:id", h.updateCompany)
		}

		admin := api.Group("/admin", h.adminIdentity)
		{
			admin.POST("/set-user-id", h.setUserId)
		}
	}

	return router
}
