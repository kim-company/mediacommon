package mpegts

import (
	"testing"

	"github.com/asticode/go-astits"
	"github.com/stretchr/testify/require"
)

func TestCodecDVBTeletextMarshal(t *testing.T) {
	codec := CodecDVBTeletext{
		Descriptors: []*astits.Descriptor{
			{
				Tag:    astits.DescriptorTagTeletext,
				Length: 5,
				Teletext: &astits.DescriptorTeletext{
					Items: []*astits.DescriptorTeletextItem{
						{
							ISO639LanguageCode:    "eng",
							TeletextType:          0x02,
							TeletextMagazineNumber: 1,
							TeletextPageNumber:     0x00,
						},
					},
				},
			},
		},
	}

	es, err := codec.marshal(256)
	require.NoError(t, err)
	require.Equal(t, uint16(256), es.ElementaryPID)
	require.Equal(t, astits.StreamTypePrivateData, es.StreamType)
	require.Len(t, es.ElementaryStreamDescriptors, 1)
	require.Equal(t, astits.DescriptorTagTeletext, es.ElementaryStreamDescriptors[0].Tag)
	require.NotNil(t, es.ElementaryStreamDescriptors[0].Teletext)
	require.Len(t, es.ElementaryStreamDescriptors[0].Teletext.Items, 1)
}
