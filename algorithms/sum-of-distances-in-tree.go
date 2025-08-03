package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
    graph := make([][]int, n)
    count := make([]int, n)
    result := make([]int, n)
    
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }
    
    var dfs1 func(node, parent int)
    dfs1 = func(node, parent int) {
        count[node] = 1
        result[node] = 0
        
        for _, child := range graph[node] {
            if child != parent {
                dfs1(child, node)
                count[node] += count[child]
                result[node] += result[child] + count[child]
            }
        }
    }
    
    var dfs2 func(node, parent int)
    dfs2 = func(node, parent int) {
        for _, child := range graph[node] {
            if child != parent {
                result[child] = result[node] - count[child] + (n - count[child])
                dfs2(child, node)
            }
        }
    }
    
    dfs1(0, -1)
    dfs2(0, -1)
    
    return result
}