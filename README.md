# autoinstall

autoinstall is a command-line tool that detects the technologies used in the current project and helps you automatically pull project dependencies based on the package manager being used.


## build

- get packages

```sh
go mod tidy
```

- build via just

Please ensure you have installed [just](https://github.com/casey/just) before building.

```sh
just release # building via just
```
