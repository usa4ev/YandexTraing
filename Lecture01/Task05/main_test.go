package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T){
	type args [5]string
	tests := []struct{
		name string
		args args
		want [2]string
	}{
		{
			name: "ok",
			args: [5]string{"89", "20", "41", "1", "11"},
			want: [2]string{"2","3"},
		},
		{
			name: "ok",
			args: [5]string{"80", "20", "160", "2", "20"},
			want: [2]string{"1","20"},
		},
		{
			name: "ok 1_2",
			args: [5]string{"161", "20", "81", "2", "1"},
			want: [2]string{"3","1"},
		},
		{
			name: "ok 1",
			args: [5]string{"1", "20", "41", "1", "11"},
			want: [2]string{"1","1"},
		},
		{
			name: " ok equal",
			args: [5]string{"41", "20", "41", "1", "11"},
			want: [2]string{"1","11"},
		},
		{
			name: "p unknown",
			args: [5]string{"11", "1", "1", "1", "1"},
			want: [2]string{"0","1"},
		},
		{
			name: "p & n unknown",
			args: [5]string{"11", "2", "1", "1", "1"},
			want: [2]string{"0","0"},
		},
		{
			name: "input invalid",
			args: [5]string{"3", "2", "2", "2", "1"},
			want: [2]string{"-1","-1"},
		},
		{
			name: "input invalid",
			args: [5]string{"3", "2", "5", "1", "2"},
			want: [2]string{"1","0"},
		},
		{
			name: "input invalid",
			args: [5]string{"6", "2", "5", "1", "2"},
			want: [2]string{"0","0"},
		},
		{
			name: "input invalid",
			args: [5]string{"1", "1", "1", "1", "1"},
			want: [2]string{"1","1"},
		},
	}

	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			f,_ := os.Create("input.txt")
			fmt.Fprintf(f, "%s %s %s %s %s", tt.args[0], tt.args[1], tt.args[2], tt.args[3], tt.args[4])
		
			main()

			f,_ = os.Open("output.txt")
			var p,n string
			fmt.Fscanf(f, "%s %s", &p, &n)

			if p != tt.want[0] || n != tt.want[1]{
				t.Errorf("wrong p or n, got p=%v, n=%v; want p=%v, n=%v", p,n, tt.want[0], tt.want[1])
			}
		})
	}
}