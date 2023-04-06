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
			args: args{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			name: "std2",
			args: args{5, 4, 3, 2, 1},
			want: 0,
		},
		{
			name: "std3",
			args: args{1, 5, 1, 5, 1},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fin, _ := os.Create("input.txt")
			defer fin.Close()

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
