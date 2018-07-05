package fundamental

import "testing"

func BenchmarkBTDepthRec(b *testing.B) {
	bt := NewBinaryTree([]Comparable{
		ComparableInt(20),
		ComparableInt(1),
		ComparableInt(2),
		ComparableInt(4),
		ComparableInt(-1),
		ComparableInt(5),
		ComparableInt(18),
		ComparableInt(-6),
		ComparableInt(35),
		ComparableInt(23),
		ComparableInt(12),
		ComparableInt(16),
		ComparableInt(9),
		ComparableInt(-12),
		ComparableInt(17),
		ComparableInt(13),
		ComparableInt(-16),
		ComparableInt(42),
		ComparableInt(3)})

	for n := 0; n < b.N; n++ {
		bt.DepthRec()
	}
}

func BenchmarkBTDepthIte(b *testing.B) {
	bt := NewBinaryTree([]Comparable{
		ComparableInt(20),
		ComparableInt(1),
		ComparableInt(2),
		ComparableInt(4),
		ComparableInt(-1),
		ComparableInt(5),
		ComparableInt(18),
		ComparableInt(-6),
		ComparableInt(35),
		ComparableInt(23),
		ComparableInt(12),
		ComparableInt(16),
		ComparableInt(9),
		ComparableInt(-12),
		ComparableInt(17),
		ComparableInt(13),
		ComparableInt(-16),
		ComparableInt(42),
		ComparableInt(3)})

	for n := 0; n < b.N; n++ {
		bt.DepthIte()
	}
}
