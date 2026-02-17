package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type keyPair struct {
	key   string
	value string
}

func TestUnique(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []keyPair
		Output []string
	}{
		{
			Name:   "empty array",
			Output: []string{},
		},
		{
			Name: "No duplicate entries",
			Input: []keyPair{
				{
					key:   "1",
					value: "1",
				},
				{
					key:   "2",
					value: "2",
				},
				{
					key:   "3",
					value: "3",
				},
				{
					key:   "4",
					value: "4",
				},
			},
			Output: []string{"1", "2", "3", "4"},
		},
		{
			Name: "Duplicate entries at the end",
			Input: []keyPair{
				{
					key:   "1",
					value: "1",
				},
				{
					key:   "2",
					value: "2",
				},
				{
					key:   "3",
					value: "3",
				},
				{
					key:   "4",
					value: "4",
				},
				{
					key:   "2",
					value: "2",
				},
				{
					key:   "3",
					value: "3",
				},
			},
			Output: []string{"1", "2", "3", "4"},
		},
		{
			Name: "Duplicate entries at the start",
			Input: []keyPair{
				{
					key:   "1",
					value: "1",
				},
				{
					key:   "1",
					value: "1",
				},
				{
					key:   "3",
					value: "3",
				},
				{
					key:   "4",
					value: "4",
				},
				{
					key:   "2",
					value: "2",
				},
				{
					key:   "3",
					value: "3",
				},
			},
			Output: []string{"1", "3", "4", "2"},
		},
		{
			Name: "Duplicate entries different values",
			Input: []keyPair{
				{
					key:   "1",
					value: "1",
				},
				{
					key:   "1",
					value: "1",
				},
				{
					key:   "3",
					value: "3",
				},
				{
					key:   "4",
					value: "4",
				},
				{
					key:   "2",
					value: "42",
				},
				{
					key:   "3",
					value: "43",
				},
			},
			Output: []string{"1", "43", "4", "42"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			orderedMap := GetOrderedMap()
			for _, v := range test.Input {
				orderedMap.Add(v.key, v.value)
			}
			assert.Equal(t, test.Output, orderedMap.GetValuesOrdered())
		})
	}
}
