# Go Bible Study

I wrote this to assist my daily Bible study routine.

When run, you get a random:
- Psalm
- chapter from the Old Testament
- chapter from the New Testament
- Proverb

Each will be printed to stdout, and also will be opened on biblegateway.com in your browser.

## Build instructions

`go build biblestudy.go bible.go browser.go`

After that I manually copy and paste to a scripts directory I have added to PATH.

## TO DO
Flags for the following
- Enable or disable opening in browser
- Choose Bible translation to read on biblegateway