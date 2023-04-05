package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	type args []string

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok1",
			args: args{"1", "0", "0"},
			want: "0",
		},
		{
			name: "ok2",
			args: args{"1", "2", "3"},
			want: "7",
		},
		{
			name: "no solution",
			args: args{"1", "2", "-3"},
			want: "NO SOLUTION",
		},
		{
			name: "many solutions",
			args: args{"0", "4", "2"},
			want: "MANY SOLUTIONS",
		},
		{
			name: "neg x",
			args: args{"0", "-3", "2"},
			want: "NO SOLUTION",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := os.Create("input.txt")
			fmt.Fprintf(f, "%s\n%s\n%s", tt.args[0], tt.args[1], tt.args[2])

			main()

			f, _ = os.Open("output.txt")

			buf := bufio.NewReader(f)
			get,_ := buf.ReadString('\n')

			if get != tt.want {
				t.Errorf("Wrong get: %v, want %v", get, tt.want)
			}
		})
	}
}
