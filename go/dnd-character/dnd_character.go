package dndcharacter

import (
	"math"
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
	return int(math.Floor((float64(score-10) / 2)))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	var scores [6]int
	for i := 0; i < 6; i++ {
		scores[i] = rand.Intn(6) + 1
	}
	sort.Ints(scores[:])
	return scores[3] + scores[4] + scores[5]
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	// You find your character's constitution modifier by subtracting 10 from your character's constitution, divide by 2 and round down
	constitution := Ability()
	hitpoints := Modifier(constitution) + 10
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    hitpoints,
	}
}
