package main

type Dictionary map[string]string

func (d Dictionary) Search(word string) (meaning string) {
	if meaning, exists := d[word]; exists {
		return meaning
	}
	return "uh oh, word not found in the dictionary"
}

func TypeAliasUsage() {
	dictionary := Dictionary{"volition": "power of using one's will."}
	dictionary.Search("volition")
}
