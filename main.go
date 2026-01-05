package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	KNIGHTS_TRAINING_BOW          = "bow-and-arrow"
	KNIGHTS_TRAINING_SCROLL       = "scroll"
	KNIGHTS_TRAINING_SHIELD       = "sword-and-shield"
	KNIGHTS_TRAINING_HORSESHOE    = "horseshoe"
	TOURNAMENT_SWORD_SHIELD_BLANK = "sword-shield-blank"
	TOURNAMENT_SWORD_BLANK_SHIELD = "sword-blank-shield"
	TOURNAMENT_SHIELD_BLANK_SWORD = "shield-blank-sword"
	TOURNAMENT_SHIELD_SWORD_BLANK = "shield-sword-blank"
	TOURNAMENT_BLANK_SHIELD_SWORD = "blank-shield-sword"
	TOURNAMENT_BLANK_SWORD_SHIELD = "blank-sword-shield"
	BREWHOUSE_HOPS                = "hops"
	BREWHOUSE_WATER               = "water"
	BREWHOUSE_BARLEY              = "barley"
	BREWHOUSE_BARM                = "barm"
	LAMMAS_YELLOW                 = "yellow"
	LAMMAS_PURPLE                 = "purple"
	LAMMAS_RED                    = "red"
	LAMMAS_GREEN                  = "green"
	ATTACK_CARD_LADDER            = "ladder"
	ATTACK_CARD_ARROWS            = "arrows"
	ATTACK_CARD_BALLISTA          = "ballista"
	ATTACK_CARD_BATTERING_RAM     = "battering_ram"
	ATTACK_CARD_SIEGE_TOWER       = "siege_tower"
	ATTACK_CARD_CATAPAULT         = "catapault"
)

type DomainCard struct {
	ID                   int    `toml:"id"`
	KnightsTraining      string `toml:"knights_training"`
	StValentinesFestival int    `toml:"st_valentines_festival"`
	Michaelmas           int    `toml:"michaelmas"`
	Tournament           string `toml:"tournament"`
	Brewhouse            string `toml:"brewhouse"`
	Lammas               string `toml:"lammas"`
}

type AttackCard struct {
	ID    int    `toml:"id"`
	Type  string `toml:"type"`
	North string `toml:"north"`
	East  string `toml:"east"`
	South string `toml:"south"`
	West  string `toml:"west"`
	Skull string `toml:"skull"`
}

type Components struct {
	PathCards          []string     `toml:"path_cards"`
	DomainCards        []DomainCard `toml:"domain_cards"`
	AttackCards        []AttackCard `toml:"attack_cards"`
	FinalEscaladeCards []AttackCard `toml:"final_escalade_cards"`
	LargePlunderTiles  []string     `toml:"large_plunder_tiles"`
	SmallPlunderTiles  []string     `toml:"small_plunder_tiles"`
}

type Output struct {
	PathCardsByRound         map[int][]string
	AttackCardsByRound       map[int][]AttackCard
	FinalEscaladeCardByRound map[int]AttackCard
	DomainCards              []DomainCard
	PlunderTilesByRound      map[int]string
	FirstSeed                uint64
	SecondSeed               uint64
}

func (o *Output) AssignPathCardsByRound(rng *rand.Rand, pathCards []string) {
	shuffle(rng, pathCards)
	o.PathCardsByRound = map[int][]string{
		1: {pathCards[0], pathCards[1]},
		2: {pathCards[2], pathCards[3]},
		3: {pathCards[4], pathCards[5]},
		4: {pathCards[6], pathCards[7]},
		5: {pathCards[8], pathCards[9]},
	}
}

func (o *Output) AssignAttackCardsByRound(rng *rand.Rand, finalEscaladeCards []AttackCard, attackCards []AttackCard) {
	o.AttackCardsByRound = make(map[int][]AttackCard)
	o.FinalEscaladeCardByRound = make(map[int]AttackCard)
	for i := range 5 {
		shuffle(rng, finalEscaladeCards)
		o.FinalEscaladeCardByRound[i] = finalEscaladeCards[0]
		shuffle(rng, attackCards)

		o.AttackCardsByRound[i] = append(o.AttackCardsByRound[i], attackCards[0:i+1]...)
	}
}

func (o *Output) AssignDomainCards(rng *rand.Rand, domainCards []DomainCard) {
	for range 5 {
		shuffle(rng, domainCards)
		o.DomainCards = append(o.DomainCards, domainCards...)
	}
	o.DomainCards = o.DomainCards[:102]
}

func (o *Output) AssignPlunderTilesByRound(rng *rand.Rand, largeTiles []string, smallTiles []string) {
	shuffle(rng, largeTiles)
	shuffle(rng, smallTiles)
	o.PlunderTilesByRound = map[int]string{
		1: largeTiles[0],
		2: largeTiles[1],
		3: smallTiles[0],
		4: smallTiles[1],
		5: smallTiles[2],
	}
}

func ComputeOutput(seed1 uint64, seed2 uint64, components Components) Output {

	rng := rand.New(rand.NewPCG(seed1, seed2))
	output := Output{
		FirstSeed:  seed1,
		SecondSeed: seed2,
	}

	output.AssignPathCardsByRound(rng, components.PathCards)
	output.AssignAttackCardsByRound(rng, components.FinalEscaladeCards, components.AttackCards)
	output.AssignDomainCards(rng, components.DomainCards)
	output.AssignPlunderTilesByRound(rng, components.LargePlunderTiles, components.SmallPlunderTiles)

	return output
}

func main() {
	var seed1, seed2 uint64
	flag.Uint64Var(&seed1, "first_seed", rand.Uint64(), "the first seed number used for shuffling")
	flag.Uint64Var(&seed2, "second_seed", rand.Uint64(), "the second seed number used for shuffling")
	flag.Parse()

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

	output := ComputeOutput(seed1, seed2, components)

	fmt.Printf("%#v\n", output)
	// f, err = os.Create("output.html")

	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// funcMap := template.FuncMap{
	// 	"add": func(a, b int) int { return a + b },
	// 	"fixedSizeString": func(a string, finalLength int) string {
	// 		length := len([]rune(a))
	// 		return a + strings.Repeat(".", finalLength-length)
	// 	},
	// }

	// tmpl, err := template.New("html_template.html.tmpl").Funcs(funcMap).ParseFiles("html_template.html.tmpl")
	// if err != nil {
	// 	panic(err)
	// }

	// err = tmpl.Execute(f, output)
	// if err != nil {
	// 	panic(err)
	// }

	// f.Close()

	// f, err = os.Create("output.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// tmpl, err = template.New("forum_template.tmpl").Funcs(funcMap).ParseFiles("forum_template.tmpl")
	// if err != nil {
	// 	panic(err)
	// }

	// err = tmpl.Execute(f, output)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Fprintf(f, "%#v\n%#v,%#v\n", output.AttackCardsByRound, output.DomainCards, output.PathCardsByRound)

}

func shuffle[T comparable](rng *rand.Rand, item []T) {
	rng.Shuffle(len(item), func(i, j int) {
		item[i], item[j] = item[j], item[i]
	})
}
