# Notes while Coding

## Printing Stuff

* `fmt.Print` - Just print stuff
* `fmt.Fprintf` - Print to a file handler in a formated way
* `fmt.Printf` - Print in a formated way (don't forget that you need line returns)
   * `fmt.Printf("Hello %v\n", "Someone")`
   * If you use `%q` instead of `%v` it will be shown quoted
* `fmt.Println` - Print with a line return at the end
* `fmt.Sprintf` - Like Printf, but returns a string you can use

----

* `log.Println` - Print with line returns and prepended timestamps


## Variables
* declare `var varname type`
* and the adding a value of the correct `type` via `varname = value`
* inside a function you can use the walrus operator `:=` to do both at the same time: `varname := value`




