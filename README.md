A simple Go package for generating ANSI terminal color codes.  
Useful for adding colors to terminal output in CLI applications.

---


## 📦 Installation

To use this package in your Go project:

```bash
go get github.com/swansestudio/colorset
```

Make sure you have Go modules enabled.

---

## 🌈 Functions

### `ColorRand() string`

Returns a random ANSI foreground color code from the basic 8-color palette (30–37):


| Code | Color    |
|------|----------|
| 30   | Black    |
| 31   | Red      |
| 32   | Green    |
| 33   | Yellow   |
| 34   | Blue     |
| 35   | Magenta  |
| 36   | Cyan     |
| 37   | White    |

Useful when you're in a hurry and don't care about specific colors, but just want to visually separate outputs in the terminal.

Example:
```go
fmt.Printf("%sRandom colored text%s\n", cs.ColorRand(), cs.ColorReset())
```

---

### `ColorSet(color string) string`

Returns a specific ANSI foreground color escape sequence based on a short identifier.

| Code | Color        | ANSI Code     |
|------|--------------|---------------|
| `"r"` | Red          | `\x1b[31m`    |
| `"g"` | Green        | `\x1b[32m`    |
| `"y"` | Yellow       | `\x1b[33m`    |
| `"b"` | Blue         | `\x1b[34m`    |
| `"p"` | Magenta      | `\x1b[35m`    |
| `"c"` | Cyan         | `\x1b[36m`    |
| `"w"` | White        | `\x1b[37m`    |
| `"_"` | Reset        | `\x1b[0m`    |

Example:
```go
fmt.Println(cs.ColorSet("g") + "This is green text" + cs.ColorReset())
```

---

### `ColorReset() string`

Resets all terminal formatting (including color). Use after applying any color escape sequence.

Equivalent to `cs.ColorSet("_")`.

Example:
```go
fmt.Println(cs.ColorSet("r") + "Red Text" + cs.ColorReset())
```

---

## 💡 Usage Examples

### 🔹 Option 1: Direct function calls (for clarity)

```go
package main

import (
	"fmt"
	cs "github.com/swansestudio/colorset"
)

func main() {
	fmt.Println(cs.ColorSet("g") + "Success!" + cs.ColorSet("_"))
	fmt.Println(cs.ColorSet("r") + "Error!" + cs.ColorReset())
	fmt.Println(cs.ColorRand() + "Random colored line!" + cs.ColorReset())
}
```

---

### 🔹 Option 2: Use shorthand variables (for brevity)

Define shortcuts at the top of your file:

```go
var (
    crd = cs.ColorRand()  // random color from 30 to 37

    cr = cs.ColorSet("r")  // red output
    _ = cr
    cg = cs.ColorSet("g")  // green output
    _ = cg
    cy = cs.ColorSet("y")  // yellow output
    _ = cy
    cb = cs.ColorSet("b")  // blue output
    _ = cb
    cc = cs.ColorSet("c")  // cyan output
    _ = cc
    cp = cs.ColorSet("p")  // magenta/purple output
    _ = cp
    cw = cs.ColorSet("w")  // white output
    _ = cw
    cres = cs.ColorSet("_")  // reset color
    _ = cres
)
```
or use `ColorSet("_")` as reset
```go
var (
	cres = cs.ColorReset()
)
```

Then write cleaner lines:

```go
func main() {
	fmt.Println(cg + "All good!" + cres)
	fmt.Println(cr + "Something went wrong!" + cres)
	fmt.Println(crd + "Fancy random message" + cres)
}
```

---

## 📄 License

MIT License – see [LICENSE](LICENSE) for details.
