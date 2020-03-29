# stdedit
`stdedit` is a tool which edits standard input with your favorite editor.

## Description
Edits standard input then output result to standard out.<br>
It uses an editor that configured to `$STDEDIT`(Default editor is `vim`).<br>
Not supported GUI editor.<br>

## Installation
`go get github.com/kita127/stdedit`

## Usage
`echo hello | stdedit | pbcopy`<br>

Then It opens editor and puts "hello".<br>
Next, you edited ,saved and closed, then the results are copied into your clipboard.<br>

## License
This software is released under the MIT License, see LICENSE.
