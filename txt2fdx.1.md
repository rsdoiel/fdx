%txt2fdx(1) | 1.0.0 e20f3bd
% R. S. Doiel
% 2023-05-20

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

-license
: display license

-version
: display version

-i
: read input for filename

-o
: write output to filename

-newline
: add a trailing newline

# EXAMPLES

Convert *screenplay.txt* into *screenplay.fdx*.

~~~
	txt2fdx -i screenplay.txt -o screenplay.fdx
~~~

Or alternatively

~~~
    cat screenplay.txt | txt2fdx > screenplay.fdx
~~~


