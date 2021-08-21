# Notes while Coding

## Printing Stuff
* `fmt.Print` - Just print stuff
* `fmt.Fprintf` - Print to a io.writer / file descriptor in a formated way
* `fmt.Printf` - Print in a formated way (don't forget that you need line returns)
   * `fmt.Printf("Hello %v\n", "Someone")`
   * If you use `%q` instead of `%v` it will be shown quoted
* `fmt.Println` - Print with a line return at the end
* `fmt.Sprintf` - Like Printf, but returns a string you can use
* `log.Println` - Print with line returns and prepended timestamps


## Variables
* declare `var varname type`
* and the adding a value of the correct `type` via `varname = value`
* inside a function you can use the walrus operator `:=` to do both at the same time: `varname := value`


## String quotes vs backticks
Copied from a [stackoverflow answer](https://stackoverflow.com/a/46917369).

In quotes "" you need to escape new lines, tabs and other characters that do not need to be escaped in backticks \`\`. If you put a line break in a backtick string, it is interpreted as a `\n` character, see [golang ref string literals](https://golang.org/ref/spec#String_literals).

Thus, if you say `\n` in a backtick string, it will be interpreted as the literal backslash and character n.


## Serving stuff with the net/http server
* Serving a file:
`ServeFile(w, r, name)` - https://pkg.go.dev/net/http#ServeFile
   * You should not link directly to an `index.html` file. It makes longer URLs. You can't dynamically link stuff.
   * If you use `r.URL.Path` as the `name` you will expose your root directory!

* `ServeTLS(l, handler, certFile, keyFile)` - https://pkg.go.dev/net/http#ServeTLS

## path/filepath package
* `filepath.Abs()` - Abs returns an absolute representation of path.
* `filepath.Join(elem ...string) string` - Join joins any number of path elements into a single path.


## os package
* `of.Getwd()` - Returns the working directory


