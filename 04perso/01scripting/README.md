# Cobra usage
*The commands are to be added under the cmd directory*
**Source:** [sbstjn.com](https://sbstjn.com/create-golang-cli-application-with-cobra-and-goxc.html)

## Initialization
> ``cobra ``
## main.go
main.go basically serves only to initialize cobra

## Note:
```bash
go build -o heft github.com/MuchChaca/GoLangTraining/04perso/01scripting/cli
```

## Build
* **With ``go build``**
> ``go build -o heft github.com/MuchChaca/GoLangTraining/04perso/01scripting/cli``
* **With ``goxc``**
> ``go get -v github.com/laher/goxc``
:point_right: create a make file to handle the build (cf. website from Source)

### Basic Crossbuiling
set the GOOS and GOARCH to the desired values:

| $GOOS		| $GOARCH	|
| ----------|-----------|
| android	|	arm		|
| darwin		|	386		|
| darwin		|	amd64		|
| darwin		|	arm		|
| darwin		|	arm64		|
| dragonfly	|	amd64		|
| freebsd	|	386		|
| freebsd	|	amd64		|
| freebsd	|	arm		|
| linux		|	386		|
| linux		|	amd64		|
| linux		|	arm		|
| linux		|	arm64		|
| linux		|	ppc64		|
| linux		|	ppc64le	|
| linux		|	mips		|
| linux		|	mipsle	|
| linux		|	mips64	|
| linux		|	mips64le	|
| netbsd		|	386		|
| netbsd		|	amd64		|
| netbsd		|	arm		|
| openbsd	|	386		|
| openbsd	|	amd64		|
| openbsd	|	arm		|
| plan9		|	386		|
| plan9		|	amd64		|
| solaris	|	amd64		|
| windows	|	386		|
| windows	|	amd64		|