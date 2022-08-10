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
	Width  int
	Height int
	Pieces [][]*Piece
}

func (j *Jigsaw) Check() bool {
	if j.Height == 0 || j.Width == 0 {
		return false
	}
	if len(j.Pieces) != j.Height {
		return false
	}
	for x := 0; x < len(j.Pieces); x++ {
		if len(j.Pieces[x]) != j.Width {
			return false
		}
		for y := 0; y < len(j.Pieces[x]); y++ {
			if j.Pieces[x][y] == nil {
				return false
			}
			// 上边缘
			if x == 0 && j.Pieces[x][y].Top != BorderTypeDirect {
				return false
			}
			// 左边缘
			if y == 0 && j.Pieces[x][y].Left != BorderTypeDirect {
				return false
			}
			// 右边缘
			if y == len(j.Pieces[x])-1 && j.Pieces[x][y].Right != BorderTypeDirect {
				return false
			}
			// 下边缘
			if x == len(j.Pieces)-1 && j.Pieces[x][y].Bottom != BorderTypeDirect {
				return false
			}
			// 内部相邻
			if x < len(j.Pieces)-1 && y < len(j.Pieces[x])-1 &&
				(!j.checkPieces(j.Pieces[x][y].Right, j.Pieces[x][y+1].Left) || !j.checkPieces(j.Pieces[x][y].Bottom, j.Pieces[x+1][y].Top)) {
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
	j := Jigsaw{
		Width:  2,
		Height: 2,
		Pieces: [][]*Piece{
			{
				&Piece{
					ID:     1,
					Left:   BorderTypeDirect,
					Top:    BorderTypeDirect,
					Right:  BorderTypeBulge,
					Bottom: BorderTypeBulge,
				},
				&Piece{
					ID:     2,
					Left:   BorderTypeIndent,
					Top:    BorderTypeDirect,
					Right:  BorderTypeDirect,
					Bottom: BorderTypeIndent,
				},
			},
			{
				&Piece{
					ID:     3,
					Left:   BorderTypeDirect,
					Top:    BorderTypeIndent,
					Right:  BorderTypeBulge,
					Bottom: BorderTypeDirect,
				},
				&Piece{
					ID:     4,
					Left:   BorderTypeIndent,
					Top:    BorderTypeBulge,
					Right:  BorderTypeDirect,
					Bottom: BorderTypeDirect,
				},
			},
		},
	}

	r := j.Check()
	fmt.Println(r)
}
