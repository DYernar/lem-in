package main

type Room struct {
	Name string 
	X int
	Y int
	Connections []*Room
	Start bool
	End bool
	Visited bool
	Distance int
}

type AllData struct {
	AntNum int
	Rooms []*Room
}

type Node struct {
	Val Room
	Next *Node
}