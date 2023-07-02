package rollee

import "sync"

type ID = int

// We suppose L is always valid with len (l.Values) >= 1).
type List struct {
	ID     ID
	Values []int
}

func Fold(initialValue int, f func(int, int) int, l List) map[ID]int {
	result := make(map[ID]int)
	// It's one unique fold, in this case I can initialize directly the current value
	currentValue := initialValue
	// For each value I use the function f in parameters
	for _, value := range l.Values {
		currentValue = f(currentValue, value)
	}
	// I attribute the value for the ID of the list directly in my map and return it
	result[l.ID] = currentValue
	return result
}

func FoldChan(initialValue int, f func(int, int) int, ch chan List) map[ID]int {
	result := make(map[ID]int)
	var currentValue int
	// This function is similar to the Fold function
	// I check at every channel if the ID exists already or not in my map
	// If Yes : I continue from the last value that it had
	// If No : I restart from the initial value
	for listChannel := range ch {
		if _, ok := result[listChannel.ID]; ok {
			currentValue = result[listChannel.ID]
			for _, value := range listChannel.Values {
				currentValue = f(currentValue, value)
			}
		} else {
			currentValue = initialValue
			for _, value := range listChannel.Values {
				currentValue = f(currentValue, value)
			}
		}
		result[listChannel.ID] = currentValue
	}
	return result
}

func FoldChanX(initialValue int, f func(int, int) int, chs ...chan List) map[ID]int {
	result := make(map[ID]int)

	// Use the sync.WaitGroup to wait until every goroutine has finished and avoid deadlock
	var wg sync.WaitGroup
	// The number of channels is added in the sync.waitGroup object
	wg.Add(len(chs))

	// For each channel, we create a new go routine for all of them and use the FoldChan function
	for _, ch := range chs {
		go func(ch chan List) {
			defer wg.Done()
			channelResult := FoldChan(initialValue, f, ch)
			for id, value := range channelResult {
				result[id] += value
			}
		}(ch)
	}
	// Wait for all goroutines and return the result
	wg.Wait()
	return result
}
