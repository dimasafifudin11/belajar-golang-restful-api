package main

import (
	"MuhammadDimasAfifudin/belajar-golang-restful-api/app"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/controller"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/helper"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/middleware"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/repository"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
