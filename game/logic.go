package game

import (
	"math/rand"
	"time"
)

var choices = []string{"rock", "paper", "scissors"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetComputerChoice() string {
	return choices[rand.Intn(len(choices))]
}

func DecideWinner(player, computer string) string {
	if player == computer {
		return "draw"
	}

	if (player == "rock" && computer == "scissors") ||
		(player == "scissors" && computer == "paper") ||
		(player == "paper" && computer == "rock") {
		return "win"
	}

	return "lose"
}
