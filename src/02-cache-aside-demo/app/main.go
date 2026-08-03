package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

// simulateSourceOfTruth mimics reading file metadata from a relational
// database, as described in docs/study-guide/10-caching-and-cdn.md.
func simulateSourceOfTruth(id string) map[string]string {
	time.Sleep(300 * time.Millisecond) // simulated database latency
	return map[string]string{
		"id":   id,
		"name": fmt.Sprintf("document-%s.txt", id),
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		id = "123"
	}
	key := "file:" + id

	cached, err := rdb.Get(ctx, key).Result()
	if err == nil {
		w.Header().Set("X-Cache", "HIT")
		w.Write([]byte(cached))
		return
	}

	data := simulateSourceOfTruth(id)
	payload, _ := json.Marshal(data)
	rdb.Set(ctx, key, payload, 5*time.Minute)

	w.Header().Set("X-Cache", "MISS")
	w.Write(payload)
}

func main() {
	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379"})

	http.HandleFunc("/file", fileHandler)
	log.Println("cache-aside demo listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
