// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/17ac39c7f9266bc303baa029f90194aecb1c3b7c

// Gets stats for anomaly detection job model snapshot upgrades that are in
// progress.
package getmodelsnapshotupgradestats

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	jobidMask = iota + 1

	snapshotidMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetModelSnapshotUpgradeStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	jobid      string
	snapshotid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetModelSnapshotUpgradeStats type alias for index.
type NewGetModelSnapshotUpgradeStats func(jobid, snapshotid string) *GetModelSnapshotUpgradeStats

// NewGetModelSnapshotUpgradeStatsFunc returns a new instance of GetModelSnapshotUpgradeStats with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetModelSnapshotUpgradeStatsFunc(tp elastictransport.Interface) NewGetModelSnapshotUpgradeStats {
	return func(jobid, snapshotid string) *GetModelSnapshotUpgradeStats {
		n := New(tp)

		n._jobid(jobid)

		n._snapshotid(snapshotid)

		return n
	}
}

// Gets stats for anomaly detection job model snapshot upgrades that are in
// progress.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-get-job-model-snapshot-upgrade-stats.html
func New(tp elastictransport.Interface) *GetModelSnapshotUpgradeStats {
	r := &GetModelSnapshotUpgradeStats{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetModelSnapshotUpgradeStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask|snapshotidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("anomaly_detectors")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "jobid", r.jobid)
		}
		path.WriteString(r.jobid)
		path.WriteString("/")
		path.WriteString("model_snapshots")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "snapshotid", r.snapshotid)
		}
		path.WriteString(r.snapshotid)
		path.WriteString("/")
		path.WriteString("_upgrade")
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodGet
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetModelSnapshotUpgradeStats) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "ml.get_model_snapshot_upgrade_stats")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "ml.get_model_snapshot_upgrade_stats")
		if reader := instrument.RecordRequestBody(ctx, "ml.get_model_snapshot_upgrade_stats", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_model_snapshot_upgrade_stats")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetModelSnapshotUpgradeStats query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getmodelsnapshotupgradestats.Response
func (r GetModelSnapshotUpgradeStats) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_model_snapshot_upgrade_stats")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetModelSnapshotUpgradeStats) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_model_snapshot_upgrade_stats")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the GetModelSnapshotUpgradeStats query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetModelSnapshotUpgradeStats headers map.
func (r *GetModelSnapshotUpgradeStats) Header(key, value string) *GetModelSnapshotUpgradeStats {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for the anomaly detection job.
// API Name: jobid
func (r *GetModelSnapshotUpgradeStats) _jobid(jobid string) *GetModelSnapshotUpgradeStats {
	r.paramSet |= jobidMask
	r.jobid = jobid

	return r
}

// SnapshotId A numerical character string that uniquely identifies the model snapshot. You
// can get information for multiple
// snapshots by using a comma-separated list or a wildcard expression. You can
// get all snapshots by using `_all`,
// by specifying `*` as the snapshot ID, or by omitting the snapshot ID.
// API Name: snapshotid
func (r *GetModelSnapshotUpgradeStats) _snapshotid(snapshotid string) *GetModelSnapshotUpgradeStats {
	r.paramSet |= snapshotidMask
	r.snapshotid = snapshotid

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
//  -  Contains wildcard expressions and there are no jobs that match.
//  -  Contains the _all string or no identifiers and there are no matches.
//  -  Contains wildcard expressions and there are only partial matches.
//
// The default value is true, which returns an empty jobs array when there are
// no matches and the subset of results
// when there are partial matches. If this parameter is false, the request
// returns a 404 status code when there are
// no matches or only partial matches.
// API name: allow_no_match
func (r *GetModelSnapshotUpgradeStats) AllowNoMatch(allownomatch bool) *GetModelSnapshotUpgradeStats {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}
