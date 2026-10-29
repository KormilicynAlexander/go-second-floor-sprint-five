package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "negative size", size: -5, want: 0},
		{name: "zero size", size: 0, want: 0},
		{name: "one element", size: 1, want: 1},
		{name: "many elements", size: 1000, want: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.want {
				t.Fatalf("len(generateRandomElements(%d)) = %d, want %d", tt.size, len(data), tt.want)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive number", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "maximum first", data: []int{9, 3, 2}, want: 9},
		{name: "maximum middle", data: []int{2, 9, 3}, want: 9},
		{name: "maximum last", data: []int{2, 3, 9}, want: 9},
		{name: "equal elements", data: []int{4, 4, 4}, want: 4},
		{name: "zero values", data: []int{0, 0, 0}, want: 0},
		{name: "negative values", data: []int{-8, -2, -5}, want: -2},
		{name: "mixed signs", data: []int{-8, 0, 6, -2}, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "fewer elements than chunks", data: []int{3, 9, 2}, want: 9},
		{name: "one element per chunk", data: []int{1, 2, 3, 8, 5, 6, 7, 4}, want: 8},
		{name: "equal elements", data: []int{4, 4, 4, 4, 4, 4, 4, 4}, want: 4},
		{name: "zero values", data: []int{0, 0, 0, 0, 0, 0, 0, 0}, want: 0},
		{name: "short negative slice", data: []int{-8, -2, -5}, want: -2},
		{name: "negative values", data: []int{-9, -7, -4, -2, -8, -6, -5, -3}, want: -2},
		{name: "maximum in remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, want: 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}

	// Проверяем максимум в каждой позиции, в том числе на границах частей.
	for position := 0; position < CHUNKS*2+3; position++ {
		data := make([]int, CHUNKS*2+3)
		data[position] = 100
		if got := maxChunks(data); got != 100 {
			t.Errorf("maximum at index %d: got %d, want 100", position, got)
		}
	}
}
