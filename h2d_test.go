// Copyright ©2016 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hplot_test

import (
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/go-hep/hbook"
	"github.com/go-hep/hplot"
	"github.com/go-hep/hplot/internal/cmpimg"
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/stat/distmv"
)

func TestH2D(t *testing.T) {
	ExampleNewH2D(t)

	want, err := ioutil.ReadFile("testdata/h2d_plot_golden.png")
	if err != nil {
		t.Fatal(err)
	}

	got, err := ioutil.ReadFile("testdata/h2d_plot.png")
	if err != nil {
		t.Fatal(err)
	}

	if ok, err := cmpimg.Equal("png", got, want); !ok || err != nil {
		t.Fatalf("error: testdata/h2d_plot.png differ with reference file: %v\n", err)
	}
}

func ExampleNewH2D(t *testing.T) {
	h := hbook.NewH2D(100, -10, 10, 100, -10, 10)

	const npoints = 10000

	dist, ok := distmv.NewNormal(
		[]float64{0, 1},
		mat64.NewSymDense(2, []float64{4, 0, 0, 2}),
		rand.New(rand.NewSource(1234)),
	)
	if !ok {
		t.Fatalf("error creating distmv.Normal")
	}

	v := make([]float64, 2)
	// Draw some random values from the standard
	// normal distribution.
	for i := 0; i < npoints; i++ {
		v = dist.Rand(v)
		h.Fill(v[0], v[1], 1)
	}

	p, err := plot.New()
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	p.Title.Text = "Hist-2D"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	p.Add(hplot.NewH2D(h, nil))
	p.Add(plotter.NewGrid())
	err = p.Save(10*vg.Centimeter, 10*vg.Centimeter, "testdata/h2d_plot.png")
	if err != nil {
		t.Fatal(err)
	}
}
