package flacmeta

import (
	. "launchpad.net/gocheck"
	"testing"
	"os"
	"fmt"
)

func Test(t *testing.T) { TestingT(t) }
type S struct{}
var _ = Suite(&S{})

func (s *S) TestFLACParseMetadataBlockHeader1(c *C) {
	f, err := os.Open("testdata/44100-16-mono.flac")
	if err != nil {
		fmt.Println("FATAL:", err)
		os.Exit(-1)
	}
	defer f.Close()

	metadata := new(FLACMetadata)
	metadata.ReadFLACMetadata(f)

	streaminfo := FLACStreaminfo{
		&FLACMetadataBlockHeader{0, 34, false, 0},
		&FLACStreaminfoBlock{4096, 4096, 11, 14, 44100, 1, 16, 1014300, "e5ccc967ced6c111530e5c79e33c969e"},
		true}
	c.Check(metadata.FLACStreaminfo.Header, Equals, streaminfo.Header)
	c.Check(metadata.FLACStreaminfo.Data, Equals, streaminfo.Data)
	c.Check(metadata.FLACStreaminfo.IsPopulated, Equals, streaminfo.IsPopulated)

	comment := FLACVorbisComment{
		&FLACMetadataBlockHeader{4, 57, false, 0},
		&FLACVorbisCommentBlock{"reference libFLAC 1.2.1 20070917", 1, []string{"ARTIST=GoGoGo"}},
		true}
	c.Check(metadata.FLACVorbisComment.Header, Equals, comment.Header)
	c.Check(metadata.FLACVorbisComment.Data, Equals, comment.Data)
	c.Check(metadata.FLACVorbisComment.IsPopulated, Equals, comment.IsPopulated)

	pad := FLACPadding{&FLACMetadataBlockHeader{1, 8175, true, 0}, nil, true}
	c.Check(pad, Equals, metadata.FLACPadding)

	stb := FLACSeektable{&FLACMetadataBlockHeader{3, 54, false, 3},
		[]*FLACSeekpointBlock{&FLACSeekpointBlock{0, 0, 4096},
			&FLACSeekpointBlock{438272, 1177, 4096},
			&FLACSeekpointBlock{880640, 2452, 4096}},
		true}
	
	c.Check(metadata.FLACSeektable.Header, Equals, stb.Header)
	c.Check(metadata.FLACSeektable.Data, Equals, stb.Data)
	c.Check(metadata.FLACSeektable.IsPopulated, Equals, stb.IsPopulated)

}

func (s *S) TestFLACParseMetadataBlockHeader2(c *C) {
	f, err := os.Open("testdata/mutagen/silence-44-s.flac")
	if err != nil {
		fmt.Println("FATAL:", err)
		os.Exit(-1)
	}
	defer f.Close()

	metadata := new(FLACMetadata)
	metadata.ReadFLACMetadata(f)

	cb := FLACCuesheet{&FLACMetadataBlockHeader{5, 588, false, 0},
		&FLACCuesheetBlock{"1234567890123\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00",
			0x15888,
			true,
			[]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			0x4,
			nil},
		true}
	c.Check(metadata.FLACCuesheet.Header, Equals, cb.Header)
	c.Check(metadata.FLACCuesheet.Data, Equals, cb.Data)
	c.Check(metadata.FLACCuesheet.IsPopulated, Equals, cb.IsPopulated)
}
