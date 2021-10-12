package iteration

import "math/rand"

func Repeat(character string, times int) string {
	var repeated string
	for i:=0; i < times; i++ {
		repeated += character
	}
	return repeated
}

func startStaticDials() {
	staticPool := []int{1,2,3,4,5}
	for len(staticPool) > 0 {
		idx := rand.Intn(len(staticPool))
		_ = staticPool[idx]
		staticPool = append(staticPool[:idx], staticPool[idx+1:]...)
		//d.startDial(task)
		//d.removeFromStaticPool(idx)
	}
}
