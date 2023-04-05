package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func Test_executeTemp(t *testing.T) {
	type args struct {
		troom int
		tcond int
	}
	tests := []struct {
		name       string
		args       args
		wantFreeze int
		wantHeat   int
	}{
		{
			name:       "basic",
			args:       args{20, 10},
			wantFreeze: 10,
			wantHeat:   20,
		},
		{
			name:       "negative",
			args:       args{-10, -20},
			wantFreeze: -20,
			wantHeat:   -10,
		},
		{
			name:       "positive to negative",
			args:       args{20, -10},
			wantFreeze: -10,
			wantHeat:   20,
		},
		{
			name:       "even",
			args:       args{20, 20},
			wantFreeze: 20,
			wantHeat:   20,
		},
		{
			name:       "even negative",
			args:       args{-20, -20},
			wantFreeze: -20,
			wantHeat:   -20,
		},
		{
			name:       "all zero",
			args:       args{0, 0},
			wantFreeze: 0,
			wantHeat:   0,
		},
		{
			name:       "positive to zero",
			args:       args{10, 0},
			wantFreeze: 0,
			wantHeat:   10,
		},
		{
			name:       "negative to zero",
			args:       args{-10, 0},
			wantFreeze: -10,
			wantHeat:   0,
		},
		{
			name:       "zero to neagtive",
			args:       args{0, -10},
			wantFreeze: -10,
			wantHeat:   0,
		},

		{
			name:       "extreme",
			args:       args{50, -50},
			wantFreeze: -50,
			wantHeat:   50,
		},

		{
			name:       "reverse extreme",
			args:       args{-50, 50},
			wantFreeze: -50,
			wantHeat:   50,
		},
	}

	// freeze
	mode := freeze
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", mode, tt.name), func(t *testing.T) {
			if got := executeTemp(tt.args.troom, mode, tt.args.tcond); got != tt.wantFreeze {
				t.Errorf("executeTemp() = %v, want %v", got, tt.wantFreeze)
			}
		})
	}

	// heat
	mode = heat
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", mode, tt.name), func(t *testing.T) {
			if got := executeTemp(tt.args.troom, mode, tt.args.tcond); got != tt.wantHeat {
				t.Errorf("executeTemp() = %v, want %v", got, tt.wantHeat)
			}
		})
	}

	// fan
	mode = fan
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", mode, tt.name), func(t *testing.T) {
			if got := executeTemp(tt.args.troom, mode, tt.args.tcond); got != tt.args.troom {
				t.Errorf("executeTemp() = %v, want %v", got, tt.args.troom)
			}
		})
	}

	// auto
	mode = auto
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", mode, tt.name), func(t *testing.T) {
			if got := executeTemp(tt.args.troom, mode, tt.args.tcond); got != tt.args.tcond {
				t.Errorf("executeTemp() = %v, want %v", got, tt.args.tcond)
			}
		})
	}

	// freeze
	mode = freeze
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v with file", mode, tt.name), func(t *testing.T) {
			err := os.WriteFile("input.txt", []byte(fmt.Sprintf("%v %v\n%v",tt.args.troom, tt.args.tcond, mode)), os.ModePerm)
			if err != nil {
				t.Errorf("failed to write file: %v", err)
			}

			main()

			f, err := os.OpenFile("output.txt", os.O_RDONLY, os.ModePerm)
			if err != nil {
				t.Errorf("failed to open file: %v", err)
			}

			buf := bufio.NewReader(f)
			str, _ := buf.ReadString('\n')
			got, _ := strconv.Atoi(str)

			if got != tt.wantFreeze {
				t.Errorf("executeTemp() = %v, want %v", got, tt.wantFreeze)
			}
		})
	}

		// heat
		mode = heat
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%v %v with file", mode, tt.name), func(t *testing.T) {
				err := os.WriteFile("input.txt", []byte(fmt.Sprintf("%v %v\n%v",tt.args.troom, tt.args.tcond, mode)), os.ModePerm)
				if err != nil {
					t.Errorf("failed to write file: %v", err)
				}
	
				main()
	
				f, err := os.OpenFile("output.txt", os.O_RDONLY, os.ModePerm)
				if err != nil {
					t.Errorf("failed to open file: %v", err)
				}
	
				buf := bufio.NewReader(f)
				str, _ := buf.ReadString('\n')
				got, _ := strconv.Atoi(str)
	
				if got != tt.wantHeat {
					t.Errorf("executeTemp() = %v, want %v", got, tt.wantHeat)
				}
			})
		}

		// auto
		mode = auto
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%v %v with file", mode, tt.name), func(t *testing.T) {
				err := os.WriteFile("input.txt", []byte(fmt.Sprintf("%v %v\n%v",tt.args.troom, tt.args.tcond, mode)), os.ModePerm)
				if err != nil {
					t.Errorf("failed to write file: %v", err)
				}
	
				main()
	
				f, err := os.OpenFile("output.txt", os.O_RDONLY, os.ModePerm)
				if err != nil {
					t.Errorf("failed to open file: %v", err)
				}
	
				buf := bufio.NewReader(f)
				str, _ := buf.ReadString('\n')
				got, _ := strconv.Atoi(str)
	
				if got != tt.args.tcond {
					t.Errorf("executeTemp() = %v, want %v", got, tt.args.tcond)
				}
			})
		}

		// fan
		mode = fan
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%v %v with file", mode, tt.name), func(t *testing.T) {
				err := os.WriteFile("input.txt", []byte(fmt.Sprintf("%v %v\n%v",tt.args.troom, tt.args.tcond, mode)), os.ModePerm)
				if err != nil {
					t.Errorf("failed to write file: %v", err)
				}
	
				main()
	
				f, err := os.OpenFile("output.txt", os.O_RDONLY, os.ModePerm)
				if err != nil {
					t.Errorf("failed to open file: %v", err)
				}
	
				buf := bufio.NewReader(f)
				str, _ := buf.ReadString('\n')
				got, _ := strconv.Atoi(str)
	
				if got != tt.args.troom {
					t.Errorf("executeTemp() = %v, want %v", got, tt.args.troom)
				}
			})
		}
}
