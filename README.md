# selpg
SELect PaGes

Sometimes one needs to extract only a specified range of
pages from an input text file. This program allows the user to do
that.

# Install
```bash
go get -u github.com/Andiedie/selpg
go install github.com/Andiedie/selpg
```

Run it:
```
$GOPATH/bin/selpg
```

# Usage
```
usage: selpg [flags] path
  -e int
        [REQUIRED] end page number (default -1)
  -f    use '\f' to paging instead of line
  -l int
        line number per page (default 76)
  -s int
        [REQUIRED] start page number (default -1)
```

# Example
```
$ cat a.txt
1
2
3
```

```
$ selpg -s 1 -e 1 a.txt
1
2
3

$ selpg -s 1 -e 2 -l 2 a.txt
1
2

$ selpg -s 1 -e 2 -l 2 < a.txt
1
2

$ cat a.txt | selpg -s 1 -e 2 -l 2 > b.txt

# b.txt becomes:
# 1
# 2

$ selpg -s 1 -e 1 -l 2 a.txt | wc
      2       2       6
```

If the file (or other input) is paging by `\f` instead of number of lines, use `-f` flags.

Pages will be be split up by `\f` and `-l` no longer takes effect.

# LICENSE
WTFPL
