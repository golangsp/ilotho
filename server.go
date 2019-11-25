package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type winner struct {
	Name string `json:"name,omitempty"`
}

func main() {
	http.HandleFunc("/raffle", raffle)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func raffle(rw http.ResponseWriter, r *http.Request) {

	part := strings.Split(strings.Replace(r.FormValue("participants"), "\r\n", "\n", -1), "\n")
	raf := strings.Split(r.FormValue("raffled"), ",")
	win := winner{}

	part = difference(part, raf)

	if len(part) > 0 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(part), func(i, j int) { part[i], part[j] = part[j], part[i] })
		win.Name = part[rand.Intn(len(part))]
		fmt.Printf("O ganhador foi... %s!!!\n", strings.ToUpper(win.Name))
	}

	for i := 0; i < 2; i++ {
		time.Sleep(1 * time.Second)
	}

	json.NewEncoder(rw).Encode(win)
}

func difference(a, b []string) []string {

	for _, ea := range b {
		for i, eb := range a {
			if ea == eb {
				// Remove the element at index i from a.
				copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
				a[len(a)-1] = ""     // Erase last element (write zero value).
				a = a[:len(a)-1]     // Truncate slice.
				break
			}
		}
	}

	return a
}
