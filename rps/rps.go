package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡BIEEEEEN le ganaste a MEGABOT!",
	"Increible",
	"Tenes dominado a MegaBot, te felicito...",
}

var loseMessages = []string{
	"¡Casi.. la proxima será!",
	"Megabot es dificil de ganar... intentalo de nuevo",
	"¡Destruido por MegaBot!",
}

var drawMessages = []string{
	"Las grandes mentes piensan igual..",
	"Nadie gana, intentalo de nuevo...",
	"Sos un Robot",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {

	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {

	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "MegaBot eligió PIEDRA"

	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "MegaBot eligió PAPEL"

	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "MegaBot eligió TIJERA"
	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"

		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "¡El jugador gana!"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "¡La computadora gana!"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
