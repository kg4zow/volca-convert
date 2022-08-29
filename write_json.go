// volca-convert - write_json.go
// John Simpson <jms1@jms1.net> 2022-08-27
//
// Write voices from memory to JSON file
//
// Note that I'm generating the output manually in order to ensure that the
// parameters in the output appear in the same order that they do in the
// Volca FM2.

package main

import (
    "fmt"
    "os"
    "strings"
)

///////////////////////////////////////////////////////////////////////////////
//
// Create the JSON-safe version of a name, per RFC 8259 section 7 pp 1
// - double quotes and backslashes must be escaped using backslash
//   (or "reverse solidus", as the RFC calls it)
// - characters outside of 20-7E are replaced with spaces

func json_safe_name( name string ) string {
    var output string

    for _ , c := range name {
        if ( c == '\\' ) {
            output += "\\\\"
        } else if ( c == '"' ) {
            output += "\\\""
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

func generate_json( pretty bool ) string {

    var i_voice string
    var i_vparm string
    var i_oparm string
    var nl      string
    var sep     = ","
    var f_item  = "%s%s:%d"
    var f_name  = "%s%s:\"%s\""
    var f_oph   = "%s\"OP%d\":{%s"
    var f_allh  = "%s\"ALL\":{%s"

    if ( pretty ) {
        i_voice = "  "
        i_vparm = "    "
        i_oparm = "      "
        nl      = "\n"
        sep     = " ,\n"
        f_item  = "%s%-6s : %2d"
        f_name  = "%s%-6s : \"%s\""
        f_oph   = "%s\"OP%d\"  : {%s"
        f_allh  = "%s\"ALL\"  : {%s"
    }

    ////////////////////////////////////////
    // Build objects for each voice

    var vdata  []string

    for _, v := range voices {

        ////////////////////////////////////////
        // Start list of voice parameters

        var vparms []string

        safe_name := json_safe_name( v.name )

        vparms = append( vparms , fmt.Sprintf( f_name , i_vparm , "\"NAME\"" , safe_name ) )
        vparms = append( vparms , fmt.Sprintf( f_item , i_vparm , "\"ALGO\"" , v.param["ALGO"] ) )
        vparms = append( vparms , fmt.Sprintf( f_item , i_vparm , "\"LFOR\"" , v.param["LFOR"] ) )
        vparms = append( vparms , fmt.Sprintf( f_item , i_vparm , "\"LPMD\"" , v.param["LPMD"] ) )

        ////////////////////////////////////////
        // Build objects for each operator

        for op := 0 ; op < 6 ; op ++ {
            ////////////////////////////////////////
            // Build list of parameters

            var pdata []string

            prefix := fmt.Sprintf( "OP%d." , op + 1 )
            for _ , f := range opf {
                qf   := "\"" + f + "\""
                item := fmt.Sprintf( f_item , i_oparm , qf , v.param[ prefix + f ] )
                pdata = append( pdata , item )
            }

            ////////////////////////////////////////
            // Assemble the object for the operator
            // and add it to the list of voice parameters

            otext := fmt.Sprintf( f_oph , i_vparm , op + 1 , nl )
            otext += strings.Join( pdata , sep )
            otext += fmt.Sprintf( "%s%s}" , nl , i_vparm )

            vparms = append( vparms , otext )
        }

        ////////////////////////////////////////
        // Build object for "ALL"

        var aldata []string

        for _ , f := range allf {
            qf := "\"" + f + "\""
            item := fmt.Sprintf( f_item , i_oparm , qf , v.param[ "ALL." + f ] )
            aldata = append( aldata , item )
        }

        ////////////////////////////////////////
        // Assemble the object for "ALL" parameters
        // and add it to the list of voice parameters

        altext := fmt.Sprintf( f_allh , i_vparm , nl )
        altext += strings.Join( aldata , sep )
        altext += fmt.Sprintf( "%s%s}" , nl , i_vparm )

        vparms = append( vparms , altext )

        ////////////////////////////////////////
        // Build object for the voice

        vtext := fmt.Sprintf( "%s{%s" , i_voice , nl )
        vtext += strings.Join( vparms , sep )
        vtext += fmt.Sprintf( "%s%s}" , nl , i_voice ) ;

        ////////////////////////////////////////
        // Add the voice to the list of voices

        vdata = append( vdata , vtext )
    }

    if ( len( vdata ) < 1 ) {
        return "[]\n"
    }

    return fmt.Sprintf( "[%s%s%s]\n" ,
        nl , strings.Join( vdata , sep ) , nl )
}

///////////////////////////////////////////////////////////////////////////////

func write_json( filename string , pretty bool ) {
    text := generate_json( pretty )

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
