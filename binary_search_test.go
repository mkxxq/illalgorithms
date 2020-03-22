package illalgorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 1},
			want: 0,
		},
		{
			name: "case 2",
			args: args{array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 0},
			want: -1,
		},
		{
			name: "case 3",
			args: args{array: []int{}, target: 0},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
