package main

// AkamaiPayload is a Golang representation of the Cloudmonitor JSON datastructure
type AkamaiPayload struct {
	// Content Provider ID
	CP string `json:"cp"`
	// Defines the format of the payload (?)
	Format  string `json:"format"`
	ID      string `json:"id"`
	Start   string `json:"start"`
	Type    string `json:"type"`
	Version string `json:"version"`

	Geo     GeoStruct     `json:"geo"`
	Message MessageStruct `json:"message"`
	NetPerf NetPerfStruct `json:"netPerf"`
	Network NetworkStruct `json:"network"`
	ReqHdr  ReqHdrStruct  `json:"reqHdr"`
	RespHdr RespHdrStruct `json:"respHdr"`
}

// GeoStruct is used for storing the JSON subfields
type GeoStruct struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Lat     string `json:"lat"`
	Long    string `json:"long"`
	Region  string `json:"region"`
}

// MessageStruct is used for storing the JSON subfields
type MessageStruct struct {
	UA        string `json:"UA"`
	Bytes     string `json:"bytes"`
	CliIP     string `json:"cliIP"`
	FwdHost   string `json:"fwdHost"`
	Proto     string `json:"proto"`
	ProtoVer  string `json:"protoVer"`
	ReqHost   string `json:"reqHost"`
	ReqMethod string `json:"reqMethod"`
	ReqPath   string `json:"reqPath"`
	ReqPort   string `json:"reqPort"`
	RespCT    string `json:"respCT"`
	RespLen   string `json:"respLen"`
	Status    string `json:"status"`
}

// NetPerfStruct is used for storing the JSON subfields
type NetPerfStruct struct {
	Asnum        string `json:"asnum"`
	CacheStatus  string `json:"cacheStatus"`
	DownloadTime string `json:"downloadTime"`
	EdgeIP       string `json:"edgeIP"`
	FirstByte    string `json:"firstByte"`
	LastByte     string `json:"lastByte"`
	LastMileRTT  string `json:"lastMileRTT"`
}

// NetworkStruct is used for storing the JSON subfields
type NetworkStruct struct {
	Asnum       string `json:"asnum"`
	EdgeIP      string `json:"edgeIP"`
	Network     string `json:"network"`
	NetworkType string `json:"networkType"`
}

// ReqHdrStruct is used for storing the JSON subfields
type ReqHdrStruct struct {
	Cookie string `json:"cookie"`
}

// RespHdrStruct is used for storing the JSON subfields
type RespHdrStruct struct {
	Server  string `json:"server"`
	ContEnc string `json:"contEnc"`
}
