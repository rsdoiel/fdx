
Installation
------------

fdx is highly experimental. At this time you must install it from source code. It is probably broken and certainly will contain many bugs.


Install from source
-------------------

## Requirements

- Golang >= 1.20
- Pandoc >= 3
- GNU Make
- Git

## Steps

1. Clone the Git repository for the project
2. change directory into the cloned project
3. Run `make`, `make test` and `make install`

Here's what that looks like for me.

~~~
git clone https://github.com/rsdoiel/fdx src/github.com/rsdoiel/fdx
cd src/github.com/rsdoiel/fdx
make
make test
make install
~~~

By default it will install the programs in `$HOME/bin`. `$HOME/bin` needs
to be included in your `PATH`. E.g.

~~~
export PATH="$HOME/bin:$PATH"
~~~

Can be added to your `.profile`, `.bashrc` or `.zshrc` file depending on your system's shell.


