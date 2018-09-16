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
	//ErrUnmatchedRow occurs when the length of the slice provided to set a row slice
	// does not match the length of the cols of the matrix to mutate
	ErrUnmatchedRow = errors.New("row provided does not match matrix cols")
	//ErrUnmatchedCol occurs when the length of the slice provided to set a row slice
	// does not match the length of the cols of the matrix to mutate
	ErrUnmatchedCol = errors.New("col provided does not match matrix rows")
	// ErrRowOutOfBound ...
	ErrRowOutOfBound = errors.New("row out out bound")
	// ErrColOutOfBound ...
	ErrColOutOfBound = errors.New("col out out bound")
)

// Matrix is backed by a flat slice of elements
// The example above is a 3 by 3 matrix
// +---+---+---+
// | 4 | 5 | 8 |
// +---+---+---+
// | 9 | 11| 13|
// +---+---+---+
// | -2| 3 | 91|
// +---+---+---+
// rows = 3, cols = 3, step = 3 and elements slice(cap=3*3)
type Matrix struct {
	rows int
	cols int

	// flattened matrix data elements[i*step+j] -> Get(i, j) -> column major
	elements []float64

	// actual offset between rows
	step int
}

// NewMatrix creates and returns a pionter to a new matrix
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

// Array returns the compressed array representing the matrix m
func (m *Matrix) Array() []float64 {
	return m.elements
}

// ColSlice returns a single column of the matrix grid
func (m *Matrix) ColSlice(j int) []float64 {
	c := make([]float64, m.rows)
	for i := 0; i < m.rows; i++ {
		c[i] = m.elements[i*m.step+j]
	}
	return c
}

// RowSlice returns a single row of the matrix grid
func (m *Matrix) RowSlice(i int) []float64 {
	c := make([]float64, m.cols)
	copy(c, m.elements[i*m.step:i*m.step+m.cols])
	return c
}

// SumRow adds the elememts in a row `i` of the grid
func (m *Matrix) SumRow(i int) float64 {
	var sum float64
	for j := 0; j < m.cols; j++ {
		sum += m.elements[i*m.step+j]
	}
	return sum
}

// SumCol adds the elements in col `j` of the grid
func (m *Matrix) SumCol(j int) float64 {
	var sum float64
	for i := 0; i < m.rows; i++ {
		sum += m.elements[i*m.step+j]
	}
	return sum
}

// SumDiag add the elements in the diagonal
// this works only for square matrixes
func (m *Matrix) SumDiag() (float64, error) {
	var sum float64
	if m.rows != m.cols {
		return sum, ErrDifferentMatrixDimension
	}
	for i := 0; i < m.cols; i++ {
		sum += m.elements[i*m.step+i]
	}
	return sum, nil
}

// Fill fills all the cells in a matrix with the value given
// analogous to a delete of all values
func (m *Matrix) Fill(v float64) {
	for i := 0; i < m.NumElements(); i++ {
		m.elements[i] = v
	}
}

// SetRowSlice sets a single row of the matrix with the given slice
func (m *Matrix) SetRowSlice(row int, s []float64) error {
	if len(s) != m.cols {
		return ErrUnmatchedRow
	}
	if row >= m.rows || row < 0 {
		return ErrRowOutOfBound
	}
	for i := 0; i < m.cols; i++ {
		m.elements[row*m.step+i] = s[i]
	}
	return nil
}

// SetColSlice set a single col of the matrix m to the elements
// of the given slice s
func (m *Matrix) SetColSlice(col int, s []float64) error {
	if len(s) != m.rows {
		return ErrUnmatchedCol
	}
	if col >= m.cols || col < 0 {
		return ErrColOutOfBound
	}
	for i := 0; i < m.rows; i++ {
		m.elements[i*m.step+col] = s[i]
	}
	return nil
}

// Get returns the element in the ith row and jth column
func (m *Matrix) Get(i, j int) float64 {
	return m.RowSlice(i)[j]
}

// Set sets the scalar element val in the ith row and jth col
func (m *Matrix) Set(i, j int, val float64) {
	m.elements[i*m.step+j] = val
}

// ScaleBy multiplies each element of the matrix by the given scalar value
func (m *Matrix) ScaleBy(s float64) {
	for i := 0; i < m.NumElements(); i++ {
		m.elements[i] *= s
	}
}

// Add creates and returns a new matrix that is the result of adding
// matrix m to n, the two matrices must have equal dimensions
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

// Multiply creates and returns a matrix whose result is the multiplication
// of matrices m and n. The cols of m must be equal to the rows n
func (m *Matrix) Multiply(n *Matrix) (*Matrix, error) {
	if m.cols != n.rows {
		return nil, ErrColumnNotEqualRow
	}

	o := NewMatrix(m.rows, n.cols)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < n.cols; j++ {
			sum := float64(0)
			for k := 0; k < m.cols; k++ {
				sum += m.elements[i*m.step+k] * n.Get(k, j)
			}
			o.elements[i*o.step+j] = sum
		}
	}
	return o, nil
}

// Transpose ...
func (m *Matrix) Transpose() *Matrix {
	o := NewMatrix(m.cols, m.rows)
	for i := 0; i < o.rows; i++ {
		for j := 0; j < o.cols; j++ {
			o.elements[i*o.step+j] = m.Get(j, i)
		}
	}
	return o
}

// Identity ...
func Identity(s int) *Matrix {
	o := NewMatrix(s, s)
	for i := 0; i < o.rows; i++ {
		o.elements[i*o.step+i] = 1
	}
	return o
}

// Determinant of A, denoted by |A| for a given 2x2 matrix
// (a b)
// (c d)
// is given as ab - bc
// however for larger  matrices `minors` and `cofactors` are required
// Minor
// if A is an (n * n)-matrix, then the minor M(i, j), for each
// (i, j), is the determinant of the (n-1 * n-1)-matrix obtained by
// deleting the ith row and the jth column

func det2by2(m *Matrix) (float64, error) {
	if m.rows != 2 || m.cols != 2 {
		return 0.0, errors.New("2 x 2 matrix dimension violated")
	}
	return m.elements[0]*m.elements[3] - m.elements[1]*m.elements[2], nil
}

func det3by3(m *Matrix) (float64, error) {
	if m.rows != 3 || m.cols != 3 {
		return 0.0, errors.New("3 x 3 matrix dimension violated")
	}
	return 0.0, nil
}

// ShedRow returns a new matrix with the given row removed
func (m *Matrix) ShedRow(row int) (*Matrix, error) {
	if row >= m.rows || row < 0 {
		return nil, ErrRowOutOfBound
	}
	o := NewMatrix(m.rows-1, m.cols)
	j := []float64{}
	for i := 0; i < m.NumElements(); i++ {
		if i >= (row*m.step) && i < (row*m.step+m.cols) {
			continue
		}
		j = append(j, m.elements[i])
	}
	copy(o.elements, j)
	return o, nil
}
