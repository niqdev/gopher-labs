package myspinner

import (
	"time"

	"github.com/briandowns/spinner"
)

func Run() {
	// 35
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond) // Build our new spinner
	s.Color("green", "bold")
	s.Suffix = "  :appended text"
	s.FinalMSG = "Complete!\nNew line!\nAnother one!\n"
	s.Start()                   // Start the spinner
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
}
