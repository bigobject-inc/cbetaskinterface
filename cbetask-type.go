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

package cbetask

import (
	ttType "github.com/bigobject-inc/stailib/tt"
)

type Event struct {
	EventType    string `json:"EventType"`
	ToInstanceID int    `json:"ToInstanceID"` // instance that receive the event

	ToSequenceName string `json:"ToSequenceName"` // instance name that receive the event

	Subjects []ttType.Node `json:"Subjects"` //  instance name pair with list of globalIDs. map[varName]global_id

	Location string `json:"Location"` // the location where the event occurs, area/cctv name

	Timestamp   int   `json:"Timestamp"`   // timestamp when the event occurs
	Microsecond int64 `json:"Microsecond"` // timestamp in microsecond
}

/*
func (e *Event) getEventType() string {
	return e.eType
}

func (e *Event) setInstanceID(i string) {
	e.eType = i
}

func (e *Event) getToInstanceID() int {
	return e.toInstanceID
}

func (e *Event) setToInstanceID(i int) {
	e.toInstanceID = i
}

func (e *Event) getToInstanceName() string {
	return e.toSequenceName
}

func (e *Event) setToInstanceName(i string) {
	e.toSequenceName = i
}

func (e *Event) getSubjects() []ttType.Node {
	return e.subjects
}

func (e *Event) setSubjects(i []ttType.Node) {
	e.subjects = i
}

func (e *Event) getLocation() string {
	return e.location
}

func (e *Event) setLocation(i string) {
	e.location = i
}

func (e *Event) getTimestamp() int {
	return e.timestamp
}

func (e *Event) setTimestamp(i int) {
	e.timestamp = i
}

func (e *Event) getMicrosecond() int64 {
	return e.microsecond
}

func (e *Event) setMicrosecond(i int64) {
	e.microsecond = i
}
*/
