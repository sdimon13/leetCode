package number_of_islands

func visitIsland(grid [][]byte, y, x int) {
	rowsSize, colsSize := len(grid), len((grid)[0])

	if y < 0 || y >= rowsSize || x < 0 || x >= colsSize || (grid)[y][x] == '0' {
		return
	}

	(grid)[y][x] = '0'

	visitIsland(grid, y-1, x)
	visitIsland(grid, y+1, x)
	visitIsland(grid, y, x-1)
	visitIsland(grid, y, x+1)
}

func numIslands(grid [][]byte) int {

	islandCount := 0

	for y, row := range grid {
		for x, value := range row {
			if value == '1' {
				visitIsland(grid, y, x)
				islandCount++
			}
		}
	}

	return islandCount
}
