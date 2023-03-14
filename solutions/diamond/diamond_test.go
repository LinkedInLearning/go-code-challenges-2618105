package diamond

import (
	"fmt"
	"testing"
)

func ExampleDiamond() {
	d := Diamond('A')
	fmt.Println(d)
	// Output:
	//    A
}

func TestDiamond(t *testing.T) {
	d := Diamond('D')

	expectedDiamond := "   A   \n" +
		"  B B  \n" +
		" C   C \n" +
		"D     D\n" +
		" C   C \n" +
		"  B B  \n" +
		"   A   "

	if d != expectedDiamond {
		t.Errorf("Expected: %s, got: %s", expectedDiamond, d)
	}
}
