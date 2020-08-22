package main

import "C"

import (
	"fmt"
)

//export PrintMessage
func PrintMessage() {
	fmt.Println("Hello from Go")
}

func main(){
	// farther you should entet in console next command
	// go build -o [name of file.o].o -buildmode=c-shared [name of go.file].go
	// create a .c file include [name of go file].h and write some code
	// compile with command 
	// gcc -o [name of compile file ] [name of file.c ]o.c [link library] ./[library name].o
	// activate with ./[name of compile file]
}
