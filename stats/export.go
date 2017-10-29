// Copyright 2017, OpenCensus Authors
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

package stats

import "sync"

var (
	exportersMu sync.Mutex // guards exporters
	exporters   = make(map[Exporter]struct{})
)

// Exporter exports the collected records as view data.
type Exporter interface {
	Export(viewData *ViewData)
}

// RegisterExporter registers an exporter.
// Collected data will be reported via all the
// registered exporters. Once you don't want data
// to be expoter on the registered exporter, use
// UnregisterExporter.
func RegisterExporter(e Exporter) {
	exportersMu.Lock()
	defer exportersMu.Unlock()

	exporters[e] = struct{}{}
}

// UnregisterExporter unregisters an exporter.
func UnregisterExporter(e Exporter) {
	exportersMu.Lock()
	defer exportersMu.Unlock()

	delete(exporters, e)
}
