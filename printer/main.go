package printer

import (
	"fmt"
	"image"

	mc "github.com/Tnze/go-mc/net"
)

type Printer struct {
	client mc.RCONClientConn
}

func NewPrinter(host, password string) *Printer {
	client, err := mc.DialRCON(host, password)

	if err != nil {
		panic(err)
	}

	return &Printer{client}
}

func (b *Printer) Draw(source image.Image, horizontal, vertical, depth int) error {
	img := quantify(source)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := height; y != 0; y-- {
		for x := 0; x < width; x++ {
			color := img.At(x, y)

			err := b.setBlock(woolColorMap[color], horizontal+x, vertical+height-y, depth)

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *Printer) Remove(width, height, horizontal, vertical, depth int) error {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			err := b.setBlock("air", horizontal+x, vertical+y, depth)

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *Printer) setBlock(blockType string, x, y, z int) error {
	if err := b.execCmd(fmt.Sprintf("setblock %d %d %d %s", x, y, z, blockType)); err != nil {
		return err
	}
	_, err := b.readResp()
	return err
}

func (b *Printer) execCmd(cmd string) error {
	return b.client.Cmd(cmd)
}

func (b *Printer) readResp() (string, error) {
	return b.client.Resp()
}

func (b *Printer) Close() error {
	return b.client.Close()
}

func quantify(src image.Image) image.Image {
	dst := image.NewPaletted(
		image.Rect(0, 0, src.Bounds().Dx(), src.Bounds().Dy()),
		woolPalette,
	)

	return dither["FloydSteinberg"].Quantize(src, dst, len(woolPalette), true, true)
}
