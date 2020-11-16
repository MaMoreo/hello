package maps

import "errors"

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var ErrorNotFound error = errors.New("could not find the word you were looking for")

// Dictionary of Strings
type Dictionary map[string]string

// Search in a dicctionary
func (d Dictionary) Search(word string) (string, error) {
	//map lookup. It can return 2 values. The second value is a boolean which indicates if the key was found successfully.
	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

//Add a word and a definition to a dictionary
func (d Dictionary) Add(word, definition string) {
	//An interesting property of maps is that you can modify
	// them without passing them as a pointer. This is because map is a reference type.
	d[word] = definition
}
