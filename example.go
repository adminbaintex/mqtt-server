// Copyright (c) 2014 The gomqtt Authors. All rights reserved.
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

package main

import (
	"fmt"
	"net"

	"github.com/escrichov/mqtt-server/server"
	"github.com/gomqtt/stream"
)

var done = make(chan struct{})

type ExampleHandler struct{}

func (handler *ExampleHandler) ServeMQTT(info *server.ConnInfo, s stream.Stream) {
	fmt.Println("New connection", info.Conn.RemoteAddr())

	close(done)
}

func main() {

	server := server.NewServer(&ExampleHandler{})

	server.LaunchTCPConfiguration("localhost:1337")
	server.LaunchWSConfiguration("localhost:1338")

	c, _ := net.Dial("tcp", "localhost:1337")
	c.Close()

	<-done
	server.Stop()

	// Output:
	// new connection
}
