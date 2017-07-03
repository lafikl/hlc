package hlc_test

import (
	"log"

	"github.com/lafikl/hlc"
)

func Example() {
	clock := hlc.NewWithPT(hlc.PT{Seconds: true})
	ts := clock.Now()

	// updates the clock time
	clock.Update(ts)

	// check ordering
	if ts.Less(clock.Now()) != true {
		log.Println("ts is less than current clock time")
	}
}
