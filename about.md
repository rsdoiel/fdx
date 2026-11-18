---
title: fdx
abstract: |-
  A Go package and command line programs for working with fdx XML files used by Screen Writing software.
authors:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903


maintainer:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903

repository_code: https://github.com/rsdoiel/fdx
version: 1.0.5
license_url: https://www.gnu.org/licenses/agpl-3.0.txt

programming_language:
  - Go >= 1.27.3

keywords:
  - GitHub
  - text markup
  - screen play

date_released: 2026-11-18
---

About this software
===================

## fdx 1.0.5

A tooling and command line release; the package API is unchanged. fdx2txt and txt2fdx follow the workspace conventions: exit codes 2 (bad command line, including a surplus argument, which used to be ignored), 65 (not an fdx file), 66 (missing input), 73, 74 and 77 instead of 1 for every failure; -h, -l and -v beside the long options; input and output may be given as arguments; a failed run no longer creates or empties the output file. A script that tested for exit status 1 must test for non-zero. The licence is AGPL-3.0-or-later throughout: the source headers still carried the BSD text that LICENSE replaced in v1.0.3. Every exported symbol has a documentation block. Requires fountain v1.0.4. See CHANGES.md.

## Authors

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)




## Maintainers

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)


A Go package and command line programs for working with fdx XML files used by Screen Writing software.

- [License](https://www.gnu.org/licenses/agpl-3.0.txt)
- [Code Repository](https://github.com/rsdoiel/fdx)
  - [Issue Tracker](https://github.com/rsdoiel/fdx/issues)

## Programming languages

- Go >= 1.27.3




## Software Requirements

- Go >= 1.27.3
- github.com/rsdoiel/fountain >= 1.0.4
- gopkg.in/yaml.v3 >= 3.0.1


## Software Suggestions

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3


