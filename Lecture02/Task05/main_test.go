package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	type (
		args []int
	)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "std1",
			args: args{10, 20, 15, 10, 30, 5, 1},
			want: 6,
		},
		{
			name: "std2",
			args: args{15, 15, 10},
			want: 1,
		},
		{
			name: "std3",
			args: args{10, 15, 20},
			want: 0,
		},
		{
			name: "all even",
			args: args{15, 15, 15},
			want: 0,
		},
		{
			name: "no solution",
			args: args{10, 15, 15},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fin, _ := os.Create("input.txt")
			defer fin.Close()

			fmt.Fprintf(fin, "%d\n", len(tt.args))

			for _, v := range tt.args {
				fmt.Fprintf(fin, "%d\n", v)
			}

			main()

			fout, _ := os.Open("output.txt")
			defer fout.Close()

			var got int

			fmt.Fscan(fout, &got)

			if tt.want != got {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
