# `GREG`
Greg is a very limited cli tool for using "grep like" string searching in windows. The only thing it does is read from a file or stdin and print the lines that match the given text. 

## TODO
- [ ] Handle wrong inputs and wirte tests
- [ ] Add option to not print the line number

## Usage
```shell
<some command> | greg <searchtext>
```

## Installation
**Use powershell as administrator to install greg**

### Install
```shell
iex(iwr -Uri "https://raw.githubusercontent.com/DiniFarb/greg/main/install.ps1" -UseBasicParsing)
```
### Uninstall
```shell
iex(iwr -Uri "https://raw.githubusercontent.com/DiniFarb/greg/main/uninstall.ps1" -UseBasicParsing)
```
