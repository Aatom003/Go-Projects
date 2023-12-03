
---

# Domain Checker

This tool is a simple domain checker written in Go (Golang) that performs checks for MX, SPF, and DMARC records of given domains.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Overview

The `main.go` file contains a Go program that checks the existence of MX, SPF, and DMARC records for a given domain. It uses Go's `net` package to perform DNS lookups for MX, TXT (SPF), and DMARC records.

## Prerequisites

To run this tool, you need:

- Go (Golang) installed on your system
- Access to the internet for performing DNS lookups

## Installation

To install the tool, follow these steps:

1. Clone this repository.

   ```bash
   git clone https://github.com/Aatom003/emailVerifier.git
   ```

2. Navigate to the directory containing the `main.go` file.

## Usage

To use the tool:

1. Open a terminal.
2. Navigate to the directory containing the `main.go` file.
3. Run the program using the following command:
   ```bash
   go run main.go
   ```
4. Enter the domains you want to check when prompted.

## Example

Suppose you want to check the MX, SPF, and DMARC records for domains like `example.com`, `google.com`, and `domain.com`. You can input these domains when prompted, and the program will perform the necessary checks and display the results in the terminal.

Example output:

```
Domain, hasMx, hasSPF, sprRecord, hasDMARC, dmarcRecord
example.com,true,true,v=spf1 include:_spf.example.com ~all,false,
google.com,true,true,v=spf1 include:_spf.google.com ~all,true,v=DMARC1; p=reject; rua=mailto:mailauth-reports@google.com
domain.com,true,false,,false,
```

## Contributing

Contributions are welcome! If you want to improve this tool or fix any issues, feel free to open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
