package pdfRendering

import (
	"bytes"
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

type PdfDocument struct {
	PageNums   int
	PageImages string
}

func InitPdfRendering() *PdfDocument {
	return &PdfDocument{
		PageNums:   0,
		PageImages: "",
	}
}

func (pdfDocument *PdfDocument) ConvertPdfToImage(data []byte) {
	ctx, e := pdfcpu.Read(bytes.NewReader(data), &pdfcpu.Configuration{})
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(ctx)
	// p , _ :=ctx.ExtractPage(0)
	// p.Crop(p.LinearizationObjs,&pdfcpu.ImportBox{})
	// pdfcpu.WMImage
}
