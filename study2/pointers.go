package main

func playWithPointers() {
	variable := "Sergio"
	var pointer *string
	pointer = &variable
	*pointer = "Martinez"
	println(variable)

	variable2 := "James"
	pointer2 := &variable2
	variable2 = "Bond"
	println(variable2)
	println(*pointer2)
}
