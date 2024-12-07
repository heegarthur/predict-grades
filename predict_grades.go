package grades_bot

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadData(filename string) map[string]int {
	data := make(map[string]int)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return data
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err == nil {
				data[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return data
}

func saveData(filename string, data map[string]int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for key, value := range data {
		line := fmt.Sprintf("%s:%d\n", key, value)
		writer.WriteString(line)
	}
	writer.Flush()
}

func updateData(filename string, data map[string]int) {
	var key string
	var value int
	fmt.Println("Enter the key to update or add:")
	fmt.Scanln(&key)
	fmt.Println("Enter the value:")
	fmt.Scanln(&value)
	data[key] = value
	saveData(filename, data)
	fmt.Println("Data updated successfully.")
}

func predict() {
	filename := "customizebot.txt"
	scores := loadData(filename)

	var learning_duration, subject, difficulty, material string

	for {
		fmt.Println("\nONLY SMALL LETTERS\nSPACES ARE '_'\nVery Good❌\nvery_good✅")

		fmt.Println("\nHow much have you learned?\nnot\na bit\nmoderate\nnormal\ngood\nvery good")
		fmt.Scanln(&learning_duration)
		if _, exists := scores[learning_duration]; !exists {
			fmt.Println("Invalid input. Please choose from the given options.")
			continue
		}
		break
	}

	for {
		fmt.Println("\nWhich subject?")
		fmt.Scanln(&subject)
		if _, exists := scores[subject]; !exists {
			fmt.Println("Invalid input. Please enter a valid subject.")
			continue
		}
		break
	}

	for {
		fmt.Println("\nHow difficult?\neasy\nnormal\ndifficult\nvery difficult")
		fmt.Scanln(&difficulty)
		if _, exists := scores[difficulty]; !exists {
			fmt.Println("Invalid input. Please choose from the given options.")
			continue
		}
		break
	}

	for {
		fmt.Println("\nHow much material?\nlittle\nnormal amount\nmuch\nvery much")
		fmt.Scanln(&material)

		if _, exists := scores[material]; !exists {
			fmt.Println("Invalid input. Please choose from the given options.")
			continue
		}
		break
	}

	final_prediction := 0
	if scores[difficulty] == 0 {
		totalScore := scores[learning_duration] + scores[subject] + scores[material]
		final_prediction = totalScore / 3
	} else {
		totalScore := scores[learning_duration] + scores[subject] + scores[difficulty] + scores[material]
		final_prediction = totalScore / 4
	}

	if final_prediction <= 3 {
		fmt.Println("\nNo! Your grade:", final_prediction, "\nYou're COOKED, go learn now")
	} else if final_prediction > 3 && final_prediction < 5 {
		fmt.Println("Bad:", final_prediction, "That's just bad, you can do this, go learn now")
	} else if final_prediction >= 5 && final_prediction < 6 {
		fmt.Println("\nNot bad, but can be better, start now and then do something fun\n", final_prediction)
	} else if final_prediction >= 6 && final_prediction < 7 {
		fmt.Println("Nice, plan when you're going to learn and then chill (if that's according to your schedule)\n", final_prediction)
	} else if final_prediction >= 7 && final_prediction < 9 {
		fmt.Println("\nGood, keep going like this:", final_prediction)
	} else {
		fmt.Println("Great, you're doing this\nGo on like this:", final_prediction)
	}
}

func main_predict() {
	filename := "customizebot.txt"
	scores := loadData(filename)
	fmt.Println("Choose an option:\n1. Predict\n2. Update Data")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		predict()
	} else if choice == 2 {
		updateData(filename, scores)
	} else {
		fmt.Println("Invalid choice.")
	}
}
