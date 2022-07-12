package pdfRendering

import (
	"fmt"
	"log"
	"time"

	"github.com/klippa-app/go-pdfium"
	// "github.com/klippa-app/go-pdfium/multi_threaded"
	// "github.com/klippa-app/go-pdfium/multi_threaded"
	"github.com/klippa-app/go-pdfium/multi_threaded"
	"github.com/klippa-app/go-pdfium/requests"
)

type PdfDocument struct {
	PageNums   int
	PageImages string
}

var pool pdfium.Pool
var instance pdfium.Pdfium

func InitPdfRendering() *PdfDocument {
	return &PdfDocument{
		PageNums:   0,
		PageImages: "",
	}
}

func (pdfDocument *PdfDocument) ConvertPdfToImage(data []byte) {
	pool = multi_threaded.Init(multi_threaded.Config{
		MinIdle:  1, // Makes sure that at least x workers are always available
		MaxIdle:  1, // Makes sure that at most x workers are ever available
		MaxTotal: 1, // Maxium amount of workers in total, allows the amount of workers to grow when needed, items between total max and idle max are automatically cleaned up, while idle workers are kept alive so they can be used directly.
		Command: multi_threaded.Command{
			BinPath: "go",                                                      // Only do this while developing, on production put the actual binary path in here. You should not want the Go runtime on production.
			Args:    []string{"run", "examples/multi_threaded/worker/main.go"}, // This is a reference to the worker package, this can be left empty when using a direct binary path.
		},
	})

	var err error
	instance, err = pool.GetInstance(time.Second * 30)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := instance.OpenDocument(&requests.OpenDocument{
		File: &data,
	})
	if err != nil {
		// return 0, err
	}

	// Always close the document, this will release its resources.
	defer instance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
		Document: doc.Document,
	})

	pageCount, err := instance.FPDF_GetPageCount(&requests.FPDF_GetPageCount{Document: doc.Document})
	if err != nil {
		// return 0, err
	}
	fmt.Println(pageCount)
	// pdf := gopdf.GoPdf{

	// }
	// pdf.Read(data)
	// fmt.Print(pdf.)

	// doc, err := fitz.NewFromMemory(data)
	// if err != nil {
	// 	panic(err)
	// }
	// pdfDocument.PageNums = doc.NumPage()
	// var listPageImage []string

	// for n := 0; n < doc.NumPage(); n++ {
	// 	img, err := doc.ImageDPI(n, 72)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	buf := new(bytes.Buffer)
	// 	err = jpeg.Encode(buf, img, nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	listPageImage[0] = base64.StdEncoding.EncodeToString(buf.Bytes())
	// }
	// result, _ := json.Marshal(listPageImage)
	// pdfDocument.PageImages = string(result)
}
