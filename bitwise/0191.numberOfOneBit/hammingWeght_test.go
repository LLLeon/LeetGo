/*************************************************************************
+Author   : chenhuijia@deepglint.com
+Data     : 2020-03-15
+************************************************************************/

package hamming

import "testing"

func TestHammingWeight(t *testing.T) {
	a := 3 // 0011 -> 2
	b := 4 // 0100 -> 1
	c := 5 // 0101 -> 2

	a = HammingWeight(a)
	b = HammingWeight(b)
	c = HammingWeight(c)

	t.Logf("a -> %d\n", a)
	t.Logf("b -> %d\n", b)
	t.Logf("c -> %d\n", c)
}
