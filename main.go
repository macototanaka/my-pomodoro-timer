package main

import (
	"fmt"
	"time"
)

func main() {
	var input string
	timerRunning := false

	for {
		fmt.Println("Enter 'start' to start the timer or 'stop' to stop it:")
		fmt.Scanln(&input)

		if input == "start" && !timerRunning {
			timerRunning = true
			go startPomodoroTimer()
		} else if input == "stop" && timerRunning {
			timerRunning = false
			fmt.Println("Timer stopped.")
		} else {
			fmt.Println("Invalid input or timer already running.")
		}
	}
}

func startPomodoroTimer() {
	for i := 0; i < 4; i++ {
		runTimer(25, "Work")
		runTimer(5, "Break")
	}
	fmt.Println("Pomodoro session complete!")
}

func runTimer(minutes int, timerType string) {
	fmt.Printf("%s timer started for %d minutes.\n", timerType, minutes)
	time.Sleep(time.Duration(minutes) * time.Minute)
	fmt.Printf("%s timer ended.\n", timerType)
}
