package rlm

import (
	"reflect"
	"testing"
)

func TestNewVec2N(t *testing.T) {
	type args struct {
		v []float32
	}
	tests := []struct {
		name string
		args args
		want Vec2
	}{
		{
			"all axes provided",
			args{[]float32{1, 2}},
			Vec2{X: 1, Y: 2},
		},
		{
			"X provided",
			args{[]float32{1}},
			Vec2{X: 1, Y: 1},
		},
		{
			"none provided",
			args{[]float32{}},
			Vec2{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVec2N(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVec2N() = %v, want %v", got, tt.want)
			}
		})
	}
}
