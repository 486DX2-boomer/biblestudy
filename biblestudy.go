package main

import (
	"flag"
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
	sb.WriteString("&version=")
	sb.WriteString("KJV")
	// sb.WriteString(bibleVersion)

	url := sb.String()
	// Fix blank spaces
	url = strings.Replace(url, " ", "%20", -1)
	Open(url)
}

func getPsalmsOrProverb(b *Book, browserToggle bool) {
	r := rand.Intn(b.Chapters)
	if r == 0 {
		r = 1
	}

	if browserToggle {
		OpenInBrowser(b.Name, r)
	}

	fmt.Println(b.Name, r)
}

func getRandomChapter(testament int, bible []Book, browserToggle bool) {
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

		if browserToggle {
			OpenInBrowser(b.Name, r)
		}

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

		if browserToggle {
			OpenInBrowser(b.Name, r)
		}

		fmt.Println(b.Name, r)
	}
}

func main() {

	browserToggle := flag.Bool("browser", false, "Open selected verses on biblegateway.com in your default browser.")
	// I attempted to implement a flag for choosing a translation when opening biblegateway, but it doesn't work on their end
	// bibleVersion := flag.String("version", "KJV", "Choose which translation to open on biblegateway.com. For example: KJV, DRA, WYC")
	flag.Parse()

	bible, Psalms, Proverbs := InitializeBible()
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bible study for today:")

	getPsalmsOrProverb(&Psalms, *browserToggle)

	getRandomChapter(OldTestament, bible, *browserToggle)

	getRandomChapter(NewTestament, bible, *browserToggle)

	getPsalmsOrProverb(&Proverbs, *browserToggle)

}
