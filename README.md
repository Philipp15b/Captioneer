# Captioneer

With Captioneer you can easily create your own [Closed Captions for Source-based Games](https://developer.valvesoftware.com/wiki/Closed_Captions), like for Team Fortress 2.

## Prerequisites

Make sure you have [installed the Source SDK](https://developer.valvesoftware.com/wiki/SDK_Installation)!

## Usage

### 1. Configuration

Captioneer is very easy to use. Just edit the `captions.txt` file and add captions in the following format:

    name: value

with one caption per line.

You can also use variables, to define them just add a `$` before the name, like so:

    $my.variable: my value

To use variables in your captions, write

    $(my.variable)

and the variable value will be inserted.

All lines beginning with a `//` or `#` will be treated as comments and ignored.

You can find an example already in [`captions.txt`](https://github.com/Philipp15b/Captioneer/blob/master/captions.txt).

### 2. Compiling

> `captioneer filename [--language="english"]`

Run `captioneer` with the filename input path and optionally the language on the command line.
It will then create a `closecaption_language.dat` in the current directory.
