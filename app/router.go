package app

import (
	"blogspot/constant"
	"blogspot/flow"
	"blogspot/handler"
	"blogspot/parser"
	"blogspot/repository"
	"blogspot/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var jsonPresenter = util.NewJsonPresenter()
var db = repository.NewConnectMysqlDb()

var userParser = parser.NewUserParser()
var userRepo = repository.NewUserRepository(db)
var userFlow = flow.NewUserFlow(userRepo)
var userHandler = handler.NewUserHandler(userParser, jsonPresenter, userFlow)

var tagParser = parser.NewTagParser()
var tagRepo = repository.NewTagRepository(db)
var tagFlow = flow.NewTagFlow(tagRepo)
var tagHandler = handler.NewTagHandler(tagParser, jsonPresenter, tagFlow)

var categoryParser = parser.NewCategoryParser()
var categoryRepo = repository.NewCategoryRepository(db)
var categoryFlow = flow.NewCategoryFlow(categoryRepo)
var categoryHandler = handler.NewCategoryHandler(categoryParser, jsonPresenter, categoryFlow)

var blogParser = parser.NewBlogParser()
var blogRepo = repository.NewBlogRepository(db)
var blogFlow = flow.NewBlogFlow(blogRepo, tagRepo, categoryRepo)
var blogHandler = handler.NewBlogHandler(blogParser, jsonPresenter, blogFlow)

func UserTypeMiddleware(userTypes []int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check user access rights based on the userTypes parameter
		// user := getCurrentUser(r) // Implement your logic to get the current user

		// Check if the user has one of the specified user types
		if checkUserType(userTypes, r) {
			// User has a valid user type, call the next handler
			next.ServeHTTP(w, r)
		} else {
			// User does not have a valid user type, return an unauthorized response
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func checkUserType(access []int, r *http.Request) bool {
	var userType string

	if len(access) == 0 {
		return true
	}

	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "userType" {
			userType = cookie.Value
		}
	}

	if userType == "" {
		return false
	}

	userTypeInt, _ := strconv.Atoi(userType)

	for _, accesses := range access {
		if accesses == userTypeInt {
			return true
		}
	}

	return false
}

func SetupRouter() http.Handler {
	r := mux.NewRouter()

	//User Routes
	r.Handle("/user", UserTypeMiddleware([]int{}, http.HandlerFunc(userHandler.Register))).Methods("POST")
	r.Handle("/user", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(userHandler.Update))).Methods("PUT")
	r.Handle("/login", UserTypeMiddleware([]int{}, http.HandlerFunc(userHandler.Login))).Methods("POST")
	r.Handle("/logout", UserTypeMiddleware([]int{constant.AdminType, constant.UserType}, http.HandlerFunc(userHandler.Logout))).Methods("POST")

	//Blog Routes
	r.Handle("/blog/{id}", UserTypeMiddleware([]int{}, http.HandlerFunc(blogHandler.GetBlog))).Methods("GET")
	r.Handle("/user/{id}", UserTypeMiddleware([]int{}, http.HandlerFunc(blogHandler.GetByUser))).Methods("GET")
	r.Handle("/hidden/blog", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(blogHandler.GetHidden))).Methods("GET")
	r.Handle("/blog", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(blogHandler.Insert))).Methods("POST")
	r.Handle("/blog", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(blogHandler.Update))).Methods("PUT")
	r.Handle("/blog/hidden/{id}", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(blogHandler.Hidden))).Methods("PUT")
	r.Handle("/blog/{id}", UserTypeMiddleware([]int{constant.UserType}, http.HandlerFunc(blogHandler.Delete))).Methods("DELETE")

	//Category Routes
	r.Handle("/category", UserTypeMiddleware([]int{constant.AdminType}, http.HandlerFunc(categoryHandler.Insert))).Methods("POST")

	//Tag Routes
	r.Handle("/tag", UserTypeMiddleware([]int{constant.AdminType}, http.HandlerFunc(tagHandler.Insert))).Methods("POST")

	return r
}
