package repo

import (
	"database/sql"
	"log"
	"revyu/entities"
)

type Db struct {
	db *sql.DB
}

func (d *Db) Create(id int, name string, state string) error {
	var err error
	_, err = d.db.Exec("INSERT INTO cities(id, name, state) VALUES($1, $2, $3)", id, name, state)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *Db) Delete(id int) error {
	_, err := d.db.Exec("DELETE FROM cities WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *Db) Update(id int, name string, state string) error {
	_, err := d.db.Exec("UPDATE cities SET name = $1, state = $2 WHERE id = $3", name, state, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *Db) List() ([]entities.City, error) {
	rows, err := d.db.Query("SELECT id, name, state FROM cities")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var cities []entities.City
	for rows.Next() {
		var city entities.City
		err := rows.Scan(&city.Id, &city.Name, &city.State)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cities = append(cities, city)

	}
	return cities, nil
}
