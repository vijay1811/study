package sorting

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	arr := RandomArray(100000, 100000)
	tests := []struct {
		name  string
		args  args
		setup func() []int
		want  []int
	}{
		{name: "Success",
			args: args{
				arr: []int{5, 2, 4, 6, 1, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{name: "Success: nothing"},
		{name: "Success: one element",
			args: args{
				arr: []int{1},
			},
			want: []int{1},
		},
		{name: "Success: random array",
			args: args{
				arr: CopySlice(arr),
			},
			setup: func() []int {
				cArr := CopySlice(arr)
				sort.Ints(cArr)
				return cArr
			},
			//want: []int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.want = tt.setup()
			}
			InsertionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.want, tt.args.arr) {
				fmt.Println(tt.want)
				fmt.Println(tt.args.arr)
				t.Error("sort failed")
			}
		})
	}
}
