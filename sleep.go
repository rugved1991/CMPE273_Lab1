package main

import "time"
import "fmt"

func Sleep(x int) {

	<-time.After(time.Second * time.Duration(x))

}

func main() {

	t1 := time.Now().Second()
	fmt.Println(t1)

	Sleep(10)

	t2 := time.Now().Second()
	fmt.Println(t2)

	if t2-t1 == 10 {

		fmt.Println("Done")
	}
}