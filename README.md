# phpunisher
Finds code pieces, that looks like viruses/trojans inside php source code.
Based on great [php-parser](https://github.com/z7zmey/php-parser) library.

# usage
```
~# phpunisher [flags] [/path/to/php-code-dir]
```

# flags
```
-mask string
   scan masks, use ';' as separator (ie: *.php;*.inc) (default "*.php*")

-score float
   minimal score to threat file as suspect (default 0)

-verbose
   show scan details for found suspects

-workers int
   workers count (scan parallelism) (default 2)
```

# scanners

- **array-call** finds function calls from array elements
- **array-operations** notifies if array operations amount is over 20% of all operations
- **bad-func** scans againts 'bad function' list (based on [this article](https://habr.com/en/company/modesco/blog/472092))
- **eval-expr** scans for eval expression
- **single-include** notifies if whole file is single include instruction
- **bad-string** notifies if string literal has more than two escaped symbols

