package set3

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChal21(t *testing.T) {

	var mt mt19337

	mt.init(0)

	fmt.Println(mt.extractNumber())

}

func TestChal22(t *testing.T) {
	rand.Seed(time.Now().Unix())

	seconds := rand.Intn(1000-40) + 40
	fmt.Printf("Waiting for %d seconds...\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)

	var mt mt19337
	mt.init(uint32(int(time.Now().Unix())))

	seconds = rand.Intn(1000-40) + 40
	fmt.Printf("Seeded and waiting for %d seconds...\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)

	a := mt.extractNumber()

	now := int(time.Now().Unix())

	for i := now - 2000; i < now; i++ {
		mt.init(uint32(i))
		if mt.extractNumber() == a {
			fmt.Printf("Found seed: %d\n", i)
		}
	}
}
