package database

// TranslationsGetter .
type TranslationsGetter interface {
	GetTranslations(name string, suraNumber, startVerse, numberOfVerses int) ([]Translation, error)
}
