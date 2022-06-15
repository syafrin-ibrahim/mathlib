package mathlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddNumber(t *testing.T) {
	sample := []struct {
		name            string
		input           []int
		expectedOutCome []string
	}{
		{name: "01", input: []int{5, 3, 2, 1}, expectedOutCome: []string{"ganjil", "ganjil", "genap", "ganjil"}},
		{name: "02", input: []int{2, 6, 1, 3}, expectedOutCome: []string{"genap", "genap", "ganjil", "ganjil"}},
		{name: "03", input: []int{7, 3, 3, 2}, expectedOutCome: []string{"ganjil", "ganjil", "ganjil", "genap"}},
		{name: "04", input: []int{10, 1, 4, 3}, expectedOutCome: []string{"genap", "ganjil", "genap", "ganjil"}},
	}
	assert := assert.New(t)
	for _, sample := range sample {
		t.Run(sample.name, func(t *testing.T) {
			res := OddNumber(sample.input...)
			assert.Equal(sample.expectedOutCome, res, "they sould be equal")
		})
	}
}

func BenchmarkOddNumber(b *testing.B) {
	sample := []struct {
		name  string
		param []int
	}{
		{
			name:  "01",
			param: []int{5, 3, 2, 1},
		},
		{
			name:  "02",
			param: []int{2, 6, 1, 3},
		},
		{
			name:  "03",
			param: []int{7, 3, 3, 2},
		},
		{
			name:  "04",
			param: []int{10, 1, 4, 3},
		},
	}

	for _, test := range sample {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				OddNumber(test.param...)
			}

		},
		)
	}
}
