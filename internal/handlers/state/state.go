package state

import "sync"

var (
	muLast      sync.RWMutex
	lastMenuMsg = map[int64]int{}
)

func SetLastMenuMessage(chatID int64, messageID int) {
	muLast.Lock()
	defer muLast.Unlock()
	lastMenuMsg[chatID] = messageID
}

func GetLastMenuMessage(chatID int64) (int, bool) {
	muLast.RLock()
	defer muLast.RUnlock()
	id, ok := lastMenuMsg[chatID]
	return id, ok
}

func ClearLastMenuMessage(chatID int64) {
	muLast.Lock()
	defer muLast.Unlock()
	delete(lastMenuMsg, chatID)
}
