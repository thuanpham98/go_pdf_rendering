package pdfRendering

import (
	"image/png"
	"log"
	"os"
	"time"

	"github.com/klippa-app/go-pdfium"
	// "github.com/klippa-app/go-pdfium/multi_threaded"
	// "github.com/klippa-app/go-pdfium/multi_threaded"

	"github.com/klippa-app/go-pdfium/requests"
	"github.com/klippa-app/go-pdfium/single_threaded"
)

type PdfDocument struct {
	pool       pdfium.Pool
	instance   pdfium.Pdfium
	PageNums   int
	PageImages string
}

func InitPdfRendering() *PdfDocument {
	_pool := single_threaded.Init(single_threaded.Config{})

	_instance, err := _pool.GetInstance(time.Second * 30)
	if err != nil {
		log.Fatal(err)
	}
	return &PdfDocument{
		pool:       _pool,
		instance:   _instance,
		PageNums:   0,
		PageImages: "",
	}
}

func (pdfDocument *PdfDocument) ConvertPdfToImage(data []byte) error {

	doc, _ := pdfDocument.instance.OpenDocument(&requests.OpenDocument{File: &data})
	log.Fatal(doc.Document)
	defer pdfDocument.instance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
		Document: doc.Document,
	})
	res, err := pdfDocument.instance.RenderPageInDPI(&requests.RenderPageInDPI{
		DPI: 72,
		Page: requests.Page{
			ByIndex: &requests.PageByIndex{
				Document: doc.Document,
				Index:    0,
			},
		},
	})
	if err != nil {
		return err
	}

	// Write the output to a file.
	f, err := os.Create("./example.pdf.png")
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, res.Result.Image)
	if err != nil {
		return err
	}

	return nil
}
