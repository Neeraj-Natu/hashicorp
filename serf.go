package hashicorp

import "sync"

//OneAndOnlyNumber that is contained in the database
type OneAndOnlyNumber struct {
	num        int
	generation int
	numMutex   sync.RWMutex
}

//InitNumber the oneAndOnlyNumber
func InitNumber(val int) *OneAndOnlyNumber {
	return &OneAndOnlyNumber{
		num: val,
	}
}

func (n *OneAndOnlyNumber) setValue(newVal int) {
	n.numMutex.Lock()
	defer n.numMutex.Unlock()
	n.num = newVal
	n.generation = n.generation + 1
}

func (n *OneAndOnlyNumber) getValue() (int, int) {
	n.numMutex.RLock()
	defer n.numMutex.RUnlock()
	return n.num, n.generation
}

func (n *OneAndOnlyNumber) notifyValue(curVal int, curGeneration int) bool {
	if curGeneration > n.generation {
		n.numMutex.Lock()
		defer n.numMutex.Unlock()
		n.generation = curGeneration
		n.num = curVal
		return true
	}
	return false
}

//MembersToNotify for the gossip protocol
const MembersToNotify = 2
