package main

func grayCode(n int) []int {
    if n == 0 {
        return []int{0}
    }
    
    result := []int{0, 1}
    
    for i := 2; i <= n; i++ {
        size := len(result)
        powerOf2 := 1 << (i - 1)
        
        for j := size - 1; j >= 0; j-- {
            result = append(result, powerOf2+result[j])
        }
    }
    
    return result
}