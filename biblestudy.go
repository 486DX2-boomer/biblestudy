package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func OpenInBrowser(bookname string, chapter int) {

	var sb strings.Builder
	sb.WriteString("https://www.biblegateway.com/passage/?search=")
	sb.WriteString(bookname)
	sb.WriteString("+")
	sb.WriteString(strconv.Itoa(chapter))
	sb.WriteString("&version=KJV")

	url := sb.String()
	// Fix blank spaces
	url = strings.Replace(url, " ", "%20", -1)
	Open(url)
}

func getPsalmsOrProverb(b *Book) {
	r := rand.Intn(b.Chapters)
	if r == 0 {
		r = 1
	}

	OpenInBrowser(b.Name, r)

	fmt.Println(b.Name, r)
}

func getRandomChapter(testament int, bible []Book) {
	if testament == OldTestament {
		r := rand.Intn(39)
		b := bible[r]
		for b.Name == "Psalms" || b.Name == "Proverbs" {
			r = rand.Intn(39)
			b = bible[r]
		}
		r = rand.Intn(b.Chapters)

		// For books that are one chapter in length
		if r == 0 {
			r = 1
		}

		OpenInBrowser(b.Name, r)

		fmt.Println(b.Name, r)
	}

	if testament == NewTestament {
		r := rand.Intn(27) + 39
		b := bible[r]
		r = rand.Intn(b.Chapters)

		// For books that are one chapter in length
		if r == 0 {
			r = 1
		}

		OpenInBrowser(b.Name, r)

		fmt.Println(b.Name, r)
	}
}

func main() {
	bible, Psalms, Proverbs := InitializeBible()
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bible study for today:")

	// Get random Psalms
	getPsalmsOrProverb(&Psalms)

	// Get random Old Testament chapter
	getRandomChapter(OldTestament, bible)

	// Get random New Testament chapter
	getRandomChapter(NewTestament, bible)

	// Get random Proverb
	getPsalmsOrProverb(&Proverbs)

}
