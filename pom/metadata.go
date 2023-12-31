package pom

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"io"
)

type Versioning struct {
	Latest   string   `xml:"latest"`
	Release  string   `xml:"release"`
	Versions []string `xml:"versions>version"`
}

/* XML file defining state of versions for a package. */
type Metadata struct {
	GroupId    string     `xml:"groupId"`
	ArtifactId string     `xml:"artifactId"`
	Version    string     `xml:"version"`
	Versioning Versioning `xml:"versioning"`
}

/* For reading Metadata from downloaded XML file. */
func MetadataFromReader(reader io.ReadCloser) (meta *Metadata, err error) {
	defer func() {
		cerr := reader.Close()
		if err == nil {
			err = cerr
		}
	}()
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Entity = xml.HTMLEntity
	err = decoder.Decode(&meta)
	return
}

/* There are multiple values that could indicate latest version. */
func (m *Metadata) GetLatest() string {
	if m.Versioning.Latest != "" {
		return m.Versioning.Latest
	} else if m.Versioning.Release != "" {
		return m.Versioning.Release
	} else {
		return m.Version
	}
}
