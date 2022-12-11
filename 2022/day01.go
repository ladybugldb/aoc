package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var fileLines []string
	var elfCalorieCount []int

	//Pass in file name
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	//Open file
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	//Parse file by line
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	elf := 0           //Number of elfves
	totalCalories := 0 //temp total per elf
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())

		//If not a new line snack is for the same elf, if it is a new line move on to next elf
		if fileScanner.Text() != "" {

			//covert string to text
			tempCalories, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			totalCalories = totalCalories + tempCalories
		} else {
			//append total for current elf
			elfCalorieCount = append(elfCalorieCount, totalCalories)

			//reset temp total
			totalCalories = 0

			//move on to next elf
			elf++
		}
	}
	readFile.Close()

	//find elves with most calories

	//get number from input
	numOfTopElves, err := strconv.Atoi(os.Args[2])
	totalCalories = 0

	//sort array of calorie counts
	sort.Ints(elfCalorieCount)

	//total the top x counts
	for i := 0; i < numOfTopElves; i++ {
		totalCalories = totalCalories + elfCalorieCount[len(elfCalorieCount)-i-1]
	}
	fmt.Println("Top", numOfTopElves, "elf as the most caloires with a toal of", totalCalories)

}
