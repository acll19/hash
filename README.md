# hash

A bash command to generate non cryptographic hashes from a strings

## Usage

```bash
hash -l <length> <string>
```

Where:

- `-l <length>`: Optional. Specifies the length of the generated hash. Default is 8 characters.
- `<string>`: The input string to be hashed.

## Example

```bash
hash -l 12 "Hello, World!"
```

This command will generate a non cryptographic hash of the string "Hello, World!" with a length of 12 characters.

## Building from source

You need go installed on your machine to build the `hash` command from source. You can download and install go from the official website: https://golang.org/dl/

To build the `hash` command from source, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/hash/hash.git
   ```
2. Navigate to the project directory:
   ```bash
   cd hash
   ```
3. Build the project using go:
   ```bash
   go build -o hash main.go
   ```

````
4. The `hash` executable will be generated in the current directory. You can move it to a directory in your PATH for easier access:

Make sure you don't have any other executable named `hash` in your PATH to avoid conflicts. If so, consider renaming the executable or using a different name for the command.

```bash
   mv hash /usr/local/bin/
````
