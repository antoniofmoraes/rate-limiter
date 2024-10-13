package middlewares

import (
	"net"
	"net/http"

	"github.com/antoniofmoraes/rate-limiter/internals/services"
)

func RateLimiterMiddleware(s *services.RateLimiterService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var allowed bool
		var err error

		token := r.Header.Get("API_KEY")
		if token != "" {
			allowed, err = s.IsAllowed(token, false)
		} else {
			allowed, err = s.IsAllowed(getIp(r), true)
		}

		if !allowed {
			http.Error(w, err.Error(), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getIp(r *http.Request) string {
	realIp := r.Header.Get("X-Real-Ip")
	if realIp != "" {
		return realIp
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	return ip
}
