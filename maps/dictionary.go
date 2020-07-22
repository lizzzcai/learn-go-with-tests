package main

import "errors"

// Dictionary store definitions to words.
type Dictionary map[string]string

var (
	// ErrNotFound means the definition could not be found for the given word
	ErrNotFound = errors.New("could not find the word you were looking for")
	// ErrWordExists means you are trying to add a word that is already known
	ErrWordExists       = errors.New("word exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

// An interesting property of maps is that you can modify them without passing them as a pointer. This is because map is a reference type.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a word and definition into the dictionary.
func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
