# Native-opener

A tool to register custom URI handlers on Linux, Windows and macOS.

## Demo

Cf. [demo project](https://github.com/Minizilla/Native-opener-demo) 

## How it works

- Native-opener saves in your OS config file that it should call a Go wrapper when it receives `myapp://myfile.pdf`

- This wrapper handles extracting arguments like `myfile.pdf` and passes them to the application you specified during URI registration

## Usage

### 1. Compile the tools

```bash
go build -o native-opener . && go build -o uri-wrapper ./uriwrapper
```

### 2. Register a URI handler

```bash
./native-opener myapp /path/to/your/program/to/execute
```

### 3. Use the URI

When someone clicks on `myapp://myfile.pdf`, your program will be launched with `myfile.pdf` as an argument.

