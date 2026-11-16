package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	redisAddr := os.Getenv("REDIS_ADDR")

	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	defer rdb.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestIncrementor(w, r, rdb)
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestIncrementor(w http.ResponseWriter, r *http.Request, rdb *redis.Client) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	requestCount, errRedis := rdb.Incr(r.Context(), "requestCount").Result()

	if errRedis != nil {
		log.Println("An error occurred while setting the request count in Redis", errRedis)
	}

	_, errPrint := fmt.Fprintln(w, "Request count: ", requestCount)

	if errPrint != nil {
		log.Println("An error occured while responding", errPrint)
		return
	}

	log.Println("Responded successfully, request count is now: ", requestCount)
}
