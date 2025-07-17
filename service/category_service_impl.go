package service

import (
	"MuhammadDimasAfifudin/belajar-golang-restful-api/exception"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/helper"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/model/domain"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/model/web"
	"MuhammadDimasAfifudin/belajar-golang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepsoitory repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepsoitory,
		DB:                 DB,
		Validate:           validate,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, req web.CategoryCreateRequest) (_ web.CategoryResponse) {
	err := s.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: req.Name,
	}

	category = s.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Update(ctx context.Context, req web.CategoryUpdateRequest) (_ web.CategoryResponse) {
	err := s.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = req.Name
	category = s.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	s.CategoryRepository.Delete(ctx, tx, category)
}

func (s *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) (_ web.CategoryResponse) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := s.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) (_ []web.CategoryResponse) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := s.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}
