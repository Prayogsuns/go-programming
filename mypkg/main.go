// Package mysum provides sum tools for myself
// Its a DIY work for me
package mysum

// MySum adds all the numbers provided together
func MySum(ni ...int) int {
	mysum := 0
	for _, v := range ni {
		mysum += v
	}
	return mysum
}


