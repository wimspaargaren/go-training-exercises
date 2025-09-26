package main

// func TestAdd(t *testing.T) {
// 	tests := []struct {
// 		name     string
//nolint:dupword
// 		counter  counter
// 		addition int
// 		expected int
// 	}{
// 		{
// 			name: "Add 1 to 0",
// 			counter: counter{
// 				val: 0,
// 			},
// 			addition: 1,
// 			expected: 1,
// 		},
// 		{
// 			name: "Add 21 to 21",
// 			counter: counter{
// 				val: 21,
// 			},
// 			addition: 21,
// 			expected: 42,
// 		},
// 		{
// 			name: "Add negative numbers",
// 			counter: counter{
// 				val: 0,
// 			},
// 			addition: -1,
// 			expected: -1,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			test.counter.Add(test.addition)
// 			actual := test.counter.val
// 			assert.Equal(t, test.expected, actual)
// 		})
// 	}
// }
