%{app_name}(1) | {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} 

# SYNOPSIS

{app_name} [OPTIONS]

# DESCRIPTION

{app_name} is a command line program that reads a plain text file file
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

# EXAMPLES

Convert *screenplay.txt* into *screenplay.fdx*.

~~~
	{app_name} -i screenplay.txt -o screenplay.fdx
~~~

Or alternatively

~~~
    cat screenplay.txt | {app_name} > screenplay.fdx
~~~


