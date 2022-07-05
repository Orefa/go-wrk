package main

import (
	"sync"
)

func SingleNode(toCall string) []byte {
	responseChannel := make(chan *Response, *totalCalls)

	benchTime := NewTimer()
	benchTime.Reset()
	//TODO check ulimit
	wg := &sync.WaitGroup{}
	numPer := *totalCalls / *numConnections
	left := *totalCalls % *numConnections
	// numConnections > totalCalls
	nuc := *numConnections
	if numPer == 0 {
		nuc = *totalCalls
	}
	for i := 0; i < nuc; i++ {
		addNum := 0
		if i < left {
			addNum = 1
		}
		go StartClient(
			toCall,
			*headers,
			*requestBody,
			*method,
			*disableKeepAlives,
			responseChannel,
			wg,
			numPer+addNum,
		)

		wg.Add(1)
	}

	wg.Wait()

	result := CalcStats(
		responseChannel,
		benchTime.Duration(),
	)
	return result
}
