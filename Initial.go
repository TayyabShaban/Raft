package main

import (
	"fmt"
	"strconv"
)

//node struct
type node struct {
	name      string
	state     string
	term      int
	totalNode int
	dataValue string
	didVote   bool
	//timer     time.Time
}

func main() {
	//main
	fmt.Println("main started")
	var nod [3]node
	//var i int =0
	for i := 0; i < 3; i++ {
		nod[i] = node{"node" + strconv.Itoa(i), "follower", 0, 3, "", false}
		fmt.Println(nod[i])
	}

}
