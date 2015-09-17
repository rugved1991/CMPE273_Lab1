package main
 
import "fmt"
 
func main(){
	var num int
	fmt.Println("Enter an number: ")
	fmt.Scanf("%d", &num)
	fmt.Println(fibonacci(num))
}

func fibonacci(i int) int{
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacci(i-1) + fibonacci(i-2)
	}

}