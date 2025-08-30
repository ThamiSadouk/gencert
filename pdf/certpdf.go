package pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"os"
	"path"
	"training/gencert/cert"
)

type PdfSaver struct {
	OutputDir   string // dossier dans lequel il y aura tous les pdf générés
	Orientation string
}

//type Option func(*PdfSaver)
//
//func WithOrientation(o string) Option {
//	return func(s *PdfSaver) { s.orientation = o }
//}

func New(outputDir string, orientation string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir:   outputDir,
		Orientation: orientation,
	}

	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	orient := gofpdf.OrientationLandscape

	//if p.Orientation == "portrait" {
	//	return handlePortraitMode(cert)
	//}

	pdf := gofpdf.New(orient, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// BACKGROUND
	background(pdf)

	// HEADER
	header(pdf, &cert)
	pdf.Ln(30)

	// BODY
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	// Student name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	// Participation
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// Date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	// FOOTER
	footer(pdf)

	// saveFile
	filename := fmt.Sprintf("%s.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to %s\n", path)
	return nil
}

func handleLanscapeMode()
func handlePortraitMode(cert cert.Cert) error {
	fmt.Printf("Handling Portrait mode for %s\n", cert.LabelTitle)
	orient := gofpdf.OrientationPortrait

	pdf := gofpdf.New(orient, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// BACKGROUND
	background(pdf)

	// HEADER
	header(pdf, &cert)
	pdf.Ln(30)

	// BODY
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	// Student name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	// Participation
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// Date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	// FOOTER
	footer(pdf)

	// saveFile
	filename := fmt.Sprintf("%s.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to %s\n", path)
	return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions(
		"img/background.png",
		0,
		0,
		pageWidth,
		pageHeight,
		false,
		opts,
		0,
		"")

}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"
	pdf.ImageOptions(
		filename,
		x+margin,
		20,
		imageWidth,
		0,
		false,
		opts, 0,
		"")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(
		filename,
		x-margin,
		20,
		imageWidth,
		0,
		false,
		opts,
		0,
		"")

	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	imageWidth := 50.0
	pageWidth, pageHeight := pdf.GetPageSize()
	x := pageWidth - imageWidth - 20.0
	y := pageHeight - imageWidth - 10.0
	filename := "img/stamp.png"

	pdf.ImageOptions(
		filename,
		x,
		y,
		imageWidth,
		0,
		false,
		opts, 0,
		"")
}
