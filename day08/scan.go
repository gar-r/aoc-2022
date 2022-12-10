package main

func scanVisibility(trees []*Tree) {
	h := trees[0]
	h.Visible = true
	for i := 1; i < len(trees); i++ {
		t := trees[i]
		if t.Height > h.Height {
			t.Visible = true
			h = t
		}
	}
}

func scanScenicScore(trees []*Tree) int {
	score := 0
	h := trees[0].Height
	for i := 1; i < len(trees); i++ {
		score++
		if trees[i].Height >= h {
			return score
		}
	}
	return score
}

func reverse(trees []*Tree) []*Tree {
	l := len(trees)
	res := make([]*Tree, l)
	for i := 0; i < l; i++ {
		res[l-1-i] = trees[i]
	}
	return res
}
