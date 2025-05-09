package dndcharacter

import (
	"math/rand"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	diff := score - 10
	if diff < 0 {
		diff += diff % 2
	}
	return diff / 2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	rolls := []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}
	sort.Ints(rolls)
	return rolls[1] + rolls[2] + rolls[3]
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	con := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: con,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(con),
	}
}
