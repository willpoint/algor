package ds

import (
	"reflect"
	"testing"
)

func TestMatrix_NewMatrix(t *testing.T) {

}

func TestNewMatrix(t *testing.T) {
	type args struct {
		rows int
		cols int
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		{"1by2", args{1, 2}, &Matrix{1, 2, make([]float64, 1*2), 2}},
		{"5by9", args{5, 9}, &Matrix{5, 9, make([]float64, 5*9), 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {

	tests := []struct {
		name   string
		matrix *Matrix
		want   int
	}{
		{"one", NewMatrix(5, 9), 5},
		{"two", NewMatrix(9, 5), 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Rows(); got != tt.want {
				t.Errorf("Matrix.Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Cols(t *testing.T) {

	tests := []struct {
		name   string
		matrix *Matrix
		want   int
	}{
		{"one", NewMatrix(5, 9), 9},
		{"two", NewMatrix(9, 5), 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Cols(); got != tt.want {
				t.Errorf("Matrix.Cols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_NumElements(t *testing.T) {

	tests := []struct {
		name   string
		matrix *Matrix
		want   int
	}{
		{"one", NewMatrix(5, 9), 5 * 9},
		{"two", NewMatrix(9, 5), 5 * 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.NumElements(); got != tt.want {
				t.Errorf("Matrix.NumElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Array(t *testing.T) {

	m := NewMatrix(2, 3)
	m.SetRowSlice(0, []float64{7, 6, -11})
	m.SetRowSlice(1, []float64{4, 1, -2})
	mArray := []float64{7, 6, -11, 4, 1, -2}

	n := NewMatrix(3, 2)
	n.SetRowSlice(0, []float64{1, 3})
	n.SetRowSlice(1, []float64{0, -1})
	n.SetRowSlice(2, []float64{7, 5})
	nArray := []float64{1, 3, 0, -1, 7, 5}

	o := NewMatrix(1, 6)
	o.SetRowSlice(0, []float64{1, 3, 11, 3, -1, -8.5})
	oArray := []float64{1, 3, 11, 3, -1, -8.5}

	tests := []struct {
		name   string
		matrix *Matrix
		want   []float64
	}{
		{"2by3", m, mArray},
		{"3by2", n, nArray},
		{"1by6", o, oArray},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.matrix.Array(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Array() = %v, want %v", got, tt.want)
			}
		})
	}

}
