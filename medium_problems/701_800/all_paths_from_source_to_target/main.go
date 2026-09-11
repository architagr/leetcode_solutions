package allpathsfromsourcetotarget

func allPathsSourceTarget(graph [][]int) [][]int {

	allPaths := make([][]int, 0)
	current := make([]int, 0)
	dfs(0, graph, &allPaths, &current)
	return allPaths
}

func dfs(node int, graph [][]int, ans *[][]int, current *[]int) {
	*current = append(*current, node)
	defer func() {
		*current = (*current)[:len((*current))-1]
	}()
	if node == len(graph)-1 {
		*ans = append(*ans, append([]int{}, (*current)...))
		return
	}

	for _, n := range graph[node] {
		dfs(n, graph, ans, current)
	}

}
