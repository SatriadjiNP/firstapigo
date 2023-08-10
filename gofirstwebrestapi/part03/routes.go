package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"findAllCategories",
		"GET",
		"/categories",
		findAllCategoriesHandler,
	},
	Route{
		"findCategory",
		"GET",
		"/category/{id}",
		findCategoryHandler,
	},
	Route{
		"AddCategoryHandler",
		"POST",
		"/category",
		AddCategoryHandler,
	},
	Route{
		"DeleteCategoryHandler",
		"DELETE",
		"/category/{id}",
		DeleteCategoryHandler,
	},
	Route{
		"UpdateCategoryHandler",
		"POST",
		"/category/{id}",
		UpdateCategoryHandler,
	},
}
