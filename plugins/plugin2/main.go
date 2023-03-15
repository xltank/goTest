package main

import "fmt"

func main() {
	fmt.Println("plugins/plugin1/main.go/main()")
}

func init() {
	fmt.Println("plugins/plugin2/main.go/init()")
}

func Go() {
	fmt.Println("plugins/plugin2/main.go/Go()")
}
