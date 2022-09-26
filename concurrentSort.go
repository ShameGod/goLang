package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Please enter a list of integers separated with a white space")
	br := bufio.NewReader(os.Stdin)
	input, _, _ := br.ReadLine()
	s := strings.Split(string(input), " ")
	var numbers []int

	for _, v := range s {
		intV, _ := strconv.Atoi(v)
		numbers = append(numbers, intV)
	}

	m := int64(len(numbers) / 4)

	var wg = sync.WaitGroup{}
	sortedSlices := make(map[int][]int)

	wg.Add(4)

	go sortSlice(numbers[0:m], sortedSlices, &wg)
	go sortSlice(numbers[m:2*m], sortedSlices, &wg)
	go sortSlice(numbers[2*m:3*m], sortedSlices, &wg)
	go sortSlice(numbers[3*m:], sortedSlices, &wg)
	wg.Wait()

	sortMapAndPrint(sortedSlices)
}

func sortSlice(slice []int, sortedSlices map[int][]int, wg *sync.WaitGroup) {
	sort.Ints(slice)
	sortedSlices[slice[0]] = slice
	wg.Done()
}

func sortMapAndPrint(mapToSort map[int][]int) {
	slice := make([]int, 4)
	var finalSlice []int

	for k, _ := range mapToSort {
		slice = append(slice, k)
	}

	sort.Ints(slice)

	for _, v := range slice {
		finalSlice = append(finalSlice, mapToSort[v]...)
	}

	fmt.Print(finalSlice)

}
