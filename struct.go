package main
import ("fmt")

type Person struct {
	Name string
	Age int
}

func main() {
	p := Person{Name: "Noornabi", Age: 23}
	fmt.Println("Person Name:", p.Name)
	fmt.Println("Person Age:", p.Age)

	// Modifying struct fields
	p.Age = 24
	fmt.Println("Updated Person Age:", p.Age)

	// Using pointers to struct
	pPointer := &p
	fmt.Println("Person Name via pointer:", pPointer.Name)
	fmt.Println("Person Age via pointer:", pPointer.Age)

	// Modifying struct fields via pointer
	pPointer.Age = 25
	fmt.Println("Updated Person Age via pointer:", p.Age) // This will reflect the change in the original struct
}