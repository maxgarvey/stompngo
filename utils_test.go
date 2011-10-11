//
// Copyright © 2011 Guy M. Allard
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package stomp

import (
	"net"
	"os"
	"testing"
)

var test_login = "guest"
var test_passcode = "guest"
var test_headers = Headers{"login", test_login, "passcode", test_passcode}

func openConn(t *testing.T) (n net.Conn, err os.Error) {
	h, p := hostAndPort()
	n, err = net.Dial("tcp", net.JoinHostPort(h, p))
	if err != nil {
		t.Errorf("Unexpected net.Dial error: %v\n", err)
	}
	return n, err
}

func closeConn(t *testing.T, n net.Conn) (err os.Error) {
	err = n.Close()
	if err != nil {
		t.Errorf("Unexpected n.Close() error: %v\n", err)
	}
	return err
}

// Host and port for Dial
func hostAndPort() (string, string) {
	h := os.Getenv("STOMP_HOST")
	if h == "" {
		h = "localhost"
	}
	p := os.Getenv("STOMP_PORT")
	if p == "" {
		p = "51613"
	}
	return h, p
}
