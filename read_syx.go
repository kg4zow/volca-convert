// volca-convert - read_syx.go
// John Simpson <jms1@jms1.net> 2022-08-27
//
// Read a SYX file into memory.

package main

import (
    "bytes"
    "fmt"
    "os"
)

///////////////////////////////////////////////////////////////////////////////
//
// Read a SYX file.
//
// Note that there are two kinds of SYX files, one containing a single voice
// and one containing 32 voices (aka a "cartridge", since the cartridges on
// the DX7 held 32 voices).
//
// - In files containing a single voice, every parameter within the voice
//   uses a full byte, even if it only needs one bit. This means that the
//   parameters use a total of 155 bytes.
// - In files containing 32 voices, some of the parameters "share" bytes
//   with others, so the final size of each voice is 128 bytes.
// - Voice names can use any ASCII character, however the Volca FM/FM2 are
//   only able to _show_ certain characters.

func read_syx( filename string ) {
    var buf = make( []byte , 8192 )

    ////////////////////////////////////////
    // Open the file

    file, err := os.Open( filename )
    if err != nil {
        fmt.Printf( "ERROR: open(\"%s\"): %s\n" , filename , err )
        os.Exit( 1 )
    }
    defer file.Close()

    ////////////////////////////////////////
    // The SYX files we're dealing with aren't supposed to be larger than
    // about 4K, so it should be safe to read it all into memory at once.

    bytes_read, err := file.Read( buf )
    if ( bytes_read < 6 ) {
        fmt.Printf( "ERROR: \"%s\" reading header: bytes_read=%d err=%s\n" ,
            filename , bytes_read , err )
        os.Exit( 1 )
    }

    ////////////////////////////////////////
    // Examine the contents, call the correct parser

    if ( bytes.Compare( buf[0:6] , SYX_h1 ) == 0 ) {
        v := parse_syx155( buf[6:161] )
        voices = append( voices , v )
    } else if ( bytes.Compare( buf[0:6] , SYX_h32 ) == 0 ) {
        for n := 0 ; n < 32 ; n++ {
            a := 128 * n + 6
            b := a + 128
            v := parse_syx128( buf[a:b] )
            voices = append( voices , v )
        }
    } else {
        fmt.Printf( "ERROR: \"%s\" does not have a recognized header\n" ,
            filename )

        fmt.Printf( "  file  = '%s'\n" , bytes2hex( buf[0:6] ) ) ;
        fmt.Printf( "  SYX1  = '%s'\n" , bytes2hex( SYX_h1   ) ) ;
        fmt.Printf( "  SYX32 = '%s'\n" , bytes2hex( SYX_h32  ) ) ;

        os.Exit( 1 )
    }
}

///////////////////////////////////////////////////////////////////////////////

func parse_syx155( b []byte ) Voice {
    var v Voice
    v.param = make( map[string]byte )

    if len( b ) < 155 {
        fmt.Printf( "ERROR: parse_syx155(): input (%d bytes) smaller than 155 bytes" ,
            len( b ) )
        os.Exit( 1 )
    }

    ////////////////////////////////////////
    // Read operator parameter blocks

    for opn := 0 ; opn < 6 ; opn ++ {
        op_loc := 21 * ( 5 - opn )
        prefix := fmt.Sprintf( "OP%d." , opn + 1 )

        v.param[ prefix + "EGR1" ] = b[ op_loc +  0 ]
        v.param[ prefix + "EGR2" ] = b[ op_loc +  1 ]
        v.param[ prefix + "EGR3" ] = b[ op_loc +  2 ]
        v.param[ prefix + "EGR4" ] = b[ op_loc +  3 ]
        v.param[ prefix + "EGL1" ] = b[ op_loc +  4 ]
        v.param[ prefix + "EGL2" ] = b[ op_loc +  5 ]
        v.param[ prefix + "EGL3" ] = b[ op_loc +  6 ]
        v.param[ prefix + "EGL4" ] = b[ op_loc +  7 ]
        v.param[ prefix + "LSBP" ] = b[ op_loc +  8 ]
        v.param[ prefix + "LSLD" ] = b[ op_loc +  9 ]
        v.param[ prefix + "LSRD" ] = b[ op_loc + 10 ]
        v.param[ prefix + "LSLC" ] = b[ op_loc + 11 ]
        v.param[ prefix + "LSRC" ] = b[ op_loc + 12 ]
        v.param[ prefix + "ORS"  ] = b[ op_loc + 13 ]
        v.param[ prefix + "AMS"  ] = b[ op_loc + 14 ]
        v.param[ prefix + "KVS"  ] = b[ op_loc + 15 ]
        v.param[ prefix + "OLVL" ] = b[ op_loc + 16 ]
        v.param[ prefix + "OSCM" ] = b[ op_loc + 17 ]
        v.param[ prefix + "FREC" ] = b[ op_loc + 18 ]
        v.param[ prefix + "FREF" ] = b[ op_loc + 19 ]
        v.param[ prefix + "DETU" ] = b[ op_loc + 20 ]
    }

    ////////////////////////////////////////
    // Read "ALL" parameter blocks

    a_loc := 126

    v.param[ "ALL.PTR1" ] = b[ a_loc +  0 ]
    v.param[ "ALL.PTR2" ] = b[ a_loc +  1 ]
    v.param[ "ALL.PTR3" ] = b[ a_loc +  2 ]
    v.param[ "ALL.PTR4" ] = b[ a_loc +  3 ]
    v.param[ "ALL.PTL1" ] = b[ a_loc +  4 ]
    v.param[ "ALL.PTL2" ] = b[ a_loc +  5 ]
    v.param[ "ALL.PTL3" ] = b[ a_loc +  6 ]
    v.param[ "ALL.PTL4" ] = b[ a_loc +  7 ]
    v.param[ "ALGO"     ] = b[ a_loc +  8 ]
    v.param[ "ALL.FDBK" ] = b[ a_loc +  9 ]
    v.param[ "ALL.OKS"  ] = b[ a_loc + 10 ]
    v.param[ "LFOR"     ] = b[ a_loc + 11 ]
    v.param[ "ALL.LFOD" ] = b[ a_loc + 12 ]
    v.param[ "LPMD"     ] = b[ a_loc + 13 ]
    v.param[ "ALL.LAMD" ] = b[ a_loc + 14 ]
    v.param[ "ALL.LFOK" ] = b[ a_loc + 15 ]
    v.param[ "ALL.LFOW" ] = b[ a_loc + 16 ]
    v.param[ "ALL.MSP"  ] = b[ a_loc + 17 ]
    v.param[ "ALL.TRSP" ] = b[ a_loc + 18 ]
    v.name = string( b[ (a_loc+19):(a_loc+29) ] )

    ////////////////////////////////////////
    // Done

    return v
}

///////////////////////////////////////////////////////////////////////////////

func parse_syx128( b []byte ) Voice {
    var v Voice
    v.param = make( map[string]byte )

    if len( b ) < 128 {
        fmt.Printf( "ERROR: parse_syx128(): input (%d bytes) smaller than 128 bytes\n" ,
            len( b ) )
        os.Exit( 1 )
    }

    ////////////////////////////////////////
    // Read operator parameter blocks

    for opn := 0 ; opn < 6 ; opn ++ {
        op_loc := 17*(5-opn)
        prefix := fmt.Sprintf( "OP%d." , opn + 1 )

        v.param[ prefix + "EGR1" ] =   b[ op_loc +  0 ]
        v.param[ prefix + "EGR2" ] =   b[ op_loc +  1 ]
        v.param[ prefix + "EGR3" ] =   b[ op_loc +  2 ]
        v.param[ prefix + "EGR4" ] =   b[ op_loc +  3 ]
        v.param[ prefix + "EGL1" ] =   b[ op_loc +  4 ]
        v.param[ prefix + "EGL2" ] =   b[ op_loc +  5 ]
        v.param[ prefix + "EGL3" ] =   b[ op_loc +  6 ]
        v.param[ prefix + "EGL4" ] =   b[ op_loc +  7 ]
        v.param[ prefix + "LSBP" ] =   b[ op_loc +  8 ]
        v.param[ prefix + "LSLD" ] =   b[ op_loc +  9 ]
        v.param[ prefix + "LSRD" ] =   b[ op_loc + 10 ]

        v.param[ prefix + "XX11" ] = ( b[ op_loc + 11 ] & 0b01110000 ) >> 4
        v.param[ prefix + "LSRC" ] = ( b[ op_loc + 11 ] & 0b00001100 ) >> 2
        v.param[ prefix + "LSLC" ] = ( b[ op_loc + 11 ] & 0b00000011 )

        v.param[ prefix + "DETU" ] = ( b[ op_loc + 12 ] & 0b01111000 ) >> 3
        v.param[ prefix + "ORS"  ] = ( b[ op_loc + 12 ] & 0b00000111 )

        v.param[ prefix + "XX13" ] = ( b[ op_loc + 13 ] & 0b01100000 ) >> 5
        v.param[ prefix + "KVS"  ] = ( b[ op_loc + 13 ] & 0b00011100 ) >> 2
        v.param[ prefix + "AMS"  ] = ( b[ op_loc + 13 ] & 0b00000011 )

        v.param[ prefix + "OLVL" ] =   b[ op_loc + 14 ]

        v.param[ prefix + "XX15" ] = ( b[ op_loc + 15 ] & 0b01000000 ) >> 6
        v.param[ prefix + "FREC" ] = ( b[ op_loc + 15 ] & 0b00111110 ) >> 1
        v.param[ prefix + "OSCM" ] = ( b[ op_loc + 15 ] & 0b00000001 )

        v.param[ prefix + "FREF" ] =   b[ op_loc + 16 ]
    }

    ////////////////////////////////////////
    // Read voice-global parameter blocks

    a_loc := 102

    v.param[ "ALL.PTR1" ] =   b[ a_loc +  0 ]
    v.param[ "ALL.PTR2" ] =   b[ a_loc +  1 ]
    v.param[ "ALL.PTR3" ] =   b[ a_loc +  2 ]
    v.param[ "ALL.PTR4" ] =   b[ a_loc +  3 ]
    v.param[ "ALL.PTL1" ] =   b[ a_loc +  4 ]
    v.param[ "ALL.PTL2" ] =   b[ a_loc +  5 ]
    v.param[ "ALL.PTL3" ] =   b[ a_loc +  6 ]
    v.param[ "ALL.PTL4" ] =   b[ a_loc +  7 ]

    v.param[ "XX08"     ] = ( b[ a_loc +  8 ] & 0b01100000 ) >> 5
    v.param[ "ALGO"     ] =   b[ a_loc +  8 ] & 0b00011111

    v.param[ "XX09"     ] = ( b[ a_loc +  9 ] & 0b01110000 ) >> 4
    v.param[ "ALL.OKS"  ] = ( b[ a_loc +  9 ] & 0b00001000 ) >> 3
    v.param[ "ALL.FDBK" ] = ( b[ a_loc +  9 ] & 0b00000111 )

    v.param[ "LFOR"     ] =   b[ a_loc + 10 ]
    v.param[ "ALL.LFOD" ] =   b[ a_loc + 11 ]
    v.param[ "LPMD"     ] =   b[ a_loc + 12 ]
    v.param[ "ALL.LAMD" ] =   b[ a_loc + 13 ]

    v.param[ "ALL.MSP"  ] = ( b[ a_loc + 14 ] & 0b01110000 ) >> 4
    v.param[ "ALL.LFOW" ] = ( b[ a_loc + 14 ] & 0b00001110 ) >> 1
    v.param[ "ALL.LFOK" ] = ( b[ a_loc + 14 ] & 0b00000001 )

    v.param[ "ALL.TRSP" ] =   b[ a_loc + 15 ]
    v.name = string( b[ (a_loc+16):(a_loc+26) ] )

    ////////////////////////////////////////
    // Done

    return v
}
