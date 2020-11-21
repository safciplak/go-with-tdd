package main

import "testing"

func assertError(t *testing.T, got, want error){
        t.Helper()

        if got != want {
                t.Errorf("got %q want %q", got, want)
        }
	
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}


func assertStrings(t *testing.T, got, want string){
        t.Helper()

        if got != want {
                t.Errorf("got %q want %q given, %q", got, want, "test")
        }
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
        got, err := dictionary.Search(word)
        if err != nil {
                t.Fatal("should find added word:", err)
        }

        if definition != got {
                t.Errorf("got %q want %q", got, definition)
        }
}
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
	
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unkown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		
		err := dictionary.Add(word, definition)
		
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)	
	})
	
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
	
		err := dictionary.Update(word, newDefinition)
		
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T){
	word := "test"
	dictionary := Dictionary{word: "test definition"}
	
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}