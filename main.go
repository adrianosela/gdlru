package main

import (
	"github.com/adrianosela/gdlru/cache"
	"github.com/adrianosela/gdlru/guardduty"
	"log"
)

func evictFunc(k, v interface{}) error {
	threatPurpose := k.(string)
	finding := v.(guardduty.Finding)

	// do a thing (e.g. commit to db)
	return nil
}

func main() {
	cache, err := cache.NewCache(100, evictFunc)
	if err != nil {
		log.Fatalf("failed to initialize cache: %s", err)
	}

	for {
		// threatPurpose := someStreamOfData.Next() // assume this exists

		// if val, ok := cache.Get(threatPurpose); ok {
		// 	finding := val.(guardduty.Finding)
		// 	finding.Inc()
		// 	cache.Put(threatPurpose, finding.Data)
		// } else {
		// 	cache.Put(threatPurpose, guardduty.NewFinding(threatPurpose))
		// }

		//  catch SIGTERM, SIGINT and call cache.Commit(), then exit
	}
	return
}
