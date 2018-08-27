package ds

import (
	"errors"
)

var (
	// ErrDifferentMatrixDimension notifies that both matrices must be of same dimension
	ErrDifferentMatrixDimension = errors.New("matrices are of different dimensions")
	// ErrColumnNotEqualRow is a multiplication error that is supported only for matrices
	// where the column on the lhs is equal to column on the rhs
	ErrColumnNotEqualRow = errors.New("column of matrix i is not equal to row of matrix j")
)

// Matrix is backed by a flat slice of elements
type Matrix struct {
	rows int
	cols int

	// flattened matrix data elements[i*step+j] -> Get(i, j) -> column major
	elements []float64

	// actual offset between rows
	step int
}

// NewMatrix create a new matrix
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		rows:     rows,
		cols:     cols,
		elements: make([]float64, rows*cols),
		step:     cols,
	}
}

// Nil returns if the matrix is nil
func (m *Matrix) Nil() bool {
	return m == nil
}

// Rows returns the number of rows
func (m *Matrix) Rows() int {
	return m.rows
}

// Cols returns the number of columns
func (m *Matrix) Cols() int {
	return m.cols
}

// NumElements return the number of elements in the grid
func (m *Matrix) NumElements() int {
	return m.rows * m.cols
}

// Arrays returns an array of slices referencing the matrix data
func (m *Matrix) Arrays() [][]float64 {
	n := make([][]float64, m.rows)
	for i := 0; i < m.rows; i++ {
		n[i] = m.elements[i*m.step : i*m.step+m.cols]
	}
	return n
}

// ColumnSlice returns a single column of the matrix grid
func (m *Matrix) ColumnSlice(j int) []float64 {
	c := make([]float64, m.rows)
	for i := 0; i < m.rows; i++ {
		c[i] = m.elements[i*m.step+j]
	}
	return c
}

// RowSlice returns a single row of the matrix grid
func (m *Matrix) RowSlice(i int) []float64 {
	return m.elements[i*m.step : i*m.step+m.cols]
}

// Get returns the element in the ith row and jth column
func (m *Matrix) Get(i, j int) float64 {
	return m.RowSlice(i)[j]
}

// Set sets the scalar element val in the ith row and jth col
func (m *Matrix) Set(i, j int, val float64) {
	m.elements[i*m.step : i*m.step+j][j] = val
}

// ScaleBy multiplies each element of the matrix by the given scalar value
func (m *Matrix) ScaleBy(s float64) {
	for i := 0; i < m.NumElements(); i++ {
		m.elements[i] *= s
	}
}

// Add creates and returns a new matrix that is the result of adding
// this matrix to the given matrix, the two matrices must be of same dimension
func (m *Matrix) Add(n *Matrix) (*Matrix, error) {
	if m.rows != n.rows || m.cols != n.cols {
		return nil, ErrDifferentMatrixDimension
	}
	o := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.NumElements(); i++ {
		o.elements[i] = m.elements[i] + n.elements[i]
	}
	return o, nil
}

// Subtract creates and returns a new matrix that is the result of subtracting
// this matrix to the given matrix, the two matrices must be of same dimension
func (m *Matrix) Subtract(n *Matrix) (*Matrix, error) {
	if m.rows != n.rows || m.cols != n.cols {
		return nil, ErrDifferentMatrixDimension
	}
	o := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.NumElements(); i++ {
		o.elements[i] = m.elements[i] - n.elements[i]
	}
	return o, nil
}
