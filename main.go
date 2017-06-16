package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main(){
    source := getSourceList(10)

    fmt.Println("bubble sort:", bubbleSort(source))
    fmt.Println("insert sort:", insertSort(source))
}


func getSourceList(num int) (arr []int){
    rand.Seed(time.Now().Unix())
    for i:=0; i< num; i++ {
        arr = append(arr, rand.Intn(30))
    }
    fmt.Println("source:", arr)
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

    return arr
}


func insertSort(arr []int) []int {
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

    return arr
}


func selectionSort(arr []int) []int {

}