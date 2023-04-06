package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	type args []int
	tests := []struct {
		name string
		args args
		x    int
		want int
	}{
		{
			name: "std1",
			args: args{1, 2, 3, 4, 5},
			x:    6,
			want: 5,
		},
		{
			name: "std2",
			args: args{5, 4, 3, 2, 1},
			x:    3,
			want: 3,
		},
		{
			name: "neg",
			args: args{-3, -1, 0, 1, 3},
			x:    -4,
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := os.Create("input.txt")
			defer f.Close()

			fmt.Fprintln(f, len(tt.args))
			for _, v := range tt.args {
				fmt.Fprintf(f, "%d ", v)
			}

			fmt.Fprintln(f, tt.x)

			main()

			f, _ = os.Open("output.txt")
			defer f.Close()

			got := 0
			fmt.Fscan(f, &got)

			if got != tt.want {
				t.Errorf("want %v, got %v", tt.want, got)
			}

		})
	}

}
