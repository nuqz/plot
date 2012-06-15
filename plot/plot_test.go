package plot

import (
	"code.google.com/p/plotinum/plt"
	"code.google.com/p/plotinum/vg"
	"code.google.com/p/plotinum/vg/veceps"
	"code.google.com/p/plotinum/vg/vecimg"
	"math/rand"
	"time"
	"testing"
)

func TestDrawImage(t *testing.T) {
	w, h := vg.Inches(4), vg.Inches(4)
	img, err := vecimg.New(w, h)
	if err != nil {
		t.Error(err)
	}
	da := plt.NewDrawArea(img, w, h)
	draw(da)
	err = img.SavePNG("test.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDrawEps(t *testing.T) {
	w, h := vg.Inches(4), vg.Inches(4)
	da := plt.NewDrawArea(veceps.New(w, h, "test"), w, h)
	draw(da)
	err := da.Canvas.(*veceps.Canvas).Save("test.eps")
	if err != nil {
		t.Fatal(err)
	}
}

// draw draws a simple test plot
func draw(da *plt.DrawArea) {
	p := plt.New()
	p.Title.Text = "Title"
	p.Y.Label.Text = "Y Label"
	vs0 := make(Values, 10)
	rand.Seed(time.Now().UnixNano())
	for i := range vs0 {
		vs0[i] = rand.Float64()*1000
	}
	vs1 := make(Values, 10)
	for i := range vs1 {
		vs1[i] = rand.NormFloat64()*200 + 500
	}
	vs2 := make(Values, 10)
	for i := range vs2 {
		vs2[i] = rand.ExpFloat64()*300
	}
	p.AddData(MakeBox(vg.Points(18), 0, vs0))
	p.AddData(MakeBox(vg.Points(18), 1, vs1))
	p.AddData(MakeBox(vg.Points(18), 2, vs2))
	p.X.Tick.Marker = plt.ConstantTicks([]plt.Tick{
		{0, "Uniform",}, {1, "Normal",}, {2, "Exponential"},
	})
	p.X.Tick.Label.Font.Size = vg.Points(12)
	p.X.Tick.Width = 0
	p.X.Tick.Length = 0
	p.X.Width = 0

	p.Y.Min = 0
	p.Y.Max = 1000
	p.Draw(da)
}