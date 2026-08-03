package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

// These thresholds are arbitrary demo values, not figures taken from the
// source study guide (see docs/study-guide/11-rate-limiting.md and
// docs/04-limitations.md).
const limit = 10
const window = time.Minute

func rateLimited(clientID string) (bool, int64) {
	key := "ratelimit:" + clientID
	count, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return false, 0
	}
	if count == 1 {
		rdb.Expire(ctx, key, window)
	}
	return count > limit, count
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	clientID := r.Header.Get("X-Client-Id")
	if clientID == "" {
		clientID = r.RemoteAddr
	}

	limited, count := rateLimited(clientID)
	if limited {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "429 Too Many Requests (count=%d, limit=%d)\n", count, limit)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "200 OK (count=%d, limit=%d)\n", count, limit)
}

func main() {
	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})

	http.HandleFunc("/resource", resourceHandler)
	log.Println("rate-limiting demo listening on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
