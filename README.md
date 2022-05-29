# go-sgr

Package **sgr** provides tools for working with **Select Graphic Rendition** (**SGR**) control sequences, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-sgr

[![GoDoc](https://godoc.org/github.com/reiver/go-sgr?status.svg)](https://godoc.org/github.com/reiver/go-sgr)

## Select Graphic Rendition (SGR)

**Select Graphic Rendition** (**SGR**) are **control sequences** of the following form:
```
CSI n "m"
```
Or in other words:
```
ESC "[" n "m"
```
Or:
```
"\x1b" "[" n "m"
```

Here `n` is number written in text (rather than binary).

For example:
```
CSI "1" "m"
```
Or in other words:
```
ESC "[" "1" "m"
```
Or:
```
"\x1b" "[" "1" "m"
```
Or:
```
"\x1b[1m"
```

The **SGR** control sequences are as follows:

| n   | name              | sequence        | string        | notes                                                                    |
|-----|-------------------|-----------------|---------------|--------------------------------------------------------------------------|
|   0 | off               | `CSI "0" "m"`   | `"\x1b[0m"`   | Cancels the effect of any preceding occurrence of SGR control sequences. |
|   1 | bold              | `CSI "1" "m"`   | `"\x1b[1m"`   |                                                                          |
|   2 | faint             | `CSI "2" "m"`   | `"\x1b[2m"`   |                                                                          |
|   3 | italicized        | `CSI "3" "m"`   | `"\x1b[3m"`   |                                                                          |
|   4 | singly underlined | `CSI "4" "m"    | `"\x1b[4m"`   |                                                                          |
|   5 | slowly blinking   | `CSI "5" "m"`   | `"\x1b[5m"`   |                                                                          |
|   6 | rapidly blinking  | `CSI "6" "n"`   | `"\x1b[6m"`   |                                                                          |
|   7 | negative          | `CSI "7" "n"`   | `"\x1b[7m"`   |                                                                          |
|   8 | concealed         | `CSI "8" "n"`   | `"\x1b[8m"`   |                                                                          |
