// volca-convert - write_csv.go
// John Simpson <jms1@jms1.net> 2022-08-27
//
// Write voices from memory to CSV file

package main

import (
    "fmt"
    "os"
)

///////////////////////////////////////////////////////////////////////////////
//
// Create the CSV-safe version of a name, per RFC 4180 section 2
// - double quotes are doubled
// - characters outside of 20-7E are replaced with spaces

func csv_safe_name( name string ) string {
    var output string

    for _ , c := range name {
        if ( c == '"' ) {
            output += "\"\""
        } else if ( c < 0x20 ) {
            output += " "
        } else if ( c > 0x7E ) {
            output += " "
        } else {
            output += string( c )
        }
    }

    return output
}

///////////////////////////////////////////////////////////////////////////////

func csv_header() string {
    h1 := ",,,"
    h2 := "\"NAME\",\"ALGO\",\"LFOR\",\"LPMD\""

    for op := 0 ; op < 6 ; op ++ {
        prefix := fmt.Sprintf( "OP%d" , op + 1 )
        for _ , f := range opf {
            h1 += fmt.Sprintf( ",\"%s\"" , prefix )
            h2 += fmt.Sprintf( ",\"%s\"" , f )
        }
    }

    for _ , f := range allf {
        h1 += ",\"ALL\""
        h2 += fmt.Sprintf( ",\"%s\"" , f )
    }

    return( h1 + "\n" + h2 + "\n" )
}

///////////////////////////////////////////////////////////////////////////////

func generate_csv( with_header bool ) string {
    var output string

    ////////////////////////////////////////
    // If a header row was requested, start with that

    if ( with_header ) {
        output += csv_header()
    }

    ////////////////////////////////////////
    // Generate a row for each voice

    for _, v := range voices {

        safe_name := csv_safe_name( v.name )

        output += fmt.Sprintf( "\"%s\",%d,%d,%d" ,
            safe_name , v.param["ALGO"] , v.param["LFOR"] , v.param["LPMD"] )

        for op := 0 ; op < 6 ; op ++ {
            prefix := fmt.Sprintf( "OP%d." , op + 1 )
            for _ , f := range opf {
                output += fmt.Sprintf( ",%d" , v.param[ prefix + f ] )
            }
        }

        for _ , f := range allf {
            output += fmt.Sprintf( ",%d" , v.param[ "ALL." + f ] )
        }

        output += "\n"
    }

    return output
}

///////////////////////////////////////////////////////////////////////////////

func write_csv( filename string , with_header bool ) {
    text := generate_csv( with_header )

    if ( filename == "" ) {
        fmt.Print( text )
    } else {
        err := os.WriteFile( filename , []byte( text ) , 0644 )
        if ( err != nil ) {
            fmt.Printf( "ERROR: writing \"%s\": %s\n" , filename , err )
            os.Exit( 1 )
        }
    }
}
