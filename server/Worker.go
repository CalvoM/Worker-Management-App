package main

type Worker struct{
	Name string `json:"name"`
	Age int8 `json:"age"`
	Salary float64 `json:"salary"`
	Email string `json:"email"`
}
type Response struct{
	Data []*Worker
}
