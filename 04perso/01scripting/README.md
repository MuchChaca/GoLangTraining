# Cobra usage
*The commands are to be added under the cmd directory*
**Source:** [sbstjn.com](https://sbstjn.com/create-golang-cli-application-with-cobra-and-goxc.html)
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