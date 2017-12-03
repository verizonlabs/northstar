/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sse

type Config struct {
	Addr                           string // Address at which this server is supposed to listen e.g :8080, localhost:8081 etc
	MaxConcurrentConectionsAllowed uint32 // Any new connection will be disconnected if it exceeds this limit.
	BufferSize                     uint32
	ConnectionQueueSize            uint32 // Number of connection that are in queue.
}
