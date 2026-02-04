package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// | card  | value | card    | value |
	// | :---: | :---: | :-----: | :---: |
	// |  ace  |  11   | eight   |   8   |
	// |  two  |   2   | nine    |   9   |
	// | three |   3   |  ten    |  10   |
	// | four  |   4   | jack    |  10   |
	// | five  |   5   | queen   |  10   |
	// |  six  |   6   | king    |  10   |
	// | seven |   7   | *other* |   0   |
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	sumOfCards := card1Value + card2Value
	switch {
	// - If you have a pair of aces you must always split them.
	case card1 == "ace" && card2 == "ace":
		return "P"
	// - If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a face card (Jack/Queen/King) or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	case sumOfCards == 21 && dealerCard != "ace" && dealerCardValue != 10:
		return "W"
	case sumOfCards == 21 && !(dealerCard != "ace" && dealerCardValue != 10):
		return "S"
	// - If your cards sum up to a value within the range [17, 20] you should always stand.
	case sumOfCards >= 17 && sumOfCards <= 20:
		return "S"
	// - If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
	case sumOfCards >= 12 && sumOfCards <= 16 && dealerCardValue < 7:
		return "S"
	case sumOfCards >= 12 && sumOfCards <= 16 && dealerCardValue >= 7:
		return "H"
		// - If your cards sum up to 11 or lower you should always hit.
	case sumOfCards < 12:
		return "H"
	}
	return ""
}
