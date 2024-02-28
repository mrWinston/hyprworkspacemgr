package grid

import (
	"reflect"
	"testing"
)

func TestIdxToCoord(t *testing.T) {
	tests := []struct {
		name      string
		idx       int
		wantCoord Coordinate
	}{
		{
			name:      "1",
			idx:       1,
			wantCoord: Coordinate{0, 0},
		},
		{
			name:      "2",
			idx:       2,
			wantCoord: Coordinate{0, 1},
		},
		{
			name:      "3",
			idx:       3,
			wantCoord: Coordinate{0, 2},
		},
		{
			name:      "5",
			idx:       5,
			wantCoord: Coordinate{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coord := IdxToCoord(tt.idx)
			if coord != tt.wantCoord {
				t.Errorf("IdxToCoord() gotY = %v, want %v", coord, tt.wantCoord)
			}
		})
	}
}

func TestNextCoordInDirection(t *testing.T) {
	type args struct {
		coord Coordinate
		dir   string
	}
	tests := []struct {
		name string
		args args
		want Coordinate
	}{
    {
    	name: "1",
    	args: args{
    		coord: Coordinate{0,1},
    		dir:   "left",
    	},
    	want: Coordinate{0,0},
    },
    {
    	name: "right",
    	args: args{
    		coord: Coordinate{1,1},
    		dir:   "right",
    	},
    	want: Coordinate{1,2},
    },
    {
    	name: "up",
    	args: args{
    		coord: Coordinate{1,1},
    		dir:   "up",
    	},
    	want: Coordinate{0,1},
    },
    {
    	name: "down",
    	args: args{
    		coord: Coordinate{1,1},
    		dir:   "down",
    	},
    	want: Coordinate{2,1},
    },
    {
    	name: "clamp",
    	args: args{
    		coord: Coordinate{2,2},
    		dir:   "right",
    	},
    	want: Coordinate{2,2},
    },
    {
    	name: "invalid dir",
    	args: args{
    		coord: Coordinate{0,0},
    		dir:   "bla",
    	},
    	want: Coordinate{0,0},
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextCoordInDirection(tt.args.coord, tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextCoordInDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
