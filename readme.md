# Go Bible Study

I wrote this to assist my daily Bible study routine.

When run, you get a random:
- Psalm
- chapter from the Old Testament
- chapter from the New Testament
- Proverb

Each will be printed to stdout.

If you use the -browser flag, it will open each in your default browser as well.

## Build instructions

`go build biblestudy.go bible.go browser.go`

That's how I do it but I just found out you can also do `go build .`

After that I manually copy and paste to a scripts directory I have added to PATH.

## TO DO
Flags for the following
- ~~Enable or disable opening in browser~~
- Choose Bible translation to read on biblegateway (Tried implementing already but doesn't work, appears to be a problem on biblegateway)
