package main

import "net/http"

// All routes will be defined here
func (app *application) routes() http.Handler {
	 
	// creating serve mux
	mux := http.NewServeMux()

	// healthcheck routes
	mux.HandleFunc("GET /v1/healthcheck",app.healthcheck)

	//blog routes
	mux.HandleFunc("GET /v1/blogs",app.getAllBlogsHandler)
	mux.HandleFunc("POST /v1/blogs",app.createBloghandler)
	mux.HandleFunc("PUT /v1/blogs/{id}",app.updateBlogHandler)
	mux.HandleFunc("DELETE /v1/blogs/{id}",app.deleteBlogHanlder)



	return mux

}

