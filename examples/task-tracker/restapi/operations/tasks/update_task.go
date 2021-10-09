// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateTaskHandlerFunc turns a function with the right signature into a update task handler
type UpdateTaskHandlerFunc func(UpdateTaskParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateTaskHandlerFunc) Handle(params UpdateTaskParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateTaskHandler interface for that can handle valid update task params
type UpdateTaskHandler interface {
	Handle(UpdateTaskParams, interface{}) middleware.Responder
}

// NewUpdateTask creates a new http.Handler for the update task operation
func NewUpdateTask(ctx *middleware.Context, handler UpdateTaskHandler) *UpdateTask {
	return &UpdateTask{Context: ctx, Handler: handler}
}

/* UpdateTask swagger:route PUT /tasks/{id} tasks updateTask

Updates the details for a task.

Allows for updating a task.
This operation requires authentication so that we know which user
last updated the task.


*/
type UpdateTask struct {
	Context *middleware.Context
	Handler UpdateTaskHandler
}

func (o *UpdateTask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateTaskParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
