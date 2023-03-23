package interfaceshexagonal

import "fmt"

// Before understanding the hexagonal architecture , it is pre-requisite to understand
// and probably use interfaces in Go.
// Currently, as per my current understanding, interface is something that two different
// modules use to talk to each other.
// So now , at first, I will study the interfaces in detail then would visit the
// hexagonal architecture.

// I have written a beautiful doc explaining what exactly interface is and how it
// can be used , all thanks to a great well written blog on medium!!

// I am all set to write and implement a interface program in Go!!

type Duck interface {
	Quack()
	Walk()
}

type Donald struct {
	Name string
}

func (d Donald) Quack() {
	fmt.Println("I Quack")
}

func (d Donald) Walk() {
	fmt.Println("I walk")
}

func StartInterface() {
	d := new(Donald)
	d.Quack()
	d.Walk()
}

// Will now write a code which would actually demonstrate how interfaces can be useful in
// Go and how they help in scaliability and re-usiability and expansion as well!
// And how the concept of run time polymorphism is used in the case of interfaces!

type CalculateSalary interface {
	TotalSalarySum() int
}

type Developer struct {
	Name            string
	BasicSalary     int
	DeveloperSalary int
	TotalSalary     int
}

type Manager struct {
	Name          string
	BasicSalary   int
	ManagerSalary int
	TotalSalary   int
}

func (d Developer) TotalSalarySum() int {
	fmt.Println("Total salary of developer is 100$")
	//d.TotalSalary = d.BasicSalary + d.DeveloperSalary
	return d.BasicSalary + d.DeveloperSalary
}

func (m Manager) TotalSalarySum() int {
	fmt.Println("Total salary of Manager is 300$")
	//m.TotalSalary = m.BasicSalary + m.ManagerSalary
	return m.BasicSalary + m.ManagerSalary
}

func EmployerTotalSalaryExpense(I []CalculateSalary) {
	fmt.Printf("Total Salary using actual interface usage is %d", I[1].TotalSalarySum()+I[0].TotalSalarySum())

}

func StartInterface2() {
	d := Developer{Name: "Developer1", BasicSalary: 100, DeveloperSalary: 200, TotalSalary: 300}
	d.TotalSalarySum()

	m := Manager{Name: "Manager1", BasicSalary: 100, ManagerSalary: 400, TotalSalary: 500}
	m.TotalSalarySum()

	/* Here is where the actual use of interface aka runtime polymorphism happens
	the EmployerTotalSalaryExpense is kind of generic code which does dynamic type casting
	during the run time based on which object was passed to it as the argument!!!
	*/
	em := []CalculateSalary{d, m}

	EmployerTotalSalaryExpense(em)
}
