package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main game loop
func gameLoop() string {
	if goblin.Dead {
		fmt.Printf("\nDRAGON 🐉  HP IS %v\nPLAYER🗡 HP IS %v \n", dragon.Hp, player.Hp)

	} else {
		fmt.Printf("\nGOBLIN 👹  HP IS %v\nPLAYER🗡 HP IS %v \n", goblin.Hp, player.Hp)
	}
	turn += 1
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	rng := random.Intn(4)
	var userChoiceInt int
	winner = ""
	// player moves first
	userTurnPrompt()
	userChoice, err := getUserInput()
	for err != nil {
		userInputErrorPrompt()
		userChoice, err = getUserInput()
	}

	userChoiceInt, err = convertInputToInt(userChoice)
	switch userChoiceInt {
	case 1:
		if !goblin.Dead {
			goblin.TakeDamage(player.Attack)
			fmt.Printf("PLAYER🗡 ATTACKED GOBLIN 👹 FOR %v DAMAGE 🗡\nGOBLIN HP IS %v\n", player.Attack, goblin.Hp)
			if goblin.Dead {
				fmt.Printf("THE GOBLIN 👹 HAS DIED THE DRAGON 🐉EMERGES 🐉🔥\n DRAGON 🐉HP: %v \n", dragon.Hp)
			}
		} else {
			dragon.TakeDamage(player.Attack)
			fmt.Printf("PLAYER🗡 ATTACKED DRAGON 🐉 FOR %v DAMAGE 🗡\nDRAGON HP IS %v\n", player.Attack, dragon.Hp)
		}

	case 2:
		player.Heal()
		fmt.Printf("PLAYER🗡 HEALED PLAYER🗡 HP IS %v\n", player.Hp)

	case 3:
		player.Forge()
		fmt.Printf("PLAYER🗡 FORGED🔨 PLAYER🗡 ATTACK IS %v\n", player.Attack)

	case 4:
		player.Rest()
		fmt.Printf("PLAYER🗡 HAS FEASTED🍖 PLAYER🗡 MAX HP IS %v\n ", player.MaxHp)

	}

	if dragon.Dead && goblin.Dead {
		winner = "PLAYER WON 🗡"
		return winner
	}

	// Monster's turn
	if !goblin.Dead {
		switch rng {
		case 1:
			player.TakeDamage(goblin.Attack)
			fmt.Printf("GOBLIN 👹 ATTACKED PLAYER🗡 FOR %v DAMAGE 👹\nPLAYER HP IS %v\n", goblin.Attack, player.Hp)

			if player.Dead {
				winner = "MONSTERS 👹"
			}
		case 2:
			goblin.Heal()
			fmt.Printf("GOBLIN👹 HEALED GOBLIN👹 HP IS %v\n", goblin.Hp)
		case 3:
			goblin.Forge()
			fmt.Printf("GOBLIN 👹  FORGED🔨 GOBLIN👹 ATTACK IS %v\n", goblin.Attack)
		case 4:
			goblin.Rest()
			fmt.Printf("GOBLIN👹 HAS FEASTED GOBLIN👹 MAX HP IS %v\n 🍖", goblin.MaxHp)
		case 5:
			player.TakeDamage(goblin.Attack)
			if player.Dead {
				winner = "MONSTERS 👹"
			}
		}
	} else {
		switch rng {
		case 1:
			player.TakeDamage(goblin.Attack)
			fmt.Printf("DRAGON🐉 ATTACKED PLAYER🗡 FOR %v DAMAGE🐉 \nPLAYER HP IS %v\n", dragon.Attack, player.Hp)
			if player.Dead {
				winner = "MONSTERS 🐉"
			}
		case 2:
			dragon.Heal()
			fmt.Printf("DRAGON🐉 HEALED DRAGON🐉 HP IS %v\n", dragon.Hp)
		case 3:
			dragon.Forge()
			fmt.Printf("DRAGON🐉 FORGED🔨 DRAGON🐉 ATTACK IS %v\n", dragon.Attack)
		case 4:
			dragon.Rest()
			fmt.Printf("DRAGON🐉 HAS FEASTED DRAGON🐉 MAX HP IS %v\n 🍖", dragon.MaxHp)
		case 5:
			fmt.Printf("DRAGON  ATTACKED DRAGON🐉 FOR %v DAMAGE 👹\nPLAYER HP IS %v\n", dragon.Attack, player.Hp)
			dragon.BreathFire(player)
			if player.Dead {
				winner = "MONSTERS 🐉"
			}
		}
	}

	if dragon.Dead && goblin.Dead {
		winner = "PLAYER WON 🗡"
		return winner
	}
	return winner
}

func userTurnPrompt() {
	fmt.Println("YOUR TURN CHOOSE AN ACTION")
	fmt.Println("1) ATTACK 🗡")
	fmt.Println("2) HEAL 💟")
	fmt.Println("3) FORGE A STRONGER WEAPON 🔨")
	fmt.Println("4) EAT TO GAIN MORE HEALTH 🍖")

}

func userInputErrorPrompt() {
	fmt.Println("❌ WRONG INPUT PLEASE INPUT A NUMBER BETWEEN 1 AND 4 ❌")

}
