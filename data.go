// volca-convert - data.go
// John Simpson <jms1@jms1.net> 2022-08-27
//
// Global data and type definitions

package main

///////////////////////////////////////////////////////////////////////////////
//
// Type definitions

////////////////////////////////////////
// Input/output file types

type FileType int

const (
    UNSET   FileType = iota
    NONE
    JSON
    SYX
    CSV
    TEXT
)

////////////////////////////////////////
// Voice data in memory

type VData map[string]byte

type Voice struct {
    name    string
    param   VData
}

///////////////////////////////////////////////////////////////////////////////
//
// Global data

var voices = make( []Voice , 0 , 32 )

////////////////////////////////////////
// Field order used by Volca FM/FM2
//
// - CSV and JSON files are written using this order.
// - TEXT file output also uses the same order, but it doesn't use these
//   lists directly.

var opf = []string{
    "EGR1" , "EGR2" , "EGR3" , "EGR4" , "EGL1" , "EGL2" , "EGL3" , "EGL4" ,
    "LSBP" , "LSLD" , "LSRD" , "LSLC" , "LSRC" , "ORS"  , "AMS"  , "KVS"  ,
    "OLVL" , "OSCM" , "FREC" , "FREF" , "DETU" ,
}

var allf = []string{
    "PTR1" , "PTR2" , "PTR3" , "PTR4" , "PTL1" , "PTL2" , "PTL3" , "PTL4" ,
    "FDBK" , "OKS"  , "LFOD" , "LAMD" , "LFOK" , "LFOW" , "MSP"  , "TRSP" ,
}
