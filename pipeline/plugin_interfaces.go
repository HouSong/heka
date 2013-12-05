/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Rob Miller (rmiller@mozilla.com)
#   Mike Trinkala (trink@mozilla.com)
#   Ben Bangert (bbangert@mozilla.com)
#
# ***** END LICENSE BLOCK *****/

package pipeline

// Input plugin interface type.
type Input interface {
	// Start listening for / gathering incoming data, populating
	// PipelinePacks, and passing them along to a Decoder or the Router as
	// appropriate.
	Run(ir InputRunner, h PluginHelper) (err error)
	// Called as a signal to the Input to stop listening for / gathering
	// incoming data and to perform any necessary clean-up.
	Stop()
}

// Heka Decoder plugin interface.
type Decoder interface {
	// Extract data loaded into the PipelinePack (usually in pack.MsgBytes)
	// and use it to populate pack.Message message objects. If decoding
	// succeeds (i.e. `err` is nil), the original pack will be mutated and
	// returned as the first item in the `packs` return slice. If there is an
	// error, `packs` should be returned as nil.
	// Returning (nil, nil) is valid in cases where the decoding failed but
	// the error should not be logged.
	Decode(pack *PipelinePack) (packs []*PipelinePack, err error)
}

// Heka Filter plugin type.
type Filter interface {
	// Starts the filter listening on the FilterRunner's provided input
	// channel. Should not return until shutdown, signaled to the Filter by
	// the closure of the input channel. Should return a non-nil error value
	// only if errors happen during start-up or if there is an unclean
	// shutdown (i.e. not due to an error processing an isolated message, in
	// that case use FilterRunner.LogError).
	Run(r FilterRunner, h PluginHelper) (err error)
}

// Heka Output plugin type.
type Output interface {
	Run(or OutputRunner, h PluginHelper) (err error)
}
