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
	currentValue := initialValue
	for _, value := range l.Values {
		currentValue = f(currentValue, value)
	}
	result[l.ID] = currentValue
	return result
}

func FoldChan(initialValue int, f func(int, int) int, ch chan List) map[ID]int {
	result := make(map[ID]int)
	var currentValue int
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
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, ch := range chs {
		go func(ch chan List) {
			defer wg.Done()
			channelResult := FoldChan(initialValue, f, ch)
			for id, value := range channelResult {
				result[id] += value
			}
		}(ch)
	}
	wg.Wait()
	return result
}
