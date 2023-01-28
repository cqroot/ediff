<div align="center">
  <h1>EDiff</h1>

  <p><i>Use your favorite editor to modify the list of strings.</i></p>

  <p>
    <a href="https://github.com/cqroot/ediff/actions">
      <img src="https://github.com/cqroot/ediff/workflows/test/badge.svg" alt="Action Status" />
    </a>
  </p>
</div>

## Usage

```go
package main

import (
	"fmt"

	"github.com/cqroot/ediff"
)

func main() {
	ed := ediff.New("vim")
	ed.AppendItems([]string{
		"item 1",
		"item 2",
		"item 3",
	})

	pairs, err := ed.Run()
	if err != nil {
		panic(err)
	}

	for _, pair := range pairs {
		fmt.Println("Prev:", pair.Prev, " Curr:", pair.Curr)
	}
}
```

This above code will use `vim` to open a temp file with the following content:

```
item 1
item 2
item 3
```

If you modify the content to:

```
obj 1
obj 2
obj 3
```

Then the output of the program is:

```
Prev: item 1  Curr: dot 1
Prev: item 2  Curr: dot 2
Prev: item 3  Curr: dot 3
```
