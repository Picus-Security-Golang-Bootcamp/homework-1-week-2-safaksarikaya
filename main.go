package main

import (
	"fmt"
	"os"
	"strings"
)

var books = [...]string{"In Search of Lost Time", "Ulysses", "Don Quixote", "One Hundred Years of Solitude", "The Great Gatsby", "Moby Dick",
	"War and Peace", "Hamlet", "The Odyssey", "Madame Bovary",
	"The Divine Comedy", "Lolita", "The Brothers Karamazov", "Crime and Punishment",
	"Wuthering Heights", "The Catcher in the Rye", "Pride and Prejudice", "The Adventures of Huckleberry Finn",
	"Anna Karenina", "Alice's Adventures in Wonderland", "The Iliad", "To the Lighthouse",
	"Catch-22", "Heart of Darkness", "The Sound and the Fury", "Nineteen Eighty Four",
	"Great Expectations", "One Thousand and One Nights", "The Grapes of Wrath", "Absalom, Absalom!",
	"Invisible Man", "To Kill a Mockingbird", "The Trial", "The Red and the Black",
	"Middlemarch", "Gulliver's Travels", "Beloved", "Mrs. Dalloway", "The Stories of Anton Chekhov", "The Stranger", "Jane Eyre", "The Aeneid",
	"Collected Fiction", "The Sun Also Rises", "David Copperfield", "Tristram Shandy",
	"Leaves of Grass", "The Magic Mountain", "A Portrait of the Artist as a Young Man",
	"Midnight's Children", "Oedipus the King", "Candide", "The Lord of the Rings",
	"The Idiot", "Les Misérables", "A Passage to India", "The Old Man and the Sea",
	"Things Fall Apart", "Emma", "For Whom the Bell Tolls", "The Complete Stories of Franz Kafka",
	"The Metamorphosis", "The Portrait of a Lady", "Frankenstein", "Pale Fire",
	"Antigone", "As I Lay Dying", "The Color Purple", "The Possessed", "Gone With the Wind", "Lord of the Flies", "Brave New World", "The Complete Tales and Poems of Edgar Allan Poe",
	"The Age of Innocence", "Dead Souls", "On the Road", "The Good Soldier", "Animal Farm", "Orlando: A Biography",
	"The Canterbury Tales", "Vanity Fair", "Under the Volcano", "The Waste Land",
	"A Farewell to Arms", "Journey to the End of The Night", "The Castle", "A Sentimental Education",
	"The Scarlet Letter", "Slaughterhouse-Five", "The Handmaid's Tale", "Charlotte's Web",
	"Native Son", "The Charterhouse of Parma", "Paradise Lost", "Gargantua and Pantagruel",
	"Poems of Emily Dickinson", "Faust", "Rebecca", "The Flowers of Evil", "Decameron"}

func main() {

	args := os.Args[1:]
	arg := ""
	for i := 0; i < len(args); i++ {
		arg += args[i]
		if i+1 == len(args) {
			break
		}
		arg += " "
	}
	if arg == "list" {
		list()
	}
	if arg[:6] == "search" {
		fmt.Println(search(arg[7:]))
	}
}

// list lists books
func list() {
	for i, arg := range books {
		fmt.Printf(" %d : %s\n", i+1, arg)
	}
}

//search search for the value entered by the use in the books
func search(bookName string) string {
	book := strings.ToLower(bookName)
	message := "Not have this book: " + bookName
	for i := 0; i < len(books); i++ {
		b := books[i]
		b = strings.ToLower(b)
		if b == book {
			message = "We have this book: " + books[i]
			break
		}
	}
	return message
}
