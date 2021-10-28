package repo

import "privateFlashCard/entities"

func IdCheck(id string, deck entities.FlashCardStruct) (interface{}, error) {
	var returnDeck interface{}
	if len(deck.Matching) > 0 {
		for _, v := range deck.Matching {
			if v.Id == id {
				returnDeck = v
				return returnDeck, nil
			}

		}
	}

	if len(deck.TrueFalse) > 0 {
		for _, v := range deck.TrueFalse {
			if v.Id == id {
				returnDeck = v
				return returnDeck, nil
			}

		}
	}

	if len(deck.Info) > 0 {
		for _, v := range deck.Info {
			if v.Id == id {
				returnDeck = v
				return returnDeck, nil
			}

		}
	}

	if len(deck.Multiple) > 0 {
		for _, v := range deck.Multiple {
			if v.Id == id {
				returnDeck = v
				return returnDeck, nil
			}

		}
	}

	if len(deck.QandA) > 0 {
		for _, v := range deck.QandA {
			if v.Id == id {
				returnDeck = v
				return returnDeck, nil
			}

		}
	}

	return returnDeck, NotFound

}
