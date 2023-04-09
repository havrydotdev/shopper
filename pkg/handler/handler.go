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
			items.POST("/", h.addItem)    // done
			items.GET("/", h.getAllItems) // done
			id := items.Group("/:id")
			{
				id.GET("/", h.getItemById) // done
				id.PUT("/", h.updateItem)
				id.DELETE("/", h.deleteItem) // in progress

				itemComments := id.Group("/comments")
				{
					itemComments.POST("/", h.addComment)
					itemComments.GET("/", h.getCommentsByItem)
				}

				itemDiscounts := id.Group("/discounts")
				{
					itemDiscounts.POST("/", h.addDiscountToItem)
				}

				ratings := id.Group("/rating")
				{
					ratings.POST("/", h.addNewRate) // done
				}
			}
		}

		comments := api.Group("/comments")
		{
			comments.GET("/", h.getAllComments)
			comments.PUT("/:id", h.updateComment)
			comments.DELETE("/:id", h.deleteComment)
		}

		users := api.Group("/users", h.adminIdentity) // done
		{
			users.DELETE("/delete", h.deleteUser)               // done
			users.PUT("/", h.updateUser)                        // done
			users.GET("/history", h.getUserHistory)             // done
			users.POST("/return/", h.returnItem)                // done
			users.GET("/notifications", h.getUserNotifications) // done
			users.PUT("/balance", h.updateUserBalance)          // done
			users.POST("/buy", h.buyItem)                       // done
		}

		companies := api.Group("/companies") // done
		{
			companies.POST("/", h.addNewCompany)    // done
			companies.GET("/:id", h.getCompanyById) // done
			companies.PUT("/:id", h.updateCompany)  // done
		}

		admin := api.Group("/admin", h.adminIdentity) // done
		{
			adminCompanies := admin.Group("/companies")
			{
				adminCompanies.POST("/companies/moderation", h.moderateCompany) // done
				adminCompanies.GET("/", h.getNotVerifiedCompanies)              // done
			}

			adminItems := admin.Group("/items")
			{
				adminItems.POST("/moderation", h.moderateItem) // done
				adminItems.GET("/", h.getNotVerifiedItems)     // done
			}

			notifications := admin.Group("/notifications")
			{
				notifications.POST("/:id", h.addNewNotification) //done
			}

			discounts := admin.Group("/discounts")
			{
				discounts.POST("/", h.addNewDiscount) // done
			}
		}
	}

	return router
}
