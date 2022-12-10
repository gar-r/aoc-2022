package main

type Forest struct {
	Trees [][]*Tree
}

type Tree struct {
	Height      int
	Visible     bool
	ScenicScore int
}

func NewForest(data [][]int) *Forest {
	f := &Forest{Trees: make([][]*Tree, len(data))}
	w := len(data[0])
	for i := range data {
		f.Trees[i] = make([]*Tree, w)
		for j := 0; j < w; j++ {
			f.Trees[i][j] = &Tree{Height: data[i][j]}
		}
	}
	f.calcVisibility()
	f.calcScenicScores()
	return f
}

func (f *Forest) CountVisible() int {
	total := 0
	for _, row := range f.rows() {
		for _, t := range row {
			if t.Visible {
				total++
			}
		}
	}
	return total
}

func (f *Forest) FindBestScore() int {
	max := 0
	for _, row := range f.rows() {
		for _, t := range row {
			if t.ScenicScore > max {
				max = t.ScenicScore
			}
		}
	}
	return max
}

func (f *Forest) calcVisibility() {
	for _, row := range f.rows() {
		scanVisibility(row)
		scanVisibility(reverse(row))
	}
	for _, col := range f.cols() {
		scanVisibility(col)
		scanVisibility(reverse(col))
	}
}

func (f *Forest) calcScenicScores() {
	for i := range f.rows() {
		for j := range f.cols() {
			row := f.row(i)
			col := f.col(j)
			scoreLeft := scanScenicScore(reverse(row[0 : j+1]))
			scoreRight := scanScenicScore(row[j:])
			scoreUp := scanScenicScore(reverse(col[0 : i+1]))
			scoreDown := scanScenicScore(col[i:])
			f.Trees[i][j].ScenicScore = scoreLeft * scoreRight * scoreUp * scoreDown
		}
	}
}

func (f *Forest) rows() [][]*Tree {
	return f.Trees
}

func (f *Forest) row(n int) []*Tree {
	return f.Trees[n]
}

func (f *Forest) cols() [][]*Tree {
	w := len(f.Trees[0])
	res := make([][]*Tree, w)
	for i := 0; i < w; i++ {
		res[i] = f.col(i)
	}
	return res
}

func (f *Forest) col(n int) []*Tree {
	h := len(f.Trees)
	res := make([]*Tree, h)
	for i := 0; i < h; i++ {
		res[i] = f.Trees[i][n]
	}
	return res
}
