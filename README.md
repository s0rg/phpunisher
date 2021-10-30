[![Build](https://github.com/s0rg/phpunisher/workflows/ci/badge.svg)](https://github.com/s0rg/phpunisher/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/phpunisher)](https://goreportcard.com/report/github.com/s0rg/phpunisher)
[![Maintainability](https://api.codeclimate.com/v1/badges/a495e449a4b9190b6571/maintainability)](https://codeclimate.com/github/s0rg/phpunisher/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a495e449a4b9190b6571/test_coverage)](https://codeclimate.com/github/s0rg/phpunisher/test_coverage)
[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/s0rg/phpunisher/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/phpunisher)](go.mod)
[![Release](https://img.shields.io/github/v/release/s0rg/phpunisher)](https://github.com/s0rg/phpunisher/releases/latest)

# phpunisher

Finds code pieces, that looks like viruses/trojans inside php source code.
Based on great [php-parser](https://github.com/z7zmey/php-parser) library.

# usage
```
~# cd /to/your/php/code
~# phpunisher -report .                  # to see report
~# phpunisher . | xargs -d "\n" -n 1 rm  # to remove suspicios
```

# flags
```
-mask string
  	scan masks, use ';' as separator (default "*.php*")
-report
  	show report for found suspects
-score float
  	minimal score to threat file as suspect
-version
  	show version
-workers int
  	workers count (scan parallelism)
```

# scanners

- **array-call** finds function calls from array elements
- **array-ops** notifies if array operations amount is over 20% of all operations
- **bad-php** invalid/malformed php code
- **escapes** notifies if string literal has more than two escaped symbols
- **evals** scans for eval expression
- **funcs** scans againts 'bad function' list (based on [this article](https://habr.com/en/company/modesco/blog/472092))
- **long-str** notifies if string literal rather long (>64 chars) and does not contains any spaces (encoded blobs)
- **single-include** notifies if whole file is single include instruction
