// This simple test application create live H264 (or WebM - see commented lines)
// content from test source, decode it and display.
package main

import (
	"fmt"
	"os"

	"github.com/lijo-jose/glib"
	"github.com/lijo-jose/gst"
)

func checkElem(e *gst.Element, name string) {
	if e == nil {
		fmt.Fprintln(os.Stderr, "can't make element: ", name)
		os.Exit(1)
	}
}

func main() {
	src := gst.ElementFactoryMake("videotestsrc", "Test source")
	checkElem(src, "videotestsrc")
	src.SetProperty("do-timestamp", true)
	src.SetProperty("pattern", 18) // ball

	fmt.Println("Hej")
	//encType := "mpeg2enc"
	// encType := "x264enc"
	encType := "vp8enc"
	enc := gst.ElementFactoryMake(encType, "Video encoder")
	checkElem(enc, encType)
	fmt.Println("Hej")

	//muxType := "webmux"
	muxType := "matroskamux"
	mux := gst.ElementFactoryMake(muxType, "Muxer")
	checkElem(mux, muxType)
	mux.SetProperty("streamable", true)

	demux := gst.ElementFactoryMake("matroskademux", "Matroska demuxer")
	checkElem(demux, "matroskademux")

	//decType := "mpeg2dec"
	//decType := "avdec_h264"
	decType := "vp8dec"
	dec := gst.ElementFactoryMake(decType, "Video decoder")
	checkElem(dec, decType)

	sink := gst.ElementFactoryMake("autovideosink", "Video display element")
	checkElem(sink, "autovideosink")

	pl := gst.NewPipeline("MyPipeline")

	pl.Add(src, enc, mux, demux, dec, sink)

	src.Link(enc, mux, demux)
	fmt.Println("Nice1" )

	demux.ConnectNoi("pad-added", cbPadAddedNew, dec.GetStaticPad("sink"))
	fmt.Println("Nice 22" )

	dec.Link(sink)
	fmt.Println("Nice 333" )

	pl.SetState(gst.STATE_PLAYING)

	glib.NewMainLoop(nil).Run()
}

// Callback function for "pad-added" event
func cbPadAddedNew(dec_sink_pad, demux_new_pad *gst.Pad) {
	fmt.Println("New pad:", demux_new_pad.GetName())
	if demux_new_pad.CanLink(dec_sink_pad) {
		if demux_new_pad.Link(dec_sink_pad) != gst.PAD_LINK_OK {
			fmt.Fprintln(os.Stderr, "link error")
		}
	} else {
		fmt.Fprintln(os.Stderr, "can't link it with:", dec_sink_pad.GetName())
	}
}
