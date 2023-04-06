package server

import (
	"challenges_9/config"
	"challenges_9/module/controller"
	repository "challenges_9/module/repository/book"
	service "challenges_9/module/service/book"
)

type Ctrs struct {
	BookCtr controller.BookController
}

func InitControllers() Ctrs {
	//why err if i do this?
	//dataStore := config.NewDBPostges()
	dataStore := config.NewDBPostgesGormConn()
	bookRepo := repository.NewBookRepositoryImpl(dataStore)
	bookServ := service.NewBookServiceImpl(bookRepo)
	bookCtr := controller.NewBookControllerImpl(bookServ)

	return Ctrs{
		BookCtr: bookCtr,
	}
}
