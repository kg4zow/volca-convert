// volca-convert - util.go
// John Simpson <jms1@jms1.net> 2022-08-27
//
// Utility functions

package main

import (
    "fmt"
)

///////////////////////////////////////////////////////////////////////////////
//
// Convert a []byte to a string of hex values. Useful for debug messages.

func bytes2hex( b []byte ) string {
    rv := ""

    for i , c := range b {
        if ( i > 0 ) {
            rv += " "
        }
        rv += fmt.Sprintf( "%02X" , c )
    }

    return rv
}

///////////////////////////////////////////////////////////////////////////////
//
// Convert a string to a string of hex values. Useful for debug messages.

func string2hex( s string ) string {
    rv := ""

    for i , c := range []byte( s ) {
        if ( i > 0 ) {
            rv += " "
        }
        rv += fmt.Sprintf( "%02X" , c )
    }

    return rv
}
