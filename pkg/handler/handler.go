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

				ratings := id.Group("/rating")
				{
					ratings.POST("/", h.addNewRate)
				}
			}
		}

		comments := api.Group("/comments")
		{
			comments.GET("/", h.getAllComments)
			comments.POST("/", h.addComment)
			comments.PUT("/:comment_id", h.updateComment)
			comments.DELETE("/:comment_id", h.deleteComment)
		}

		discounts := api.Group("/discounts")
		{
			discounts.GET("/", h.getAllDiscounts)
			discounts.POST("/", h.addNewDiscount)
		}

		users := api.Group("/users") // done
		{
			users.DELETE("/delete", h.deleteUser)               // done
			users.PUT("/", h.updateUser)                        // done
			users.GET("/history", h.getUserHistory)             // done
			users.GET("/return/:item_id", h.returnItem)         // in progress
			users.GET("/notifications", h.getUserNotifications) // done
		}

		companies := api.Group("/companies") // done
		{
			companies.POST("/", h.addNewCompany)    // done
			companies.GET("/:id", h.getCompanyById) // done
			companies.PUT("/:id", h.updateCompany)  // done
		}

		admin := api.Group("/admin", h.adminIdentity) // done
		{
			admin.POST("/set-user-id", h.setUserId)                // done
			admin.POST("/companies/moderation", h.moderateCompany) // done
			admin.POST("/items/moderation")                        // implement after item routes are done

			notifications := admin.Group("/notifications")
			{
				notifications.POST("/:id", h.addNewNotification) //done
			}
		}
	}

	return router
}
