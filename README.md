# stdedit

`stdedit` is a tool which edits standard input with your favorite editor.

## Description

Edit stdin with your favorite editor and the result is out to stdout.<br>
It uses the editor provided via its argument or the environment variable `$STDEDIT` (Default is `vim`).<br>
GUI editor is not supported.

## Installation

```sh
go get github.com/kita127/stdedit
```

## Usage

```sh
echo hello | stdedit vim | pbcopy
```

This opens the editor and puts "hello".<br>
After you save changes and close the editor, the results are copied into your clipboard.

You can specify which editor to use via `$STDEDIT`.

```sh
export STDEDIT=vim
echo hello | stdedit | pbcopy
```

The argument is prioritized over the environment variable.

## License

This software is released under the MIT License, see [LICENSE](LICENSE).
