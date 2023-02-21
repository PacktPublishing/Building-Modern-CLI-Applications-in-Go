package main

import (
	"github.com/marianina8/Chapter-6/examples"
)

func main() {
	examples.CreateCommandUsingStruct()
	examples.CreateCommandUsingCommandFunction()
	examples.RunMethod()
	examples.StartMethod()
	examples.OutputMethod()
	examples.CombinedOutputMethod()
	examples.Pagination()
	examples.Limit()
	examples.Timeout()
	examples.HandlingDoesNotExistErrors()
	examples.HandlingOtherErrors()
	examples.Panic()
	examples.HTTPTimeout()
	examples.HTTPError()
}
