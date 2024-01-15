package main

import (
	"fmt"
)

func main() {
	stringOfAgents := GetLatestList()

	// work directly with list of user agents
	useragents := turnToJson(stringOfAgents)

	fmt.Println("useragents:", useragents)

	fileName := generateFileName()

	err := saveIntoFile(fileName,stringOfAgents)

	if err != nil {
		fmt.Println("Error : ",err)
	}
}