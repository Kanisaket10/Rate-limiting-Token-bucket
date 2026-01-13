package main

import(
	"golang.org/x/time/rate"
	"net/http"
)

func rateLimter(next http.HandlerFunc) http.HandlerFunc {
	limiter := rate.NewLimiter(2, 4)
	return func(writer http.ResponseWriter, request *http.Request) {
		if !limiter.Allow() {
			message := Message{
				Status: "Request Failed",
				Body:   "Too many requests, try again later.",	
			}
			 w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(writer).Encode(&message)
			return
		}else {
			next(writer, request)
		}
	}
}