package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const num = 50000

func main() {
	fmt.Println("MaxProcs:", runtime.GOMAXPROCS(4)) // 8

	sequentialTest()
	concurrentTest()
}

func sequentialTest() {
	fmt.Println("\nsequentialTest ...")

	source := getSourceList(num)
	s1 := deepCopySlice(source)
	s2 := deepCopySlice(source)
	s3 := deepCopySlice(source)

	quickSort(s1, nil)
	mergeSort(s2, nil)
	insertSort(s3, nil)

	fmt.Println("\nGoroutine num:", runtime.NumGoroutine())

	//time.Sleep(time.Second*10)
}

func concurrentTest() {
	fmt.Println("\nconcurrentTest ...")
	wg := sync.WaitGroup{}

	source := getSourceList(num)

	s1 := deepCopySlice(source)
	s2 := deepCopySlice(source)
	s3 := deepCopySlice(source)

	wg.Add(1)
	go func() {
		quickSort(s1, nil)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		mergeSort(s2, nil)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		insertSort(s3, nil)
		wg.Done()
	}()

	fmt.Println("\nGoroutine num:", runtime.NumGoroutine())

	wg.Wait()
}

func printTime(s string, arr []int, t time.Time) {
	fmt.Printf("\n%v :", s)
	fmt.Printf(" %v", time.Now().Sub(t))
}

func getSourceList(num int) (arr []int) {
	t := time.Now()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		arr = append(arr, rand.Intn(num))
	}
	printTime("Source", arr, t)
	return
}

func deepCopySlice(arr []int) []int {
	l := len(arr)
	result := make([]int, l, l)
	for i := 0; i < l; i++ {
		result[i] = arr[i]
	}
	return result
}

func bubbleSort(arr []int, ch chan []int) []int {
	t := time.Now()
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	printTime("Bubble sort", arr, t)
	return arr
}

func insertSort(arr []int, ch chan []int) []int {
	t := time.Now()
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	printTime("Insertion sort", arr, t)
	return arr
}

func selectionSort(arr []int, ch chan []int) []int {
	t := time.Now()
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	printTime("Selection sort", arr, t)
	return arr
}

func quickSort(arr []int, ch chan []int) []int {
	t := time.Now()
	doQuickSort(arr, 0, len(arr)-1)
	printTime("Quick sort", arr, t)
	return arr
}
func doQuickSort(arr []int, start int, end int) {
	//fmt.Println(arr, start, end)
	if end <= start {
		return
	}

	std := arr[start]
	index := start
	i := start
	j := end
	for i < j {
		for ; j > i; j-- {
			if arr[j] < std {
				arr[index] = arr[j]
				index = j
				break
			}
		}
		for ; i < j; i++ {
			if arr[i] > std {
				arr[index] = arr[i]
				index = i
				break
			}
		}
	}
	arr[index] = std

	if start < index {
		doQuickSort(arr, start, index)
	}
	if index < end {
		doQuickSort(arr, index+1, end)
	}
}

func mergeSort(arr []int, ch chan []int) []int {
	t := time.Now()
	doMergeSort(arr, 0, len(arr)-1)
	printTime("Merge sort", arr, t)
	return arr
}
func doMergeSort(arr []int, start int, end int) []int {
	if start >= end {
		return arr
	}

	mid := (start + end) / 2

	doMergeSort(arr, start, mid)
	doMergeSort(arr, mid+1, end)
	mergeSlices(arr, start, mid, end)

	return arr
}
func mergeSlices(arr []int, start int, mid int, end int) {
	length := end - start + 1
	temp := make([]int, length, length)
	i := start
	j := mid + 1
	for index := 0; index < length; index++ {
		if i <= mid && (j > end || arr[i] <= arr[j]) {
			temp[index] = arr[i]
			i++
		} else if j <= end && (i > mid || arr[i] > arr[j]) {
			temp[index] = arr[j]
			j++
		}
	}

	for k := 0; k < length; k++ {
		arr[start+k] = temp[k]
	}
}
