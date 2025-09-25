package mpegts

import (
	"fmt"

	"github.com/asticode/go-astits"
)

// Codec is a MPEG-TS codec.
type Codec interface {
	IsVideo() bool

	isCodec()
	marshal(pid uint16) (*astits.PMTElementaryStream, error)
}

// CodecDVBTeletext is a DVB Teletext codec (ETSI EN 300 472).
// Payload is passed through as PES data (one or more teletext data units).
type CodecDVBTeletext struct {
	Descriptors []*astits.Descriptor
}

func (CodecDVBTeletext) IsVideo() bool { return false }

func (*CodecDVBTeletext) isCodec() {}

func (c CodecDVBTeletext) marshal(pid uint16) (*astits.PMTElementaryStream, error) {
	if len(c.Descriptors) == 0 {
		return nil, fmt.Errorf("DVB Teletext descriptors are missing")
	}

	for _, desc := range c.Descriptors {
		if desc == nil {
			return nil, fmt.Errorf("DVB Teletext descriptor is nil")
		}
		if desc.Tag != astits.DescriptorTagTeletext {
			return nil, fmt.Errorf("invalid descriptor tag %#x for DVB Teletext", desc.Tag)
		}
	}

	return &astits.PMTElementaryStream{
		ElementaryPID: pid,
		StreamType:    astits.StreamTypePrivateData,
		ElementaryStreamDescriptors: cloneDescriptors(c.Descriptors),
	}, nil
}
