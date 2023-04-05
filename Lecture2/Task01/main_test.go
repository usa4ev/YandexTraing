package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T){
	type args []int
	tests := []struct{
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{1,4,8},
			want: "YES",
		},
		{
			name: "oklong",
			args: args{1,4,5,6,8},
			want: "YES",
		},
		{
			name: "okneg",
			args: args{-10,-8,-5,-4,-1},
			want: "YES",
		},
		{
			name: "nonegeq",
			args: args{-10,-8,-5,-5,-1},
			want: "NO",
		},
		{
			name: "noneg",
			args: args{-10,-8,-9,-1},
			want: "NO",
		},
		{
			name: "noeg",
			args: args{1, 4, 4},
			want: "NO",
		},
		{
			name: "no",
			args: args{1, 4, 2},
			want: "NO",
		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			f,_ := os.Create("input.txt")
			for _,v := range tt.args{
				fmt.Fprintf(f, "%d ", v)
			}
			
			main()

			f,_ = os.Open("output.txt")
			got := ""
			fmt.Fscan(f, &got)

			if got != tt.want{
				t.Errorf("want %v, got %v", tt.want, got)
			}

		})
	}

	
}
