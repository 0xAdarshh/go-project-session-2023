package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const PrintTableHeader = "index \tinput \tresult \n"
const PrintFormattedInputAndResult = "%d \t%d \t%d \n"
const PrintTotalTimeTook = "\nTotal time took = %f"

func main() {
	tStart := time.Now()

	wg := sync.WaitGroup{}
	inputs := getSampleInputList(15000)
	var results []int
	for _, input := range inputs {

		wg.Add(1)
		go func(opInput int) {
			res := verySlowAssHeavyOperation(opInput)
			results = append(results, res)
			wg.Done()
		}(input)
	}

	wg.Wait()
	fmt.Printf(PrintTableHeader)
	for i, result := range results {
		fmt.Printf(PrintFormattedInputAndResult, i, result, inputs[i])
	}

	tEnd := time.Now()
	fmt.Printf(PrintTotalTimeTook, tEnd.Sub(tStart).Seconds())
}

func verySlowAssHeavyOperation(input int) int {
	time.Sleep(1 * time.Second)
	return (input * 12) + 1
}

func getSampleInputList(numberOfInputs int) []int {
	inputs := make([]int, numberOfInputs)
	for i, _ := range inputs {
		inputs[i] = rand.Intn(999)
	}
	return inputs
}
