
# OUTPUT

Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.

## Instructions

- Your project must be written in Go.
- The code must respect the good practices.
- You must follow the same instructions as in the first subject while writing the result into a file.

The file must be named by using the flag --output=<fileName.txt>, in which --output is the flag and <fileName.txt> is the file name which will contain the output.

   - The flag must have exactly the same format as above, any other formats must return the following usage message:
```javascript
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard
```

## Usage/Examples

```javascript

$ go run . --output=banner.txt "hello" standard
$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
$ go run . --output=banner.txt "Hello There!" shadow
$ cat -e banner.txt
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
$


```



## Allowed packages

   Only the standard Go packages are allowed

