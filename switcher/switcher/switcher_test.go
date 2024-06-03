package switcher

import (
	"reflect"
	"testing"
)

// ListNodeToSlice
func ListNodeToSlice(head *ListNode) (values []int) {
	ref := head
	for ref != nil {
		values = append(values, ref.Value)
		ref = ref.Next
	}
	return
}

func TestSwap(t *testing.T) {
	type testCase struct {
		name     string
		input    *ListNode
		expected *ListNode
	}
	cases := []testCase{
		{
			name:     "case: nil ListNode",
			input:    nil,
			expected: nil,
		},
		{
			name:     "case: 1 node",
			input:    &ListNode{Value: 1},
			expected: &ListNode{Value: 1},
		},
		{
			name:     "case: 2 nodes",
			input:    &ListNode{Value: 1, Next: &ListNode{Value: 2}},
			expected: &ListNode{Value: 2, Next: &ListNode{Value: 1}},
		},
		{
			name:     "case: 3 nodes",
			input:    &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{Value: 3}}},
			expected: &ListNode{Value: 2, Next: &ListNode{Value: 1, Next: &ListNode{Value: 3}}},
		},
		{
			name: "case: 4 nodes",
			input: &ListNode{Value: 1, Next: &ListNode{Value: 2, Next: &ListNode{
				Value: 3, Next: &ListNode{Value: 4},
			}}},
			expected: &ListNode{Value: 2, Next: &ListNode{Value: 1, Next: &ListNode{
				Value: 4, Next: &ListNode{Value: 3},
			}}},
		},
	}

	// run cases
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			// ...

			// act
			output := Swap(tc.input)

			// assert
			output_vals := ListNodeToSlice(output)
			expected_vals := ListNodeToSlice(tc.expected)
			if !reflect.DeepEqual(expected_vals, output_vals) {
				t.Errorf("ListNodes not equal. Expected %v - Received %v", expected_vals, output_vals)
			}
		})
	}
}
