package middleware

//
//import (
//	"context"
//	"github.com/waretool/go-common/domain"
//	"github.com/waretool/go-common/logger"
//	"github.com/waretool/go-common/service"
//	"net/http"
//	"regexp"
//)
//
//func AuthZ(requiredRole domain.Role) domain.Middleware {
//	return func(h http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			claims := r.Context().Value("claims").(domain.Claims)
//			if claims == nil {
//				w.WriteHeader(http.StatusUnauthorized)
//				return
//			}
//
//			if role, ok := claims["subject"]; !ok {
//				logger.Infof("missing required claims in jwt")
//				w.WriteHeader(http.StatusForbidden)
//				return
//			}
//
//			h.ServeHTTP(w, r)
//
//			re := regexp.MustCompile(`(?i)bearer`)
//			jwt := re.ReplaceAllString(header[0], "")
//			if claims, valid := service.NewJwtService().Valid(jwt); !valid {
//				w.WriteHeader(http.StatusUnauthorized)
//				return
//			} else {
//				ctx := context.WithValue(r.Context(), "claims", claims)
//				h.ServeHTTP(w, r.WithContext(ctx))
//			}
//		})
//	}
//}
