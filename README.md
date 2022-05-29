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

| n   | name                | sequence        | string         | notes                                                                    |
|-----|---------------------|-----------------|----------------|--------------------------------------------------------------------------|
|   0 | off                 | `CSI "0" "m"`   | `"\x1b[0m"`    | Cancels the effect of any preceding occurrence of SGR control sequences. |
|   1 | bold                | `CSI "1" "m"`   | `"\x1b[1m"`    |                                                                          |
|   2 | faint               | `CSI "2" "m"`   | `"\x1b[2m"`    |                                                                          |
|   3 | italicized          | `CSI "3" "m"`   | `"\x1b[3m"`    |                                                                          |
|   4 | singly underlined   | `CSI "4" "m"`   | `"\x1b[4m"`    |                                                                          |
|   5 | slowly blinking     | `CSI "5" "m"`   | `"\x1b[5m"`    |                                                                          |
|   6 | rapidly blinking    | `CSI "6" "m"`   | `"\x1b[6m"`    |                                                                          |
|   7 | negative            | `CSI "7" "m"`   | `"\x1b[7m"`    |                                                                          |
|   8 | concealed           | `CSI "8" "m"`   | `"\x1b[8m"`    |                                                                          |
|   9 | font, default       | `CSI "9" "m"`   | `"\x1b[9m"`    | default font                                                             |
|  10 | font, alternative 1 | `CSI "10" "m"`  | `"\x1b[10m"`   | first alternative font                                                   |
|  11 | font, alternative 2 | `CSI "11" "m"`  | `"\x1b[11m"`   | second alternative font                                                  |
|  12 | font, alternative 3 | `CSI "12" "m"`  | `"\x1b[12m"`   | third alternative font                                                   |
|  13 | font, alternative 4 | `CSI "13" "m"`  | `"\x1b[13m"`   | fourth alternative font                                                  |
|  14 | font, alternative 5 | `CSI "14" "m"`  | `"\x1b[14m"`   | fifth alternative font                                                   |
|  15 | font, alternative 6 | `CSI "15" "m"`  | `"\x1b[15m"`   | sixth alternative font                                                   |
|  16 | font, alternative 7 | `CSI "16" "m"`  | `"\x1b[16m"`   | seventh alternative font                                                 |
|  17 | font, alternative 8 | `CSI "17" "m"`  | `"\x1b[17m"`   | eighth alternative font                                                  |
|  18 | font, alternative 9 | `CSI "18" "m"`  | `"\x1b[18m"`   | nineth alternative font                                                  |
