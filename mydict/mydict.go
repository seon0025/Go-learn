package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errCantUpdate = errors.New("can't update")
var errWordExists = errors.New("already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, error := d.Search(word)
	switch error {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, error := d.Search(word)
	switch error {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string){
	delete(d, word)
}