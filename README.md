# phpunisher
Finds code pieces, that looks like viruses/trojans inside php source code.

# usage
```
~# phpunisher [flags] [/path/to/php-code-dir]
```

# flags
```
-m string
   scan masks, use ';' as separator (ie: *.php;*.inc) (default "*.php*")

-s float
   minimal score to threat file as suspect (default 0)
  
-v 
   show scan details for found suspects
  
-w int
   workers count (scan parallelism) (default 2)
```

# scanners

- **array-call** finds function calls from array elements
- **array-operations** notifies if array operations amount is over 20% of all operations
- **bad-func** scans for a 'bad functions' (like str_rot13 and base64_decode)
- **eval-expr** scans for eval expression
- **single-include** notifies if whole file is single include instruction
- **bad-string** notifies if string literal has more than two escaped symbols

