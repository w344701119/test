package grid

import (
	"fmt"
	"testing"
)

func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row := len(grid)
	nums := len(grid[0])
	for m := 0; m < row; m++ {
		for n := 0; n < nums; n++ {
			if m == 0 && n == 0 {
				continue
			} else if m == 0 {
				grid[m][n] = grid[m][n-1] + grid[m][n]
			} else if n == 0 {
				grid[m][n] = grid[m-1][n] + grid[m][n]
			} else {
				grid[m][n] = GetMin(grid[m][n-1]+grid[m][n], grid[m-1][n]+grid[m][n])
			}
			fmt.Println(grid)
		}
	}
	return grid[row-1][nums-1]
}

func GetMin(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func TestMinPathSum(t *testing.T) {
	//var grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	var grid = [][]int{{1, 2, 3}, {4, 5, 6}}
	re := MinPathSum(grid)
	fmt.Println(re)
}

//扩展列出路径
//定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
type Grid struct {
	X int
	Y int
	V int
}
type GridObj struct {
	Value    int
	GridPath []Grid
}

func MinPathSumExt(grid [][]int) *GridObj {
	var gridList [][]GridObj
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	row := len(grid)
	nums := len(grid[0])
	gridList = make([][]GridObj, row)
	for m := 0; m < row; m++ {
		for n := 0; n < nums; n++ {
			if m == 0 && n == 0 {
				gridList[m] = append(gridList[m], GridObj{grid[m][n], []Grid{{m, n, grid[m][n]}}})
				continue
			} else if m == 0 {

				prePath := append(gridList[m][n-1].GridPath, Grid{m, n, grid[m][n]})
				grid[m][n] = grid[m][n-1] + grid[m][n]
				gridList[m] = append(gridList[m], GridObj{grid[m][n], prePath})

			} else if n == 0 {
				tmpPath := append(gridList[m-1][n].GridPath, Grid{m, n, grid[m][n]})
				grid[m][n] = grid[m-1][n] + grid[m][n]
				gridList[m] = append(gridList[m], GridObj{grid[m][n], tmpPath})
			} else {
				obj := GetMinExt(MinObj{grid[m][n-1] + grid[m][n], m, n - 1}, MinObj{grid[m-1][n] + grid[m][n], m - 1, n})

				tmpPath := append(gridList[obj.M][obj.N].GridPath, Grid{m, n, grid[m][n]})
				grid[m][n] = obj.Value
				gridList[m] = append(gridList[m], GridObj{grid[m][n], tmpPath})

			}
		}
	}
	return &gridList[row-1][nums-1]
}

type MinObj struct {
	Value int
	M     int
	N     int
}

func GetMinExt(x, y MinObj) MinObj {
	if x.Value > y.Value {
		return y
	} else {
		return x
	}
}

func TestMinPathSumExt(t *testing.T) {
	var grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	re := MinPathSumExt(grid)
	fmt.Println(re)
}
