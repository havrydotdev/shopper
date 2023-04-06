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
		auth.POST("/sign-up", h.signUp) // done
		auth.POST("/sign-in", h.signIn) // done
	}

	api := router.Group("/api", h.userIdentity) // done
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

		users := api.Group("/user/:id") // implement after notifications
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
			notifications.POST("/", h.addNewNotification) // implement after companies
		}

		companies := api.Group("/companies")
		{
			companies.POST("/", h.addNewCompany)    // done
			companies.GET("/:id", h.getCompanyById) // done
			companies.PUT("/:id", h.updateCompany)  // implement next
		}

		admin := api.Group("/admin", h.adminIdentity) // done
		{
			admin.POST("/set-user-id", h.setUserId)                // done
			admin.POST("/companies/moderation", h.moderateCompany) // done
			admin.POST("/items/moderation")                        // implement after item routes are done
		}
	}

	return router
}
