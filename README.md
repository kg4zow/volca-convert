# volca-convert

John Simpson `<jms1@jms1.net>` 2022-08-27

This started as a Perl script which read SYX files and printed the parameter values in a plain text format. Once I had this working, I decided to re-write it in Go, as a way to become more familiar with the language.

I also decided to write multiple input and output functions, to make the tool more useful as a generic "converter" tool.

# Usage

If you run the program with no command line arguments (or with the "`-h`" option), it will print a usage message which explains the available options and what they do. Currently the message looks like this:

```plain
volca-convert [options] INFILE [OUTFILE]

Convert a Volca FM/FM2 (or DX7) "patch" file (a set of FM synthesis parameters
which configure what kind of sound is made) from one format to another.

Input file types: SYX, NONE, (coming soon) JSON, CSV

Output file types: TEXT, CSV, JSON, (coming soon) SYX

-i ___  Specify the type of INFILE. This is needed if INFILE doesn't end
        with '.json', '.syx', or '.csv'.

-o ___  Specify the type of OUTFILE. This may needed if OUTFILE doesn't end
        with '.json', '.syx', or '.csv'. If the program can't tell what kind
        of file to write, it will write TEXT by default.

-s      Generate "simple" output.
        - TEXT  don't include the voice's name in hex.
        - CSV   don't include the header rows.
        - JSON  don't include any extra indentation to make the file easier
                for humans to read/edit.

You can use '-i none' to not read any input file, which is useful if you need
to create a CSV file with just the headers. If you do this, no input filename
is needed, and the first filename on the command line will be used as the
output filename.
```

## Examples

### Show the parameters in a SYX file

Using the default file that comes with [Dexed](https://asb2m10.github.io/dexed/)...

```
$ volca-convert Dexed_01.syx
[Say Again.] ALGO 31  LFOR 35  LPMD  0    NAME 53 61 79 20 41 67 61 69 6E 2E
  OP1
    EGR1 70  EGR2 40  EGR3 49  EGR4 99    EGL1 99  EGL2 92  EGL3  0  EGL4  0
    LSBP  0  LSLD  0  LSRD  0  LSLC  0    LSRC  0  ORS   7  AMS   0  KVS   0
    OLVL  0  OSCM 99  FREC  1  FREF  0    DETU  0
  OP2
    EGR1 25  EGR2 64  EGR3 49  EGR4 99    EGL1 50  EGL2 99  EGL3  0  EGL4  0
    LSBP  0  LSLD  0  LSRD  0  LSLC  0    LSRC  0  ORS   7  AMS   0  KVS   0
    OLVL  0  OSCM 99  FREC  0  FREF  0    DETU  0
  OP3
    EGR1 15  EGR2 64  EGR3 49  EGR4 99    EGL1 44  EGL2 99  EGL3  0  EGL4  0
    LSBP  0  LSLD  0  LSRD  0  LSLC  0    LSRC  0  ORS   7  AMS   0  KVS   0
    OLVL  0  OSCM 99  FREC  2  FREF  0    DETU  0
  OP4
    EGR1 13  EGR2 64  EGR3 49  EGR4 99    EGL1 46  EGL2 99  EGL3  0  EGL4  0
    LSBP  0  LSLD  0  LSRD  0  LSLC  0    LSRC  0  ORS   7  AMS   0  KVS   0
    OLVL  0  OSCM 99  FREC  0  FREF  0    DETU  0
  OP5
    EGR1 10  EGR2 64  EGR3 49  EGR4 99    EGL1 46  EGL2 99  EGL3  0  EGL4  0
    LSBP  0  LSLD  0  LSRD  0  LSLC  0    LSRC  0  ORS   7  AMS   0  KVS   0
    OLVL  0  OSCM 99  FREC  2  FREF  0    DETU  0
  OP6
    EGR1  7  EGR2 64  EGR3 45  EGR4 99    EGL1 45  EGL2 99  EGL3  0  EGL4  0
    LSBP  0  LSLD  0  LSRD  0  LSLC  0    LSRC  0  ORS   7  AMS   0  KVS   0
    OLVL  0  OSCM 99  FREC  0  FREF  0    DETU  0
  ALL
    PTR1 99  PTR2 99  PTR3 99  PTR4 99    PTL1 50  PTL2 50  PTL3 50  PTL4 50
    FDBK  1  OKS   7  LFOD  0  LAMD  0    LFOK  3  LFOW  0  MSP   1  TRSP 24

(not showing the other 31 voices)
```

As you can see, this shows the values of all of the parameters, using the same groups and order that the Volca FM/FM2 use.

Note that the `ALGO`, `LFOR`, and `LPMD` parameters are shown on the same line with the name. This is because these three parameters are controlled directly (using dedicated knobs on the Volca) rather than through the Edit menus.

### Convert a SYX file to JSON

```
$ volca-convert input.syx output.json
```

### Convert a SYX file to CSV.

```
$ volca-convert input.syx output.csv
```

### CSV Header Rows

The CSV files produced by this program use the first *two* rows as headers.

Print *just* the CSV header rows.

```
$ volca-convert -i none -o csv
```

Save the CSV header rows to a file. This can be useful if you're starting a
new file.

```
$ volca-convert -i none new.csv
```
