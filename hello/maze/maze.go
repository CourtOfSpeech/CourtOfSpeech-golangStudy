package main

import (
	"fmt"
	"os"
)

//读取maze.in文件
func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var row, col int
	//读取文件,这里读取第一行，得到这是多少行多少列的迷宫
	fmt.Fscanf(file, "%d %d", &row, &col)
	//创建二维slice
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			//这里是一行一行的读数据
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	//返回这个二维的slice
	return maze
}

//点结构
type ponit struct {
	i, j int
}

//探索的4个位置
var dirs = [4]ponit{
	{-1, 0}, //上
	{0, -1}, //左
	{1, 0},  //下
	{0, 1},  //右
}

//取得下一个节点的位置
func (p ponit) add(r ponit) ponit {
	return ponit{p.i + r.i, p.j + r.j}
}

//返回该位置的值是否为0
//p 坐标
//grid 迷宫数组
//int 返回的迷宫数组值
//bool 返回的该坐标是否越界的标识
func (p ponit) at(grid [][]int) (int, bool) {
	//向上或者向下都越界了
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	//向左或者向有都越界了
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

//广度优先算法走迷宫
//maze要走的迷宫
//start 开始的位置
//end	结束的位置
func walk(maze [][]int, start, end ponit) [][]int {
	//建一个二维的slice来放走的路径这些
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//把起点加进队列里面,队列里面主要放我们要走的位置
	Q := []ponit{start}

	//退出循环的条件是：1.走到终点、2.队列没有值
	for len(Q) > 0 {
		//得到要探索的位置，即队列的头
		cur := Q[0]
		//新的队列需要去掉刚才已经取得的队列数据
		Q = Q[1:]

		//如果到达终点了，也要退出
		if cur == end {
			break
		}

		//取得的队列，即要探索的位置，需要从上、下、左、右4个方向去探索
		for _, dir := range dirs {
			//下一次需要探索的位置
			next := cur.add(dir)

			//要探索下个节点需要满足的条件
			val, ok := next.at(maze)
			if !ok || val == 1 {
				//不满足条件, !ok 表示越界， val==1表示撞强
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				//不满足条件, !ok 表示越界， val!=0表示已经走过
				continue
			}

			if next == start {
				//表示回到原点
				continue
			}

			//把上面不满足条件的过滤后，接下来是满足条件的，所以开始探索
			curSteps, _ := cur.at(steps)
			//1.改变数据，记录走的步数
			steps[next.i][next.j] = curSteps + 1

			//2.将下一个节点，加入到队列里去
			Q = append(Q, next)

		}
	}

	//返回我们探索所走的路径即步骤
	return steps

}

func main() {
	//读取maze.in文件
	maze := readMaze("maze.in")

	//看看文件有没有读取正确
	// for _, row := range maze {
	// 	for _, val := range row {
	// 		fmt.Printf("%d ", val)
	// 	}
	// 	fmt.Println()
	// }

	//fmt.Println(len(maze) - 1)
	//fmt.Println(len(maze[0]) - 1)

	start := ponit{0, 0}
	end := ponit{len(maze) - 1, len(maze[0]) - 1}

	//广度优先算法走迷宫
	steps := walk(maze, start, end)

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	//查看走了多少步
	if step, ok := end.at(steps); !ok {
		fmt.Println("error")
	} else {
		fmt.Printf("walk %d steps", step)
	}

}
