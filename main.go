package grades_bot

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {

		var choice string
		println("Enter a number:\n1. Predict your grade\n2. Motivation\n3. Choose a study method\n4. Quit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			main_predict()
		case "2":
			generate_motivation()

		case "3":
			choose_studyway()
		case "4":
			quitDirectly()

		default:
			println("Invalid choice: ", choice, " try again.")
		}
	}
}

func quitDirectly() {
	for i := 0; i < 100; i++ {
		print("-")
		time.Sleep(50 * time.Millisecond)
	}
	println("\nstart to learn NOW")
	os.Exit(0)

}
