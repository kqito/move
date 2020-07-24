# move

![Release](https://img.shields.io/github/v/release/kqito/move)
![License](https://img.shields.io/github/license/kqito/move)

Commands to move or copy multiple files or directories at once.


## Features
- Multiple files and directories can be moved or copied together at once.
- Easy to select files and directories.
- Easy to use.


## Installation
You can install the package from npm.
```
go get github.com/kqito/move
```


## Usage
### General
You can move the specified source as follows. (It behaves like the `mv` command)

```shell
$ move operation/dir target/dir
```

### Run as `cp` command
You can also copy the specified sources to a specified directory, as in the `cp` command.

```shell
$ move operation/dir target/dir -c
```


## License
[MIT © kqito](./LICENSE)
