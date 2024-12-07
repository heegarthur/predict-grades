package grades_bot

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_motivation() {
	motivation_data := []string{"trade offer: now 120 minutes of learning in exchange for a pass",
		"not learning when you have to is not REALLY enjoying, doing nothing when you don't have to and when you have good grades is REALLY enjoying",
		"it's not too late to learn, in 2 hours you can still get a MUCH higher grade: so start now and REALLY go on until you can't",
		"keep going until you can't learn anymore, not until you don't feel like it",
		"if you don't feel like learning, you probably do feel like getting bad grades",
		"drink water, go to the bathroom and maybe briefly outside if possible and then learn: that goes better",
		"stupid things (school) make programming and fun things like this more enjoyable",
		"reward yourself with good grades after learning 'but that's super lame' well then with rest and programming and stuff 'pfff' okay also with friends 'nerd' okay also with having time for everything",
		"you can study better unexpectedly with the methods chooser here"}
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	random_index := r.Intn(len(motivation_data))
	random_quote := motivation_data[random_index]
	println(random_quote)
}
func choose_studyway() {
	var vak_type string
	fmt.Println("Choose a number:")
	fmt.Println("1 - Creative subject: drawing, drama etc.")
	fmt.Println("2 - Mathematical subject: physics, mathematics etc.")
	fmt.Println("3 - Language: German, Spanish etc.")
	fmt.Println("4 - Other things: history, geography etc.")
	fmt.Scanln(&vak_type)
	random_study_manier_kies(vak_type)
}

func random_study_manier_kies(vak_type string) {
	var studiemethoden []string

	switch vak_type {
	case "1":
		studiemethoden = []string{
			"Put on music and practice",
			"Make a timelapse of yourself while practicing, so you don't sit on your phone",
			"just do it and when you take a break you have to do as many push-ups as minutes before the previous break/beginning",
		}
	case "2":
		studiemethoden = []string{
			"Do 1 more push-up per sum",
			"set a timer from random 30 minutes to 2 hours (don't look) and do exercises until the timer runs out",
			"do 1 hour of exercises and then do 1 hour of a random mission here: https://heegarthur.github.io/ivoarthur/tasks.html",
		}
	case "3":
		studiemethoden = []string{
			"Make a short story with grammar and words",
			"test yourself by covering and guessing words/grammar, 1 push-up for 1 mistake",
			"make a list in studygo: https://studygo.com/nl/learn/create-list",
		}
	case "4":
		studiemethoden = []string{
			"summarize everything",
			"mark important things and read everything 3 times",
			"ask an AI to summarize the text so you can understand it easily, for example in programming language if you program and game language if you game etc",
		}
	default:
		fmt.Println("Invalid choice. Try again.")
		return
	}

	// Choose a random study method
	randomIndex := rand.Intn(len(studiemethoden))

	// Print the selected study method
	fmt.Println("You can try this study method: ", studiemethoden[randomIndex])
}
