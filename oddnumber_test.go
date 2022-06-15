package mathlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddNumber(t *testing.T) {
	sample := []struct {
		name            string
		input           int
		expectedOutCome string
	}{
		{name: "01", input: 1, expectedOutCome: "ini bilangan ganjil"},
		{name: "02", input: 2, expectedOutCome: "ini bilangan genap"},
		{name: "03", input: 5, expectedOutCome: "ini bilangan ganjil"},
	}
	assert := assert.New(t)
	for _, sample := range sample {
		t.Run(sample.name, func(t *testing.T) {
			res := OddNumber(sample.input)
			assert.Equal(sample.expectedOutCome, res, "they sould be equal")
		})
	}
}

func BenchmarkOddNumber(b *testing.B) {
	sample := []struct {
		name  string
		param int
	}{
		{
			name:  "01",
			param: 1,
		},
		{
			name:  "02",
			param: 2,
		},
		{
			name:  "03",
			param: 3,
		},
		{
			name:  "04",
			param: 4,
		},
	}

	for _, test := range sample {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				OddNumber(test.param)
			}

		},
		)
	}
}
