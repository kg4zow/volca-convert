// volca-convert - write_syx.go
// John Simpson <jms1@jms1.net> 2022-09-04
//
// Write voices from memory to SYX file

package main

import (
    "fmt"
    "os"
)

///////////////////////////////////////////////////////////////////////////////
//
// Generate SYX data for one voice

func generate_syx155() []byte {
    output := make( []byte , 163 )  // 6 + 155 + 1

    ////////////////////////////////////////
    // Start with header

    copy( output[0:6] , SYX_h1 )

    ////////////////////////////////////////
    // Add voice operator data

    v := voices[0]

    for opn := 0 ; opn < 6 ; opn ++ {
        op_loc := 6 + 21 * ( 5 - opn )
        prefix := fmt.Sprintf( "OP%d." , opn + 1 )

        output[ op_loc +  0 ] = v.param[ prefix + "EGR1" ]
        output[ op_loc +  1 ] = v.param[ prefix + "EGR2" ]
        output[ op_loc +  2 ] = v.param[ prefix + "EGR3" ]
        output[ op_loc +  3 ] = v.param[ prefix + "EGR4" ]
        output[ op_loc +  4 ] = v.param[ prefix + "EGL1" ]
        output[ op_loc +  5 ] = v.param[ prefix + "EGL2" ]
        output[ op_loc +  6 ] = v.param[ prefix + "EGL3" ]
        output[ op_loc +  7 ] = v.param[ prefix + "EGL4" ]
        output[ op_loc +  8 ] = v.param[ prefix + "LSBP" ]
        output[ op_loc +  9 ] = v.param[ prefix + "LSLD" ]
        output[ op_loc + 10 ] = v.param[ prefix + "LSRD" ]
        output[ op_loc + 11 ] = v.param[ prefix + "LSLC" ]
        output[ op_loc + 12 ] = v.param[ prefix + "LSRC" ]
        output[ op_loc + 13 ] = v.param[ prefix + "ORS"  ]
        output[ op_loc + 14 ] = v.param[ prefix + "AMS"  ]
        output[ op_loc + 15 ] = v.param[ prefix + "KVS"  ]
        output[ op_loc + 16 ] = v.param[ prefix + "OLVL" ]
        output[ op_loc + 17 ] = v.param[ prefix + "OSCM" ]
        output[ op_loc + 18 ] = v.param[ prefix + "FREC" ]
        output[ op_loc + 19 ] = v.param[ prefix + "FREF" ]
        output[ op_loc + 20 ] = v.param[ prefix + "DETU" ]
    }

    ////////////////////////////////////////
    // Add voice "ALL" data

    a_loc := 6 + 126

    output[ a_loc +  0 ] = v.param[ "ALL.PTR1" ]
    output[ a_loc +  1 ] = v.param[ "ALL.PTR2" ]
    output[ a_loc +  2 ] = v.param[ "ALL.PTR3" ]
    output[ a_loc +  3 ] = v.param[ "ALL.PTR4" ]
    output[ a_loc +  4 ] = v.param[ "ALL.PTL1" ]
    output[ a_loc +  5 ] = v.param[ "ALL.PTL2" ]
    output[ a_loc +  6 ] = v.param[ "ALL.PTL3" ]
    output[ a_loc +  7 ] = v.param[ "ALL.PTL4" ]
    output[ a_loc +  8 ] = v.param[ "ALGO"     ]
    output[ a_loc +  9 ] = v.param[ "ALL.FDBK" ]
    output[ a_loc + 10 ] = v.param[ "ALL.OKS"  ]
    output[ a_loc + 11 ] = v.param[ "LFOR"     ]
    output[ a_loc + 12 ] = v.param[ "ALL.LFOD" ]
    output[ a_loc + 13 ] = v.param[ "LPMD"     ]
    output[ a_loc + 14 ] = v.param[ "ALL.LAMD" ]
    output[ a_loc + 15 ] = v.param[ "ALL.LFOK" ]
    output[ a_loc + 16 ] = v.param[ "ALL.LFOW" ]
    output[ a_loc + 17 ] = v.param[ "ALL.MSP"  ]
    output[ a_loc + 18 ] = v.param[ "ALL.TRSP" ]

    ////////////////////////////////////////
    // Add voice name.
    // First add spaces in case v.NAME is less than 10 bytes.

    copy( output[ (a_loc+19):(a_loc+29) ] , "          " )
    copy( output[ (a_loc+19):(a_loc+29) ] , v.name       )

    ////////////////////////////////////////
    // Calculate and add checksum byte

    cs := 0
    for n := 0 ; n < 155 ; n ++ {
        cs += int( output[ 6 + n ] )
    }

    output[161] = byte( ( ^cs + 1 ) & 0x7F )

    ////////////////////////////////////////
    // Add SYSEX end of message marker

    output[162] = 0xF7

    ////////////////////////////////////////
    // fin

    return output
}

///////////////////////////////////////////////////////////////////////////////
//
// Generate SYX data for 32 voices

func generate_syx128() []byte {
    output := make( []byte , 4104 ) // 6 + 4096 + 1

    ////////////////////////////////////////
    // Start with header

    copy( output[0:6] , SYX_h32 )

    ////////////////////////////////////////
    // Process voices

    for vn := 0 ; vn < 32 ; vn ++ {
        v_loc := 6 + 128 * vn

        ////////////////////////////////////////
        // Add voice operator data

        v := voices[vn]

        for opn := 0 ; opn < 6 ; opn ++ {
            op_loc := v_loc + 17 * ( 5 - opn )
            prefix := fmt.Sprintf( "OP%d." , opn + 1 )

            output[ op_loc +  0 ] = v.param[ prefix + "EGR1" ]
            output[ op_loc +  1 ] = v.param[ prefix + "EGR2" ]
            output[ op_loc +  2 ] = v.param[ prefix + "EGR3" ]
            output[ op_loc +  3 ] = v.param[ prefix + "EGR4" ]
            output[ op_loc +  4 ] = v.param[ prefix + "EGL1" ]
            output[ op_loc +  5 ] = v.param[ prefix + "EGL2" ]
            output[ op_loc +  6 ] = v.param[ prefix + "EGL3" ]
            output[ op_loc +  7 ] = v.param[ prefix + "EGL4" ]
            output[ op_loc +  8 ] = v.param[ prefix + "LSBP" ]
            output[ op_loc +  9 ] = v.param[ prefix + "LSLD" ]
            output[ op_loc + 10 ] = v.param[ prefix + "LSRD" ]

            XX11 := v.param[ prefix + "XX11" ]
            LSRC := v.param[ prefix + "LSRC" ]
            LSLC := v.param[ prefix + "LSLC" ]
            output[ op_loc + 11 ] = ( XX11 << 4 ) | ( LSRC << 2 ) | LSLC

            DETU := v.param[ prefix + "DETU" ]
            ORS  := v.param[ prefix + "ORS"  ]
            output[ op_loc + 12 ] = ( DETU << 3 ) | ORS

            XX13 := v.param[ prefix + "XX13" ]
            KVS  := v.param[ prefix + "KVS"  ]
            AMS  := v.param[ prefix + "AMS"  ]
            output[ op_loc + 13 ] = ( XX13 << 5 ) | ( KVS << 2 ) | AMS

            output[ op_loc + 14 ] = v.param[ prefix + "OLVL" ]

            XX15 := v.param[ prefix + "XX15" ]
            FREC := v.param[ prefix + "FREC" ]
            OSCM := v.param[ prefix + "OSCM" ]
            output[ op_loc + 15 ] = ( XX15 << 6 ) | ( FREC << 1 ) | OSCM

            output[ op_loc + 16 ] = v.param[ prefix + "FREF" ]
        }

        ////////////////////////////////////////
        // Add voice "ALL" data

        a_loc := v_loc + 102

        output[ a_loc +  0 ] = v.param[ "ALL.PTR1" ]
        output[ a_loc +  1 ] = v.param[ "ALL.PTR2" ]
        output[ a_loc +  2 ] = v.param[ "ALL.PTR3" ]
        output[ a_loc +  3 ] = v.param[ "ALL.PTR4" ]
        output[ a_loc +  4 ] = v.param[ "ALL.PTL1" ]
        output[ a_loc +  5 ] = v.param[ "ALL.PTL2" ]
        output[ a_loc +  6 ] = v.param[ "ALL.PTL3" ]
        output[ a_loc +  7 ] = v.param[ "ALL.PTL4" ]

        XX08 := v.param[ "XX08" ]
        ALGO := v.param[ "ALGO" ]
        output[ a_loc +  8 ] = ( XX08 << 5 ) | ALGO

        XX09 := v.param[ "XX09"     ]
        OKS  := v.param[ "ALL.OKS"  ]
        FDBK := v.param[ "ALL.FDBK" ]
        output[ a_loc +  9 ] = ( XX09 << 4 ) | ( OKS << 3 ) | FDBK

        output[ a_loc + 10 ] = v.param[ "LFOR"     ]
        output[ a_loc + 11 ] = v.param[ "ALL.LFOD" ]
        output[ a_loc + 12 ] = v.param[ "LPMD"     ]
        output[ a_loc + 13 ] = v.param[ "ALL.LAMD" ]

        MSP  := v.param[ "ALL.MSP"  ]
        LFOW := v.param[ "ALL.LFOW" ]
        LFOK := v.param[ "ALL.LFOK" ]
        output[ a_loc + 14 ] = ( MSP << 4 ) | ( LFOW << 1 ) | LFOK

        output[ a_loc + 15 ] = v.param[ "ALL.TRSP" ]

        ////////////////////////////////////////
        // Add voice name.
        // First add spaces in case v.NAME is less than 10 bytes.

        copy( output[ (a_loc+16):(a_loc+26) ] , "          " )
        copy( output[ (a_loc+16):(a_loc+26) ] , v.name       )
    }

    ////////////////////////////////////////
    // Calculate and add checksum byte

    cs := 0
    for n := 0 ; n < 4096 ; n ++ {
        cs += int( output[ n + 6 ] )
    }

    output[4102] = byte( ( ^cs + 1 ) & 0x7F )

    ////////////////////////////////////////
    // Add SYSEX end of message marker

    output[4103] = 0xF7

    ////////////////////////////////////////
    // fin

    return output
}

///////////////////////////////////////////////////////////////////////////////

func write_syx( filename string ) {
    var contents []byte

    ////////////////////////////////////////
    // We can only write SYX files with 1 or 32 voices

    nv := len( voices )
    if ( nv == 1 ) {
        contents = generate_syx155()
    } else if ( nv == 32 ) {
        contents = generate_syx128()
    } else {
        fmt.Printf( "ERROR: cannot write SYX file with %d voices\n" , nv )
        os.Exit( 1 )
    }

    ////////////////////////////////////////
    // Do the deed

    if ( filename == "" ) {
        fmt.Print( contents )
    } else {
        err := os.WriteFile( filename , contents , 0644 )
        if ( err != nil ) {
            fmt.Printf( "ERROR: writing \"%s\": %s\n" , filename , err )
            os.Exit( 1 )
        }
    }
}
