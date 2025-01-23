package main

import (
	"net/http"

	"github.com/half-blood-prince-2710/blog/server-go/internals/data"
)

func (app *application)  getAllBlogsHandler(w http.ResponseWriter, r *http.Request){
	
}

func (app *application)  createBloghandler(w http.ResponseWriter, r *http.Request){
	blog := &data.Blog{
		ID: 1,
		Title: "A good boy",
		Body:"Manish gupta is the best boy",
		Author: "Manish Gupta"

	}
	

}

func (app *application)  updateBlogHandler(w http.ResponseWriter, r *http.Request){
	
}

func (app *application)  deleteBlogHanlder(w http.ResponseWriter, r *http.Request){
	
}

