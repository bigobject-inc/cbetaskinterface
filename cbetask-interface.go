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
	"os"

	cbeutility "bitbucket.org/bigobject/cbe-utility/v3"
	cbe "github.com/bigobject-inc/stailib/cbe"
	ttType "github.com/bigobject-inc/stailib/tt"
)

type Task interface {
	Init(*os.File, *cbeutility.CBEutility, func(string, ...ttType.Node) error) error
	Prepare(cbe.CBE, Event) error
	Process() error
	GetName() string
	GetVersion() string
}
