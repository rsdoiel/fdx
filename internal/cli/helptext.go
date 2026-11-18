package cli

/** Fdx2txtHelpText is the fdx2txt(1) page, shown by `fdx2txt -help`. The {app_name},
 * {version}, {release_hash} and {release_date} tokens are replaced by fdx.FmtHelp.
 * The Makefile writes the page to fdx2txt.1.md.
 *
 * Example:
 *   fmt.Println(fdx.FmtHelp(cli.Fdx2txtHelpText, "fdx2txt", fdx.Version, fdx.ReleaseDate, fdx.ReleaseHash))
 */
const Fdx2txtHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — read an fdx file and write plain text

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads an fdx (Final Draft XML) file and writes the screenplay as plain text suitable for reading on the console.

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

    {app_name} -i screenplay.fdx -o screenplay.txt

Or alternatively

    cat screenplay.fdx | {app_name} >screenplay.txt
`

/** Txt2fdxHelpText is the txt2fdx(1) page, shown by `txt2fdx -help`. The {app_name},
 * {version}, {release_hash} and {release_date} tokens are replaced by fdx.FmtHelp.
 * The Makefile writes the page to txt2fdx.1.md.
 *
 * Example:
 *   fmt.Println(fdx.FmtHelp(cli.Txt2fdxHelpText, "txt2fdx", fdx.Version, fdx.ReleaseDate, fdx.ReleaseHash))
 */
const Txt2fdxHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — read a plain text (Fountain) screenplay and write an fdx file

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads a screenplay in plain text, in Fountain markup, and writes an fdx (Final Draft XML) file.
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

    {app_name} -i screenplay.txt -o screenplay.fdx

Or alternatively

    cat screenplay.txt | {app_name} >screenplay.fdx
`
