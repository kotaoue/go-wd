# go-wd
Get the same working directory path at 'go run' and after 'go build'

## Usage
```Go
func main() {
	dir, err := wd.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(dir)
}
```

## Links
* [kotaoue/go-wd-tester](https://github.com/kotaoue/go-wd-tester)