package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T){
	type args [3]int
	tests := []struct{
		name string
		args args
		want int
	}{
		{
			name:"test1",
			args: args{10, 5, 2},
			want: 4,
		},
		{
			name:"test2",
			args: args{13, 5, 3},
			want: 3,
		},
		{
			name:"test3",
			args: args{14, 5, 3},
			want: 4,
		},
		{
			name:"test4",
			args: args{13, 5, 2},
			want: 6,
		},
		{
			name:"test5",
			args: args{1, 1, 1},
			want: 1,
		},
		{
			name:"test6",
			args: args{1, 1, 2},
			want: 0,
		},
	}
	
	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			f,_ := os.Create("input.txt")
			fmt.Fprintf(f, "%d %d %d", tt.args[0], tt.args[1], tt.args[2])

			main()

			f,_ = os.Open("output.txt")
			var got int
			fmt.Fscanf(f, "%d", &got)

			if got != tt.want{
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
	
}