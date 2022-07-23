package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ruang-arah/backend/model/web"
	"ruang-arah/backend/pkg/service"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIdCtxKey        = "user_id"
	roleCtxKey          = "role"
)

var (
	roleVal   string
	userIdVal int32
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(web.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    "Need Authorization header",
			})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(web.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    "Invalid Authorization header",
			})
			return
		}

		userId, userRole, err := service.ParseToken(r.Context(), headerParts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(web.WebResponse{
				Code:    http.StatusUnauthorized,
				Message: "UNAUTHORIZED",
				Data:    err.Error(),
			})
			return
		}

		newCtx := context.WithValue(r.Context(), userIdCtxKey, userId)
		ctx := context.WithValue(newCtx, roleCtxKey, userRole)

		roleVal = ctx.Value(roleCtxKey).(string)
		userIdVal = ctx.Value(userIdCtxKey).(int32)

		next.ServeHTTP(w, r)
	})
}

func AdminMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if roleVal != "admin" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(web.WebResponse{
				Code:    http.StatusForbidden,
				Message: http.StatusText(http.StatusForbidden),
				Data:    "Forbidden access",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetUserId() (int32, error) {
	if userIdVal == 0 {
		return 0, fmt.Errorf("user id is not set")
	}

	return userIdVal, nil
}
