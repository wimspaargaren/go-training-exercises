package main

// func TestUnique(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		Input  []int
// 		Output []int
// 	}{
// 		{
// 			Name:   "empty array",
// 			Output: []int{},
// 		},
// 		{
// 			Name:   "negative numbers",
// 			Input:  []int{-1, -1},
// 			Output: []int{-1},
// 		},
// 		{
// 			Name:   "postive numbers",
// 			Input:  []int{1, 2, 3, 4},
// 			Output: []int{1, 2, 3, 4},
// 		},
// 		{
// 			Name:   "duplicate spread out",
// 			Input:  []int{1, 42, 2, 3, 4, 42},
// 			Output: []int{1, 42, 2, 3, 4},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			out := Unique(test.Input)
// 			assert.Equal(t, test.Output, out)
// 		})
// 	}
// }
