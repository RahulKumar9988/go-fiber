package enum

type BookStatus string

const (
	Read    BookStatus = "read"
	Reading BookStatus = "reading"
	ToRead  BookStatus = "to_read"
)
