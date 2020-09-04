package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"unsafe"

	"github.com/lijo-jose/glib"
)

type PadLinkReturn C.GstPadLinkReturn

const (
	PAD_LINK_OK              = PadLinkReturn(C.GST_PAD_LINK_OK)
	PAD_LINK_WRONG_HIERARCHY = PadLinkReturn(C.GST_PAD_LINK_WRONG_HIERARCHY)
	PAD_LINK_WAS_LINKED      = PadLinkReturn(C.GST_PAD_LINK_WAS_LINKED)
	PAD_LINK_WRONG_DIRECTION = PadLinkReturn(C.GST_PAD_LINK_WRONG_DIRECTION)
	PAD_LINK_NOFORMAT        = PadLinkReturn(C.GST_PAD_LINK_NOFORMAT)
	PAD_LINK_NOSCHED         = PadLinkReturn(C.GST_PAD_LINK_NOSCHED)
	PAD_LINK_REFUSED         = PadLinkReturn(C.GST_PAD_LINK_REFUSED)
)

func (p PadLinkReturn) String() string {
	switch p {
	case PAD_LINK_OK:
		return "PAD_LINK_OK"
	case PAD_LINK_WRONG_HIERARCHY:
		return "PAD_LINK_WRONG_HIERARCHY"
	case PAD_LINK_WAS_LINKED:
		return "PAD_LINK_WAS_LINKED"
	case PAD_LINK_WRONG_DIRECTION:
		return "PAD_LINK_WRONG_DIRECTION"
	case PAD_LINK_NOFORMAT:
		return "PAD_LINK_NOFORMAT"
	case PAD_LINK_NOSCHED:
		return "PAD_LINK_NOSCHED"
	case PAD_LINK_REFUSED:
		return "PAD_LINK_REFUSED"
	}
	panic("Wrong value of PadLinkReturn variable")
}

type PadDirection C.GstPadDirection

const (
	PAD_UNKNOWN = PadDirection(C.GST_PAD_UNKNOWN)
	PAD_SRC     = PadDirection(C.GST_PAD_SRC)
	PAD_SINK    = PadDirection(C.GST_PAD_SINK)
)

func (p PadDirection) g() C.GstPadDirection {
	return C.GstPadDirection(p)
}

func (p PadDirection) String() string {
	switch p {
	case PAD_UNKNOWN:
		return "PAD_UNKNOWN"
	case PAD_SRC:
		return "PAD_SRC"
	case PAD_SINK:
		return "PAD_SINK"
	}
	panic("Wrong value of PadDirection variable")
}

type Pad struct {
	GstObj
}

func (p *Pad) g() *C.GstPad {
	return (*C.GstPad)(p.GetPtr())
}

func (p *Pad) AsPad() *Pad {
	return p
}

func (p *Pad) CanLink(sink_pad *Pad) bool {
	return C.gst_pad_can_link(p.g(), sink_pad.g()) != 0
}

func (p *Pad) Link(sink_pad *Pad) PadLinkReturn {
	return PadLinkReturn(C.gst_pad_link(p.g(), sink_pad.g()))
}

type PadProbeType C.GstprobeType

const (
	PAD_PROBE_TYPE_INVALID = PadProbeType(C.GST_PAD_PROBE_TYPE_INVALID)
	PAD_PROBE_TYPE_IDLE = PadProbeType(C.GST_PAD_PROBE_TYPE_IDLE)
	PAD_PROBE_TYPE_BLOCK = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK)

	PAD_PROBE_TYPE_BUFFER = PadProbeType(C.GST_PAD_PROBE_TYPE_BUFFER)
	PAD_PROBE_TYPE_BUFFER_LIST = PadProbeType(C.GST_PAD_PROBE_TYPE_BUFFER_LIST)

	PAD_PROBE_TYPE_EVENT_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_DOWNSTREAM)
	PAD_PROBE_TYPE_EVENT_UPSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_UPSTREAM)
	PAD_PROBE_TYPE_EVENT_FLUSH = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_FLUSH)

	PAD_PROBE_TYPE_QUERY_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_DOWNSTREAM)
	PAD_PROBE_TYPE_QUERY_UPSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_UPSTREAM)

	PAD_PROBE_TYPE_PUSH = PadProbeType(C.GST_PAD_PROBE_TYPE_PUSH)
	PAD_PROBE_TYPE_PULL = PadProbeType(C.GST_PAD_PROBE_TYPE_PULL)

	PAD_PROBE_TYPE_BLOCKING = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCKING)
	PAD_PROBE_TYPE_DATA_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_DOWNSTREAM)
	PAD_PROBE_TYPE_DATA_UPSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_UPSTREAM)
	PAD_PROBE_TYPE_DATA_BOTH = PadProbeType(C.GST_PAD_PROBE_TYPE_DATA_BOTH)
	PAD_PROBE_TYPE_BLOCK_DOWNSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK_DOWNSTREAM)
	PAD_PROBE_TYPE_BLOCK_UPSTREAM = PadProbeType(C.GST_PAD_PROBE_TYPE_BLOCK_UPSTREAM)
	PAD_PROBE_TYPE_EVENT_BOTH = PadProbeType(C.GST_PAD_PROBE_TYPE_EVENT_BOTH)
	PAD_PROBE_TYPE_QUERY_BOTH = PadProbeType(C.GST_PAD_PROBE_TYPE_QUERY_BOTH)
	PAD_PROBE_TYPE_ALL_BOTH = PadProbeType(C.GST_PAD_PROBE_TYPE_ALL_BOTH)
	PAD_PROBE_TYPE_SCHEDULING = PadProbeType(C.GST_PAD_PROBE_TYPE_SCHEDULING)

)

func (p *PadProbeType) String() string  {
	switch p {
	case PAD_PROBE_TYPE_INVALID:
		return "PAD_PROBE_TYPE_INVALID"
	case PAD_PROBE_TYPE_IDLE:
		return "PAD_PROBE_TYPE_IDLE"
	case PAD_PROBE_TYPE_BLOCK:
		return "PAD_PROBE_TYPE_BLOCK"

	case PAD_PROBE_TYPE_BUFFER:
		return "PAD_PROBE_TYPE_BUFFER"
	case PAD_PROBE_TYPE_BUFFER_LIST:
		return "PAD_PROBE_TYPE_BUFFER_LIST"
	case PAD_PROBE_TYPE_EVENT_DOWNSTREAM:
		return "PAD_PROBE_TYPE_EVENT_DOWNSTREAM"
	case PAD_PROBE_TYPE_EVENT_UPSTREAM:
		return "PAD_PROBE_TYPE_BLOCK"
	case PAD_PROBE_TYPE_EVENT_FLUSH:
		return "PAD_PROBE_TYPE_EVENT_FLUSH"

	case PAD_PROBE_TYPE_QUERY_DOWNSTREAM:
		return "PAD_PROBE_TYPE_QUERY_DOWNSTREAM"
	case PAD_PROBE_TYPE_QUERY_UPSTREAM:
		return "PAD_PROBE_TYPE_QUERY_UPSTREAM"

	case PAD_PROBE_TYPE_PUSH:
		return "PAD_PROBE_TYPE_PUSH"
	case PAD_PROBE_TYPE_PULL:
		return "PAD_PROBE_TYPE_PULL"

	case PAD_PROBE_TYPE_BLOCKING:
		return "PAD_PROBE_TYPE_BLOCKING"
	case PAD_PROBE_TYPE_DATA_DOWNSTREAM:
		return "PAD_PROBE_TYPE_DATA_DOWNSTREAM"
	case PAD_PROBE_TYPE_DATA_UPSTREAM:
		return "PAD_PROBE_TYPE_DATA_UPSTREAM"
	case PAD_PROBE_TYPE_DATA_BOTH:
		return "PAD_PROBE_TYPE_DATA_BOTH"
	case PAD_PROBE_TYPE_BLOCK_DOWNSTREAM:
		return "PAD_PROBE_TYPE_BLOCK_DOWNSTREAM"
	case PAD_PROBE_TYPE_BLOCK_UPSTREAM:
		return "PAD_PROBE_TYPE_BLOCK_UPSTREAM"
	case PAD_PROBE_TYPE_EVENT_BOTH:
		return "PAD_PROBE_TYPE_EVENT_BOTH"
	case PAD_PROBE_TYPE_QUERY_BOTH:
		return "PAD_PROBE_TYPE_QUERY_BOTH"
	case PAD_PROBE_TYPE_ALL_BOTH:
		return "PAD_PROBE_TYPE_ALL_BOTH"
	case PAD_PROBE_TYPE_SCHEDULING:
		return "PAD_PROBE_TYPE_SCHEDULING"
	}
	panic("Wrong value of PAD_PROBE_TYPE variable")
}

func (p PadProbeType) g() C.GstprobeType {
	return C.GstprobeType(p)
}

func (p *Pad) AddProbe(pad *Pad) uint64 {
	return C.gst_pad_add_probe()
}

type GhostPad struct {
	Pad
}

func (p *GhostPad) g() *C.GstGhostPad {

	return (*C.GstGhostPad)(p.GetPtr())
}

func (p *GhostPad) AsGhostPad() *GhostPad {
	return p
}

func (p *GhostPad) SetTarget(new_target *Pad) bool {
	return C.gst_ghost_pad_set_target(p.g(), new_target.g()) != 0
}

func (p *GhostPad) GetTarget() *Pad {
	r := new(Pad)
	r.SetPtr(glib.Pointer(C.gst_ghost_pad_get_target(p.g())))
	return r
}

func (p *GhostPad) Construct() bool {
	return C.gst_ghost_pad_construct(p.g()) != 0
}

func NewGhostPad(name string, target *Pad) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new(s, target.g())))
	return p
}

func NewGhostPadNoTarget(name string, dir PadDirection) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new_no_target(s, dir.g())))
	return p
}
