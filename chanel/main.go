package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		message <- data
	}

	go sayHelloTo("made suande")
	go sayHelloTo("made suande 2")
	go sayHelloTo("made suande 3")

	var message1 = <- message
	fmt.Println(message1)

	var message2 = <-message
    fmt.Println(message2)

    var message3 = <-message
    fmt.Println(message3)
}