package graph

func dfs(current int, visited []bool, isConnected [][]int) {
	visited[current] = true
	for next := 0; next < len(isConnected); next++ {
		if isConnected[current][next] == 0 {
			continue
		}
		if !visited[next] {
			dfs(next, visited, isConnected)
		}
	}
}

func NumberOfProvinces(isConnected [][]int) int {
	count := 0
	visited := make([]bool, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			count++
			dfs(i, visited, isConnected)
		}
	}
	return count
}
