package routes

import (
	"../controllers"
	"net/http"
)

type CategoryRoutes interface {
	Routes() []*Route
}

type categoryRoutesImpl struct {
	categoriesController controllers.CategoriesController
}

func NewCategoryRoutes(categoriesController controllers.CategoriesController) *categoryRoutesImpl {
	return &categoryRoutesImpl{categoriesController}
}

func (r *categoryRoutesImpl) Routes() []*Route {
	return []*Route{
		&Route{
			Path:    "/categories",
			Method:  http.MethodPost,
			Handler: r.categoriesController.PostCategory,
		},
		&Route{
			Path:    "/categories",
			Method:  http.MethodGet,
			Handler: r.categoriesController.GetCategories,
		},
		&Route{
			Path:    "/categories/{category_id}",
			Method:  http.MethodGet,
			Handler: r.categoriesController.GetCategory,
		},
		&Route{
			Path:    "/categories/{category_id}",
			Method:  http.MethodPut,
			Handler: r.categoriesController.PutCategory,
		},
	}
}