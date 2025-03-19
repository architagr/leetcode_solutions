package allpathsfromsourceleadtodestination

var (
	white int8 = 0
	gray  int8 = 1
	black int8 = 2
)

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	aj := buildAj(n, edges)
	c := make([]int8, n)
	return dfs(aj, source, destination, c)
}

func dfs(aj map[int][]int, source, destination int, c []int8) bool {
	if c[source] != white {
		return c[source] == black
	}
	if len(aj[source]) == 0 {
		return source == destination
	}
	c[source] = gray
	for _, n := range aj[source] {
		if !dfs(aj, n, destination, c) {
			return false
		}
	}
	c[source] = black
	return true
}
func buildAj(n int, edges [][]int) map[int][]int {
	aj := make(map[int][]int, n)
	for _, e := range edges {
		aj[e[0]] = append(aj[e[0]], e[1])
	}
	return aj
}
