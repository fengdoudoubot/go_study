package main
import(
	_"fmt"
	"testing"
)

func TestAddUpper(t *testing.T)  {
	res := addUpper(10)
	if res != 55{
		t.Fatalf("err")
	}

	t.Logf("ok")
}
