package main

import (
	"fmt"
	mc "github.com/Tnze/go-mc/net"
	"image"
	"image/jpeg"
	"os"
)

type woolColor struct {
	Text string
	Hex  [4]uint32
}

func (c *woolColor) RGBA() (r, g, b, a uint32) {
	return c.Hex[0], c.Hex[1], c.Hex[2], c.Hex[3]
}

var WHITE woolColor = woolColor{"white_wool", [4]uint32{0xE9, 0xEC, 0xEC, 0xFF}}
var ORANGE woolColor = woolColor{"orange_wool", [4]uint32{0xF0, 0x76, 0x13, 0xFF}}
var MAGENTA woolColor = woolColor{"magenta_wool", [4]uint32{0xBD, 0x44, 0xB3, 0xFF}}
var LIGHT_BLUE woolColor = woolColor{"light_blue_wool", [4]uint32{0x3A, 0xAF, 0xD9, 0xFF}}
var YELLOW woolColor = woolColor{"yellow_wool", [4]uint32{0xF8, 0xC6, 0x27, 0xFF}}
var LIME woolColor = woolColor{"lime_wool", [4]uint32{0x70, 0xB9, 0x19, 0xFF}}
var PINK woolColor = woolColor{"pink_wool", [4]uint32{0xED, 0x8D, 0xAC, 0xFF}}
var GRAY woolColor = woolColor{"gray_wool", [4]uint32{0x3E, 0x44, 0x47, 0xFF}}
var LIGHT_GRAY woolColor = woolColor{"light_gray_wool", [4]uint32{0x8E, 0x8E, 0x86, 0xFF}}
var CYAN woolColor = woolColor{"cyan_wool", [4]uint32{0x15, 0x89, 0x91, 0xFF}}
var PURPLE woolColor = woolColor{"purple_wool", [4]uint32{0x79, 0x2A, 0xAC, 0xFF}}
var BLUE woolColor = woolColor{"blue_wool", [4]uint32{0x35, 0x39, 0x9D, 0xFF}}
var BROWN woolColor = woolColor{"brown_wool", [4]uint32{0x72, 0x47, 0x28, 0xFF}}
var GREEN woolColor = woolColor{"green_wool", [4]uint32{0x54, 0x6D, 0x1B, 0xFF}}
var RED woolColor = woolColor{"red_wool", [4]uint32{0xA1, 0x27, 0x22, 0xFF}}
var BLACK woolColor = woolColor{"black_wool", [4]uint32{0x14, 0x15, 0x19, 0xFF}}
var AIR woolColor = woolColor{"air", [4]uint32{0x00, 0x00, 0x00, 0x00}}

type Builder struct {
	client mc.RCONClientConn
}

func NewBuilder() *Builder {
	client, err := mc.DialRCON("localhost:25575", "3712697845864520877")

	if err != nil {
		panic(err)
	}

	return &Builder{client}
}

func (b *Builder) setBlock(blockType string, x, y, z int) (string, error) {
	if err := b.execCmd(fmt.Sprintf("setblock %d %d %d %s", x, y, z, blockType)); err != nil {
		return "", err
	}
	return b.readResp()
}

func (b *Builder) execCmd(cmd string) error {
	return b.client.Cmd(cmd)
}

func (b *Builder) readResp() (string, error) {
	return b.client.Resp()
}

func (b *Builder) Close() error {
	return b.client.Close()
}

func (b *Builder) Draw(filePath string, horizontal, vertical, depth int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := height; y != 0; y-- {
		for x := 0; x < width; x++ {
			r, _, _, _ := img.At(x, y).RGBA()

			color := BLACK.Text

			if r < 25000 {
				color = WHITE.Text
			}

			str, err := b.setBlock(color, horizontal+x, vertical+y, depth)

			if err != nil {
				return err
			}

			fmt.Println("setblock", str)
		}
	}
	return nil
}

func (b *Builder) Remove(width, height, horizontal, vertical, depth int) error {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			str, err := b.setBlock(AIR.Text, horizontal+x, vertical+y, depth)

			if err != nil {
				return err
			}

			fmt.Println("setblock", str)
		}
	}
	return nil
}

func main() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	builder := NewBuilder()
	defer builder.Close()

	horizontal := 853
	vertical := 68
	depth := 127

	if err := builder.Draw("./data/image.jpg", horizontal, vertical, depth); err != nil {
		panic(err)
	}

	if err := builder.Remove(100, 100, horizontal, vertical, depth); err != nil {
		panic(err)
	}

}
