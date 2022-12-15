package packet

import (
	"fmt"
	"testing"
)

func TestList_Compare(t *testing.T) {

	tests := []struct {
		left     Packet
		right    Packet
		expected int
	}{
		{
			left:     Parse("[1,1,3,1,1]"),
			right:    Parse("[1,1,5,1,1]"),
			expected: -1,
		},
		{
			left:     Parse("[[1],[2,3,4]]"),
			right:    Parse("[[1],4]"),
			expected: -1,
		},
		{
			left:     Parse("[9]"),
			right:    Parse("[[8,7,6]]"),
			expected: 1,
		},
		{
			left:     Parse("[[4,4],4,4]"),
			right:    Parse("[[4,4],4,4,4]"),
			expected: -1,
		},
		{
			left:     Parse("[7,7,7,7]"),
			right:    Parse("[7,7,7]"),
			expected: 1,
		},
		{
			left:     Parse("[]"),
			right:    Parse("[3]"),
			expected: -1,
		},
		{
			left:     Parse("[[[]]]"),
			right:    Parse("[[]]"),
			expected: 1,
		},
		{
			left:     Parse("[1,[2,[3,[4,[5,6,7]]]],8,9]"),
			right:    Parse("[1,[2,[3,[4,[5,6,0]]]],8,9]"),
			expected: 1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			actual := test.left.Compare(test.right)
			if actual != test.expected {
				t.Errorf("expected: %d, got: %d\n", test.expected, actual)
			}
		})
	}

}
