package fenwick

import "testing"

func TestTree(t *testing.T) {
	type query struct {
		index, value int
	}
	tests := []struct {
		name string
		size int
		qs   []query
		want []int
	}{
		{
			name: "normal case [1 0 0]",
			size: 3,
			qs:   []query{{0, 1}},
			want: []int{1, 1, 1, 0, 0, 0},
		},
		{
			name: "normal case [1 2 0]",
			size: 3,
			qs:   []query{{0, 1}, {1, 2}},
			want: []int{1, 3, 3, 2, 2, 0},
		},
		{
			name: "normal case [1 2 3]",
			size: 3,
			qs:   []query{{0, 1}, {1, 2}, {2, 3}},
			want: []int{1, 3, 6, 2, 5, 3},
		},
		{
			name: "normal case [1 0 3]",
			size: 3,
			qs:   []query{{0, 1}, {1, 2}, {2, 3}, {1, -2}},
			want: []int{1, 1, 4, 0, 3, 3},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewTree(tt.size)
			for _, q := range tt.qs {
				tree.Add(q.index, q.value)
			}

			got := tree.Query(0, 0)
			if got != tt.want[0] {
				t.Errorf("case %d. got %d; want %d", i, got, tt.want[0])
			}
			got = tree.Query(0, 1)
			if got != tt.want[1] {
				t.Errorf("case %d. got %d; want %d", i, got, tt.want[1])
			}
			got = tree.Query(0, 2)
			if got != tt.want[2] {
				t.Errorf("case %d. got %d; want %d", i, got, tt.want[2])
			}
			got = tree.Query(1, 1)
			if got != tt.want[3] {
				t.Errorf("case %d. got %d; want %d", i, got, tt.want[3])
			}
			got = tree.Query(1, 2)
			if got != tt.want[4] {
				t.Errorf("case %d. got %d; want %d", i, got, tt.want[4])
			}
			got = tree.Query(2, 2)
			if got != tt.want[5] {
				t.Errorf("case %d. got %d; want %d", i, got, tt.want[5])
			}
		})
	}
}
