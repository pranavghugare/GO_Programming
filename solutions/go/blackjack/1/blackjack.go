package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
    switch  {
        case card== "ace":
        return 11;
    	case card== "two":
     	return 2;
    	case card== "three":
        return 3;
    	case card== "four":
     	return 4;
    	case card== "five":
        return 5;
    	case card== "six":
     	return 6;
    	case card== "seven":
        return 7;
        case card== "eight":
     	return 8;
    	case card=="nine":
        return 9;
		case (card== "ten" || card== "jack" || card== "queen" || card=="king"):
     	return 10;
    	default:
     	return 0;
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")
   playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case playerScore == 21:
		if dealerScore < 10 {
			return "W"
		}
		return "S"
	case playerScore >= 17 && playerScore <= 20:
		return "S"
	case playerScore >= 12 && playerScore <= 16:
		if dealerScore >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
