package main

import (
	"fmt"
	"time"
	"flag"
	"os"
	"os/signal"
)

func main() {
	flag.Parse()
	//musiclib := NewSongList()
	//fmt.Println("Song: ", musiclib.SongNo)

	//for i := int32(0); i < musiclib.SongNo; i++ {
	//	fmt.Println("Track: ", musiclib.Songs[i])
	//}

        //for i := int32(0); i < 10; i++ {
        //        fmt.Println("Random: ", musiclib.NextTrack())
        //}


        // Use BCM pin numbering
        // Echo pin
        // Trigger pin
	h := NewUSProbe(3, 2, 15)
	Play := NewOMXPlayer()
	//defer Play.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	defer signal.Stop(quit)

	for true {
                //fmt.Println("Probing")
		h.Probe()
		fmt.Printf("Time %v: Distance: %2.3f, Min: %2.3f, Max: %2.3ff, Raw: %2.3f  \n", h.TriggeredUntil, h.Distance, h.DistanceLow, h.DistanceHigh, h.DistanceRaw)
		if (h.TriggeredUntil.After(time.Now())){
			//fmt.Println("Playing")
			Play.Start()
		} else {
			Play.Stop()	
		}
		//fmt.Println("Check signals")
		select {
		case <-time.After(1 * time.Millisecond):
			//fmt.Println("Pause over")
		case <-quit:
			fmt.Println("Catch ctrl-c, exiting!")
			Play.Close()
			return
		}
		time.Sleep(time.Duration(250) * time.Millisecond)
	}

}
