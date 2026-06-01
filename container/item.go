package container

func RegisterItem(i Item) bool {
	if _, exists := ItemRegister[i.ID]; exists {
		return false
	}
	ItemRegister[i.ID] = i

	return true
}
