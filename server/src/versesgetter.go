package main

type versesGetter interface {
	getVerses(suraNumber, startVerse, numberOfVerses int) ([]verse, error)
}
