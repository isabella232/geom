/*
Copyright [2015] Alex Davies-Moore

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

package ewkb

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/devork/geom"
	"github.com/stretchr/testify/assert"
)

func TestEncodePoint(t *testing.T) {
	datasets := []struct {
		data     geom.Geometry
		expected string
	}{
		{&geom.Point{geom.Hdr{geom.XYZM, 27700}, geom.Coordinate{1, 1, 1, 1}}, "00e000000100006c343ff00000000000003ff00000000000003ff00000000000003ff0000000000000"},
		{&geom.Point{geom.Hdr{geom.XYZM, 0}, geom.Coordinate{1, 1, 1, 1}}, "0000000bb93ff00000000000003ff00000000000003ff00000000000003ff0000000000000"},
		{&geom.Point{geom.Hdr{geom.XYZ, 27700}, geom.Coordinate{1, 1, 1}}, "00a000000100006c343ff00000000000003ff00000000000003ff0000000000000"},
		{&geom.Point{geom.Hdr{geom.XYZ, 0}, geom.Coordinate{1, 1, 1}}, "00000003e93ff00000000000003ff00000000000003ff0000000000000"},
		{&geom.Point{geom.Hdr{geom.XYM, 27700}, geom.Coordinate{1, 1, 1}}, "006000000100006c343ff00000000000003ff00000000000003ff0000000000000"},
		{&geom.Point{geom.Hdr{geom.XYM, 0}, geom.Coordinate{1, 1, 1}}, "00000007d13ff00000000000003ff00000000000003ff0000000000000"},
		{&geom.Point{geom.Hdr{geom.XY, 27700}, geom.Coordinate{1, 1}}, "002000000100006c343ff00000000000003ff0000000000000"},
		{&geom.Point{geom.Hdr{geom.XY, 0}, geom.Coordinate{1, 1}}, "00000000013ff00000000000003ff0000000000000"},
	}

	for _, dataset := range datasets {
		var w = new(bytes.Buffer)
		err := Encode(dataset.data, w)

		if err != nil {
			t.Fatalf("Failed to encode Point geometry: err = %s", err)
		}
		data := hex.EncodeToString(w.Bytes())

		assert.Equal(t, dataset.expected, strings.ToLower(data))

	}

}

func TestEncodeLineString(t *testing.T) {
	datasets := []struct {
		data     geom.Geometry
		expected string
	}{
		{&geom.LineString{geom.Hdr{geom.XY, 27700}, []geom.Coordinate{{30, 10}, {10, 30}, {40, 40}}}, "002000000200006c3400000003403e00000000000040240000000000004024000000000000403e00000000000040440000000000004044000000000000"},
	}

	for _, dataset := range datasets {
		var w = new(bytes.Buffer)
		err := Encode(dataset.data, w)

		if err != nil {
			t.Fatalf("Failed to encode LineString geometry: err = %s", err)
		}
		data := hex.EncodeToString(w.Bytes())

		assert.Equal(t, dataset.expected, strings.ToLower(data))

	}

}

func TestEncodePolygon(t *testing.T) {
	datasets := []struct {
		data     geom.Geometry
		expected string
	}{
		{
			&geom.Polygon{
				geom.Hdr{geom.XY, 27700},
				[]geom.LinearRing{{
					[]geom.Coordinate{
						{30, 10},
						{40, 40},
						{20, 40},
						{10, 20},
						{30, 10},
					},
				},
				},
			},
			"002000000300006c340000000100000005403e0000000000004024000000000000404400000000000040440000000000004034000000000000404400000000000040240000000000004034000000000000403e0000000000004024000000000000"},
	}

	for _, dataset := range datasets {
		var w = new(bytes.Buffer)
		err := Encode(dataset.data, w)

		if err != nil {
			t.Fatalf("Failed to encode Polygon geometry: err = %s", err)
		}
		data := hex.EncodeToString(w.Bytes())

		assert.Equal(t, dataset.expected, strings.ToLower(data))

	}

}

func TestEncodeMultiPoint(t *testing.T) {
	datasets := []struct {
		data     geom.Geometry
		expected string
	}{
		{
			&geom.MultiPoint{geom.Hdr{geom.XY, 27700}, []geom.Point{
				geom.Point{geom.Hdr{geom.XY, 0}, geom.Coordinate{10, 40}},
				geom.Point{geom.Hdr{geom.XY, 0}, geom.Coordinate{40, 30}},
				geom.Point{geom.Hdr{geom.XY, 0}, geom.Coordinate{20, 20}},
				geom.Point{geom.Hdr{geom.XY, 0}, geom.Coordinate{30, 10}},
			}},
			"002000000400006c340000000400000000014024000000000000404400000000000000000000014044000000000000403e0000000000000000000001403400000000000040340000000000000000000001403e0000000000004024000000000000",
		},
	}

	for _, dataset := range datasets {
		var w = new(bytes.Buffer)
		err := Encode(dataset.data, w)

		if err != nil {
			t.Fatalf("Failed to encode MultiPoint geometry: err = %s", err)
		}
		data := hex.EncodeToString(w.Bytes())

		assert.Equal(t, dataset.expected, strings.ToLower(data))

	}

}

func TestEncodeMultiLineString(t *testing.T) {

	datasets := []struct {
		data     geom.Geometry
		expected string
	}{
		{
			&geom.MultiLineString{geom.Hdr{geom.XY, 27700}, []geom.LineString{
				geom.LineString{geom.Hdr{geom.XY, 0}, []geom.Coordinate{{10, 10}, {20, 20}, {10, 40}}},
				geom.LineString{geom.Hdr{geom.XY, 0}, []geom.Coordinate{{40, 40}, {30, 30}, {40, 20}, {30, 10}}},
			}},
			"002000000500006c340000000200000000020000000340240000000000004024000000000000403400000000000040340000000000004024000000000000404400000000000000000000020000000440440000000000004044000000000000403e000000000000403e00000000000040440000000000004034000000000000403e0000000000004024000000000000",
		},
	}

	for _, dataset := range datasets {
		var w = new(bytes.Buffer)
		err := Encode(dataset.data, w)

		if err != nil {
			t.Fatalf("Failed to encode MultiPoint geometry: err = %s", err)
		}
		data := hex.EncodeToString(w.Bytes())

		assert.Equal(t, dataset.expected, strings.ToLower(data))

	}
}

func TestEncodeMultiPolygon(t *testing.T) {

	datasets := []struct {
		data     geom.Geometry
		expected string
	}{
		{
			&geom.MultiPolygon{geom.Hdr{geom.XY, 27700}, []geom.Polygon{
				geom.Polygon{
					geom.Hdr{geom.XY, 0},
					[]geom.LinearRing{{
						[]geom.Coordinate{
							{40, 40},
							{20, 45},
							{45, 30},
							{40, 40},
						},
					}},
				},
				geom.Polygon{
					geom.Hdr{geom.XY, 0},
					[]geom.LinearRing{
						{
							[]geom.Coordinate{
								{20, 35}, {10, 30}, {10, 10}, {30, 5}, {45, 20}, {20, 35},
							},
						},
						{
							[]geom.Coordinate{
								{30, 20}, {20, 15}, {20, 25}, {30, 20},
							},
						},
					},
				},
			}},
			"002000000600006c34000000020000000003000000010000000440440000000000004044000000000000403400000000000040468000000000004046800000000000403e0000000000004044000000000000404400000000000000000000030000000200000006403400000000000040418000000000004024000000000000403e00000000000040240000000000004024000000000000403e0000000000004014000000000000404680000000000040340000000000004034000000000000404180000000000000000004403e00000000000040340000000000004034000000000000402e00000000000040340000000000004039000000000000403e0000000000004034000000000000",
		},
	}

	for _, dataset := range datasets {
		var w = new(bytes.Buffer)
		err := Encode(dataset.data, w)

		if err != nil {
			t.Fatalf("Failed to encode MultiPolygon geometry: err = %s", err)
		}
		data := hex.EncodeToString(w.Bytes())

		assert.Equal(t, dataset.expected, strings.ToLower(data))

	}
}

func TestEncodeGeometryCollection(t *testing.T) {

	datasets := []struct {
		data     geom.Geometry
		expected string
	}{
		{
			&geom.GeometryCollection{geom.Hdr{geom.XY, 27700}, []geom.Geometry{
				&geom.Point{
					geom.Hdr{geom.XY, 0},
					geom.Coordinate{4, 6},
				},
				&geom.LineString{
					geom.Hdr{geom.XY, 0},
					[]geom.Coordinate{{4, 6}, {7, 10}},
				},
			}},
			"002000000700006c340000000200000000014010000000000000401800000000000000000000020000000240100000000000004018000000000000401c0000000000004024000000000000",
		},
	}

	for _, dataset := range datasets {
		var w = new(bytes.Buffer)
		err := Encode(dataset.data, w)

		if err != nil {
			t.Fatalf("Failed to encode MultiPolygon geometry: err = %s", err)
		}
		data := hex.EncodeToString(w.Bytes())

		assert.Equal(t, dataset.expected, strings.ToLower(data))

	}
}
