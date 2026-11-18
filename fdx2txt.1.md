%fdx2txt(1) user manual | version 1.0.5 58deef0
% R. S. Doiel
% 2026-11-18

# NAME

fdx2txt — read an fdx file and write plain text

# SYNOPSIS

fdx2txt [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

fdx2txt reads an fdx (Final Draft XML) file and writes the screenplay as plain text suitable for reading on the console.

# OPTIONS

-h, -help
: display this help page

-l, -license
: display the license

-v, -version
: display the version

-i, -input FILENAME
: read from FILENAME; the default, and -, is standard input

-o, -output FILENAME
: write to FILENAME; the default, and -, is standard output

-nl, -newline
: add a trailing newline (off by default)

-quiet
: suppress error messages; the exit status still says what happened

The input and output may also be given as the first and second argument.

# EXIT STATUS

0
: success

2
: bad command line: an unknown option, a bad option value, a surplus argument

65
: the input is not a valid fdx file

66
: an input file does not exist

73
: the output file cannot be created

74
: reading or writing failed part way

77
: the operating system refused access to a file

# EXAMPLES

Convert *screenplay.fdx* into *screenplay.txt*.

    fdx2txt -i screenplay.fdx -o screenplay.txt

Or alternatively

    cat screenplay.fdx | fdx2txt >screenplay.txt

