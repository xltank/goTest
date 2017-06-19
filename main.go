package main

import (
    "fmt"
    "math/rand"
    "time"
)

const num = 100000

func main(){
    source := getSourceList(num)
    //source := []int{5,1,2,3,4,4,5}

    bubbleSort(source)
    insertSort(source)
    selectionSort(source)
    quickSort(deepCopySlice(source))
    mergeSort(deepCopySlice(source))
}


func getSourceList(num int) (arr []int){
    t := time.Now();
    rand.Seed(time.Now().UnixNano())
    for i:=0; i< num; i++ {
        arr = append(arr, rand.Intn(num))
    }
    fmt.Println("source:        ", arr, "---", time.Now().Sub(t).Nanoseconds())
    return
}

func deepCopySlice(arr []int) []int{
    l := len(arr)
    result := make([]int, l, l)
    for i:=0; i<l; i++ {
        result[i] = arr[i]
    }
    return result
}


func bubbleSort(arr []int) ([]int) {
    t:= time.Now()
    arr = deepCopySlice(arr)
    //fmt.Println(arr)
    length := len(arr)
    for i:=0; i<length; i++ {
        for j:=0; j< length-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println("Bubble sort:", arr, "---", time.Now().Sub(t).Nanoseconds())
    return arr
}


func insertSort(arr []int) []int {
    t := time.Now()
    arr = deepCopySlice(arr)
    //fmt.Println(arr)
    length := len(arr)
    for i:=1; i<length; i++{
        for j:=i; j>0; j-- {
            if arr[j] < arr[j-1] {
                arr[j], arr[j-1] = arr[j-1], arr[j]
            }else{
                break
            }
        }
    }
    fmt.Println("Insert sort:", arr, "---", time.Now().Sub(t).Nanoseconds())
    return arr
}


func selectionSort(arr []int) []int {
    t := time.Now()
    arr = deepCopySlice(arr)
    length := len(arr)
    for i:=0; i< length-1; i++ {
        minIndex := i
        for j:=i+1; j< length; j++ {
            if arr[minIndex] > arr[j] {
                minIndex = j
            }
        }
        arr[minIndex], arr[i] = arr[i], arr[minIndex]
    }
    fmt.Println("Selection sort:", arr, "---", time.Now().Sub(t).Nanoseconds())
    return arr
}

func quickSort(arr []int) []int {
    t := time.Now()
    doQuickSort(arr, 0, len(arr)-1)
    fmt.Println("Quick sort:", arr, "---", time.Now().Sub(t).Nanoseconds())
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
        for ; j> i; j-- {
            if arr[j] < std {
                arr[index] = arr[j]
                index = j
                break
            }
        }
        for ; i< j; i++ {
            if arr[i] > std {
                arr[index] = arr[i]
                index = i
                break
            }
        }
    }
    arr[index] = std

    //fmt.Println(index)
    if start < index {
        doQuickSort(arr, start, index)
    }
    if index < end {
        doQuickSort(arr, index+1, end)
    }
}


func mergeSort(arr []int) []int {
    t:= time.Now()
    doMergeSort(arr, 0, len(arr)-1)
    fmt.Println("Merge sort:", arr, "---", time.Now().Sub(t).Nanoseconds())
    return arr
}
func doMergeSort(arr []int, start int, end int) []int{
    if start >= end {
        return arr
    }

    mid := (start + end)/2

    doMergeSort(arr, start, mid)
    doMergeSort(arr, mid+1, end)
    mergeSlices(arr, start, mid, end)

    return arr
}
func mergeSlices(arr []int, start int, mid int, end int){
    length := end-start+1
    //fmt.Println(start, mid, end, length)
    temp := make([]int, length, length)
    i:= start
    j:= mid+1
    for index:=0; index<length; index++ {
        //fmt.Println("index, i, j:", index, i, j)
        if i<= mid && (j > end || arr[i] <= arr[j]) {
            temp[index] = arr[i]
            i++
        }else if j<= end && (i > mid || arr[i] > arr[j]) {
            temp[index] = arr[j]
            j++
        }
    }

    for k:=0; k<length; k++{
        arr[start+k] = temp[k]
    }
}