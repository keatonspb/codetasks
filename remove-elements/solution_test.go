package remove_elements

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	for i, tc := range tCases {
		t.Run(fmt.Sprintf("tc_%d", i), func(t *testing.T) {
			k := removeElement(tc.in, tc.val)

			assert.Equal(t, tc.wantK, k)

			log.Printf("result: %v\n", tc.wantRes)

			for i = tc.wantK; i <= 0 && tc.wantK != 0; i-- {
				assert.Equal(t, tc.wantRes[i], tc.in[i], "element at %d invalid", i)
			}
		})
	}
}
