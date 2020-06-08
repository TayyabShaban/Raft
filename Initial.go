package Raft

import (
	"time"
)

//node struct
type node struct {
	name      string
	state     string
	term      int
	totalNode int
	dataValue string
	didVote   bool
	timer     time.Time
}

func main() {
	//main

}
