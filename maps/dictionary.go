package maps

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string { return string(e) }

type Dictionary map[string] string

func (d Dictionary) Search(word string) (string, error) {
	difinition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return difinition, nil
}

func (d Dictionary) Add(word, difinition string) error  {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = difinition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, difinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = difinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
