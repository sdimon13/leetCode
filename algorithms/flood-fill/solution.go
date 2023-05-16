package flood_fill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	sourceColor := image[sr][sc]
	if sourceColor != newColor {
		dfs(image, sr, sc, sourceColor, newColor)
	}
	return image
}

func dfs(image [][]int, x int, y int, sourceColor int, newColor int) {
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[0]) || image[x][y] != sourceColor {
		return
	}
	image[x][y] = newColor
	dfs(image, x-1, y, sourceColor, newColor)
	dfs(image, x+1, y, sourceColor, newColor)
	dfs(image, x, y-1, sourceColor, newColor)
	dfs(image, x, y+1, sourceColor, newColor)
}
