package remove_elements

type tCase struct {
	in      []int
	val     int
	wantRes []int
	wantK   int
}

var tCases = []tCase{
	{
		in:      []int{3, 2, 2, 3},
		val:     3,
		wantRes: []int{2, 2, 0, 0},
		wantK:   2,
	},
	{
		in:      []int{0, 1, 2, 2, 3, 0, 4, 2},
		val:     2,
		wantRes: []int{0, 1, 4, 0, 3, 0, 0, 0},
		wantK:   5,
	},
	{
		in:      []int{0, 1, 2, 2, 3, 0, 4, 2},
		val:     10,
		wantRes: []int{0, 1, 2, 2, 3, 0, 4, 2},
		wantK:   8,
	},
	{
		in:      []int{},
		val:     0,
		wantRes: []int{},
		wantK:   0,
	},
}
