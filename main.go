package main

import "head_first_go/gadget"

//define an interface type
type Player interface {
	Play(string) // require a Play method with a string parameter
	Stop()       // Also require a stop method
}

func playList(device Player, songs []string) {
	// playList accepts any Player, not just a TapePlayer
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{} // player variable holds any player
	playList(player, mixtape)               // Pass a TapePlayer to playList
	player = gadget.TapeRecorder{}
	playList(player, mixtape) // Pass a TapeRecorder to playList
}
