// Copyright (c) 2016 The Constantin Karataev. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd.
//
// desc
// https://wiki.wmtransfer.com/projects/webmoney/wiki/Interface_X17

package webmoney

import (
	"encoding/xml"
)

type ContractCreateRequest struct {
	XMLName		 xml.Name  `xml:"contract.request"`
	SignWmid	 string		 `xml:"sign_wmid"`
	Contract   Contract
}

type Contract struct {
	Name			 string    `xml:"name"`
	Ctype      int8      `xml:"ctype"`
	Text			 string    `xml:"text"`
	Sign       string    `xml:"sign"`
	AccessList string    `xml:"accesslist"`
	Wmids      []string  `xml:"accesslist>wmid"`

}

type ContractCreateResponse struct {
	XMLName   xml.Name  `xml:"contract.response"`

	Retval     int8     `xml:"retval"`
	Retdesc    string   `xml:"retdesc"`
	ContractId int64    `xml:"contractid"`
}


type ContractGetInfoReqeust struct {
	XMLName		xml.Name		`xml:"contract.request"`

	Wmid				string		`xml:"wmid"`
	ContractId	string		`xml:"contractid"`
	Mode				string    `xml:"mode"`
	Sign        string		`xml:"sign"`
}

type ContractGetInfoResponse struct {
	XMLName   xml.Name					`xml:"contract.response"`

	Retval       int8						`xml:"retval"`
	Retdesc      string					`xml:"retdesc"`
	ContractInfo []ContractInfo `xml:"contractinfo>row"`
}

type ContractInfo struct {
	XMLName			xml.Name	`xml:"row"`

	ContractId	int64			`xml:"contractid,attr"`
	Wmid				string		`xml:"wmid,attr"`
	AcceptDate	string		`xml:"acceptdate,attr"`
}

func(w *WmClient) CreateContract(c ContractCreateRequest) (ContractCreateResponse) {

}

func(w *WmClient) GetContractInfo(c ContractGetInfoReqeust) (ContractGetInfoResponse) {

}


