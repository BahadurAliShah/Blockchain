package main

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {
	emplys := []employee{
		employee{"Bob", 20, "Manager"},
		employee{"Alice", 30, "CEO"},
		employee{"John", 40, "CTO"},
	}
	comp := company{"Google", emplys}

	println("Company: ", comp.companyName)
	println("Employees:")
	for _, v := range comp.employees {
		println("\tName:", v.name, "Salary:", v.salary, "Position:", v.position)
	}
}
