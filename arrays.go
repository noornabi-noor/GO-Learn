package main
import ("fmt")

func main() {
	var a [5]int
	fmt.Println("Enter 5 numbers:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scan(&a[i])
	}
	fmt.Println("You entered:")
	for i := 0; i < 5; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}