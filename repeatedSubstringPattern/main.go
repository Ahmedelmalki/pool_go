package main

import "fmt"

func main(){
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}

func repeatedSubstringPattern(s string) bool {
    n := len(s)
    for i := 1; i <= n/2; i++ {
        if n%i == 0 {
            sub := s[:i]
            repeated := true
            for j := i; j < n; j += i {
                if s[j:j+i] != sub {
                    repeated = false
                    break
                }
            }
            if repeated {
                return true
            }
        }
    }
    return false
}
