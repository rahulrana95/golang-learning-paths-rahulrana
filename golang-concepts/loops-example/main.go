package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	daysMap := make(map[int]string)
	daysMap[0] = "Monday"
	daysMap[1] = "Tuesday"
	daysMap[2] = "Wednesday"
	daysMap[3] = "Thursday"
	daysMap[4] = "Friday"
	daysMap[5] = "Saturday"

	fmt.Println(days)
	fmt.Println(daysMap)

	// iterating over arrays
	for i := range days {
		fmt.Println(days[i])
	}

}
