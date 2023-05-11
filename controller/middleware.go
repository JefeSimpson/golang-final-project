package controller

import (
	"context"
	"net/http"
	"strings"
)

func (c *Controller) middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 {
			w.WriteHeader(401)
			return
		}

		username, err := c.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			w.WriteHeader(401)
			return
		}

		newToken, err := c.services.Authorization.RefreshToken(headerParts[1])
		if err != nil {
			w.WriteHeader(401)
			return
		}

		w.Header().Set("Authorization", newToken)

		ctx := context.WithValue(r.Context(), "username", username)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
