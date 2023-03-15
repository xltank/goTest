package main

/**
The plugin should be compiled with the same flags as the main application.
If the application is compiled using the IDE already, then add the -gcflags="all=-N -l" to the above go build ... command.
go build -buildmode=plugin -gcflags="all=-N -l" -o test.so
*/

import (
	"fmt"
	"plugin"
)

func main() {
	p1, err := plugin.Open("../plugins/plugin1/main.so")
	if err != nil {
		fmt.Println("open plugin1 error:", err.Error())
		return
	}
	go1, err := p1.Lookup("Go")
	if err != nil {
		fmt.Println("lookup plugin1/Go error:", err.Error())
		return
	}

	go1.(func())()

}
