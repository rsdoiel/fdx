%txt2fdx(1) user manual | version 1.0.4 2410edd
% R. S. Doiel
% 2026-10-20

# NAME

txt2fdx — read a plain text (Fountain) screenplay and write an fdx file

# SYNOPSIS

txt2fdx [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

txt2fdx reads a screenplay in plain text, in Fountain markup, and writes an fdx (Final Draft XML) file.
A scene heading ending in a number between hash marks, such as #12A#, becomes a Scene Heading with that number.

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
: add a trailing newline (the default; -newline=false leaves it off)

-quiet
: suppress error messages; the exit status still says what happened

The input and output may also be given as the first and second argument.

# EXIT STATUS

0
: success

2
: bad command line: an unknown option, a bad option value, a surplus argument

65
: the input cannot be parsed as a Fountain screenplay

66
: an input file does not exist

70
: the fdx file could not be produced (a bug)

73
: the output file cannot be created

74
: reading or writing failed part way

77
: the operating system refused access to a file

# EXAMPLES

Convert *screenplay.txt* into *screenplay.fdx*.

    txt2fdx -i screenplay.txt -o screenplay.fdx

Or alternatively

    cat screenplay.txt | txt2fdx >screenplay.fdx

