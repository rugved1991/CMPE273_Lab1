package fibo
import "testing"
func Test_fibo (t *testing.T){
var i int
i = fibonacci(10)
if i != 55 {
	t.Error("expected 2, got ", i)
}

}