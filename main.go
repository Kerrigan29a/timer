package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/gobuffalo/packr"
	flag "github.com/spf13/pflag"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func parseArgs(mute, repeat *bool) []time.Duration {
	flag.Usage = func() {
		fmt.Println("")
		fmt.Printf("Usage:  %s [OPTIONS] duration...\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Printf("Version: %s\n", version)
		fmt.Printf("Commit: %s\n", commit)
		fmt.Printf("Date: %s\n", date)
		fmt.Println("")
	}
	flag.BoolVarP(mute, "mute", "m", false, "Mute sound")
	flag.BoolVarP(repeat, "repeat", "r", false, "Repeat")
	flag.Parse()
	durations := make([]time.Duration, flag.NArg(), flag.NArg())
	for i, arg := range flag.Args() {
		duration, err := time.ParseDuration(arg)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		durations[i] = duration
	}
	return durations
}

func main() {
	var mute bool
	var repeat bool
	durations := parseArgs(&mute, &repeat)

	i := 1
	var done chan struct{}
	var err error
	for {
		if repeat {
			fmt.Printf("Iteration %d\n", i)
		}
		for _, duration := range durations {
			done, err = handleDuration(duration, mute)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if !repeat {
			break
		}
		i++
	}
	if done != nil {
		<-done
	}
}

func handleDuration(duration time.Duration, mute bool) (chan struct{}, error) {
	ticker := time.NewTicker(1 * time.Second)
	count := 0 * time.Second
	for range ticker.C {
		remaining := (duration - count).String()
		remainingLen := len(remaining)
		if remainingLen < 9 {
			remainingLen = 9
		}

		width, _, err := terminal.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			width = 80
		}

		fmt.Print("\r")
		if count >= duration {
			fmt.Printf("%*s", remainingLen-1, "DONE")
		} else {
			fmt.Printf("%*s", remainingLen-1, remaining)
		}
		fmt.Print(" ")
		fmt.Print(composeBar(float64(count), float64(duration), width-2-remainingLen))
		if count >= duration {
			break
		}
		count += 1 * time.Second
	}
	fmt.Println("")
	ticker.Stop()
	if !mute {
		return playSong()
	}
	return nil, nil
}

func composeBar(count, duration float64, width int) string {
	b := bytes.Buffer{}
	b.WriteRune('[')
	base := width - 2
	var limit int
	if duration > 0 {
		limit = int(math.Round(count / duration * float64(base)))
	} else {
		limit = base
	}
	for i := 0; i < base; i++ {
		if i < limit {
			b.WriteRune('#')
		} else {
			b.WriteRune(' ')
		}
	}
	b.WriteRune(']')
	return b.String()
}

func playSong() (chan struct{}, error) {
	/* Load sound */
	box := packr.NewBox("./resources")
	f, err := box.Open("gong.wav")
	if err != nil {
		return nil, fmt.Errorf("unable to open the audio file: %s", err.Error())
	}
	defer f.Close()
	s, format, err := wav.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("unable to decode the audio file: %s", err.Error())
	}

	/* Init speaker */
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return nil, fmt.Errorf("unable to init the speaker: %s", err.Error())
	}

	/* Play sound */
	done := make(chan struct{})
	//s = beep.Take(format.SampleRate.N(time.Second*10), s)
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))

	return done, nil
}
