package mydict

import "errors"

//Dictionary type
// you can add method to the type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExits = errors.New("That word already exists ")
var errCantUpdate = errors.New("Can't update non-exisiting word")
var errNoWord = errors.New("The word does not exist")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExits
	}
	return nil
}

//Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	// 만약에 단어가 없다면? 에러 리턴
	// 만약에 단어가 있다면? 삭제
	_, err := d.Search(word)
	if err != nil {
		return errNoWord
	}
	delete(d, word)
	return nil
}
