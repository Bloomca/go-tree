## Go tree

Tree CLI utility.

Usage:

```sh
go run main.go . # to show only folder
go run main.go . -f # to show with files
```

## Examples

```sh
go run main.go . -f
├───main.go (1798b)
├───main_test.go (1865b)
├───README.md (1022b)
└───testdata
	├───project
	│	├───file.txt (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───a_lorem
	│	│	├───dolor.txt (empty)
	│	│	├───gopher.png (70372b)
	│	│	└───ipsum
	│	│		└───gopher.png (70372b)
	│	├───css
	│	│	└───body.css (28b)
	│	├───empty.txt (empty)
	│	├───html
	│	│	└───index.html (57b)
	│	├───js
	│	│	└───site.js (10b)
	│	└───z_lorem
	│		├───dolor.txt (empty)
	│		├───gopher.png (70372b)
	│		└───ipsum
	│			└───gopher.png (70372b)
	├───zline
	│	├───empty.txt (empty)
	│	└───lorem
	│		├───dolor.txt (empty)
	│		├───gopher.png (70372b)
	│		└───ipsum
	│			└───gopher.png (70372b)
	└───zzfile.txt (empty)
```


```sh
go run main.go .
└───testdata
	├───project
	├───static
	│	├───a_lorem
	│	│	└───ipsum
	│	├───css
	│	├───html
	│	├───js
	│	└───z_lorem
	│		└───ipsum
	└───zline
		└───lorem
			└───ipsum
```

## Tests

Run this command from this directory:

```sh
go test -v
```

## License

MIT