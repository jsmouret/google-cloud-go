// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package data_test

import (
	"context"

	data "cloud.google.com/go/analytics/data/apiv1alpha"
	datapb "google.golang.org/genproto/googleapis/analytics/data/v1alpha"
)

func ExampleNewAlphaAnalyticsDataClient() {
	ctx := context.Background()
	c, err := data.NewAlphaAnalyticsDataClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleAlphaAnalyticsDataClient_RunReport() {
	// import datapb "google.golang.org/genproto/googleapis/analytics/data/v1alpha"

	ctx := context.Background()
	c, err := data.NewAlphaAnalyticsDataClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datapb.RunReportRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RunReport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAlphaAnalyticsDataClient_RunPivotReport() {
	// import datapb "google.golang.org/genproto/googleapis/analytics/data/v1alpha"

	ctx := context.Background()
	c, err := data.NewAlphaAnalyticsDataClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datapb.RunPivotReportRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RunPivotReport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAlphaAnalyticsDataClient_BatchRunReports() {
	// import datapb "google.golang.org/genproto/googleapis/analytics/data/v1alpha"

	ctx := context.Background()
	c, err := data.NewAlphaAnalyticsDataClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datapb.BatchRunReportsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.BatchRunReports(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAlphaAnalyticsDataClient_BatchRunPivotReports() {
	// import datapb "google.golang.org/genproto/googleapis/analytics/data/v1alpha"

	ctx := context.Background()
	c, err := data.NewAlphaAnalyticsDataClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &datapb.BatchRunPivotReportsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.BatchRunPivotReports(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}