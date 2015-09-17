package perimeter
import "testing"
func Test_perimeterRect (t *testing.T) {

	r := Rect{length: 5, width: 3}
	var x Shape
	x = r 

	Rectperi :=  x.peri()
	if Rectperi != 16 {

		t.Error("expected 16, got ", Rectperi)

	}

}

func Test_perimeterCircle (t *testing.T) {

	c := Circ{r: 5}
	var x Shape
	x = c 

	CPeri := x.peri()
	if CPeri != 31.41592653589793 {

		t.Error("expected 31.41, got ", CPeri)

	}

}