// Copyright (C) 2013-2018 by Maxim Bublis <b@codemonkey.ru>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package uuid provides implementations of the Universally Unique Identifier
// (UUID), as specified in RFC-4122,
//
// RFC-4122[1] provides the specification for versions 1, 3, 4, and 5.
//
// DCE 1.1[2] provides the specification for version 2, but version 2 support
// was removed from this package in v4 due to some concerns with the
// specification itself. Reading the spec, it seems that it would result in
// generating UUIDs that aren't very unique. In having read the spec it seemed
// that our implementation did not meet the spec. It also seems to be at-odds
// with RFC 4122, meaning we would need quite a bit of special code to support
// it. Lastly, there were no Version 2 implementations that we could find to
// ensure we were understanding the specification correctly.
//
// [1] https://tools.ietf.org/html/rfc4122
// [2] http://pubs.opengroup.org/onlinepubs/9696989899/chap5.htm#tagcjh_08_02_01_01
package uuid

import (
	"encoding/hex"
)

// StringHex returns a hex string representation of the UUID:
// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.
func (u UUID) HexString() string {
	return hex.EncodeToString(u.Bytes())
}
