// dynasty2048
// 2048朝代版
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	ROWS = 5
	COLS = 5
	//难度级别（每次自动填写的朝代个数）
	LEVEL = 3
)

//操作缩写
const (
	UP    = "U"
	DOWN  = "D"
	LEFT  = "L"
	RIGHT = "R"
	QUIT  = "Q"
)

//商、周、秦、汉、晋、隋、唐、宋、元、明、清
var dynastyMap map[int]string
var valuesArray [ROWS][COLS]int
var namesArray [ROWS][COLS]string
var rdTimes int = 0

//得分
var score int = 0

//游戏是否结束
var gameover bool = false

//剩余空格
var remaidCeil int = ROWS * COLS

//方向选择
var choice string

func init() {
	fmt.Println("-----------2048朝代版-----------")
	fmt.Printf("棋盘尺寸（SIZE）：%d ，难度系数（LEVEL）: %d\n", ROWS, LEVEL)
	fmt.Println("请输入U(p)、D(own)、L(eft)、R(ight),Q(uit)，控制‘朝代’上、下、左、右移动，或者退出")
	//初始化dynastyMap
	dynastyMap = map[int]string{0: "  ", 2: "商", 4: "周", 8: "秦", 16: "汉", 32: "晋", 64: "隋", 128: "唐", 256: "宋", 512: "元", 1024: "明", 2048: "清"}

	//map排序
	values := make([]int, len(dynastyMap))
	i := 0
	for value, _ := range dynastyMap {
		values[i] = value
		i++
	}

	fmt.Println("得分规则：")
	sort.Ints(values)
	for _, value := range values {
		fmt.Printf("%s:%d ", dynastyMap[value], value)
	}

	fmt.Println()

	//初始化valuesArray
	randInput()
	printGrid()
}

func main() {

	var direct string
	for direct != QUIT {
		fmt.Scanln(&direct)

		direct = strings.ToUpper(direct)
		if direct == "" {
			continue
		} else if direct == QUIT {
			break
		}

		move(direct)
		randInput()
		printGrid()

		if gameover == true {
			break
		}

		direct = ""

	}

}

//随机向空格填入朝代
func randInput() {
	//重置remaidCeil=0，以便重新计算剩余空格
	remaidCeil = 0
	//存放剩余空格的位置坐标row+col
	var remaidCeilArray = [25]string{}
	var index int = 0
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {

			if valuesArray[row][col] == 0 {
				remaidCeil += 1
				remaidCeilArray[index] = strconv.Itoa(row) + strconv.Itoa(col)
				index++
			}
		}

	}

	//判断游戏是否结束
	if remaidCeil < LEVEL {
		gameover = true
		fmt.Println("游戏结束")
		return
	}
	//向空格填入朝代
	for t := 0; t < LEVEL; {
		idx := randNum2(remaidCeil)
		rowCol := remaidCeilArray[idx]

		r := rowCol[0]
		c := rowCol[1]
		rStr := string(r)
		cStr := string(c)

		if rStr != "-" && cStr != "-" {
			row, _ := strconv.Atoi(rStr)
			col, _ := strconv.Atoi(cStr)
			valuesArray[row][col] = randNum()
			remaidCeilArray[idx] = "--"
			t++
		}

	}
}

//打印网格
func printGrid() {

	//打印得分
	fmt.Println("SCORE:", score)

	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			value := valuesArray[row][col]
			fmt.Printf("[%s]", dynastyMap[value])
		}
		fmt.Println()
	}
}

//2 or 4，商or周
func randNum() (rd int) {
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	rd = 2 * (rand.Intn(2) + 1)
	time.Sleep(1 * time.Nanosecond)
	return
}

//随机生成[0,n)的整数
func randNum2(n int) (rd int) {
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	rd = rand.Intn(n)
	time.Sleep(3 * time.Nanosecond)
	return
}

//移动
func move(direct string) {
	switch direct {
	case UP:
		moveUp()
		mergeUp()
		moveUp()

	case DOWN:
		moveDown()
		mergeDown()
		moveDown()

	case LEFT:
		moveLeft()
		mergeLeft()
		moveLeft()
	case RIGHT:
		moveRight()
		mergeRight()
		moveRight()
	default:
		fmt.Printf("无效的移动方向:%s\n", direct)
	}
}

//向左移动
func moveLeft() {
	for row := 0; row < ROWS; row++ {
		//假设当前行值为0的网格的最小下标值
		var minIdx int = 0
		if valuesArray[row][0] != 0 {
			minIdx += 1
		}

		for col := 1; col < COLS; col++ {
			value := valuesArray[row][col]
			//判断是否需要移动
			if value != 0 {
				if minIdx < col {
					valuesArray[row][minIdx] = value
					valuesArray[row][col] = 0
					minIdx++
				} else if minIdx == col {
					minIdx++
				}

			}
		}
	}

}

//向右移动
func moveRight() {
	for row := 0; row < ROWS; row++ {
		//假设当前行值为0的网格的最大下标值
		var maxIdx int = COLS - 1
		if valuesArray[row][COLS-1] != 0 {
			maxIdx -= 1
		}

		for col := COLS - 2; col >= 0; col-- {
			value := valuesArray[row][col]
			//判断是否需要移动
			if value != 0 {
				if maxIdx > col {
					valuesArray[row][maxIdx] = value
					valuesArray[row][col] = 0
					maxIdx--
				} else if maxIdx == col {
					maxIdx--
				}

			}
		}
	}

}

//向上移动
func moveUp() {
	for col := 0; col < COLS; col++ {
		//假设当前列值为0的网格的最小下标
		var minIdx int = 0
		if valuesArray[0][col] != 0 {
			minIdx++
		}

		for row := 1; row < ROWS; row++ {
			//判断是否需要移动
			value := valuesArray[row][col]
			if value != 0 {
				if minIdx < row {
					valuesArray[minIdx][col] = value
					valuesArray[row][col] = 0
					minIdx++
				} else if minIdx == row {
					minIdx++
				}
			}

		}

	}
}

//向下移动
func moveDown() {
	for col := 0; col < COLS; col++ {
		//假设当前列值为0的网格的最大下标
		var maxIdx int = ROWS - 1
		if valuesArray[ROWS-1][col] != 0 {
			maxIdx--
		}

		for row := ROWS - 2; row >= 0; row-- {
			//判断是否需要移动
			value := valuesArray[row][col]
			if value != 0 {
				if maxIdx > row {
					valuesArray[maxIdx][col] = value
					valuesArray[row][col] = 0
					maxIdx--
				} else if maxIdx == row {
					maxIdx--
				}
			}

		}

	}
}

//向左合并
func mergeLeft() {
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS-1; col++ {
			value := valuesArray[row][col]
			if value != 0 {
				nextValue := valuesArray[row][col+1]
				if nextValue == 0 {
					continue
				} else if value == nextValue {
					valuesArray[row][col] = 2 * value
					valuesArray[row][col+1] = 0
					score += value
				}
			} else {
				continue
			}

		}
	}

}

//向右合并
func mergeRight() {
	for row := 0; row < ROWS; row++ {
		for col := COLS - 1; col > 0; col-- {
			value := valuesArray[row][col]
			if value != 0 {
				nextValue := valuesArray[row][col-1]
				if nextValue == 0 {
					continue
				} else if value == nextValue {
					valuesArray[row][col] = 2 * value
					valuesArray[row][col-1] = 0
					score += value
				}
			} else {
				continue
			}

		}
	}

}

//向上合并
func mergeUp() {
	for col := 0; col < COLS; col++ {
		for row := 0; row < ROWS-1; row++ {
			value := valuesArray[row][col]
			if value != 0 {
				nextValue := valuesArray[row+1][col]
				if nextValue == 0 {
					continue
				} else if value == nextValue {
					valuesArray[row][col] = 2 * value
					valuesArray[row+1][col] = 0
					score += value
				}
			} else {
				continue
			}
		}
	}
}

//向下合并
func mergeDown() {
	for col := 0; col < COLS; col++ {
		for row := ROWS - 1; row > 0; row-- {
			value := valuesArray[row][col]
			if value != 0 {
				nextValue := valuesArray[row-1][col]
				if nextValue == 0 {
					continue
				} else if value == nextValue {
					valuesArray[row][col] = 2 * value
					valuesArray[row-1][col] = 0
					score += value
				}
			} else {
				continue
			}
		}
	}
}
