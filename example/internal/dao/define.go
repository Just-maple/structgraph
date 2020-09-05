package dao

type User interface {
	GetUserByID()
}

type Book interface {
	GetBookByID()
}
