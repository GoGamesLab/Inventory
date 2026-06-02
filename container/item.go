package container

func RegisterNewItem(i Item) (Item, bool) {
	i.ID = ItemID(len(ItemRegister) + 1)
	ItemRegister[i.ID] = i

	return i, true
}
