package repositories

import "context"

type Lightbulbs struct {
	data map[string]bool
}

func New() *Lightbulbs {
	return &Lightbulbs{
		data: map[string]bool{},
	}
}

func (l *Lightbulbs) Get(ctx context.Context) (map[string]bool, error) {
	return l.data, nil
}

func (l *Lightbulbs) GetById(ctx context.Context, name string) (bool, error) {
	return l.data[name], nil
}

func (l *Lightbulbs) Create(ctx context.Context, name string, value bool) error {
	l.data[name] = value
	return nil
}

func (l *Lightbulbs) Update(ctx context.Context, name string, value bool) error {
	l.data[name] = value
	return nil
}

func (l *Lightbulbs) Delete(ctx context.Context, name string) error {
	delete(l.data, name)
	return nil
}
