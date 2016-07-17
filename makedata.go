package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	f, e := os.OpenFile("arr.txt", os.O_WRONLY|os.O_CREATE, 0777)
	fmt.Println(e)
	for a := 0; a < 99000; a++ {
		_, e = f.Write([]byte(strconv.Itoa(rand.Int())))
	}
}
