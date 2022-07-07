package storage

import (
	"math/rand"
)

type WordsOfWisdom struct {
	book [][]byte
}

func NewWordsOfWisdom() *WordsOfWisdom {
	wow := &WordsOfWisdom{}
	wow.initBook()

	return wow
}

func (w *WordsOfWisdom) initBook() {
	w.book = [][]byte{
		[]byte("You are perfect because of your imperfections."),
		[]byte("Do what inspires you. Life is too short not to love the job you do every day."),
		[]byte("Complaining will not get anything done."),
		[]byte("At the end of your day, you’ve done your best. Even if you haven’t accomplished all that’s on your list. You’ve given it you’re all."),
		[]byte("You don’t need to have it figured all out. Taking the wrong path is part of the process."),
		[]byte("Never lose yourself because of someone else. You are perfect just the way you are."),
		[]byte("Trust your gut. If you ever feel it's not right, then it's not."),
		[]byte("A smile is a free way to brighten someone’s day."),
	}
}

func (w *WordsOfWisdom) GetWordOfWisdom() []byte {
	i := rand.Int() % len(w.book)
	return w.book[i]
}
