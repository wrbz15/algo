package sort

import (
	"testing"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{
				num: []int{2,8,6,3,4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.num)
			fmt.Printf("%+v\n", tt.args.num)
		})
	}
}