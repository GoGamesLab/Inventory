package container

type ItemID uint16

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type Item struct {
	ID   ItemID
	Name string
}

type Container[V Numeric] struct {
	Items map[ItemID]V
}

type Inventory = Container[uint8]

func NewInventory() *Inventory {
	return &Inventory{
		Items: make(map[ItemID]uint8),
	}
}

type Storage = Container[float32]

func NewStorage() *Storage {
	return &Storage{
		Items: make(map[ItemID]float32),
	}
}
