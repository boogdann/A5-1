package matrix

type BinMatrix [][]byte

func New(bits []byte, prev, rows, cols int) BinMatrix {
	m := make(BinMatrix, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]byte, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m[i][j] = bits[prev*rows*cols+i*cols+j]
		}
	}
	return m
}

func (m BinMatrix) Rank() int {
	var rank, r, c int
	rowCount := len(m)
	colCount := len(m[0])

	for r < rowCount && c < colCount {
		p := pivot(m, c, r)
		if p != -1 {
			rank++
			if p != r {
				swapRows(m, r, p)
			}
			eliminateRow(m, r, c)
			r++
		}
		c++
	}

	return rank
}

func pivot(m BinMatrix, pivot, row int) int {
	for i := row; i < len(m); i++ {
		if m[i][pivot] == 1 {
			return i
		}
	}
	return -1
}

func swapRows(m BinMatrix, r1, r2 int) {
	for i := 0; i < len(m[0]); i++ {
		m[r1][i], m[r2][i] = m[r2][i], m[r1][i]
	}
}

func eliminateRow(m BinMatrix, r, c int) {
	for i := r + 1; i < len(m); i++ {
		if m[i][c] == 1 {
			for j := c; j < len(m[0]); j++ {
				m[i][j] ^= m[r][j]
			}
		}
	}
}
