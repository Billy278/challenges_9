package repository

import (
	model "challenges_9/module/model/book"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepositoryImpl(dt *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		DB: dt,
	}
}

func (bookRepo *BookRepositoryImpl) DoMigration() (err error) {
	if err = bookRepo.DB.AutoMigrate(&model.Book{}); err != nil {
		panic(err)
	}
	log.Println("Success Create book table")
	return
}
func (bookRepo *BookRepositoryImpl) FindByIdBook(ctx context.Context, idIn uint64) (book model.Book, err error) {
	tx := bookRepo.DB.Model(&model.Book{}).Where("id=?", idIn).Find(&book)
	if err = tx.Error; err != nil {
		return
	}
	if book.Id <= 0 {
		err = errors.New("NOT FOUND")
	}

	return book, err

}
func (bookRepo *BookRepositoryImpl) CreateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	tx := bookRepo.DB.Model(&model.Book{}).Create(&bookIn)
	if err = tx.Error; err != nil {
		return
	}

	return bookIn, err
}
func (bookRepo *BookRepositoryImpl) UpdateBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {

	tx := bookRepo.DB.Model(&model.Book{}).Where("id=?", bookIn.Id).Updates(&bookIn)
	if tx.Error != nil {
		return
	}
	if tx.RowsAffected <= 0 {
		err = errors.New("BOOK IS NOT FOUND")
		return
	}

	return bookIn, err

}

func (bookRepo *BookRepositoryImpl) FindAllBook(ctx context.Context) (books []model.Book, err error) {
	tx := bookRepo.DB.Model(&model.Book{}).Find(&books)
	if err = tx.Error; err != nil {
		return
	}
	return
}
func (bookRepo *BookRepositoryImpl) SoftDeleteBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	tx := bookRepo.DB.Model(&model.Book{}).Clauses(clause.Returning{}).Where("id", bookIn.Id).Delete(&book)
	if err = tx.Error; err != nil {
		return
	}
	return
	// clause to return data after delete
	// by default, func delete
	// di gorm akan mengupdate column deleted_at

}

func (bookRepo *BookRepositoryImpl) HardDeleteBook(ctx context.Context, bookId uint64) (book model.Book, err error) {
	tx := bookRepo.DB.Unscoped().Model(&model.Book{}).Where("id", bookId).Delete(&book)
	if err = tx.Error; err != nil {
		return
	}
	return book, err
}
