package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card rune

func (c Card) String() string {
	return string(c)
}

func (c Card) Compare(other Card) int {
	return cardRanking[c] - cardRanking[other]
}

var cards = []Card{
	'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J',
}

var cardRanking = map[Card]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'J': 3,
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	hands := parseHands(file)

	total := 0
	for i, hand := range hands {
		rank := i + 1
		fmt.Println(rank, hand.cards, hand.handType)
		fmt.Println(hand.bid)

		total += rank * hand.bid
		fmt.Println(total)
	}

	println(total)
}

func parseHands(f *os.File) []Hand {
	var hands []Hand

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		hand := Hand{
			cardCounts: make(map[Card]int),
			counts:     make(map[int][]Card),
		}

		fields := strings.Fields(sc.Text())

		if bid, err := strconv.Atoi(fields[1]); err == nil {
			hand.bid = bid
		} else {
			panic(err)
		}

		for _, c := range fields[0] {
			hand.cards = append(hand.cards, Card(c))
			hand.cardCounts[Card(c)]++

			if c == 'J' {
				hand.jokerCount++
			}
		}

		for r, c := range hand.cardCounts {
			hand.counts[c] = append(hand.counts[c], r)
		}

		hand.handType = hand.calculateHand()

		hands = append(hands, hand)
	}

	// sort the hands asc
	slices.SortFunc(hands, func(i, j Hand) int {
		if j.handType == i.handType {
			return j.Compare(i)
		}

		return int(j.handType) - int(i.handType)
	})

	return hands
}

type Hand struct {
	cards []Card

	cardCounts map[Card]int

	counts map[int][]Card

	bid int

	jokerCount int

	handType HandType
}

func (h Hand) Compare(other Hand) int {
	for i := 0; i < len(h.cards); i++ {
		if h.cards[i] != other.cards[i] {
			return h.cards[i].Compare(other.cards[i])
		}
	}

	return 0
}

type HandType int

const (
	FIVE_OF_A_KIND HandType = iota + 1
	FOUR_OF_A_KIND
	FULL_HOUSE
	THREE_OF_A_KIND
	TWO_PAIR
	ONE_PAIR
	HIGH_CARD
)

var hand_types = map[HandType]string{
	FIVE_OF_A_KIND:  "five of a kind",
	FOUR_OF_A_KIND:  "four of a kind",
	FULL_HOUSE:      "full house",
	THREE_OF_A_KIND: "three of a kind",
	TWO_PAIR:        "two pair",
	ONE_PAIR:        "one pair",
	HIGH_CARD:       "high card",
}

func (r HandType) String() string {
	return hand_types[r]
}

func (h Hand) calculateHand() HandType {
	if _, ok := h.counts[5]; ok {
		return FIVE_OF_A_KIND
	}

	if _, ok := h.counts[4]; ok {
		if h.jokerCount > 0 {
			return FIVE_OF_A_KIND
		}

		return FOUR_OF_A_KIND
	}

	_, three_ok := h.counts[3]
	two, two_ok := h.counts[2]

	if three_ok && h.jokerCount == 2 {
		return FIVE_OF_A_KIND // if two jokers, 5 of a kind
	}

	if three_ok && h.jokerCount == 3 || three_ok && h.jokerCount == 1 {
		return FOUR_OF_A_KIND // if any joker, it's always a four of a kind
	}

	if three_ok && two_ok {
		return FULL_HOUSE
	}

	if three_ok {
		return THREE_OF_A_KIND
	}

	if two_ok && len(two) == 2 {
		if h.jokerCount == 2 {
			return FOUR_OF_A_KIND
		} else if h.jokerCount == 1 {
			return FULL_HOUSE
		} else {
			return TWO_PAIR
		}
	}

	if two_ok && len(two) == 1 {
		if h.jokerCount > 0 {
			return THREE_OF_A_KIND
		} else {
			return ONE_PAIR
		}
	}

	if h.jokerCount == 1 {
		return ONE_PAIR
	}

	return HIGH_CARD
}
