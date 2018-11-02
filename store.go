package main

import (
	"database/sql"
)

type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBird(bird *Bird) error {
	_, err := store.db.Query("INSERT INTO birds (species, description) VALUES ($1, $2)", bird.Species, bird.Description)
	return err
}

func (store *dbStore) GetBirds() ([]*Bird, error) {
	rows, err := store.db.Query("SELECT species, description from birds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	birds := []*Bird{}
	for rows.Next() {
		bird := &Bird{}
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		birds = append(birds, bird)
	}
	return birds, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
