package main

const (
	OldTestament = iota
	NewTestament = iota
)

type Book struct {
	Name      string
	Chapters  int
	Testament int
}

func InitializeBible() ([]Book, Book, Book) {
	bible := make([]Book, 0)

	Psalms := Book{Name: "Psalms", Chapters: 150, Testament: OldTestament}
	Proverbs := Book{Name: "Proverbs", Chapters: 31, Testament: OldTestament}

	var b Book
	b.Name = "Genesis"
	b.Chapters = 50
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Exodus"
	b.Chapters = 40
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Leviticus"
	b.Chapters = 27
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Numbers"
	b.Chapters = 36
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Deuteronomy"
	b.Chapters = 34
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Joshua"
	b.Chapters = 24
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Judges"
	b.Chapters = 21
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Ruth"
	b.Chapters = 4
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "1 Samuel"
	b.Chapters = 31
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "2 Samuel"
	b.Chapters = 24
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "1 Kings"
	b.Chapters = 22
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "2 Kings"
	b.Chapters = 25
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "1 Chronicles"
	b.Chapters = 29
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "2 Chronicles"
	b.Chapters = 36
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Ezra"
	b.Chapters = 10
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Nehemiah"
	b.Chapters = 13
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Esther"
	b.Chapters = 10
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Job"
	b.Chapters = 42
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Psalms"
	b.Chapters = 150
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Proverbs"
	b.Chapters = 31
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Ecclesiastes"
	b.Chapters = 12
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Song of Solomon"
	b.Chapters = 8
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Isaiah"
	b.Chapters = 66
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Jeremiah"
	b.Chapters = 52
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Lamentations"
	b.Chapters = 5
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Ezekiel"
	b.Chapters = 48
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Daniel"
	b.Chapters = 12
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Hosea"
	b.Chapters = 14
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Joel"
	b.Chapters = 3
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Amos"
	b.Chapters = 9
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Obadiah"
	b.Chapters = 1
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Jonah"
	b.Chapters = 4
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Micah"
	b.Chapters = 7
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Nahum"
	b.Chapters = 3
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Habakkuk"
	b.Chapters = 3
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Zephaniah"
	b.Chapters = 3
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Haggai"
	b.Chapters = 2
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Zechariah"
	b.Chapters = 14
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Malachi"
	b.Chapters = 4
	b.Testament = OldTestament
	bible = append(bible, b)

	b.Name = "Matthew"
	b.Chapters = 28
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Mark"
	b.Chapters = 16
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Luke"
	b.Chapters = 24
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "John"
	b.Chapters = 21
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Acts"
	b.Chapters = 28
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Romans"
	b.Chapters = 16
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "1 Corinthians"
	b.Chapters = 16
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "2 Corinthians"
	b.Chapters = 13
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Galatians"
	b.Chapters = 6
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Ephesians"
	b.Chapters = 6
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Phillipians"
	b.Chapters = 4
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Colossians"
	b.Chapters = 4
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "1 Thessalonians"
	b.Chapters = 5
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "2 Thessalonians"
	b.Chapters = 4
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "1 Timothy"
	b.Chapters = 6
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "2 Timothy"
	b.Chapters = 4
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Titus"
	b.Chapters = 3
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Philemon"
	b.Chapters = 1
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Hebrews"
	b.Chapters = 13
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "James"
	b.Chapters = 5
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "1 Peter"
	b.Chapters = 5
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "2 Peter"
	b.Chapters = 3
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "1 John"
	b.Chapters = 5
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "2 John"
	b.Chapters = 1
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "3 John"
	b.Chapters = 1
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Jude"
	b.Chapters = 1
	b.Testament = NewTestament
	bible = append(bible, b)

	b.Name = "Revelation"
	b.Chapters = 22
	b.Testament = NewTestament
	bible = append(bible, b)

	return bible, Psalms, Proverbs
}
