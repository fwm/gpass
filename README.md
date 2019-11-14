# gpass: Password generating CLI tool

## Installation
Installed with `go install github.com/fwm/gpass`

## Usage
Once invoked as `gpass [flags]` it generates a password and copies it to clipboard

```
Usage of gpass:
  -a, --alpha        use alphabetical symbols (default true)
  -l, --length int   password length to generate (default 10)
  -n, --numeric      use numeric symbols (default true)
  -s, --special      use special symbols
  -v, --verbose      print out the generated password
```
