// volca-convert - read_json.go
// John Simpson <jms1@jms1.net> 2022-09-04
//
// Read a JSON file into memory.

package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "encoding/json"
)

////////////////////////////////////////
// The JSON "Unmarshal" function needs a data structure matching
// the format of the file itself.

type JOpData struct {
    EGR1    int `json:"EGR1"`
    EGR2    int `json:"EGR2"`
    EGR3    int `json:"EGR3"`
    EGR4    int `json:"EGR4"`
    EGL1    int `json:"EGL1"`
    EGL2    int `json:"EGL2"`
    EGL3    int `json:"EGL3"`
    EGL4    int `json:"EGL4"`
    LSBP    int `json:"LSBP"`
    LSLD    int `json:"LSLD"`
    LSRD    int `json:"LSRD"`
    LSLC    int `json:"LSLC"`
    LSRC    int `json:"LSRC"`
    ORS     int `json:"ORS"`
    AMS     int `json:"AMS"`
    KVS     int `json:"KVS"`
    OLVL    int `json:"OLVL"`
    OSCM    int `json:"OSCM"`
    FREC    int `json:"FREC"`
    FREF    int `json:"FREF"`
    DETU    int `json:"DETU"`
}

type JAllData struct {
    PTR1    int `json:"PTR1"`
    PTR2    int `json:"PTR2"`
    PTR3    int `json:"PTR3"`
    PTR4    int `json:"PTR4"`
    PTL1    int `json:"PTL1"`
    PTL2    int `json:"PTL2"`
    PTL3    int `json:"PTL3"`
    PTL4    int `json:"PTL4"`
    FDBK    int `json:"FDBK"`
    OKS     int `json:"OKS"`
    LFOD    int `json:"LFOD"`
    LAMD    int `json:"LAMD"`
    LFOK    int `json:"LFOK"`
    LFOW    int `json:"LFOW"`
    MSP     int `json:"MSP"`
    TRSP    int `json:"TRSP"`
}

type JVoice struct {
    NAME    string      `json:"NAME"`
    ALGO    int         `json:"ALGO"`
    LFOR    int         `json:"LFOR"`
    LPMD    int         `json:"LPMD"`
    OP1     JOpData     `json:"OP1"`
    OP2     JOpData     `json:"OP2"`
    OP3     JOpData     `json:"OP3"`
    OP4     JOpData     `json:"OP4"`
    OP5     JOpData     `json:"OP5"`
    OP6     JOpData     `json:"OP6"`
    ALL     JAllData    `json:"ALL"`
}

///////////////////////////////////////////////////////////////////////////////
//
// Read a JSON file into memory.

func read_json( filename string ) {

    ////////////////////////////////////////
    // Open the file

    file, err := os.Open( filename )
    if err != nil {
        fmt.Printf( "ERROR: open(\"%s\"): %s\n" , filename , err )
        os.Exit( 1 )
    }
    defer file.Close()

    ////////////////////////////////////////
    // Read the file's contents

    jbytes , err := ioutil.ReadAll( file )
    if err != nil {
        fmt.Printf( "ERROR: reading \"%s\": %s\n" , filename , err )
        os.Exit( 1 )
    }

    ////////////////////////////////////////
    // Parse the JSON

    var jvoices []JVoice
    json.Unmarshal( jbytes , &jvoices )

    ////////////////////////////////////////
    // Process voices from JSON

    for _ , jv := range jvoices {
        var v Voice
        v.param = make( map[string]byte )

        ////////////////////////////////////////
        // Copy the top-level fields

        v.name          =       jv.NAME
        v.param["ALGO"] = byte( jv.ALGO )
        v.param["LFOR"] = byte( jv.LFOR )
        v.param["LPMD"] = byte( jv.LPMD )

        ////////////////////////////////////////
        // Copy the operator value fields

        v.param["OP1.EGR1"] = byte( jv.OP1.EGR1 )
        v.param["OP1.EGR2"] = byte( jv.OP1.EGR2 )
        v.param["OP1.EGR3"] = byte( jv.OP1.EGR3 )
        v.param["OP1.EGR4"] = byte( jv.OP1.EGR4 )
        v.param["OP1.EGL1"] = byte( jv.OP1.EGL1 )
        v.param["OP1.EGL2"] = byte( jv.OP1.EGL2 )
        v.param["OP1.EGL3"] = byte( jv.OP1.EGL3 )
        v.param["OP1.EGL4"] = byte( jv.OP1.EGL4 )
        v.param["OP1.LSBP"] = byte( jv.OP1.LSBP )
        v.param["OP1.LSLD"] = byte( jv.OP1.LSLD )
        v.param["OP1.LSRD"] = byte( jv.OP1.LSRD )
        v.param["OP1.LSLC"] = byte( jv.OP1.LSLC )
        v.param["OP1.LSRC"] = byte( jv.OP1.LSRC )
        v.param["OP1.ORS" ] = byte( jv.OP1.ORS  )
        v.param["OP1.AMS" ] = byte( jv.OP1.AMS  )
        v.param["OP1.KVS" ] = byte( jv.OP1.KVS  )
        v.param["OP1.OLVL"] = byte( jv.OP1.OLVL )
        v.param["OP1.OSCM"] = byte( jv.OP1.OSCM )
        v.param["OP1.FREC"] = byte( jv.OP1.FREC )
        v.param["OP1.FREF"] = byte( jv.OP1.FREF )
        v.param["OP1.DETU"] = byte( jv.OP1.DETU )

        v.param["OP2.EGR1"] = byte( jv.OP2.EGR1 )
        v.param["OP2.EGR2"] = byte( jv.OP2.EGR2 )
        v.param["OP2.EGR3"] = byte( jv.OP2.EGR3 )
        v.param["OP2.EGR4"] = byte( jv.OP2.EGR4 )
        v.param["OP2.EGL1"] = byte( jv.OP2.EGL1 )
        v.param["OP2.EGL2"] = byte( jv.OP2.EGL2 )
        v.param["OP2.EGL3"] = byte( jv.OP2.EGL3 )
        v.param["OP2.EGL4"] = byte( jv.OP2.EGL4 )
        v.param["OP2.LSBP"] = byte( jv.OP2.LSBP )
        v.param["OP2.LSLD"] = byte( jv.OP2.LSLD )
        v.param["OP2.LSRD"] = byte( jv.OP2.LSRD )
        v.param["OP2.LSLC"] = byte( jv.OP2.LSLC )
        v.param["OP2.LSRC"] = byte( jv.OP2.LSRC )
        v.param["OP2.ORS" ] = byte( jv.OP2.ORS  )
        v.param["OP2.AMS" ] = byte( jv.OP2.AMS  )
        v.param["OP2.KVS" ] = byte( jv.OP2.KVS  )
        v.param["OP2.OLVL"] = byte( jv.OP2.OLVL )
        v.param["OP2.OSCM"] = byte( jv.OP2.OSCM )
        v.param["OP2.FREC"] = byte( jv.OP2.FREC )
        v.param["OP2.FREF"] = byte( jv.OP2.FREF )
        v.param["OP2.DETU"] = byte( jv.OP2.DETU )

        v.param["OP3.EGR1"] = byte( jv.OP3.EGR1 )
        v.param["OP3.EGR2"] = byte( jv.OP3.EGR2 )
        v.param["OP3.EGR3"] = byte( jv.OP3.EGR3 )
        v.param["OP3.EGR4"] = byte( jv.OP3.EGR4 )
        v.param["OP3.EGL1"] = byte( jv.OP3.EGL1 )
        v.param["OP3.EGL2"] = byte( jv.OP3.EGL2 )
        v.param["OP3.EGL3"] = byte( jv.OP3.EGL3 )
        v.param["OP3.EGL4"] = byte( jv.OP3.EGL4 )
        v.param["OP3.LSBP"] = byte( jv.OP3.LSBP )
        v.param["OP3.LSLD"] = byte( jv.OP3.LSLD )
        v.param["OP3.LSRD"] = byte( jv.OP3.LSRD )
        v.param["OP3.LSLC"] = byte( jv.OP3.LSLC )
        v.param["OP3.LSRC"] = byte( jv.OP3.LSRC )
        v.param["OP3.ORS" ] = byte( jv.OP3.ORS  )
        v.param["OP3.AMS" ] = byte( jv.OP3.AMS  )
        v.param["OP3.KVS" ] = byte( jv.OP3.KVS  )
        v.param["OP3.OLVL"] = byte( jv.OP3.OLVL )
        v.param["OP3.OSCM"] = byte( jv.OP3.OSCM )
        v.param["OP3.FREC"] = byte( jv.OP3.FREC )
        v.param["OP3.FREF"] = byte( jv.OP3.FREF )
        v.param["OP3.DETU"] = byte( jv.OP3.DETU )

        v.param["OP4.EGR1"] = byte( jv.OP4.EGR1 )
        v.param["OP4.EGR2"] = byte( jv.OP4.EGR2 )
        v.param["OP4.EGR3"] = byte( jv.OP4.EGR3 )
        v.param["OP4.EGR4"] = byte( jv.OP4.EGR4 )
        v.param["OP4.EGL1"] = byte( jv.OP4.EGL1 )
        v.param["OP4.EGL2"] = byte( jv.OP4.EGL2 )
        v.param["OP4.EGL3"] = byte( jv.OP4.EGL3 )
        v.param["OP4.EGL4"] = byte( jv.OP4.EGL4 )
        v.param["OP4.LSBP"] = byte( jv.OP4.LSBP )
        v.param["OP4.LSLD"] = byte( jv.OP4.LSLD )
        v.param["OP4.LSRD"] = byte( jv.OP4.LSRD )
        v.param["OP4.LSLC"] = byte( jv.OP4.LSLC )
        v.param["OP4.LSRC"] = byte( jv.OP4.LSRC )
        v.param["OP4.ORS" ] = byte( jv.OP4.ORS  )
        v.param["OP4.AMS" ] = byte( jv.OP4.AMS  )
        v.param["OP4.KVS" ] = byte( jv.OP4.KVS  )
        v.param["OP4.OLVL"] = byte( jv.OP4.OLVL )
        v.param["OP4.OSCM"] = byte( jv.OP4.OSCM )
        v.param["OP4.FREC"] = byte( jv.OP4.FREC )
        v.param["OP4.FREF"] = byte( jv.OP4.FREF )
        v.param["OP4.DETU"] = byte( jv.OP4.DETU )

        v.param["OP5.EGR1"] = byte( jv.OP5.EGR1 )
        v.param["OP5.EGR2"] = byte( jv.OP5.EGR2 )
        v.param["OP5.EGR3"] = byte( jv.OP5.EGR3 )
        v.param["OP5.EGR4"] = byte( jv.OP5.EGR4 )
        v.param["OP5.EGL1"] = byte( jv.OP5.EGL1 )
        v.param["OP5.EGL2"] = byte( jv.OP5.EGL2 )
        v.param["OP5.EGL3"] = byte( jv.OP5.EGL3 )
        v.param["OP5.EGL4"] = byte( jv.OP5.EGL4 )
        v.param["OP5.LSBP"] = byte( jv.OP5.LSBP )
        v.param["OP5.LSLD"] = byte( jv.OP5.LSLD )
        v.param["OP5.LSRD"] = byte( jv.OP5.LSRD )
        v.param["OP5.LSLC"] = byte( jv.OP5.LSLC )
        v.param["OP5.LSRC"] = byte( jv.OP5.LSRC )
        v.param["OP5.ORS" ] = byte( jv.OP5.ORS  )
        v.param["OP5.AMS" ] = byte( jv.OP5.AMS  )
        v.param["OP5.KVS" ] = byte( jv.OP5.KVS  )
        v.param["OP5.OLVL"] = byte( jv.OP5.OLVL )
        v.param["OP5.OSCM"] = byte( jv.OP5.OSCM )
        v.param["OP5.FREC"] = byte( jv.OP5.FREC )
        v.param["OP5.FREF"] = byte( jv.OP5.FREF )
        v.param["OP5.DETU"] = byte( jv.OP5.DETU )

        v.param["OP6.EGR1"] = byte( jv.OP6.EGR1 )
        v.param["OP6.EGR2"] = byte( jv.OP6.EGR2 )
        v.param["OP6.EGR3"] = byte( jv.OP6.EGR3 )
        v.param["OP6.EGR4"] = byte( jv.OP6.EGR4 )
        v.param["OP6.EGL1"] = byte( jv.OP6.EGL1 )
        v.param["OP6.EGL2"] = byte( jv.OP6.EGL2 )
        v.param["OP6.EGL3"] = byte( jv.OP6.EGL3 )
        v.param["OP6.EGL4"] = byte( jv.OP6.EGL4 )
        v.param["OP6.LSBP"] = byte( jv.OP6.LSBP )
        v.param["OP6.LSLD"] = byte( jv.OP6.LSLD )
        v.param["OP6.LSRD"] = byte( jv.OP6.LSRD )
        v.param["OP6.LSLC"] = byte( jv.OP6.LSLC )
        v.param["OP6.LSRC"] = byte( jv.OP6.LSRC )
        v.param["OP6.ORS" ] = byte( jv.OP6.ORS  )
        v.param["OP6.AMS" ] = byte( jv.OP6.AMS  )
        v.param["OP6.KVS" ] = byte( jv.OP6.KVS  )
        v.param["OP6.OLVL"] = byte( jv.OP6.OLVL )
        v.param["OP6.OSCM"] = byte( jv.OP6.OSCM )
        v.param["OP6.FREC"] = byte( jv.OP6.FREC )
        v.param["OP6.FREF"] = byte( jv.OP6.FREF )
        v.param["OP6.DETU"] = byte( jv.OP6.DETU )

        ////////////////////////////////////////
        // Copy the "ALL" value fields

        v.param["ALL.PTR1"] = byte( jv.ALL.PTR1 )
        v.param["ALL.PTR2"] = byte( jv.ALL.PTR2 )
        v.param["ALL.PTR3"] = byte( jv.ALL.PTR3 )
        v.param["ALL.PTR4"] = byte( jv.ALL.PTR4 )
        v.param["ALL.PTL1"] = byte( jv.ALL.PTL1 )
        v.param["ALL.PTL2"] = byte( jv.ALL.PTL2 )
        v.param["ALL.PTL3"] = byte( jv.ALL.PTL3 )
        v.param["ALL.PTL4"] = byte( jv.ALL.PTL4 )
        v.param["ALL.FDBK"] = byte( jv.ALL.FDBK )
        v.param["ALL.OKS" ] = byte( jv.ALL.OKS  )
        v.param["ALL.LFOD"] = byte( jv.ALL.LFOD )
        v.param["ALL.LAMD"] = byte( jv.ALL.LAMD )
        v.param["ALL.LFOK"] = byte( jv.ALL.LFOK )
        v.param["ALL.LFOW"] = byte( jv.ALL.LFOW )
        v.param["ALL.MSP" ] = byte( jv.ALL.MSP  )
        v.param["ALL.TRSP"] = byte( jv.ALL.TRSP )

        /////////////////////////////////////////
        // Store the finished voice

        voices = append( voices , v )
    }
}
