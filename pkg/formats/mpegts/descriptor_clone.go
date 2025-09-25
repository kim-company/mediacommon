package mpegts

import "github.com/asticode/go-astits"

func cloneDescriptors(src []*astits.Descriptor) []*astits.Descriptor {
	if len(src) == 0 {
		return nil
	}

	cloned := make([]*astits.Descriptor, len(src))
	for i, sd := range src {
		if sd == nil {
			continue
		}

		copyDesc := &astits.Descriptor{
			Tag:    sd.Tag,
			Length: sd.Length,
		}

		if sd.Unknown != nil {
			content := make([]byte, len(sd.Unknown.Content))
			copy(content, sd.Unknown.Content)
			copyDesc.Unknown = &astits.DescriptorUnknown{Content: content}
		}

		if sd.Teletext != nil {
			items := make([]*astits.DescriptorTeletextItem, len(sd.Teletext.Items))
			for j, item := range sd.Teletext.Items {
				if item == nil {
					continue
				}

				itemCopy := *item
				items[j] = &itemCopy
			}
			copyDesc.Teletext = &astits.DescriptorTeletext{Items: items}
		}

		if sd.Subtitling != nil {
			items := make([]*astits.DescriptorSubtitlingItem, len(sd.Subtitling.Items))
			for j, item := range sd.Subtitling.Items {
				if item == nil {
					continue
				}

				itemCopy := *item
				items[j] = &itemCopy
			}
			copyDesc.Subtitling = &astits.DescriptorSubtitling{Items: items}
		}

		if sd.Registration != nil {
			reg := *sd.Registration
			copyDesc.Registration = &reg
		}

		if sd.Extension != nil {
			ext := *sd.Extension
			if sd.Extension.Unknown != nil {
				data := make([]byte, len(*sd.Extension.Unknown))
				copy(data, *sd.Extension.Unknown)
				ext.Unknown = &data
			}
			copyDesc.Extension = &ext
		}

		cloned[i] = copyDesc
	}

	return cloned
}
