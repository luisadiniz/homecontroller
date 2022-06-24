package repositories

import "context"

type InMemoryDB struct {
	data map[string]bool
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: map[string]bool{},
	}
}

func (l *InMemoryDB) Get(ctx context.Context) (map[string]bool, error) {
	return l.data, nil
}

func (l *InMemoryDB) GetById(ctx context.Context, name string) (bool, error) {
	return l.data[name], nil
}

func (l *InMemoryDB) Create(ctx context.Context, name string, value bool) error {
	l.data[name] = value
	return nil
}

func (l *InMemoryDB) Update(ctx context.Context, name string, value bool) error {
	l.data[name] = value
	return nil
}

func (l *InMemoryDB) Delete(ctx context.Context, name string) error {
	delete(l.data, name)
	return nil
}
