% txt2fdx(1) txt2fdx user manual
% R. S. Doiel
% August 7, 2022

# NAME

txt2fdx

# SYNOPSIS

txt2fdx [OPTIONS]

# DESCRIPTION

txt2fdx is a command line program that reads a plain text file file
and returns a fdx file.

# OPTIONS

-help
: display help

-i, -input
: set the input filename

-l, -license
: display license

-nl, -newline
: add a trailing newline

-o, -output
: set the output filename

-quiet
: suppress error messages

-v, -version
: display version


# EXAMPLES

Convert *screenplay.txt* into *screenplay.fdx*.

~~~shell
    txt2fdx -i screenplay.txt -o screenplay.fdx
~~~

Or alternatively

~~~shell
    cat screenplay.txt | txt2fdx > screenplay.fdx
~~~


