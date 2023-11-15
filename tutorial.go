package main

import "fmt"

func main() {
	fmt.Println("Willkommen zu unserem kleinen Quiz.")

	//Durch Printf wird verhindert, dass nach dem Print in die nächste Zeile gesprungen wird. Bei Println springt das Programm nach dem print in eine neue Zeile.
	fmt.Printf("Wie ist dein Name?:")

	var name string
	//Input, der den Input in der variable speichert
	fmt.Scan(&name)

	fmt.Printf("Hi %v. Dann schauen wir mal, was Du kannst!\n", name)

	fmt.Printf("Wie alt bist Du?:")
	var age uint //unit verhindert, dass der int negativ werden kann
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Alles klar. Du bist alt genug um zu spielen. ")
	} else {
		fmt.Println("Sorry, Du bist zu jung.")
		return //beendet die main function und damit das Quiz
	}

	score := 0
	number_questions := 2

	fmt.Printf("Welche Grafikkarte ist besser? RTX 3080 oder RTX 3090?:")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2) // Problem mit .Scan ist in diesem Fall, dass nur "RTX" als Antwort erscheint, da RTX die einzige Variable ist. 3080 wäre eine eigene Variable. Um das zu verhinder 2. Var

	if answer+" "+answer2 == "RTX 3090" {
		println("Richtig!")
		score += 1
	} else if answer+" "+answer2 == "rtx 3090" {
		println("Richtig!")
		score += 1
	} else if answer+" "+answer2 == "rtx3090" {
		println("Richtig!")
		score += 1
	} else {
		println("Falsch!")
	}

	fmt.Printf("Welche Zahl ist größer? 1 oder 10?")
	var answer3 uint
	fmt.Scan(&answer3)

	if answer3 == 10 {
		println("Richtig!")
		score += 1
	} else {
		println("Falsch!")
	}

	var success_rate int = (score * 100) / number_questions

	fmt.Printf("Glückwunsch du hast %v Punkte erreicht. Damit hast du %v Prozent der Fragen korrekt beantwortet.", score, success_rate)
}
