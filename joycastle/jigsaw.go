package main

import "fmt"

type BorderType int8

const (
	BorderTypeDirect = 0
	BorderTypeIndent = -1
	BorderTypeBulge  = 1
)

type Piece struct {
	ID     int
	Left   BorderType
	Top    BorderType
	Right  BorderType
	Bottom BorderType
}

type Jigsaw struct {
	width  int
	height int
	pieces [][]*Piece
}

func NewJigsaw(width int, height int) *Jigsaw {
	j := Jigsaw{
		width:  width,
		height: height,
	}
	j.pieces = make([][]*Piece, height)
	for i := 0; i < height; i++ {
		j.pieces[i] = make([]*Piece, width)
	}
	return &j
}

func (j *Jigsaw) Add(x, y int, piece *Piece) {
	j.pieces[x][y] = piece
}

func (j *Jigsaw) Check() bool {
	if j.height == 0 || j.width == 0 {
		return false
	}
	if len(j.pieces) != j.height {
		return false
	}
	for x := 0; x < len(j.pieces); x++ {
		if len(j.pieces[x]) != j.width {
			return false
		}
		for y := 0; y < len(j.pieces[x]); y++ {
			if j.pieces[x][y] == nil {
				return false
			}
			// 上边缘
			if x == 0 && j.pieces[x][y].Top != BorderTypeDirect {
				return false
			}
			// 左边缘
			if y == 0 && j.pieces[x][y].Left != BorderTypeDirect {
				return false
			}
			// 右边缘
			if y == len(j.pieces[x])-1 && j.pieces[x][y].Right != BorderTypeDirect {
				return false
			}
			// 下边缘
			if x == len(j.pieces)-1 && j.pieces[x][y].Bottom != BorderTypeDirect {
				return false
			}
			// 内部相邻
			if x < len(j.pieces)-1 && y < len(j.pieces[x])-1 &&
				(!j.checkPieces(j.pieces[x][y].Right, j.pieces[x][y+1].Left) || !j.checkPieces(j.pieces[x][y].Bottom, j.pieces[x+1][y].Top)) {
				return false
			}
		}
	}
	return true
}

func (j Jigsaw) checkPieces(borderType BorderType, borderType2 BorderType) bool {
	return borderType+borderType2 == 0
}

func main() {
	j := NewJigsaw(2, 2)
	j.Add(0, 0, &Piece{
		ID:     1,
		Left:   BorderTypeDirect,
		Top:    BorderTypeDirect,
		Right:  BorderTypeBulge,
		Bottom: BorderTypeBulge,
	})
	j.Add(0, 1, &Piece{
		ID:     2,
		Left:   BorderTypeIndent,
		Top:    BorderTypeDirect,
		Right:  BorderTypeDirect,
		Bottom: BorderTypeIndent,
	})
	j.Add(1, 0, &Piece{
		ID:     3,
		Left:   BorderTypeDirect,
		Top:    BorderTypeIndent,
		Right:  BorderTypeBulge,
		Bottom: BorderTypeDirect,
	})
	j.Add(1, 1, &Piece{
		ID:     4,
		Left:   BorderTypeIndent,
		Top:    BorderTypeBulge,
		Right:  BorderTypeDirect,
		Bottom: BorderTypeDirect,
	})

	r := j.Check()
	fmt.Println(r)
}
