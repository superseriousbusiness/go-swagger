package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/casualjim/go-swagger/swagger/httputils"
	"github.com/casualjim/go-swagger/swagger/testing/petstore"
	"github.com/gorilla/context"
	"github.com/stretchr/testify/assert"
)

func TestContextRender(t *testing.T) {
	ct := httputils.JSONMime
	ctx := NewContext(petstore.NewAPI(t))
	request, _ := http.NewRequest("GET", "http://localhost:8080/api/pets", nil)
	request.Header.Set(httputils.HeaderAccept, ct)

	recorder := httptest.NewRecorder()
	ctx.Respond(recorder, request, []string{ct}, map[string]interface{}{"name": "hello"})
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "{\"name\":\"hello\"}\n", recorder.Body.String())

	recorder = httptest.NewRecorder()
	ctx.Respond(recorder, request, []string{ct}, errors.New("this went wrong"))
	assert.Equal(t, 500, recorder.Code)

	recorder = httptest.NewRecorder()
	assert.Panics(t, func() { ctx.Respond(recorder, request, []string{ct}, map[int]interface{}{1: "hello"}) })

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "http://localhost:8080/api/pets", nil)
	assert.Panics(t, func() { ctx.Respond(recorder, request, []string{}, map[string]interface{}{"name": "hello"}) })

}

func TestContextValidResponseFormat(t *testing.T) {
	ct := "application/json"
	ctx := NewContext(nil, nil)

	request, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Set(httputils.HeaderAccept, ct)

	// check there's nothing there
	cached, ok := context.GetOk(request, ctxResponseFormat)
	assert.False(t, ok)
	assert.Empty(t, cached)

	// trigger the parse
	mt := ctx.ResponseFormat(request, []string{ct})
	assert.Equal(t, ct, mt)

	// check it was cached
	cached, ok = context.GetOk(request, ctxResponseFormat)
	assert.True(t, ok)
	assert.Equal(t, ct, cached)

	// check if the cast works and fetch from cache too
	mt = ctx.ResponseFormat(request, []string{ct})
	assert.Equal(t, ct, mt)
}

func TestContextInvalidResponseFormat(t *testing.T) {
	ct := "application/x-yaml"
	other := "application/xml"
	ctx := NewContext(nil, nil)

	request, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Set(httputils.HeaderAccept, ct)

	// check there's nothing there
	cached, ok := context.GetOk(request, ctxResponseFormat)
	assert.False(t, ok)
	assert.Empty(t, cached)

	// trigger the parse
	mt := ctx.ResponseFormat(request, []string{other})
	assert.Empty(t, mt)

	// check it was cached
	cached, ok = context.GetOk(request, ctxResponseFormat)
	assert.False(t, ok)
	assert.Empty(t, cached)

	// check if the cast works and fetch from cache too
	mt = ctx.ResponseFormat(request, []string{other})
	assert.Empty(t, mt)
}

func TestContextValidRoute(t *testing.T) {
	ctx := NewContext(petstore.NewAPI(t))
	request, _ := http.NewRequest("GET", "http://localhost:8080/api/pets", nil)

	// check there's nothing there
	_, ok := context.GetOk(request, ctxMatchedRoute)
	assert.False(t, ok)

	matched, ok := ctx.RouteInfo(request)
	assert.True(t, ok)
	assert.NotNil(t, matched)

	// check it was cached
	_, ok = context.GetOk(request, ctxMatchedRoute)
	assert.True(t, ok)

	matched, ok = ctx.RouteInfo(request)
	assert.True(t, ok)
	assert.NotNil(t, matched)
}

func TestContextInvalidRoute(t *testing.T) {
	ctx := NewContext(petstore.NewAPI(t))
	request, _ := http.NewRequest("DELETE", "http://localhost:8080/api/pets", nil)

	// check there's nothing there
	_, ok := context.GetOk(request, ctxMatchedRoute)
	assert.False(t, ok)

	matched, ok := ctx.RouteInfo(request)
	assert.False(t, ok)
	assert.Nil(t, matched)

	// check it was cached
	_, ok = context.GetOk(request, ctxMatchedRoute)
	assert.False(t, ok)

	matched, ok = ctx.RouteInfo(request)
	assert.False(t, ok)
	assert.Nil(t, matched)
}

func TestContextValidContentType(t *testing.T) {
	ct := "application/json"
	ctx := NewContext(nil, nil)

	request, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Set(httputils.HeaderContentType, ct)

	// check there's nothing there
	_, ok := context.GetOk(request, ctxContentType)
	assert.False(t, ok)

	// trigger the parse
	mt, _, err := ctx.ContentType(request)
	assert.NoError(t, err)
	assert.Equal(t, ct, mt)

	// check it was cached
	_, ok = context.GetOk(request, ctxContentType)
	assert.True(t, ok)

	// check if the cast works and fetch from cache too
	mt, _, err = ctx.ContentType(request)
	assert.NoError(t, err)
	assert.Equal(t, ct, mt)
}

func TestContextInvalidContentType(t *testing.T) {
	ct := "application("
	ctx := NewContext(nil, nil)

	request, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	request.Header.Set(httputils.HeaderContentType, ct)

	// check there's nothing there
	_, ok := context.GetOk(request, ctxContentType)
	assert.False(t, ok)

	// trigger the parse
	mt, _, err := ctx.ContentType(request)
	assert.Error(t, err)
	assert.Empty(t, mt)

	// check it was not cached
	_, ok = context.GetOk(request, ctxContentType)
	assert.False(t, ok)

	// check if the failure continues
	_, _, err = ctx.ContentType(request)
	assert.Error(t, err)
}
