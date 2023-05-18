package middleware

import (
	"net/http"
)

func SignatureChecker(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		//TODO: write logic here
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

//func Middleware(c context.Context) {
//	authHeader := c.GetHeader("Authorization")
//	if authHeader == "" {
//		c.AbortWithStatus(http.StatusUnauthorized)
//		return
//	}
//
//	headerParts := strings.Split(authHeader, " ")
//	if len(headerParts) != 2 {
//		c.AbortWithStatus(http.StatusUnauthorized)
//		return
//	}
//
//	if headerParts[0] != "Bearer" {
//		c.AbortWithStatus(http.StatusUnauthorized)
//		return
//	}
//
//	err := parser.ParseToken(headerParts[1], SIGNING_KEY)
//	if err != nil {
//		status := http.StatusBadRequest
//		if err == auth.ErrInvalidAccessToken {
//			status = http.StatusUnauthorized
//		}
//
//		c.AbortWithStatus(status)
//		return
//	}
//}
