package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	numClients := 300
	serverAddr := "localhost:8080"
	var wg sync.WaitGroup
	wg.Add(numClients)

	successCount := 0
	var mu sync.Mutex

	var connLatencies []time.Duration
	var msgLatencies []time.Duration
	var latMu sync.Mutex

	for i := 0; i < numClients; i++ {
		go func(i int) {
			defer wg.Done()
			connStart := time.Now()
			conn, err := net.Dial("tcp", serverAddr)
			if err != nil {
				fmt.Println("Client", i, "failed to connect:", err)
				return
			}
			connLatency := time.Since(connStart)
			defer conn.Close()
			latMu.Lock()
			connLatencies = append(connLatencies, connLatency)
			latMu.Unlock()

			mu.Lock()
			successCount++
			mu.Unlock()

			message := fmt.Sprintf("Hello from client %d\n", i)
			msgStart := time.Now()
			fmt.Fprintf(conn, message)
			msgLatency := time.Since(msgStart)

			latMu.Lock()
			msgLatencies = append(msgLatencies, msgLatency)
			latMu.Unlock()

			time.Sleep(2 * time.Second) //keep connection alive for 2 s
		}(i)
	}

	wg.Wait()
	if len(connLatencies) > 0 {
		var total time.Duration
		min := connLatencies[0]
		max := connLatencies[0]
		for _, l := range connLatencies {
			total += l
			if l < min {
				min = l
			}
			if l > max {
				max = l
			}
		}
		avg := total / time.Duration(len(connLatencies))
		fmt.Printf("Connection latency (min/avg/max): %v / %v / %v\n", min, avg, max)
	}

	if len(msgLatencies) > 0 {
		var total time.Duration
		min := msgLatencies[0]
		max := msgLatencies[0]
		for _, l := range msgLatencies {
			total += l
			if l < min {
				min = l
			}
			if l > max {
				max = l
			}
		}
		avg := total / time.Duration(len(msgLatencies))
		fmt.Printf("Message latency (min/avg/max): %v / %v / %v\n", min, avg, max)
	}

	fmt.Printf("%d/%d clients connected successfully\n", successCount, numClients)
}
