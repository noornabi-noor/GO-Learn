// package main
// import ("fmt")

// func change(x *int) {
// 	*x = 10
// }

// func main() {
// 	y := 5
// 	change(&y)
// 	fmt.Println(y) // This will print 10, not 5
// }



package main
import ("fmt")


func main() {
	x := 15
	fmt.Println("Address of x:", &x)

	y := &x
	fmt.Println("Value of y (address of x):", y)
	fmt.Println("Value at address y (value of x):", *y)
}