package container

func (c *Container[V]) AddItem(id ItemID, amount V) bool {
	if c.Items == nil {
		c.Items = make(map[ItemID]V)
	}

	switch any(amount).(type) {
	case uint8:
		a := any(amount).(uint8)

		if a == 0 {
			return false
		}

		var current uint8
		if val, ok := any(c.Items[id]).(uint8); ok {
			current = val
		}

		if uint16(current)+uint16(a) > 255 {
			return false
		}

		c.Items[id] = any(current + a).(V)
		return true

	case float32:
		a := any(amount).(float32)

		if a <= 0 {
			return false
		}

		var current float32
		if val, ok := any(c.Items[id]).(float32); ok {
			current = val
		}

		c.Items[id] = any(current + a).(V)
		return true
	}

	return false
}

func (c *Container[V]) RemoveItem(id ItemID, amount V) bool {
	if c.Items == nil {
		return false
	}

	switch any(amount).(type) {
	case uint8:
		a := any(amount).(uint8)

		if a == 0 {
			return false
		}

		var currentAmount uint8
		val, existe := any(c.Items[id]).(uint8)
		if !existe {
			return false
		}
		currentAmount = val

		if currentAmount < a {
			return false
		}

		newAmount := currentAmount - a
		if newAmount == 0 {
			delete(c.Items, id)
		} else {
			c.Items[id] = any(newAmount).(V)
		}
		return true

	case float32:
		a := any(amount).(float32)

		if a <= 0 {
			return false
		}
		var currentAmount float32
		val, existe := any(c.Items[id]).(float32)
		if !existe {
			return false
		}
		currentAmount = val

		if currentAmount < a {
			return false
		}

		newAmount := currentAmount - a
		if newAmount <= 0.0001 {
			delete(c.Items, id)
		} else {
			c.Items[id] = any(newAmount).(V)
		}
		return true
	}

	return false
}

func (c *Container[V]) GetItem(id ItemID) (V, bool) {
	if c.Items == nil {
		var zero V
		return zero, false
	}
	q, ok := c.Items[id]
	return q, ok
}
