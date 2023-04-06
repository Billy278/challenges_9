package router

import (
	"challenges_9/module/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, bookCtrl controller.BookController) {
	//register all router
	group := r.Group("/book")
	group.GET("/all", bookCtrl.FindAllBookCtr)
	group.GET("/:id", bookCtrl.FindByIdBookCtr)
	group.POST("", bookCtrl.CreateBookCtr)
	group.PUT("/:id", bookCtrl.UpdateBookCtr)
	group.DELETE("/:id", bookCtrl.SoftDeleteBookCtr)
	group.DELETE("/hard/:id", bookCtrl.HardDeleteBookCtr)

}
