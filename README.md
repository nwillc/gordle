# Gordle

A Go language version of Wordle.

## Running

```shell
$ go build gordle.go
$ ./gorgle
```

![screenshot](screenshot.png)

## Notes

 - I use Red/Amber/Green rather than Gray/Gold/Green.
 - If some of this Go looks _odd_ that might be because I'm using the beta generics and my [genfuncs](https://github.com/nwillc/genfuncs) paackage.
 - I'm still not happy with the word list I created [dict.txt](./data/dict.txt). Using a dictionary API I removed 
proper names, and place names but more to do. Looking at past Wordle words though don't completely like their curation either.

