package lru

type Entry struct {
	value interface{}
	prev  *Entry
	next  *Entry
}

func NewEntry(value interface{}) *Entry {
	return &Entry{
		value: value,
	}
}

func (e *Entry) Value() interface{} {
	return e.value
}
