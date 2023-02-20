package services

import (
	"sync"

	"audiofile/services/metadata"
	"audiofile/services/transcriptreview"
)

func Run() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	metadata.Run(wg)
	transcriptreview.Run(wg)
	wg.Wait()
}
