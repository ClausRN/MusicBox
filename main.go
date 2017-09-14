package main

import (
	"fmt"
	"time"
)

func main() {

	musiclib := NewSongList()
	fmt.Println("Song: ", musiclib.SongNo)

	for i := int32(0); i < musiclib.SongNo; i++ {
		fmt.Println("Track: ", musiclib.Songs[i])
	}

        for i := int32(0); i < 10; i++ {
                fmt.Println("Random: ", musiclib.NextTrack())
        }


        // Use BCM pin numbering
        // Echo pin
        // Trigger pin
	h := NewUSProbe(3, 2)
//CYN


	for true {
		//distance := h.MeasureDistance()
                //fmt.Println(distance)
		h.Probe()
		fmt.Printf("Time %v: Distance: %2.3f, Min: %2.3f, Max: %2.3ff, Raw: %2.3f  \n", h.TriggeredUntil, h.Distance, h.DistanceLow, h.DistanceHigh, h.DistanceRaw)
		if (h.TriggeredUntil.After(time.Now())){
			fmt.Println("Playing")
		}
		time.Sleep(time.Duration(1) * time.Second)
	}

}
