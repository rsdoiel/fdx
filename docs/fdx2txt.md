% fdx2txt(1) fdx2txt user manual
% R. S. Doiel
% Aug 7, 2022

# NAME

fdx2txt

# SYNOPSIS

fdx2txt [OPTIONS]

# DESCRIPTION

fdx2txt is a command line program that reads an fdx file
and returns plain text


## OPTIONS

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

Cervert *screenplay.fdx* into *screenplay.txt*.

~~~shell
    fdx2txt -i screenplay.fdx -o screenplay.txt
~~~

Or alternatively

~~~shell
    cat screenplay.fdx | fdx2txt > screenplay.txt
~~~


