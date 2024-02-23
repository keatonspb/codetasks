package megre_sorted_arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	for i, tc := range tCases {
		t.Run(fmt.Sprintf("tc_%d", i), func(t *testing.T) {
			merge(tc.list1, tc.m, tc.list2, tc.n)

			require.Equal(t, tc.want, tc.list1)
		})
	}
}
