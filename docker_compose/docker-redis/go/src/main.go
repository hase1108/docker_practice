package main

import (
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World %d", countUp())

}

func countUp() int {

	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}

	count, err := redis.Int(c.Do("GET", "count"))
	plusCount := count
	plusCount++
	c.Do("SET", "count", plusCount)
	defer c.Close()
	return count
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
