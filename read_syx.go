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

var SYX_h1  = []byte{ 0xF0 , 0x43 , 0x00 , 0x00 , 0x01 , 0x1B }
var SYX_h32 = []byte{ 0xF0 , 0x43 , 0x00 , 0x09 , 0x20 , 0x00 }

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

    for op := 0 ; op < 6 ; op ++ {
        p := 21*(5-op)
        prefix := fmt.Sprintf( "OP%d." , op + 1 )

        v.param[ prefix + "EGR1" ] = b[p+ 0]
        v.param[ prefix + "EGR2" ] = b[p+ 1]
        v.param[ prefix + "EGR3" ] = b[p+ 2]
        v.param[ prefix + "EGR4" ] = b[p+ 3]
        v.param[ prefix + "EGL1" ] = b[p+ 4]
        v.param[ prefix + "EGL2" ] = b[p+ 5]
        v.param[ prefix + "EGL3" ] = b[p+ 6]
        v.param[ prefix + "EGL4" ] = b[p+ 7]
        v.param[ prefix + "LSBP" ] = b[p+ 8]
        v.param[ prefix + "LSLD" ] = b[p+ 9]
        v.param[ prefix + "LSRD" ] = b[p+10]
        v.param[ prefix + "LSLC" ] = b[p+11]
        v.param[ prefix + "LSRC" ] = b[p+12]
        v.param[ prefix + "ORS"  ] = b[p+13]
        v.param[ prefix + "AMS"  ] = b[p+14]
        v.param[ prefix + "KVS"  ] = b[p+15]
        v.param[ prefix + "OLVL" ] = b[p+16]
        v.param[ prefix + "OSCM" ] = b[p+17]
        v.param[ prefix + "FREC" ] = b[p+18]
        v.param[ prefix + "FREF" ] = b[p+19]
        v.param[ prefix + "DETU" ] = b[p+20]
    }

    ////////////////////////////////////////
    // Read "ALL" parameter blocks

    v.param["ALL.PTR1"] = b[126]
    v.param["ALL.PTR2"] = b[127]
    v.param["ALL.PTR3"] = b[128]
    v.param["ALL.PTR4"] = b[129]
    v.param["ALL.PTL1"] = b[130]
    v.param["ALL.PTL2"] = b[131]
    v.param["ALL.PTL3"] = b[132]
    v.param["ALL.PTL4"] = b[133]
    v.param["ALGO"    ] = b[134]
    v.param["ALL.FDBK"] = b[135]
    v.param["ALL.OKS "] = b[136]
    v.param["LFOR"    ] = b[137]
    v.param["ALL.LFOD"] = b[138]
    v.param["LPMD"    ] = b[139]
    v.param["ALL.LAMD"] = b[140]
    v.param["ALL.LFOK"] = b[141]
    v.param["ALL.LFOW"] = b[142]
    v.param["ALL.MSP "] = b[143]
    v.param["ALL.TRSP"] = b[144]
    v.name = string( b[145:155] )

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

    for op := 0 ; op < 6 ; op ++ {
        p := 17*(5-op)
        prefix := fmt.Sprintf( "OP%d." , op + 1 )

        v.param[ prefix + "EGR1" ] =   b[p+ 0]
        v.param[ prefix + "EGR2" ] =   b[p+ 1]
        v.param[ prefix + "EGR3" ] =   b[p+ 2]
        v.param[ prefix + "EGR4" ] =   b[p+ 3]
        v.param[ prefix + "EGL1" ] =   b[p+ 4]
        v.param[ prefix + "EGL2" ] =   b[p+ 5]
        v.param[ prefix + "EGL3" ] =   b[p+ 6]
        v.param[ prefix + "EGL4" ] =   b[p+ 7]
        v.param[ prefix + "LSBP" ] =   b[p+ 8]
        v.param[ prefix + "LSLD" ] =   b[p+ 9]
        v.param[ prefix + "LSRD" ] =   b[p+10]
        v.param[ prefix + "LSLC" ] = ( b[p+11] & 0b00001100 ) >> 2
        v.param[ prefix + "LSRC" ] = ( b[p+11] & 0b00000011 )
        v.param[ prefix + "ORS"  ] = ( b[p+12] & 0b01111000 ) >> 3
        v.param[ prefix + "AMS"  ] = ( b[p+12] & 0b00000111 )
        v.param[ prefix + "KVS"  ] = ( b[p+13] & 0b00011100 ) >> 2
        v.param[ prefix + "OLVL" ] = ( b[p+13] & 0b00000011 )
        v.param[ prefix + "OSCM" ] =   b[p+14]
        v.param[ prefix + "FREC" ] = ( b[p+15] & 0b00111110 ) >> 1
        v.param[ prefix + "FREF" ] = ( b[p+15] & 0b00000001 )
        v.param[ prefix + "DETU" ] =   b[p+16]
    }

    ////////////////////////////////////////
    // Read voice-global parameter blocks

    v.param["ALL.PTR1"] =   b[102]
    v.param["ALL.PTR2"] =   b[103]
    v.param["ALL.PTR3"] =   b[104]
    v.param["ALL.PTR4"] =   b[105]
    v.param["ALL.PTL1"] =   b[106]
    v.param["ALL.PTL2"] =   b[107]
    v.param["ALL.PTL3"] =   b[108]
    v.param["ALL.PTL4"] =   b[109]
    v.param["ALGO"    ] =   b[110]
    v.param["ALL.FDBK"] = ( b[111] & 0b00001000 ) >> 3
    v.param["ALL.OKS" ] = ( b[111] & 0b00000111 )
    v.param["LFOR"    ] =   b[112]
    v.param["ALL.LFOD"] =   b[113]
    v.param["LPMD"    ] =   b[114]
    v.param["ALL.LAMD"] =   b[115]
    v.param["ALL.LFOK"] = ( b[116] & 0b01110000 ) >> 4
    v.param["ALL.LFOW"] = ( b[116] & 0b00001110 ) >> 1
    v.param["ALL.MSP" ] = ( b[116] & 0b00000001 )
    v.param["ALL.TRSP"] =   b[117]
    v.name = string( b[118:128] )

    ////////////////////////////////////////
    // Done

    return v
}
