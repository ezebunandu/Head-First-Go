package main

import (
	"head_first_go/gadget"
)

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

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	// only call the Record() method if the type assertion is successful
	if ok {
		recorder.Record()
	}
}

func main() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
