package others

const (
	blank     = 0
	point     = 1
	processed = 2
	connected = 3
)

// FindDroppingBricks
// with a matrix of first row as top, any brick (mark as 1) was stable if connected to top
// now remove a brick at i, j, find how many bricks will drop
func FindDroppingBricks(matrix [][]int, x, y int) int {
	m, n := len(matrix), len(matrix[0])
	c, _ := dfs(matrix, m, n, x, y, x, y)
	return c
}

// dfs returns the count of points (from i, j) not connected to top (count, false), or (0, true) means they are connected
func dfs(matrix [][]int, m, n, i, j int, rooti, rootj int) (count int, top bool) {
	//defer func() {
	//	fmt.Printf("%v,%v: %v,%v\n", i, j, count, top)
	//}()
	//fmt.Printf("%v,%v\n", i, j)
	if i < 0 || i >= m || j < 0 || j >= n || matrix[i][j] == blank || matrix[i][j] == processed {
		count = 0
		top = false
		return
	}
	if matrix[i][j] == connected {
		count = 0
		top = true
		return
	}
	matrix[i][j] = processed
	count = 1
	c1, t1 := dfs(matrix, m, n, i+1, j, rooti, rootj)
	c2, t2 := dfs(matrix, m, n, i-1, j, rooti, rootj)
	c3, t3 := dfs(matrix, m, n, i, j+1, rooti, rootj)
	c4, t4 := dfs(matrix, m, n, i, j-1, rooti, rootj)
	if i == rooti && j == rootj {
		if !t1 {
			count += c1
		}
		if !t2 {
			count += c2
		}
		if !t3 {
			count += c3
		}
		if !t4 {
			count += c4
		}
		top = false
		return
	} else {
		if t1 || t2 || t3 || t4 || i == 0 {
			matrix[i][j] = connected
			count = 0
			top = true

		} else {
			matrix[i][j] = processed
			count, top = count+c1+c2+c3+c4, false

		}
		//fmt.Printf("%v,%v:%v,%v\n", i, j, count, top)
		return
	}
}
