package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"testing"
)

var l []int

func mergeSortAsync1000s(l []int, c chan []int) {
	defer close(c)
	if len(l) < 3 {
		if len(l) == 1 {
			c <- []int{l[0]}
			return
		} else if l[0] < l[1] {
			c <- []int{l[0], l[1]}
			return
		} else {
			c <- []int{l[1], l[0]}
			return
		}
	}
	mid := len(l) / 2
	c1 := make(chan []int, 1)
	c2 := make(chan []int, 1)
	go mergeSortAsync1000s(l[:mid], c1)
	go mergeSortAsync1000s(l[mid:], c2)
	c <- merge(<-c1, <-c2)
}

func mergeSort(l []int) []int {
	if len(l) < 3 {
		if len(l) == 1 {
			return []int{l[0]}
		} else if l[0] < l[1] {
			return []int{l[0], l[1]}
		} else {
			return []int{l[1], l[0]}
		}
		/*
		   r := make([]int, len(l))
		   smallest := 0
		   for a:=1; a<len(l);a++ {
		     if l[a] < l[smallest] {}
		   }*/
	}
	mid := len(l) / 2
	a := mergeSort(l[:mid])
	b := mergeSort(l[mid:])
	return merge(a, b)
}

func qs1000async(l []int, out chan []int) {
	mid := len(l) >> 1
	if mid > 100 {
		c1 := make(chan []int, 1)
		c2 := make(chan []int, 1)
		go qs1000async(l[:mid], c1)
		go qs1000async(l[mid:], c2)
		out <- merge(<-c1, <-c2)
	} else {
		sort.Ints(l)
		out <- l
	}
}
func qs1000async3(l []int, ch chan bool) {
	mid := len(l) >> 1
	if mid > 1000 {
		c1 := make(chan bool)
		c2 := make(chan bool)
		go qs1000async3(l[:mid], c1)
		go qs1000async3(l[mid:], c2)
		<-c1
		<-c2
	}
	sort.Ints(l)
	close(ch)
}

func qs1000async6(l []int, ch chan bool) {
	mid := len(l) >> 1
	if mid > 10000 {
		c1 := make(chan bool)
		go qs1000async6(l[:mid], c1)
		qs1000async6Lower(l[mid:])
		<-c1
		mergeInPlace3(l, mid)
	} else {
		sort.Ints(l)
	}
	close(ch)
}
func qs1000async6Lower(l []int) {
	mid := len(l) >> 1
	if mid > 10000 {
		c1 := make(chan bool)
		go qs1000async6(l[:mid], c1)
		qs1000async6(l[mid:], make(chan bool))
		<-c1
		mergeInPlace3(l, mid)
	} else {
		sort.Ints(l)
	}
}

func qs1000async5(l []int, ch chan bool) {
	mid := len(l) >> 1
	if mid > 1000 {
		c1 := make(chan bool)
		go qs1000async5(l[:mid], c1)
		qs1000async5(l[mid:], make(chan bool))
		<-c1
		mergeInPlace3(l, mid)
	} else {
		sort.Ints(l)
	}
	close(ch)
}

func qs1000async4(l []int, ch chan bool) {
	mid := len(l) >> 1
	if mid > 1000 {
		c1 := make(chan bool)
		go qs1000async4(l[:mid], c1)
		qs1000async4(l[mid:], make(chan bool))
		<-c1
		mergeInPlace(l, mid)
	} else {
		sort.Ints(l)
	}
	close(ch)
}

func qs1000async2(l []int, out chan []int) {
	counter := 0
	toMerge := make(chan []int, 2)
	for a := 0; a < len(l); a += 1000 {
		b := a + 1000
		if b > len(l) {
			b = len(l)
		}
		counter++
		go func(a, b int) {
			sort.Ints(l[a:b])
			toMerge <- l[a:b]
		}(a, b)
	}

	for counter > 1 {
		a, b := <-toMerge, <-toMerge
		go func(a, b []int) {
			toMerge <- merge(a, b)
		}(a, b)
		counter--
	}
	tmp := <-toMerge
	out <- tmp
}

func merge(left, right []int) []int {
	var i, j int
	result := make([]int, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		result[i+j] = left[i]
		i++
	}
	for j < len(right) {
		result[i+j] = right[j]
		j++
	}
	return result
}

func mergeInPlace(l []int, j int) {
	all := len(l)
	third := make([]int, 0, 250)
	var tmp int
	for i := 0; i < all; i++ {
		if l[i] <= l[j] && (len(third) == 0 || l[i] < third[0]) {
			continue
		}
		if l[j] < l[i] && (len(third) == 0 || l[j] < third[0]) {
			tmp, l[i] = l[i], l[j]
			//addToThird
			for idx := 0; idx < len(third); idx++ {
				if tmp < third[idx] {
					third = append(third[:idx], append([]int{tmp}, third[idx:]...)...)
					continue
				}
			}
			third = append(third, tmp)
			continue
		}
		tmp, l[i], third = l[i], third[0], third[1:]
	}
}

func mergeInPlace3(l []int, j int) {
	all := len(l)
	third := make([]int, 0, 150)
	var tmp int
	var idx int
	var lenthird int
	for i := 0; i < all; i++ {
		if l[i] <= l[j] && (len(third) == 0 || l[i] < third[0]) {
			continue
		}
		if l[j] < l[i] && (len(third) == 0 || l[j] < third[0]) {
			tmp, l[i] = l[i], l[j]
			//addToThird
			lenthird = len(third)
			for idx = 0; idx < lenthird; idx++ {
				if tmp < third[idx] {
					// found index of addition
					for ; idx < lenthird; idx++ {
						tmp, third[idx] = third[idx], tmp
					}
					third = append(third, tmp)
					continue
				}
			}
			third = append(third, tmp)
			continue
		}
		tmp, l[i] = l[i], third[0]
		lenthird = len(third)
		for idx = 1; idx < lenthird; idx++ {
			third[idx-1] = third[idx]
		}
		third = third[:lenthird-1]
	}
}

func mergeInPlace2(l []int, j int) {
	all := len(l)
	third := make([]int, 0, 50)
	var tmp int
	for i := 0; i < all; i++ {
		if l[i] <= l[j] && (len(third) == 0 || l[i] < third[0]) {
			continue
		}
		if l[j] < l[i] && (len(third) == 0 || l[j] < third[0]) {
			tmp, l[i] = l[i], l[j]
			//addToThird
			third = append(third, tmp)
			sort.Ints(third)
			continue
		}
		tmp, l[i], third = l[i], third[0], third[1:]
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func load() {
	l = make([]int, 10000000)
	lines, err := readLines("arr.txt") //your array file
	for i, line := range lines {
		if i <= len(l) {
			l[i], err = strconv.Atoi(line)
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

/*
func aBenchmarkMergesort(b *testing.B) {
	load()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mergeSort(l)
	}
}

func BenchmarkMergesort1000s(b *testing.B) {
	c := make(chan []int, 1)
	load()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mergeSortAsyncDivide(l, c)
		<-c
	}
}
*/
func BenchmarkQuicksort(b *testing.B) {
	load()
	b.ResetTimer()
	tmp := make([]int, len(l))
	for n := 0; n < b.N; n++ {
		for i, a := range l {
			tmp[i] = a
		}
		sort.Ints(tmp)
	}
}
func BenchmarkQuicksort1000Async4(b *testing.B) {
	load()
	b.ResetTimer()
	tmp := make([]int, len(l))
	for n := 0; n < b.N; n++ {
		for i, a := range l {
			tmp[i] = a
		}
		qs1000async4(tmp, make(chan bool))
	}
}
func BenchmarkQuicksort1000Async5(b *testing.B) {
	load()
	b.ResetTimer()
	tmp := make([]int, len(l))
	for n := 0; n < b.N; n++ {
		for i, a := range l {
			tmp[i] = a
		}
		qs1000async5(tmp, make(chan bool))
	}
}
func BenchmarkQuicksort1000Async6(b *testing.B) {
	load()
	b.ResetTimer()
	tmp := make([]int, len(l))
	for n := 0; n < b.N; n++ {
		for i, a := range l {
			tmp[i] = a
		}
		qs1000async6(tmp, make(chan bool))
	}
}
