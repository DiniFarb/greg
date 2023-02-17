# `GREG`
Greg is a very limited cli tool for using grep in windows. The only Thing it does is read from a file or stdin and print the lines that match the given text. 

## TODO
- [ ] Add Regex as search pattern
- [ ] Handle wrong inputs and wirte tests
- [ ] Add option to not print the line number

## Usage
```shell
greg -p <filepath> <searchtext> 
```
or
```shell
cat <filepath> | greg <searchtext>
```

## Installation
User powershell as administrator to install greg

# Install
```shell
iex(iwr -Uri "https://raw.githubusercontent.com/DiniFarb/greg/main/install.ps1" -UseBasicParsing)
```
# Uninstall
```shell
iex(iwr -Uri "https://raw.githubusercontent.com/DiniFarb/greg/main/uninstall.ps1" -UseBasicParsing)
```