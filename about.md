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
version: 1.0.4
license_url: https://www.gnu.org/licenses/agpl-3.0.txt

programming_language:
  - Go >= 1.26.4

keywords:
  - GitHub
  - text markup
  - screen play

date_released: 2026-10-20
---

About this software
===================

## fdx 1.0.4

Scene numbers now survive a Fountain to FDX round trip. A Fountain heading ending #12A# becomes a Scene Heading with the FDX Number attribute set to 12A and the text without the marker; a Number on a Scene Heading is written back to Fountain as #12A#. Numbered headings therefore no longer carry the number in their text in FDX output. Requires fountain v1.0.3; the go directive is now 1.26.4.

## Authors

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)




## Maintainers

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)


A Go package and command line programs for working with fdx XML files used by Screen Writing software.

- [License](https://www.gnu.org/licenses/agpl-3.0.txt)
- [Code Repository](https://github.com/rsdoiel/fdx)
  - [Issue Tracker](https://github.com/rsdoiel/fdx/issues)

## Programming languages

- Go >= 1.26.4




## Software Requirements

- Go >= 1.26.4
- github.com/rsdoiel/fountain >= 1.0.3
- gopkg.in/yaml.v3 >= 3.0.1


## Software Suggestions

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3


