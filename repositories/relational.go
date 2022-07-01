package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

type RelationalRepository struct {
	data *sql.DB
}

func NewRelationalRepository() (*RelationalRepository, error) {

	connectionString := fmt.Sprintf("host=db port=26257 dbname=mydb user=luisa password=anything sslmode=disable")

	var (
		db  *sql.DB
		err error
	)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS lightbulbs (name STRING PRIMARY KEY, isOn BOOL)"); err != nil {
		return nil, err
	}

	return &RelationalRepository{
		data: db,
	}, nil
}

func (l *RelationalRepository) Get(ctx context.Context) (map[string]bool, error) {

	lightbulbs := map[string]bool{}
	rows, err := l.data.Query("SELECT name, isOn FROM lightbulbs")
	if err != nil {
		fmt.Println("No rows were returned!")
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name string
			on   bool
		)
		err := rows.Scan(name, on)
		if err != nil {
			fmt.Println(err.Error())
		}

		lightbulbs[name] = on
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
	return lightbulbs, nil
}

func (l *RelationalRepository) GetById(ctx context.Context, name string) (bool, error) {

	sqlStatement := "SELECT name, isOn FROM lightbulbs WHERE name=$1"
	var repositoryName string
	var isOn bool

	row := l.data.QueryRow(sqlStatement, name)
	err := row.Scan(&repositoryName, &isOn)
	if err != nil {
		fmt.Println("No rows were returned!")
	}

	return isOn, nil
}

func (l *RelationalRepository) Create(ctx context.Context, name string, value bool) error {

	sqlStatement := "INSERT INTO lightbulbs (name, isOn) VALUES ($1, $2)"
	_, err := l.data.Exec(sqlStatement, name, value)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (l *RelationalRepository) Update(ctx context.Context, name string, value bool) error {

	sqlStatement := "UPDATE lightbulbs SET isOn=$2 WHERE name=$1;"
	_, err := l.data.Exec(sqlStatement, name, value)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (l *RelationalRepository) Delete(ctx context.Context, name string) error {

	sqlStatement := "DELETE FROM lightbulbs WHERE name=$1"
	_, err := l.data.Exec(sqlStatement, name)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
