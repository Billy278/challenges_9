package service

import (
	model "challenges_9/module/model/book"
	"context"
)

type BookService interface {
	CreateBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error)
	UpdateBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error)
	FindByIdBookSrv(ctx context.Context, idIn uint64) (book model.Book, err error)
	FindAllBookSrv(ctx context.Context) (books []model.Book, err error)
	SoftDeleteBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error)
	HardDeleteBookSrv(ctx context.Context, bookId uint64) (book model.Book, err error)
}
