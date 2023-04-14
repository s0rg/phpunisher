[![Build](https://github.com/s0rg/phpunisher/workflows/ci/badge.svg)](https://github.com/s0rg/phpunisher/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/phpunisher)](https://goreportcard.com/report/github.com/s0rg/phpunisher)
[![Maintainability](https://api.codeclimate.com/v1/badges/a495e449a4b9190b6571/maintainability)](https://codeclimate.com/github/s0rg/phpunisher/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a495e449a4b9190b6571/test_coverage)](https://codeclimate.com/github/s0rg/phpunisher/test_coverage)

[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/s0rg/phpunisher/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/phpunisher)](go.mod)
[![Release](https://img.shields.io/github/v/release/s0rg/phpunisher)](https://github.com/s0rg/phpunisher/releases/latest)
![Downloads](https://img.shields.io/github/downloads/s0rg/phpunisher/total.svg)

# phpunisher

Finds code pieces, that looks like viruses/trojans inside php source code.

Tested on following public malware collections:

- [https://github.com/nikicat/web-malware-collection](https://github.com/nikicat/web-malware-collection)
- [https://github.com/nbs-system/php-malware-finder](https://github.com/nbs-system/php-malware-finder)
- [https://github.com/mnutsch/Computer-Security---Malware](https://github.com/mnutsch/Computer-Security---Malware)
- [https://github.com/sarn1/example-malware-vulnerabilities](https://github.com/sarn1/example-malware-vulnerabilities)
- [https://github.com/AUCyberClub/php-malwares](https://github.com/AUCyberClub/php-malwares)
- [https://github.com/nexylan/PHPAV](https://github.com/nexylan/PHPAV)
- [https://github.com/marcocesarato/PHP-Malware-Collection](https://github.com/marcocesarato/PHP-Malware-Collection)
- [https://github.com/ollyxar/php-malware-detector](https://github.com/ollyxar/php-malware-detector)
- [https://github.com/planet-work/php-malware-scanner](https://github.com/planet-work/php-malware-scanner)
- [https://github.com/bediger4000/php-malware-analysis](https://github.com/bediger4000/php-malware-analysis)
- [https://github.com/Am0rphous/Malware](https://github.com/Am0rphous/Malware)
- [https://github.com/harsxv/malware-bucket](https://github.com/harsxv/malware-bucket)


# features

- powered by great [php-parser](https://github.com/z7zmey/php-parser) library
- selected scanners run in parrallel
- no signatures
- fully customized detection rules


# installation

- [binaries](https://github.com/s0rg/phpunisher/releases) for Linux, macOS and Windows


# usage

```
~# cd /to/your/php/code
~# phpunisher -report                  # to see report
~# phpunisher | xargs -d "\n" -n 1 rm  # to remove suspicios
```

or

```
~# phpunisher -dump-conf > my_rules.yaml
~# $EDITOR my_rules.yaml # edit to suit your needs
~# cd /to/your/php/code
~# phpunisher -conf /path/to/my_rules.yaml -report
```


# flags

```
-conf string
    load scanners config from file
-dump-conf
    dump default scanners config to stdout
-mask string
    scan masks, use ';' as separator (default "*.php*")
-report
    show report for found suspects
-score float
    minimal score to threat file as suspect
-version
    show version
-workers int
    workers count (scan parallelism) (default 2)
```


# scanners

- **array-call** finds function calls from array elements
- **array-ops** notifies if array operations amount is over 20% of all operations
- **escapes** notifies if string literal has more than two escaped symbols
- **evals** scans for eval expression
- **funcs** scans againts 'bad function' list (based on [this article](https://habr.com/en/company/modesco/blog/472092))
- **include** notifies if whole file is single include instruction
- **long-str** notifies if string literal rather long (>64 chars) and does not contains any spaces (encoded blobs)
