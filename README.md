# Captioneer

With Captioneer you can compile your own caption files which can be used to write scripts with a GUI output.

## Prerequisites

Make sure you have installed the Source SDK!

## Usage

### Configuration

Captioneer is very easy to use. Just edit the `captions.txt` file and add captions in the following format:

    name: value

with one caption per line.

You can also use variables, to define them just add a `$` before the name, like so:

    $my.variable: my value

To use variables in your captions, write

    $(my.variable)

and the variable value will be inserted.

All lines beginning with a `//` or `#` will be treated as comments and ignored.

### Compiling

> `captioneer filename [--language="english"]`

Run `captioneer` with the filename input path and optionally the language. It will then create a `closecaption_language.dat` in the current working directory.
