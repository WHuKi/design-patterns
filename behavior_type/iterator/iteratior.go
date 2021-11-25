package iterator

//
type iterator interface {
	hasNext() bool
	getNext() *user
}

type collection interface {
	createIterator() iterator
}

type user struct {
	name string
	age  int
}

// 迭代器实现
type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}
