%fdx2txt(1) | version 1.0.3 23c5825
% R. S. Doiel
% 2024-07-09 

# NAME

fdx2txt

# SYNOPSIS

fdx2txt [OPTIONS]

# DESCRIPTION

fdx2txt is a command line program that reads an fdx file
and returns plain text

# OPTIONS

-help
: display help

-license
: display license

-version
: display version


-i
: read input from file

-o
: write output to file

-newline
: add a trailing newline 

# EXAMPLES

Convert *screenplay.fdx* into *screenplay.txt*.

~~~
    fdx2txt -i screenplay.fdx -o screenplay.txt
~~~

Or alternatively

~~~
    cat screenplay.fdx | fdx2txt > screenplay.txt
~~~


