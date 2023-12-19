# Collection Package for Go

The `collection` package is a generic collection type for Go, inspired by [@discordjs/collection](https://github.com/discordjs/collection). It provides a flexible and easy-to-use data structure that allows you to store and manipulate key-value pairs.

## Features

- **Thread-Safe Operations**: The `Collection` type is equipped with a read-write mutex (`sync.RWMutex`), ensuring safe concurrent access to the collection.

- **Basic Operations**: Perform common operations like `Set`, `Get`, and `Delete` to manage key-value pairs in the collection.

- **Size and Keys**: Easily obtain the size of the collection with `Size()` and retrieve all keys using `Keys()`.

- **Iteration**: Iterate over the collection with `Each`, applying a callback function to each element.

- **Filtering and Mapping**: Use `Filter` to create a new collection with elements that satisfy a given predicate. Similarly, `Map` transforms each element using a specified function.

- **JSON Serialization**: Serialize the collection into a JSON-encoded byte slice, representing each element as an array containing a key and a value.

- **Random Access**: Access a random value from the collection using `Random`.

- **Functional Operations**: Use functions like `Every` and `Some` to check if all or some elements satisfy a given test.

- **Entry Management**: Obtain a slice of `CollectionEntry` with key-value pairs using `Entries`.

- **Sweeping**: Remove items from the collection that satisfy a provided filter function with `Sweep`.

- **First and Last Elements**: Access the first and last element in the collection using `First`, `FirstN`, `Last`, and `LastN`.

## Installation

```bash
go get github.com/yourusername/collection
```

## Example Usage

```go
package main

import (
	"fmt"
	"github.com/yourusername/collection"
)

type User struct {
	Id   int
	Name string
}

func main() {
	// Create a new collection of User type with an initial size of 100
	col := collection.New[User](100)

	// Set values in the collection
	col.Set("user1", User{1, "John"})
	col.Set("user2", User{2, "Jane"})

	// Get a value from the collection
	user := col.Get("user1")
	fmt.Println(user) // Output: {1 John}

	// Delete a value from the collection
	col.Delete("user2")

	// Iterate over the collection and print each element
	col.Each(func(key string, value User) {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	})
}
```

## Contributing

Contributions are welcome! Feel free to open issues or pull requests to improve the package.

## License

This package is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.