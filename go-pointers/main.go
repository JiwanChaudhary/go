// with big power comes great responsibility
// pointer has big power
// race condition
// nil pointer
// when to use pointer
//! when we need to directly update from memory like in the below case we are updating players health directly from memory
//! use when processing with big data so that everytime instead of copying big data, it just points to that memory location

package main

import (
	"fmt"
)

type BigData struct {
	// 500 MB data -> this is big data
}

// 8 bytes => pointer (64bit machine)
func processBigData(bd *BigData) {
	// so here if this function is called 10000 times(suppose), then we just point them(8 byte) instead of copying that 500 MB big data
}

type Player struct {
	health int
}

// this is method
func (player *Player) takeDamageFromExplosion(dmg int) {
	fmt.Println("Player is taking damage from explosion.")
	player.health -= dmg
}

// the below function is the syntactic sugar of the above function.

// this is function
func takeDamageFromExplosion(player *Player, dmg int) {
	fmt.Println("Player is taking damage from explosion.")
	player.health -= dmg
}

func main() {
	player := &Player{
		health: 100,
	}
	// 8 byte long interger (pointer)
	// directly modify the player state in memory

	fmt.Printf("before explosion %+v\n", player)
	//! takeDamageFromExplosion(player) //copy
	// giving reference to memory of that player
	player.takeDamageFromExplosion(50)
	// takeDamageFromExplosion(player, 25)
	fmt.Printf("after explosion %+v\n", player)

}
