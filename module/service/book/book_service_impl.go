package service

import (
	model "challenges_9/module/model/book"
	repository "challenges_9/module/repository/book"
	"context"
	"log"
	"time"
)

type BookServiceImpl struct {
	RepoBook repository.BookRepository
}

func NewBookServiceImpl(rpbook repository.BookRepository) BookService {
	return &BookServiceImpl{
		RepoBook: rpbook,
	}
}

func (booksrv_imp *BookServiceImpl) CreateBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	tNow := time.Now()
	bookIn.Create_at = &tNow

	book, err = booksrv_imp.RepoBook.CreateBook(ctx, bookIn)
	if err != nil {
		log.Printf("[ERROR] error Insert Book :%v\n", err)
	}
	return book, err

}
func (booksrv_imp *BookServiceImpl) UpdateBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBook.FindByIdBook(ctx, bookIn.Id)
	if err != nil {
		return book, err
	}
	book.Title = bookIn.Title
	book.Author = bookIn.Author
	book.Desc = bookIn.Desc
	tNow := time.Now()
	book.Update_at = &tNow
	book, err = booksrv_imp.RepoBook.UpdateBook(ctx, book)
	if err != nil {
		log.Printf("[ERROR] error Update Book :%v\n", err)
	}
	return book, err
}
func (booksrv_imp *BookServiceImpl) FindByIdBookSrv(ctx context.Context, idIn uint64) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBook.FindByIdBook(ctx, idIn)
	if err != nil {
		log.Printf("[ERROR] error findbook Book :%v\n", err)
		return book, err
	}
	return book, err
}
func (booksrv_imp *BookServiceImpl) FindAllBookSrv(ctx context.Context) (books []model.Book, err error) {
	books, err = booksrv_imp.RepoBook.FindAllBook(ctx)
	if err != nil {
		log.Printf("[ERROR] error findbook Book :%v\n", err)
		return books, err
	}
	return books, err
}
func (booksrv_imp *BookServiceImpl) HardDeleteBookSrv(ctx context.Context, bookId uint64) (book model.Book, err error) {
	book, err = booksrv_imp.RepoBook.HardDeleteBook(ctx, bookId)
	if err != nil {
		log.Printf("[ERROR] error deletebook Book :%v\n", err)
		return book, err
	}
	return book, err
}

func (booksrv_imp *BookServiceImpl) SoftDeleteBookSrv(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	_, err = booksrv_imp.FindByIdBookSrv(ctx, bookIn.Id)
	if err != nil {
		return book, err
	}
	book, err = booksrv_imp.RepoBook.SoftDeleteBook(ctx, bookIn)
	if err != nil {
		log.Printf("[ERROR] error deletebook Book :%v\n", err)
		return book, err
	}
	return book, err
}
