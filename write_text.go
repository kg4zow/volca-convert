// volca-convert - read_syx.go
// John Simpson <jms1@jms1.net> 2022-08-27
//
// Write voices from memory to TEXT file

package main

import (
    "fmt"
    "os"
)

///////////////////////////////////////////////////////////////////////////////

func generate_text( extras bool ) string {
    var output string

    for i , v := range voices {
        if ( i > 0 ) {
            output += "\n"
        }

        output += fmt.Sprintf( "%-12s ALGO %2d  LFOR %2d  LPMD %2d" ,
            ( "[" + v.name + "]" ) , v.param["ALGO"] , v.param["LFOR"] , v.param["LPMD"] )

        if ( extras ) {
            output += fmt.Sprintf( "    NAME %s\n" , string2hex( v.name ) )
        } else {
            output += "\n"
        }

        for op := 0 ; op < 6 ; op ++ {
            prefix := fmt.Sprintf( "OP%d" , op + 1 )
            output += fmt.Sprintf( "  %s\n" , prefix )

            output += fmt.Sprintf( "    EGR1 %2d  EGR2 %2d  EGR3 %2d  EGR4 %2d" ,
                v.param[ prefix + ".EGR1" ] ,
                v.param[ prefix + ".EGR2" ] ,
                v.param[ prefix + ".EGR3" ] ,
                v.param[ prefix + ".EGR4" ] )
            output += fmt.Sprintf( "    EGL1 %2d  EGL2 %2d  EGL3 %2d  EGL4 %2d\n" ,
                v.param[ prefix + ".EGL1" ] ,
                v.param[ prefix + ".EGL2" ] ,
                v.param[ prefix + ".EGL3" ] ,
                v.param[ prefix + ".EGL4" ] )

            output += fmt.Sprintf( "    LSBP %2d  LSLD %2d  LSRD %2d  LSLC %2d" ,
                v.param[ prefix + ".LSBP" ] ,
                v.param[ prefix + ".LSLD" ] ,
                v.param[ prefix + ".LSRD" ] ,
                v.param[ prefix + ".LSLC" ] )
            output += fmt.Sprintf( "    LSRC %2d  ORS  %2d  AMS  %2d  KVS  %2d\n" ,
                v.param[ prefix + ".LSRC" ] ,
                v.param[ prefix + ".ORS"  ] ,
                v.param[ prefix + ".AMS"  ] ,
                v.param[ prefix + ".KVS"  ] )

            output += fmt.Sprintf( "    OLVL %2d  OSCM %2d  FREC %2d  FREF %2d" ,
                v.param[ prefix + ".OLVL" ] ,
                v.param[ prefix + ".OSCM" ] ,
                v.param[ prefix + ".FREC" ] ,
                v.param[ prefix + ".FREF" ] )
            output += fmt.Sprintf( "    DETU %2d\n" ,
                v.param[ prefix + ".DETU" ] )
        }

        output += "  ALL\n"

        output += fmt.Sprintf( "    PTR1 %2d  PTR2 %2d  PTR3 %2d  PTR4 %2d" ,
            v.param[ "ALL.PTR1" ] ,
            v.param[ "ALL.PTR2" ] ,
            v.param[ "ALL.PTR3" ] ,
            v.param[ "ALL.PTR4" ] )
        output += fmt.Sprintf( "    PTL1 %2d  PTL2 %2d  PTL3 %2d  PTL4 %2d\n" ,
            v.param[ "ALL.PTL1" ] ,
            v.param[ "ALL.PTL2" ] ,
            v.param[ "ALL.PTL3" ] ,
            v.param[ "ALL.PTL4" ] )

        output += fmt.Sprintf( "    FDBK %2d  OKS  %2d  LFOD %2d  LAMD %2d" ,
            v.param[ "ALL.FDBK" ] ,
            v.param[ "ALL.OKS"  ] ,
            v.param[ "ALL.LFOD" ] ,
            v.param[ "ALL.LAMD" ] )
        output += fmt.Sprintf( "    LFOK %2d  LFOW %2d  MSP  %2d  TRSP %2d\n" ,
            v.param[ "ALL.LFOK" ] ,
            v.param[ "ALL.LFOW" ] ,
            v.param[ "ALL.MSP"  ] ,
            v.param[ "ALL.TRSP" ] )
    }

    return output
}

///////////////////////////////////////////////////////////////////////////////

func write_text( filename string , extras bool ) {

    ////////////////////////////////////////
    // Generate the output text

    text := generate_text( extras )

    ////////////////////////////////////////
    // Write the output

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
