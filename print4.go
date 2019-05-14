/**
*  
*  *****
*   ***
*    *
*/
package main

import "fmt"

func main() {
	n := 5
        for i := 0; i < n; i++ {
		for k :=0; k<i; k++ {
			fmt.Print(" ")
		}
		for j :=0; j < 2*(n-i)-1; j++ {
			fmt.Print("*")
		}
        	fmt.Println("")
        }
}
