package main

import "fmt"

func main() {
    // Test Gray Code - harder cases
    fmt.Println("Testing Gray Code:")
    fmt.Println("n=4:", grayCode(4)) // 16 numbers
    fmt.Println("n=3:", grayCode(3)) // 8 numbers
    fmt.Println("n=5 length:", len(grayCode(5))) // Should be 32
    
    // Test Sum of Distances in Tree - harder cases
    fmt.Println("\nTesting Sum of Distances in Tree:")
    
    // Larger tree with 10 nodes - linear chain
    edges1 := [][]int{{0,1},{1,2},{2,3},{3,4},{4,5},{5,6},{6,7},{7,8},{8,9}}
    fmt.Println("Linear chain 10 nodes:", sumOfDistancesInTree(10, edges1))
    
    // Star graph with center at node 5
    edges2 := [][]int{{5,0},{5,1},{5,2},{5,3},{5,4},{5,6},{5,7}}
    fmt.Println("Star graph 8 nodes:", sumOfDistancesInTree(8, edges2))
    
    // Binary tree structure
    edges3 := [][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6},{3,7},{3,8}}
    fmt.Println("Binary tree 9 nodes:", sumOfDistancesInTree(9, edges3))
    
    // Test Maximum Length of Repeated Subarray - harder cases
    fmt.Println("\nTesting Maximum Length of Repeated Subarray:")
    
    // Longer arrays with multiple matches
    nums1 := []int{1,2,3,4,5,6,7,8,9,1,2,3,4,5}
    nums2 := []int{9,1,2,3,4,5,10,11,12,13,14,15}
    fmt.Println("Long arrays with overlap:", findLength(nums1, nums2)) // Should be 5
    
    // No common subarray
    fmt.Println("No overlap:", findLength([]int{1,2,3}, []int{4,5,6})) // Should be 0
    
    // Entire arrays are the same
    same1 := []int{7,8,9,10,11}
    same2 := []int{7,8,9,10,11}
    fmt.Println("Identical arrays:", findLength(same1, same2)) // Should be 5
    
    // Complex pattern
    complex1 := []int{0,1,1,1,1,0,0,1,1,0}
    complex2 := []int{1,0,1,1,1,1,0,0,1,1}
    fmt.Println("Complex pattern:", findLength(complex1, complex2)) // Should find longest match
}