package main

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	errNamingSchemeNotFollowed = errors.New("File name doesn't follow the naming scheme.")
	errIncorrectIchanKey       = errors.New("Incorrect ichan key")
	errDimensionsDontMatch     = errors.New("Dimensions don't match")
	errImproperDimensions      = errors.New("Improper dimensions! Dimensions must be a power of 2!")
)

var (
	defaultResolutionPx = 1024
	defaultResolution   = image.Rect(0, 0, defaultResolutionPx, defaultResolutionPx)
)

var regIchans = regexp.MustCompile(`(?i)^(.+)_(a|s|g|r|da|d|ao)\.(png)$`)

func rgb2gray(r, g, b uint32) uint32 {
	return (r + g + b) / 3
}

// uint32 0x0000-0xFFFF color to float64 0.0-1.0 color, clamped (e.g. 65537 -> 1.0)
func i2f(n uint32) float64 {
	if n > 65535 {
		return 1.0
	}

	return float64(n) / 65535.0
}

// uint32 0x0000-0xFFFF color to 0-255 uint8 color, clamped (e.g. 65537 -> 255)
func i2i8(n uint32) uint8 {
	if n > 65535 {
		return 255
	}

	return uint8(n / 257) // 65535 / 255 = 257
}

// float64 0.0-1.0 color to 0-255 uint32 color, clamped (e.g. -1.0 -> 0, 2.0 -> 65535)
func f2i(n float64) uint32 {
	return uint32(clamp1(n) * 65535)
}

// float64 0.0-1.0 color to 0-255 uint8 color, clamped (e.g. -1.0 -> 0, 2.0 -> 255)
func f2i8(n float64) uint8 {
	return uint8(clamp1(n) * 255)
}

// clamp float64 to 0.0-1.0
func clamp1(n float64) float64 {
	if n < 0.0 {
		return 0.0
	}

	if n > 1.0 {
		return 1.0
	}

	return n
}

func isPowerOfTwo(n int) bool {
	return n != 0 && n&(n-1) == 0
}

func dirExists(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return info.IsDir()
}

func scanDirForIchanFiles(dirPath string) (map[string]*ichanFiles, error) {
	fileInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	ics := make(map[string]*ichanFiles, len(fileInfo))

	for _, x := range fileInfo {
		// Skip dirs
		if x.IsDir() {
			continue
		}

		filename := x.Name()
		name, chtype, _, err := decodeFileName(filename)
		if err != nil {
			continue
		}

		ic, ok := ics[name]
		// Create a new ichanFiles struct if neccesarry
		if !ok {
			ic = &ichanFiles{
				dir:  dirPath,
				name: name,
			}
			ics[name] = ic
		}

		realpath, err := filepath.Abs(path.Join(dirPath, filename))
		if err != nil {
			logwarn.Println("WTF:", err)
			continue
		}

		ic.set(chtype, realpath)
	}

	return ics, nil
}

func decodeFileName(s string) (name, chtype, ext string, err error) {
	ss := regIchans.FindStringSubmatch(s)
	if len(ss) != 4 {
		return "", "", "", errNamingSchemeNotFollowed
	}

	name = ss[1]
	chtype = strings.ToLower(ss[2])
	ext = strings.ToLower(ss[3])

	return
}

func loadImageFromFile(path string) (image.Image, error) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func writeImageAsPngFile(path string, img image.Image) error {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}

//
//  ichanFiles
//

type ichanFiles struct {
	dir  string
	name string
	a    string // Alpha
	s    string // Specular
	g    string // Glow
	r    string // Reflectivity
	d    string // Diffuse
	da   string // Diffuse Alpha
	ao   string // Ambient Occlusion
}

func (self *ichanFiles) get(k string) (string, error) {
	switch k {
	case "a":
		return self.a, nil
	case "s":
		return self.s, nil
	case "g":
		return self.g, nil
	case "r":
		return self.r, nil
	case "d":
		return self.d, nil
	case "da":
		return self.da, nil
	case "ao":
		return self.ao, nil
	default:
		return "", errIncorrectIchanKey
	}
}

func (self *ichanFiles) set(k, v string) error {
	switch k {
	case "a":
		self.a = v
	case "s":
		self.s = v
	case "g":
		self.g = v
	case "r":
		self.r = v
	case "d":
		self.d = v
	case "da":
		self.da = v
	case "ao":
		self.ao = v
	default:
		return errIncorrectIchanKey
	}

	return nil
}

func (self *ichanFiles) load() (ichans, error) {
	ic := ichans{
		dir:  self.dir,
		name: self.name,
	}

	var (
		err error
		dim = image.Rectangle{}
	)

	// loads an image and sets dim variable if it isn't set already
	f := func(x string) (image.Image, error) {
		var (
			err error
			img image.Image
		)

		if x != "" {
			img, err = loadImageFromFile(x)
			if err != nil {
				return nil, err
			}

			switch {
			case dim.Empty():
				dim = img.Bounds()
				if !(isPowerOfTwo(dim.Dx()) && isPowerOfTwo(dim.Dy())) {
					return nil, errImproperDimensions
				}

			case !dim.Eq(img.Bounds()):
				return nil, errDimensionsDontMatch
			}
		}

		return img, nil
	}

	// ASGR
	ic.a, err = f(self.a)
	if err != nil {
		goto exitWithError
	}

	ic.s, err = f(self.s)
	if err != nil {
		goto exitWithError
	}

	ic.g, err = f(self.g)
	if err != nil {
		goto exitWithError
	}

	ic.r, err = f(self.r)
	if err != nil {
		goto exitWithError
	}

	// DIF and ASG might have different resolutions
	dim = image.Rectangle{}

	ic.d, err = f(self.d)
	if err != nil {
		goto exitWithError
	}

	ic.da, err = f(self.da)
	if err != nil {
		goto exitWithError
	}

	ic.ao, err = f(self.ao)
	if err != nil {
		goto exitWithError
	}

	return ic, nil

exitWithError:
	return ichans{}, err
}

//
// ichans
//

type ichans struct {
	dir  string
	name string
	a    image.Image // Alpha
	s    image.Image // Specular
	g    image.Image // Glow
	r    image.Image // Reflectivity
	d    image.Image // Diffuse
	da   image.Image // Diffuse Alpha
	ao   image.Image // Ambient Occlusion
}

func (self *ichans) compileAsgr() image.Image {
	bounds := image.Rectangle{}

	probeBounds := func(i image.Image) {
		if i != nil && bounds.Empty() {
			bounds = i.Bounds()
		}
	}

	probeBounds(self.a)
	probeBounds(self.s)
	probeBounds(self.g)
	probeBounds(self.r)

	// Return an empty image with default resolution if there's no images to combine
	if bounds.Empty() {
		return image.NewNRGBA(defaultResolution)
	}

	img := image.NewNRGBA(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {

			f := func(i image.Image) uint8 {
				if i == nil {
					return 0
				}

				r, g, b, _ := i.At(x, y).RGBA()
				return i2i8(rgb2gray(r, g, b))
			}

			img.Set(
				x,
				y,
				color.NRGBA{
					f(self.a),
					f(self.s),
					f(self.g),
					f(self.r),
				},
			)
		}
	}

	return img
}

func (self *ichans) compileDif() image.Image {
	bounds := image.Rectangle{}

	probeBounds := func(i image.Image) {
		if i != nil && bounds.Empty() {
			bounds = i.Bounds()
		}
	}

	probeBounds(self.d)
	probeBounds(self.da)
	probeBounds(self.ao)

	// Return an empty image with default resolution if there's no images to combine
	if bounds.Empty() {
		return image.NewRGBA(defaultResolution)
	}

	img := image.NewRGBA(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var (
				imgR = 0.0
				imgG = 0.0
				imgB = 0.0
				imgA = 0.0
			)

			if self.d != nil {
				r, g, b, a := self.d.At(x, y).RGBA()

				imgR = float64(r) / 65535.0
				imgG = float64(g) / 65535.0
				imgB = float64(b) / 65535.0
				imgA = float64(a) / 65535.0
			}

			if self.da != nil {
				r, g, b, _ := self.da.At(x, y).RGBA()
				imgA *= float64(r+g+b) / 196605.0 // 65535 * 3
			}

			if self.ao != nil {
				rao, gao, bao, _ := self.ao.At(x, y).RGBA()
				ao := float64(rao+bao+gao) / 196605.0

				imgR *= ao
				imgG *= ao
				imgB *= ao
				imgA += clamp1(1 - ao)
			}

			img.Set(
				x,
				y,
				color.RGBA{
					f2i8(imgR),
					f2i8(imgG),
					f2i8(imgB),
					f2i8(imgA),
				},
			)
		}
	}

	return img
}
