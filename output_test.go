package main

import (
	"os"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/nikkon226/anarchy-solo-challenge/internal/assert"
)

func TestComputeOutput(t *testing.T) {

	f, err := os.Open("components.toml")
	if err != nil {
		panic(err)
	}

	var components Components
	_, err = toml.NewDecoder(f).Decode(&components)
	if err != nil {
		panic(err)
	}

	f.Close()

	output := ComputeOutput(1, 2, components)

	assert.Equal(t, output.PlunderTilesByRound, map[int]string{1: "material,material,food", 2: "silver,material,material", 3: "material,material", 4: "silver,material", 5: "material,food"})
	assert.Equal(t, output.PathCardsByRound, map[int][]string{1: {"Foreman", "Warrior"}, 2: {"Champion", "Scout"}, 3: {"Engineer", "Weaponsmith"}, 4: {"Mentor", "Tactician"}, 5: {"Silversmith", "Baron"}})
	assert.Equal(t,
		output.AttackCardsByRound,
		map[int][]AttackCard{
			0: {
				{ID: 11, Type: "catapault", North: "", East: "", South: "A", West: "", Skull: "true"},
			},
			1: {
				{ID: 22, Type: "siege_tower", North: "", East: "", South: "", West: "A", Skull: "false"},
				{ID: 34, Type: "arrows", North: "B", East: "C", South: "", West: "C", Skull: "false"},
			},
			2: {
				{ID: 2, Type: "ladder", North: "", East: "A", South: "", West: "", Skull: "false"},
				{ID: 35, Type: "arrows", North: "B", East: "", South: "B", West: "D", Skull: "true"},
				{ID: 15, Type: "ballista", North: "", East: "C", South: "", West: "B", Skull: "false"},
			},
			3: {
				{ID: 32, Type: "arrows", North: "C", East: "", South: "C", West: "B", Skull: "false"},
				{ID: 36, Type: "arrows", North: "B", East: "B", South: "", West: "D", Skull: "true"},
				{ID: 14, Type: "ballista", North: "B", East: "", South: "C", West: "", Skull: "false"},
				{ID: 3, Type: "ladder", North: "", East: "", South: "A", West: "", Skull: "false"},
			},
			4: {
				{ID: 8, Type: "catapault", North: "", East: "", South: "C", West: "", Skull: "false"},
				{ID: 34, Type: "arrows", North: "B", East: "C", South: "", West: "C", Skull: "false"},
				{ID: 18, Type: "ballista", North: "B", East: "", South: "A", West: "", Skull: "true"},
				{ID: 33, Type: "arrows", North: "", East: "C", South: "B", West: "C", Skull: "false"},
				{ID: 4, Type: "ladder", North: "", East: "", South: "", West: "A", Skull: "false"},
			},
		},
	)
	assert.Equal(t,
		output.FinalEscaladeCardByRound,
		map[int]AttackCard{
			0: {ID: 4, Type: "ladder", North: "", East: "A", South: "", West: "A", Skull: "false"},
			1: {ID: 1, Type: "ladder", North: "A", East: "", South: "", West: "A", Skull: "false"},
			2: {ID: 6, Type: "ladder", North: "A", East: "A", South: "", West: "", Skull: "false"},
			3: {ID: 6, Type: "ladder", North: "A", East: "A", South: "", West: "", Skull: "false"},
			4: {ID: 6, Type: "ladder", North: "A", East: "A", South: "", West: "", Skull: "false"},
		},
	)
	assert.Equal(t, len(output.DomainCards), 102)

	goldenDomainCards := []DomainCard{
		{ID: 12, KnightsTraining: "bow-and-arrow", StValentinesFestival: 2, Michaelmas: 5, Tournament: "blank-shield-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 16, KnightsTraining: "bow-and-arrow", StValentinesFestival: 4, Michaelmas: 1, Tournament: "shield-blank-sword", Brewhouse: "barm", Lammas: "green"},
		{ID: 5, KnightsTraining: "bow-and-arrow", StValentinesFestival: 6, Michaelmas: 5, Tournament: "blank-shield-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 11, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
		{ID: 20, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 17, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
		{ID: 10, KnightsTraining: "horseshoe", StValentinesFestival: 3, Michaelmas: 4, Tournament: "sword-shield-blank", Brewhouse: "hops", Lammas: "purple"},
		{ID: 14, KnightsTraining: "scroll", StValentinesFestival: 6, Michaelmas: 7, Tournament: "shield-blank-sword", Brewhouse: "water", Lammas: "purple"},
		{ID: 7, KnightsTraining: "bow-and-arrow", StValentinesFestival: 6, Michaelmas: 1, Tournament: "blank-sword-shield", Brewhouse: "barm", Lammas: "green"},
		{ID: 24, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 4, Tournament: "sword-shield-blank", Brewhouse: "hops", Lammas: "purple"},
		{ID: 4, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 6, Tournament: "shield-sword-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 9, KnightsTraining: "bow-and-arrow", StValentinesFestival: 2, Michaelmas: 1, Tournament: "blank-sword-shield", Brewhouse: "barm", Lammas: "green"},
		{ID: 13, KnightsTraining: "scroll", StValentinesFestival: 4, Michaelmas: 3, Tournament: "blank-shield-sword", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 22, KnightsTraining: "horseshoe", StValentinesFestival: 1, Michaelmas: 8, Tournament: "sword-blank-shield", Brewhouse: "barm", Lammas: "red"},
		{ID: 6, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 2, Tournament: "shield-sword-blank", Brewhouse: "water", Lammas: "red"},
		{ID: 23, KnightsTraining: "horseshoe", StValentinesFestival: 3, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 19, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 2, KnightsTraining: "scroll", StValentinesFestival: 2, Michaelmas: 7, Tournament: "blank-sword-shield", Brewhouse: "water", Lammas: "purple"},
		{ID: 3, KnightsTraining: "scroll", StValentinesFestival: 4, Michaelmas: 7, Tournament: "shield-blank-sword", Brewhouse: "water", Lammas: "purple"},
		{ID: 1, KnightsTraining: "bow-and-arrow", StValentinesFestival: 4, Michaelmas: 5, Tournament: "shield-blank-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 18, KnightsTraining: "horseshoe", StValentinesFestival: 1, Michaelmas: 4, Tournament: "sword-blank-shield", Brewhouse: "hops", Lammas: "purple"},
		{ID: 21, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 15, KnightsTraining: "scroll", StValentinesFestival: 2, Michaelmas: 3, Tournament: "blank-sword-shield", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 8, KnightsTraining: "scroll", StValentinesFestival: 6, Michaelmas: 3, Tournament: "blank-shield-sword", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 21, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 7, KnightsTraining: "bow-and-arrow", StValentinesFestival: 6, Michaelmas: 1, Tournament: "blank-sword-shield", Brewhouse: "barm", Lammas: "green"},
		{ID: 20, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 4, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 6, Tournament: "shield-sword-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 18, KnightsTraining: "horseshoe", StValentinesFestival: 1, Michaelmas: 4, Tournament: "sword-blank-shield", Brewhouse: "hops", Lammas: "purple"},
		{ID: 19, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 9, KnightsTraining: "bow-and-arrow", StValentinesFestival: 2, Michaelmas: 1, Tournament: "blank-sword-shield", Brewhouse: "barm", Lammas: "green"},
		{ID: 12, KnightsTraining: "bow-and-arrow", StValentinesFestival: 2, Michaelmas: 5, Tournament: "blank-shield-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 10, KnightsTraining: "horseshoe", StValentinesFestival: 3, Michaelmas: 4, Tournament: "sword-shield-blank", Brewhouse: "hops", Lammas: "purple"},
		{ID: 15, KnightsTraining: "scroll", StValentinesFestival: 2, Michaelmas: 3, Tournament: "blank-sword-shield", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 2, KnightsTraining: "scroll", StValentinesFestival: 2, Michaelmas: 7, Tournament: "blank-sword-shield", Brewhouse: "water", Lammas: "purple"},
		{ID: 22, KnightsTraining: "horseshoe", StValentinesFestival: 1, Michaelmas: 8, Tournament: "sword-blank-shield", Brewhouse: "barm", Lammas: "red"},
		{ID: 23, KnightsTraining: "horseshoe", StValentinesFestival: 3, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 17, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
		{ID: 24, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 4, Tournament: "sword-shield-blank", Brewhouse: "hops", Lammas: "purple"},
		{ID: 14, KnightsTraining: "scroll", StValentinesFestival: 6, Michaelmas: 7, Tournament: "shield-blank-sword", Brewhouse: "water", Lammas: "purple"},
		{ID: 11, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
		{ID: 1, KnightsTraining: "bow-and-arrow", StValentinesFestival: 4, Michaelmas: 5, Tournament: "shield-blank-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 5, KnightsTraining: "bow-and-arrow", StValentinesFestival: 6, Michaelmas: 5, Tournament: "blank-shield-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 6, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 2, Tournament: "shield-sword-blank", Brewhouse: "water", Lammas: "red"},
		{ID: 16, KnightsTraining: "bow-and-arrow", StValentinesFestival: 4, Michaelmas: 1, Tournament: "shield-blank-sword", Brewhouse: "barm", Lammas: "green"},
		{ID: 8, KnightsTraining: "scroll", StValentinesFestival: 6, Michaelmas: 3, Tournament: "blank-shield-sword", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 3, KnightsTraining: "scroll", StValentinesFestival: 4, Michaelmas: 7, Tournament: "shield-blank-sword", Brewhouse: "water", Lammas: "purple"},
		{ID: 13, KnightsTraining: "scroll", StValentinesFestival: 4, Michaelmas: 3, Tournament: "blank-shield-sword", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 23, KnightsTraining: "horseshoe", StValentinesFestival: 3, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 5, KnightsTraining: "bow-and-arrow", StValentinesFestival: 6, Michaelmas: 5, Tournament: "blank-shield-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 6, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 2, Tournament: "shield-sword-blank", Brewhouse: "water", Lammas: "red"},
		{ID: 16, KnightsTraining: "bow-and-arrow", StValentinesFestival: 4, Michaelmas: 1, Tournament: "shield-blank-sword", Brewhouse: "barm", Lammas: "green"},
		{ID: 2, KnightsTraining: "scroll", StValentinesFestival: 2, Michaelmas: 7, Tournament: "blank-sword-shield", Brewhouse: "water", Lammas: "purple"},
		{ID: 8, KnightsTraining: "scroll", StValentinesFestival: 6, Michaelmas: 3, Tournament: "blank-shield-sword", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 4, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 6, Tournament: "shield-sword-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 12, KnightsTraining: "bow-and-arrow", StValentinesFestival: 2, Michaelmas: 5, Tournament: "blank-shield-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 17, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
		{ID: 15, KnightsTraining: "scroll", StValentinesFestival: 2, Michaelmas: 3, Tournament: "blank-sword-shield", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 19, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 3, KnightsTraining: "scroll", StValentinesFestival: 4, Michaelmas: 7, Tournament: "shield-blank-sword", Brewhouse: "water", Lammas: "purple"},
		{ID: 1, KnightsTraining: "bow-and-arrow", StValentinesFestival: 4, Michaelmas: 5, Tournament: "shield-blank-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 22, KnightsTraining: "horseshoe", StValentinesFestival: 1, Michaelmas: 8, Tournament: "sword-blank-shield", Brewhouse: "barm", Lammas: "red"},
		{ID: 10, KnightsTraining: "horseshoe", StValentinesFestival: 3, Michaelmas: 4, Tournament: "sword-shield-blank", Brewhouse: "hops", Lammas: "purple"},
		{ID: 14, KnightsTraining: "scroll", StValentinesFestival: 6, Michaelmas: 7, Tournament: "shield-blank-sword", Brewhouse: "water", Lammas: "purple"},
		{ID: 24, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 4, Tournament: "sword-shield-blank", Brewhouse: "hops", Lammas: "purple"},
		{ID: 9, KnightsTraining: "bow-and-arrow", StValentinesFestival: 2, Michaelmas: 1, Tournament: "blank-sword-shield", Brewhouse: "barm", Lammas: "green"},
		{ID: 18, KnightsTraining: "horseshoe", StValentinesFestival: 1, Michaelmas: 4, Tournament: "sword-blank-shield", Brewhouse: "hops", Lammas: "purple"},
		{ID: 11, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
		{ID: 21, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 7, KnightsTraining: "bow-and-arrow", StValentinesFestival: 6, Michaelmas: 1, Tournament: "blank-sword-shield", Brewhouse: "barm", Lammas: "green"},
		{ID: 20, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 13, KnightsTraining: "scroll", StValentinesFestival: 4, Michaelmas: 3, Tournament: "blank-shield-sword", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 11, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
		{ID: 24, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 4, Tournament: "sword-shield-blank", Brewhouse: "hops", Lammas: "purple"},
		{ID: 10, KnightsTraining: "horseshoe", StValentinesFestival: 3, Michaelmas: 4, Tournament: "sword-shield-blank", Brewhouse: "hops", Lammas: "purple"},
		{ID: 6, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 2, Tournament: "shield-sword-blank", Brewhouse: "water", Lammas: "red"},
		{ID: 9, KnightsTraining: "bow-and-arrow", StValentinesFestival: 2, Michaelmas: 1, Tournament: "blank-sword-shield", Brewhouse: "barm", Lammas: "green"},
		{ID: 16, KnightsTraining: "bow-and-arrow", StValentinesFestival: 4, Michaelmas: 1, Tournament: "shield-blank-sword", Brewhouse: "barm", Lammas: "green"},
		{ID: 18, KnightsTraining: "horseshoe", StValentinesFestival: 1, Michaelmas: 4, Tournament: "sword-blank-shield", Brewhouse: "hops", Lammas: "purple"},
		{ID: 7, KnightsTraining: "bow-and-arrow", StValentinesFestival: 6, Michaelmas: 1, Tournament: "blank-sword-shield", Brewhouse: "barm", Lammas: "green"},
		{ID: 1, KnightsTraining: "bow-and-arrow", StValentinesFestival: 4, Michaelmas: 5, Tournament: "shield-blank-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 15, KnightsTraining: "scroll", StValentinesFestival: 2, Michaelmas: 3, Tournament: "blank-sword-shield", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 17, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
		{ID: 3, KnightsTraining: "scroll", StValentinesFestival: 4, Michaelmas: 7, Tournament: "shield-blank-sword", Brewhouse: "water", Lammas: "purple"},
		{ID: 12, KnightsTraining: "bow-and-arrow", StValentinesFestival: 2, Michaelmas: 5, Tournament: "blank-shield-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 20, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 2, KnightsTraining: "scroll", StValentinesFestival: 2, Michaelmas: 7, Tournament: "blank-sword-shield", Brewhouse: "water", Lammas: "purple"},
		{ID: 19, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 4, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 6, Tournament: "shield-sword-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 22, KnightsTraining: "horseshoe", StValentinesFestival: 1, Michaelmas: 8, Tournament: "sword-blank-shield", Brewhouse: "barm", Lammas: "red"},
		{ID: 8, KnightsTraining: "scroll", StValentinesFestival: 6, Michaelmas: 3, Tournament: "blank-shield-sword", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 23, KnightsTraining: "horseshoe", StValentinesFestival: 3, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 21, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 13, KnightsTraining: "scroll", StValentinesFestival: 4, Michaelmas: 3, Tournament: "blank-shield-sword", Brewhouse: "barley", Lammas: "yellow"},
		{ID: 5, KnightsTraining: "bow-and-arrow", StValentinesFestival: 6, Michaelmas: 5, Tournament: "blank-shield-sword", Brewhouse: "hops", Lammas: "yellow"},
		{ID: 14, KnightsTraining: "scroll", StValentinesFestival: 6, Michaelmas: 7, Tournament: "shield-blank-sword", Brewhouse: "water", Lammas: "purple"},
		{ID: 6, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 2, Tournament: "shield-sword-blank", Brewhouse: "water", Lammas: "red"},
		{ID: 19, KnightsTraining: "sword-and-shield", StValentinesFestival: 1, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 21, KnightsTraining: "horseshoe", StValentinesFestival: 5, Michaelmas: 8, Tournament: "shield-sword-blank", Brewhouse: "barm", Lammas: "red"},
		{ID: 20, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 6, Tournament: "sword-shield-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 4, KnightsTraining: "sword-and-shield", StValentinesFestival: 3, Michaelmas: 6, Tournament: "shield-sword-blank", Brewhouse: "barley", Lammas: "green"},
		{ID: 11, KnightsTraining: "sword-and-shield", StValentinesFestival: 5, Michaelmas: 2, Tournament: "sword-blank-shield", Brewhouse: "water", Lammas: "red"},
	}
	assert.Equal(t,
		output.DomainCards,
		goldenDomainCards,
	)

	assert.True(t, func() bool {
		for i := range 5 {
			// 0  : 23
			// 24 : 47
			// 48 : 71
			// 72 : 95
			// 96 : 101
			test := output.DomainCards[i*24 : (i+1)*24]
			unique := make(map[int]struct{})
			for _, item := range test {
				unique[item.ID] = struct{}{}
			}

			t.Logf("len(unique)=%v | len(test)=%v", len(unique), len(test))
			if len(unique) != len(test) {
				t.Log("domain card sets have duplicates")
				return false
			}
		}
		return true
	}())
}
