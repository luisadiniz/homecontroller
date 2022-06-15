package repositories

import "context"

var (
	lightbulbs = map[string]bool{}
)

func GetLightbulbs(ctx context.Context) (map[string]bool, error) {
	return lightbulbs, nil
}

func GetLightbulbById(ctx context.Context, name string) (bool, error) {
	return lightbulbs[name], nil
}

func CreateLightbulbs(ctx context.Context, name string, value bool) error {
	lightbulbs[name] = value
	return nil
}

func UpdateLightbulb(ctx context.Context, name string, value bool) error {
	lightbulbs[name] = value
	return nil
}

func DeleteLightbulb(ctx context.Context, name string) error {
	delete(lightbulbs, name)
	return nil
}
