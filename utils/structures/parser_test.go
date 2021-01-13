package structures

import (
    "reflect"
    "testing"
)

func TestParse2dIntArr(t *testing.T) {
    type args struct {
        array string
    }
    tests := []struct {
        name string
        args args
        want [][]int
    }{
        // TODO: Add test cases.
        {
            name: "Simple 2d int array test",
            args: args{"[[1,2,3],[1,2],[2,4]]"},
            want: [][]int{{1, 2, 3}, {1, 2}, {2, 4}},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Parse2dIntArr(tt.args.array); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Parse2dIntArr() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestParseIntArr(t *testing.T) {
    type args struct {
        array string
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        // TODO: Add test cases.
        {name: "Simple int array test", args: args{"[1,2,3]"}, want: []int{1, 2, 3}},
        {name: "Empty int array test", args: args{"[]"}, want: []int{}},
        {name: "Invalid int array test", args: args{"[1,2,3"}, want: nil},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ParseIntArr(tt.args.array); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ParseIntArr() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestParseStringArr(t *testing.T) {
    type args struct {
        array string
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        // TODO: Add test cases.
        {name: "Simple int array test", args: args{"[1,2,3]"}, want: []string{"1", "2", "3"}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ParseStringArr(tt.args.array); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ParseStringArr() = %v, want %v", got, tt.want)
            }
        })
    }
}
