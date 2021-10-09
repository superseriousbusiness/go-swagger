// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UploadFileHandlerFunc turns a function with the right signature into a upload file handler
type UploadFileHandlerFunc func(UploadFileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadFileHandlerFunc) Handle(params UploadFileParams) middleware.Responder {
	return fn(params)
}

// UploadFileHandler interface for that can handle valid upload file params
type UploadFileHandler interface {
	Handle(UploadFileParams) middleware.Responder
}

// NewUploadFile creates a new http.Handler for the upload file operation
func NewUploadFile(ctx *middleware.Context, handler UploadFileHandler) *UploadFile {
	return &UploadFile{Context: ctx, Handler: handler}
}

/* UploadFile swagger:route POST /upload uploads uploadFile

uploads

*/
type UploadFile struct {
	Context *middleware.Context
	Handler UploadFileHandler
}

func (o *UploadFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUploadFileParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
