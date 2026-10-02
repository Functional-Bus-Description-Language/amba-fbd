[![Tests](https://github.com/Functional-Bus-Description-Language/amba-fbd/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/Functional-Bus-Description-Language/amba-fbd/actions?query=master)

# amba-fbd

Functional Bus Description Language compiler backend for Advanced Microcontroller Bus Architecture 5 (AMBA5) specifications.

Supported targets:
- c-sync - C target with synchronous (blocking) interface functions,
- json - json target,
- python - Python target,
- vhdl-apb - VHDL target for APB.

## Installation

### go
```
go install github.com/Functional-Bus-Description-Language/amba-fbd/cmd/amba-fbd@latest
```

Go installation installs to go configured path.

### Manual

```
git clone https://github.com/Functional-Bus-Description-Language/amba-fbd.git
make
make install
```

Manual installation installs to `/usr/local/bin`.
