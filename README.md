# go-cli-sample

Sample command-line application written in Go. It takes an **action** as the first
parameter, followed by the parameters specific to that action.

```
go-cli-sample <action> [parameters...]
```

It uses only the Go standard library — there are no external dependencies.

## Prerequisites

- [Go](https://go.dev/dl/) 1.22 or later

Check your installation with:

```
go version
```

## Building

From the project root:

**Windows**

```
go build -o build/go-cli-sample.exe .
```

**Linux / macOS**

```
go build -o build/go-cli-sample .
```

This produces the `go-cli-sample` (or `go-cli-sample.exe`) executable inside the `build` folder.
`go build` creates the folder if it does not exist yet.

### Cross-compiling

Go supports cross-compilation by setting `GOOS` and `GOARCH`. For example, to
build a Linux binary from Windows (PowerShell):

```powershell
$env:GOOS = "linux"; $env:GOARCH = "amd64"; go build -o build/go-cli-sample .
```

Or from Linux/macOS targeting Windows:

```
GOOS=windows GOARCH=amd64 go build -o build/go-cli-sample.exe .
```

## Running

### Without building (development)

```
go run . <action> [parameters...]
```

Example:

```
go run . greet -name Mary
```

### Using the executable

**Windows (PowerShell)**

```powershell
.\build\go-cli-sample.exe greet -name Mary
```

**Linux / macOS**

```
./build/go-cli-sample greet -name Mary
```

## Available actions

| Action    | Description              | Parameters                                     |
|-----------|--------------------------|------------------------------------------------|
| `greet`   | Prints a greeting        | `-name <text>`, `-times <n>`, `-shout`         |
| `sum`     | Sums a list of numbers   | `-decimals <n>` (optional), then the numbers   |
| `reverse` | Reverses a text          | the text to reverse                            |
| `help`    | Lists the actions        | —                                              |

To see the parameters of a specific action:

```
go-cli-sample <action> -h
```

### Examples

```
$ go-cli-sample greet
Hello, world!

$ go-cli-sample greet -name Mary -times 2 -shout
HELLO, MARY!
HELLO, MARY!

$ go-cli-sample sum 1 2 3.5
6.50

$ go-cli-sample sum -decimals 0 10 20
30

$ go-cli-sample reverse Hello world
dlrow olleH
```

> **Note:** an action's flags must come **before** its positional arguments.
> For example, `sum -decimals 1 1 2` works, but in `sum 1 2 -decimals 1`
> `-decimals` is treated as an invalid number.

## Exit codes

| Code | Meaning                                      |
|------|----------------------------------------------|
| `0`  | Success                                      |
| `1`  | Error while running the action (e.g. `sum abc`) |
| `2`  | Missing or unknown action                    |

## Project structure

```
.
├── build/        # compiled executables (output of go build)
├── go.mod        # module definition
├── main.go       # reads the action and dispatches to the command
└── commands.go   # implementation of each action
```

## Adding a new action

1. Write a function with the signature `func(args []string) error` in `commands.go`,
   using its own `flag.NewFlagSet` to parse the parameters:

   ```go
   func cmdEcho(args []string) error {
   	fs := flag.NewFlagSet("echo", flag.ContinueOnError)
   	if err := fs.Parse(args); err != nil {
   		return ignoreHelp(err)
   	}
   	fmt.Println(strings.Join(fs.Args(), " "))
   	return nil
   }
   ```

2. Register it in the `commands` map in `main.go`:

   ```go
   "echo": {"Repeats the given text", cmdEcho},
   ```

3. Rebuild. The new action shows up in `go-cli-sample help`.
