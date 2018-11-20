package main

import (
	"container/heap"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type IntHeap []int

type LatencyData struct {
	Latency int `json:"latency"`
}

var db = &IntHeap{}

const InterestedCount = 10

func init() {
	heap.Init(db)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/latency", GetMaxLatency).Methods("GET")
	router.HandleFunc("/api/v1/latency", AddLatencyData).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

// AddLatencyData will check if the incoming latency falls under the top InterestedCount
// and adds it to the data bucket
func AddLatencyData(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var ld LatencyData
	err := decoder.Decode(&ld)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	elapsedSeconds := ld.Latency
	if db.Len() < InterestedCount {
		heap.Push(db, elapsedSeconds)
	} else {
		if (*db)[0] < elapsedSeconds {
			heap.Pop(db)
			heap.Push(db, elapsedSeconds)
		}
	}

	w.WriteHeader(http.StatusCreated)
	return

}

// Return the elements in the data bucket (db)
func GetMaxLatency(w http.ResponseWriter, r *http.Request) {
	jsonInfo, err := json.Marshal(struct{Latencies []int `json:"latencies"`}{*db})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInfo)
	return

}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (db *IntHeap) Push(x interface{}) {
	*db = append(*db, x.(int))
}

func (db *IntHeap) Pop() interface{} {
	old := *db
	n := len(old)
	x := old[n-1]
	*db = old[0 : n-1]
	return x
}
