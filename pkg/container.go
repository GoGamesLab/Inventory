package container

func (c *Container[T]) AddItem(id ItemID, amount T) (T, bool) {
	var zero T

	if c.Items == nil {
		return zero, false
	}

	if _, ok := c.GetItem(id); !ok {
		return zero, false
	}

	switch any(amount).(type) {
	case uint8:
		a := any(amount).(uint8)

		if a == 0 {
			return c.Items[id], true
		}

		var current uint8
		if val, ok := any(c.Items[id]).(uint8); ok {
			current = val
		}

		if uint16(current)+uint16(a) > 255 {
			return c.Items[id], false
		}

		c.Items[id] = any(current + a).(T)

		return c.Items[id], true
	case float32:
		a := any(amount).(float32)

		if a == 0 {
			return c.Items[id], true
		}

		if a < 0 {
			return c.Items[id], false
		}

		var current float32
		if val, ok := any(c.Items[id]).(float32); ok {
			current = val
		}

		c.Items[id] = any(current + a).(T)
		return c.Items[id], true
	}

	return c.Items[id], false
}

func (c *Container[T]) RemoveItem(id ItemID, amount T) (T, bool) {
	var zero T

	if _, ok := c.GetItem(id); !ok {
		return zero, false
	}

	if c.Items == nil {
		return zero, false
	}

	switch any(amount).(type) {
	case uint8:
		a := any(amount).(uint8)

		if a == 0 {
			return c.Items[id], true
		}

		var currentAmount uint8
		val, ok := any(c.Items[id]).(uint8)
		if !ok {
			return c.Items[id], false
		}
		currentAmount = val

		if currentAmount < a {
			return T(currentAmount), false
		}

		newAmount := currentAmount - a
		if newAmount == 0 {
			delete(c.Items, id)

			return zero, true
		} else {
			c.Items[id] = any(newAmount).(T)
		}

		return c.Items[id], true
	case float32:
		a := any(amount).(float32)

		if a == 0 {
			return T(a), true
		}

		if a < 0 {
			return T(a), false
		}

		var currentAmount float32
		val, ok := any(c.Items[id]).(float32)
		if !ok {
			return zero, false
		}
		currentAmount = val

		var newAmount float32
		if currentAmount < a {
			newAmount = 0.0
		} else {
			newAmount = currentAmount - a
		}

		if newAmount <= 0.0001 {
			delete(c.Items, id)

			return zero, true
		} else {
			c.Items[id] = any(newAmount).(T)
		}

		return c.Items[id], true
	}

	return zero, false
}

func (c *Container[T]) GetItem(id ItemID) (T, bool) {
	var zero T

	if c.Items == nil {
		return zero, false
	}

	return c.Items[id], true
}
