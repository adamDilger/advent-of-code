package main

import "testing"

func TestAdd(t *testing.T) {

	tests := []struct {
		hand string
		want HandType
	}{
		{"23456", HIGH_CARD},
		{"2J456", ONE_PAIR},
		{"2JJ56", THREE_OF_A_KIND},
		{"2JJJ6", FOUR_OF_A_KIND},
		{"2JJJJ", FIVE_OF_A_KIND},

		{"22456", ONE_PAIR},
		{"22J56", THREE_OF_A_KIND},
		{"22JJ6", FOUR_OF_A_KIND},
		{"22JJJ", FIVE_OF_A_KIND},

		{"22256", THREE_OF_A_KIND},
		{"222J6", FULL_HOUSE},
		{"222JJ", FIVE_OF_A_KIND},

		{"22226", FOUR_OF_A_KIND},
		{"2222J", FIVE_OF_A_KIND},

		{"J2345", ONE_PAIR},
		{"JJ234", THREE_OF_A_KIND},
		{"JJJ34", FOUR_OF_A_KIND},
		{"JJJJ4", FIVE_OF_A_KIND},
	}

	for _, tt := range tests {
		hand := NewHand("0", tt.hand)

		got := hand.handType
		want := tt.want

		if got != want {
			t.Errorf("%s: got %s, wanted %s", tt.hand, got, want)
		}
	}
}
