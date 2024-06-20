package routes

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/gumanzanoo/email-service-provider/internal/exceptions"
	"net/http"
)

type RouteFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func ErrorHandler(routeFunc RouteFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := routeFunc(w, r)
		if err != nil {
			if errors.Is(err, exceptions.InternalErr) {
				render.Status(r, 500)
			} else {
				render.Status(r, 400)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, status)
		if obj != nil {
			render.JSON(w, r, obj)
		}
	}
}
