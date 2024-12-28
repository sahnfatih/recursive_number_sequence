package sequence

import (
    "reflect"
    "testing"
)

func TestGenerateSequence(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected []int
    }{
        {
            name:     "Input 9",
            input:    9,
            expected: []int{2, 4, 9},
        },
        {
            name:     "Input 15",
            input:    15,
            expected: []int{2, 3, 7, 15},
        },
        {
            name:     "Input 8",
            input:    8,
            expected: []int{2, 4, 8},
        },
        {
            name:     "Input 1",
            input:    1,
            expected: []int{},
        },
        {
            name:     "Input 0",
            input:    0,
            expected: []int{},
        },
        {
            name:     "Negative input",
            input:    -5,
            expected: []int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := GenerateSequence(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("GenerateSequence(%d) = %v, want %v", 
                    tt.input, result, tt.expected)
            }
        })
    }
}

// Test performance with larger numbers
func BenchmarkGenerateSequence(b *testing.B) {
    for i := 0; i < b.N; i++ {
        GenerateSequence(1000)
    }
}