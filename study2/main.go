package main

func main() {
	/* 	// There are 3 scopes in Go
	   	// 1. Global
	   	// 2. Local
	   	// 3. Function

	   	// There are 2 types of create a slice
	   	var fruits = []string{"apple", "banana", "orange"}
	   	for _, fruit := range fruits {
	   		println(fruit)
	   	}
	   	println("")
	   	println("Option 2")
	   	for i := 0; i < len(fruits); i++ {

	   		println(fruits[i])
	   	}

	   	num := make([]int, 1)
	   	slice := []int{}
	   	println("")
	   	for i := 0; i < 5; i++ {
	   		num[i] = i
	   		slice = append(slice, i)
	   		num = append(num, i)
	   		println(num[i])
	   		println(slice[i])
	   		println("")
	   	}

	   	fmt.Printf("Cap %v \n", cap(num))
	   	fmt.Printf("Cap %v \n", cap(slice))
	   	fmt.Printf("Cap %d \n", len(num))
	   	fmt.Printf("Cap %d \n", len(slice))

	   	users := make(map[string]User)
	   	users["Sergio"] = User{Name: "Sergio", Age: 30, City: "Madrid", Country: "Spain"}

	   	programmers := make(map[int]Programmer)
	   	programmers[1] = Programmer{Person: users["Sergio"], Id: 1}

	   	for _, v := range users {
	   		v.PrintName()
	   		v.PrintAge()
	   		v.PrintAddress()
	   	}

	   	println("Today's date " + getDateTime(getDate, getTime)) */

	//playWithPointers()

	//juan := Employee{name: "Juan", age: 30}

	/* 	greetings(juan)
	   	genericFunction(10)
	   	genericFunction(10.5)
	   	genericFunction("Hello")
	   	genericFunction(true) */
	//goroutines()
	//mainChannels()
	mainWaitGroups()
}

func greetings(g Greeter) {
	g.sayHello()
}

/* func getDate() string {
	return "2020-01-01"
}

func getTime() string {
	return "12:00:00"
}

func getDateTime(getDate func() string, getTime func() string) string {
	defer println("i am a defer") // defer is executed after the function returns
	println("i am a println")
	return getDate() + " " + getTime()
}
*/
