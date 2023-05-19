%{app_name}(1) | version {version} {release_hash}
% R. S. Doiel
% {release_date} 

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS]

# DESCRIPTION

{app_name} is a command line program that reads an fdx file
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
    {app_name} -i screenplay.fdx -o screenplay.txt
~~~

Or alternatively

~~~
    cat screenplay.fdx | fdx2txt > screenplay.txt
~~~


