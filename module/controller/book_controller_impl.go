package controller

import (
	model "challenges_9/module/model/book"
	service "challenges_9/module/service/book"
	"challenges_9/pkg/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookControllerImpl struct {
	BookSrv service.BookService
}

func NewBookControllerImpl(booksrv service.BookService) BookController {
	return &BookControllerImpl{
		BookSrv: booksrv,
	}
}

func (book_ctr *BookControllerImpl) CreateBookCtr(ctx *gin.Context) {
	reqIn := model.Book{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	if reqIn.Title == "" || reqIn.Author == "" || reqIn.Desc == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: "failed to insert book",
		})
		return
	}

	book, err := book_ctr.BookSrv.CreateBookSrv(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Status: err.Error(),
		})
		return
	}
	//fmt.Println(book.Id)
	ctx.JSON(http.StatusOK, response.ResWeb{
		Code:   http.StatusOK,
		Status: "Success Create Book ",
		Data:   book,
	})
}
func (book_ctr *BookControllerImpl) UpdateBookCtr(ctx *gin.Context) {
	id, err := book_ctr.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	//bind
	reqIn := model.Book{}
	if err := ctx.ShouldBindJSON(&reqIn); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	if reqIn.Title == "" || reqIn.Author == "" || reqIn.Desc == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: "failed to Update book",
		})
		return
	}

	reqIn.Id = id
	_, err = book_ctr.BookSrv.UpdateBookSrv(ctx, reqIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResWeb{
		Code:   http.StatusOK,
		Status: "Success Update Book",
	})
}
func (book_ctr *BookControllerImpl) FindByIdBookCtr(ctx *gin.Context) {
	id, err := book_ctr.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	book, err := book_ctr.BookSrv.FindByIdBookSrv(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ResWeb{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.ResWeb{
		Code:   http.StatusOK,
		Status: "success find Book",
		Data:   book,
	})

}

func (book_ctr *BookControllerImpl) FindAllBookCtr(ctx *gin.Context) {
	books, err := book_ctr.BookSrv.FindAllBookSrv(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.ResWeb{
			Code:   http.StatusNotFound,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResWeb{
		Code:   http.StatusOK,
		Status: "success find  All Book",
		Data:   books,
	})

}
func (book_ctr *BookControllerImpl) SoftDeleteBookCtr(ctx *gin.Context) {
	id, err := book_ctr.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	bookIn := model.Book{
		Id: id,
	}
	_, err = book_ctr.BookSrv.SoftDeleteBookSrv(ctx, bookIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResWeb{
		Code:   http.StatusOK,
		Status: "success delele book",
	})

}
func (book_ctr *BookControllerImpl) HardDeleteBookCtr(ctx *gin.Context) {
	id, err := book_ctr.getIdFromParam(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	_, err = book_ctr.BookSrv.HardDeleteBookSrv(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.ResWeb{
		Code:   http.StatusOK,
		Status: "success delele book ",
	})

}

func (book_ctr *BookControllerImpl) getIdFromParam(ctx *gin.Context) (idUint uint64, err error) {
	id := ctx.Param("id")
	if id == "" {
		err = errors.New("failed id")
		ctx.JSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}
	// transform id string to uint64
	idUint, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("failed parse id")
		ctx.JSON(http.StatusBadRequest, response.ResWeb{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		})
		return
	}

	return idUint, err

}
