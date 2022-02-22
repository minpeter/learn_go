package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string


var (
	errNotFound = errors.New("not found")
	errCantUpdate = errors.New("can't update non-existent word")
	errWordExist = errors.New("word already exist")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExist
	}
	return nil
}


// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
