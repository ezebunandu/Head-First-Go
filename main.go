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

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder) // type assertion to be able to use the Record method from the TapeRecorder
	recorder.Record()
}

func main() {
	TryOut(gadget.TapeRecorder{})
}
