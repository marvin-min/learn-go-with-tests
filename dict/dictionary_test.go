package dict

import "testing"

func TestDictionary(t *testing.T) {
	d := Dictionary{"test": "this is a test string"}
	t.Run("known word", func(t *testing.T) {
		got, err := d.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertString(got, "this is a test string", t)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := d.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		d := Dictionary{}
		d.Add("test", "this is just a test")

		want := "this is just a test"
		assertDefinition(d, "test", want, t)

	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		err := dict.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(dict, word, definition, t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing key in dictionary ", func(t *testing.T) {
		key := "test"
		word := "init content"
		dict := Dictionary{key: word}
		newContent := "new content"
		dict.Update(key, newContent)
		assertDefinition(dict, key, newContent, t)
	})

	t.Run("update non existing key in dictionary ", func(t *testing.T) {
		key := "test"
		word := "init content"
		dict := Dictionary{key: word}
		newKey := "newKey"
		newContent := "new content"
		err := dict.Update(newKey, newContent)
		assertError(t, err, ErrorKeyDoeNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing key in dictionary", func(t *testing.T) {
		key := "testee"
		dict := Dictionary{key: "test content"}
		dict.Delete(key)
		_, err := dict.Search("test")
		assertError(t, err, ErrNotFound)
	})
}

func assertDefinition(d Dictionary, word, definition string, t *testing.T) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(got, definition, t)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
func assertString(got string, want string, t *testing.T) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
