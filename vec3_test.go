package rlm

import (
	"reflect"
	"testing"
)

func TestNewVec3N(t *testing.T) {
	type args struct {
		v []float32
	}
	tests := []struct {
		name string
		args args
		want Vec3
	}{
		{
			"all axes provided",
			args{[]float32{1, 2, 3}},
			Vec3{X: 1, Y: 2, Z: 3},
		},
		{
			"X Y provided",
			args{[]float32{1, 2}},
			Vec3{X: 1, Y: 2, Z: 2},
		},
		{
			"X provided",
			args{[]float32{1}},
			Vec3{X: 1, Y: 1, Z: 1},
		},
		{
			"none provided",
			args{[]float32{}},
			Vec3{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVec3N(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVec3N() = %v, want %v", got, tt.want)
			}
		})
	}
}
