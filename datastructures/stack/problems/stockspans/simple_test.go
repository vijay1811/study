package stockspans

import (
	"reflect"
	"testing"
)

func TestGetSpans(t *testing.T) {
	type args struct {
		stocks []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "success",
			args: args{
				stocks: []int{7, 4, 6, 7, 8, 6, 9},
			},
			want: []int{1, 1, 2, 4, 5, 1, 7},
		},
		{name: "success: all ascending",
			args: args{
				stocks: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{name: "success: all descending",
			args: args{
				stocks: []int{7, 6, 5, 4, 3, 2, 1},
			},
			want: []int{1, 1, 1, 1, 1, 1, 1},
		},
		{name: "success: empty",
			want: make([]int, 0),
		},
		{name: "success: one elem",
			args: args{
				stocks: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Spans(tt.args.stocks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSpans() = %v, want %v", got, tt.want)
			}
		})
	}
}
