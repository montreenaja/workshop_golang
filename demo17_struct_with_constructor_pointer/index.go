package main

import ("demo15/employee")

func main(){
	e := employee.Init(
		"Prem",
		"Adolf",
		30,
		20,
	)
	e.LeavesRemaining()

	e = employee.Init(
		"Sek",
		"Adolf",
		30,
		20,
	)
	e.LeavesRemaining()
}