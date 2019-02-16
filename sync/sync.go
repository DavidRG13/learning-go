package sync

type Counter struct {
	value int64
}

func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int64 {
	return c.value
}
