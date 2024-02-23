package megre_sorted_arrays

type tCase struct {
	list1 []int
	list2 []int
	m     int
	n     int
	want  []int
}

var tCases = []tCase{
	{
		list1: []int{1, 2, 3, 0, 0, 0},
		list2: []int{2, 5, 6},
		m:     3,
		n:     3,
		want:  []int{1, 2, 2, 3, 5, 6},
	},
}
