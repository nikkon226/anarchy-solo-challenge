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
