package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)
//ordererorgs struct that contains orderers
type OrdererOrgs struct{
	OrdererOrgs Orderer `yaml:"OrdererOrgs"`
}
// Orderer structs contain name domain and specs
type Orderer struct {
	Name string `yaml:"Name"`
	Domain string `yaml:"Domain"`
	Specs []SpecDetails `yaml:"Specs"`
}
//Specs contain hostname
type SpecDetails struct{
	Hostname string `yaml:"Hostname"`
}
type PeerOrgs struct {
	PeerOrgs []Peers `yaml:"PeerOrgs"`
}
type Peers struct{
	Name string `yaml:"Name"`
	Domain string `yaml:"Domain"`
	Template []Count `yaml:"Template"`
	Users []Count  `yaml:"Users"`
}
type Count struct{
	Count int `yaml:"Count"`
}
type Crypto struct{
	OrdererOrgs OrdererOrgs `yaml:"OrdererOrgs"`
	PeerOrgs PeerOrgs `yaml:"PeerOrgs"`
}

func main() {
	fmt.Println("%%%%%%%%%%%Creating crypto-config.yaml%%%%%%%%%%%%%%%")
	/****
	*      ORDERER PART
	 */
	var specs SpecDetails
	specs = SpecDetails{Hostname:"orderer"}
	specArray := make([]SpecDetails,0,0)
	specArray = append(specArray,specs)
	odata :=Orderer{Name:"Orderer",Domain:"example.com",Specs:specArray}
	ordererdata :=OrdererOrgs{odata}
	/****
	*      PEER PART
	 */
	var count Count
	count = Count{Count:1}
	countArray := make([]Count,0,0)
	countArray = append(countArray,count)
	peerdata:=Peers{Name:"Org1",Domain:"org1.example.com",Template:countArray,Users:countArray}
	PeersSlice := make([]Peers,0,0)
	PeersSlice  = append(PeersSlice, peerdata)
    peerdatafinal:=PeerOrgs{PeersSlice}
    finaldata := Crypto{ordererdata,peerdatafinal}

	marshalbytes, err := yaml.Marshal(&finaldata)
	if err != nil {
		fmt.Println(err.Error)
	}
	writeErr := ioutil.WriteFile("crypto-config.yaml", marshalbytes, 0644)//0644 gives user permission to read and write

	if writeErr != nil {
		fmt.Println(writeErr.Error)
	}
	fmt.Println("%%%%%%%%%%%Done creating crypto-config.yaml%%%%%%%%%%%%%%%")
}