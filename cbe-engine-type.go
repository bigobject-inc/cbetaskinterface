// ----------------------------------------------------------------------
// Copyright (c) 2021 BigObject Inc. ("BigObject")
// All Rights Reserved.
//
// Use of, copying, modifications to, and distribution of this software
// and its documentation without BigObject's
// written permission can result in the violation of U.S. Copyright
// and Patent laws. Violators will be prosecuted to the highest
// extent of the applicable laws.
//
// BIGOBJECT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
// THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
// TO THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
// PARTICULAR PURPOSE, OR NON-INFRINGEMENT.
//
// @author
// Leo Ho <leoho@bigobject.io>
// 2021/02/05
// ----------------------------------------------------------------------

package cbeengineinterface

//	ttType "github.com/bigobject-inc/stailib/tt"

type CBEEngineSetting struct {
	Address  string `json:"address"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`

	CBEPath            string  `json:"cbePath"` // path for cbe script
	LogPath            string  `json:"logPath"`
	LogLevel           string  `json:"logLevel"`
	OpenPoseConfidence float64 `json:"openPoseConfidence"`

	SimultaneousRange float64 `json:"simultaneousRange"`

	Mode string `json:"mode"` // "realtime" or "off-line"
	//	RaiseExternalFunc func(string, ...ttType.Node) error
}

// Event is a structure sending between CBE
type Event struct {
	EventType    string `json:"EventType"`
	ToInstanceID int    `json:"ToInstanceID"` // instance that receive the event

	ToInstanceName string `json:"ToInstanceName"` // instance name that receive the event

	//	Subjects []ttType.Node `json:"Subjects"` //  instance name pair with list of globalIDs. map[varName]global_id

	Location string `json:"Location"` // the location where the event occurs, area/cctv name

	Timestamp   int   `json:"Timestamp"`   // timestamp when the event occurs
	Microsecond int64 `json:"Microsecond"` // timestamp in microsecond

	EventData interface{}
}
