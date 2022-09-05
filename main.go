// volca-convert - main.go
// John Simpson <jms1@jms1.net> 2022-08-27
//
// Main program - parses the command line and calls the appropriate read/write
// functions based on what the user is asking for.

package main

import (
    "flag"
    "fmt"
    "os"
    "regexp"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//
// usage

const usage_text = `volca-convert [options] INFILE [OUTFILE]

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

Source: https://github.com/kg4zow/volca-convert

`

func usage() {
    fmt.Print( usage_text )
    os.Exit( 0 )
}

func usage_msg( msg string ) {
    fmt.Print( usage_text )
    fmt.Println( msg )
    os.Exit( 1 )
}

func fail( msg string ) {
    fmt.Println( msg )
    os.Exit( 1 )
}

///////////////////////////////////////////////////////////////////////////////

func main() {
    var infile      string
    var outfile     string

    var in_type     FileType
    var out_type    FileType
    var out_simple  bool

    ////////////////////////////////////////
    // Set up and parse command line options

    var itype string
    var otype string

    flag.StringVar( &itype    , "i" , ""    , "input type" )
    flag.StringVar( &otype    , "o" , ""    , "output type" )
    flag.BoolVar( &out_simple , "s" , false , "simple output" )

    flag.Usage = usage
    flag.Parse()

    ////////////////////////////////////////////////////////////
    // Figure out the input and output file types and names.

    is_json := regexp.MustCompile( "(?i)\\.json$" )
    is_syx  := regexp.MustCompile( "(?i)\\.syx$"  )
    is_csv  := regexp.MustCompile( "(?i)\\.csv$"  )

    infile  = flag.Arg( 0 )
    outfile = flag.Arg( 1 )

    ////////////////////////////////////////
    // Figure out the input file type.
    // - If no '-i' option was used, the program will try to detect it,
    //   based on the filename.
    // - If the filename doesn't match one of the recognized patterns, fail.

    if ( strings.EqualFold( itype , "NONE" ) ) {
        in_type = NONE
        infile  = ""
        outfile = flag.Arg(0)
    } else if ( strings.EqualFold( itype , "SYX" ) ) {
        in_type = SYX
    } else if ( strings.EqualFold( itype , "JSON" ) ) {
        in_type = JSON
    } else if ( strings.EqualFold( itype , "CSV" ) ) {
        in_type = CSV
    } else if ( infile == "" ) {
        usage()
    } else if ( is_json.MatchString( infile ) ) {
        in_type = JSON
    } else if ( is_syx.MatchString( infile ) ) {
        in_type = SYX
    } else if ( is_csv.MatchString( infile ) ) {
        in_type = CSV
    } else {
        usage_msg( "ERROR: unable to tell what kind of input file to read" )
    }

    ////////////////////////////////////////
    // Figure out the output file type.
    // - If we can't tell what kind of file to write, fail.
    // - SYX files are binary, do not write to STDOUT.

    if ( strings.EqualFold( otype , "TEXT" ) ) {
        out_type = TEXT
    } else if ( strings.EqualFold( otype , "TXT" ) ) {
        out_type = TEXT
    } else if ( strings.EqualFold( otype , "CSV" ) ) {
        out_type = CSV
    } else if ( strings.EqualFold( otype , "JSON" ) ) {
        out_type = JSON
    } else if ( strings.EqualFold( otype , "SYX" ) ) {
        out_type = SYX
    } else if ( is_json.MatchString( outfile ) ) {
        out_type = JSON
    } else if ( is_syx.MatchString( outfile ) ) {
        out_type = SYX
    } else if ( is_csv.MatchString( outfile ) ) {
        out_type = CSV
    } else {
        out_type = TEXT
    }

    ////////////////////////////////////////
    // Read/parse input file into memory

    if ( in_type == NONE ) {
        // do nothing
    } else if ( in_type == SYX ) {
        read_syx( infile )
    } else if ( in_type == JSON ) {
        read_json( infile )
    } else if ( in_type == CSV ) {
        read_csv( infile )
    } else {
        usage_msg( "ERROR: requested reader not recognized (bug)" )
    }

    ////////////////////////////////////////
    // Write memory to output file

    if ( out_type == TEXT ) {
        write_text( outfile , !out_simple )
    } else if ( out_type == CSV ) {
        write_csv( outfile , !out_simple )
    } else if ( out_type == JSON ) {
        write_json( outfile , !out_simple )
    } else if ( out_type == SYX ) {
        fmt.Println( "SYX writer not written yet" )
        os.Exit( 1 )
    } else {
        usage_msg( "ERROR: requested writer not recognized (bug)" )
    }
}
