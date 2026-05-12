package utils

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func PlayMusic() (<-chan bool, error) {
	// region Play Music
	music, err := os.Open("world.execute(me).mp3")
	if err != nil {
		return nil, err
	}

	streamer, format, err := mp3.Decode(music)
	if err != nil {
		return nil, err
	}

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return nil, err
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		streamer.Close()
		music.Close()
		done <- true
	})))

	return done, nil
}
