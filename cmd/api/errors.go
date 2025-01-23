package main

import "net/http"

func (app *application) logError (r *http.Request, err error){
	app.logger.Error(err.Error(),"method",r.Method,"url",r.URL.RequestURI())
}

func (app *application) errorResponse(w http.ResponseWriter,r *http.Request, status int, message string){
	
	env := envelope{"error":message}

	err := app.writeJSON(w,status,env,nil)
	if err!=nil{
		app.logError(r,err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter,r *http.Request, err error){
	app.logError(r,err)
	message:= "Server encountered a problem could not process your request"
	app.errorResponse(w,r,http.StatusInternalServerError,message)
}


func (app *application) notFoundResponse(w http.ResponseWriter,r *http.Request){
	message:= "Content Not found"
	app.errorResponse(w,r,http.StatusNotFound,message)
}
func (app *application) badRequestResponse(w http.ResponseWriter,r *http.Request){
	message:= "A bad request"
	app.errorResponse(w,r,http.StatusBadRequest,message)
}
