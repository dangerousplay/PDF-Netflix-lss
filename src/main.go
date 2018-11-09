package main

import (
	"bytes"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/jung-kurt/gofpdf"
	"github.com/wcharczuk/go-chart"
	"io"
	"time"
)

type Filme struct {
	nome          string
	valor         float64
	anoLancamento string
	sinopse       string
	genero        string
}

type Alocacao struct {
	id          int
	dataInicial uint64
	dataFinal   uint64
	dataEntrega uint64
	filme       Filme
}

//export generatePDF
func generatePDF(ids []int, dataInicial []uint64, dataFinal []uint64, dataEntrega []uint64, valores []float64) ([]byte, error) {
	alocacoes := arraylist.New()

	for i, v := range ids {
		aloc := Alocacao{
			id:          v,
			dataInicial: dataInicial[i],
			dataEntrega: dataEntrega[i],
			dataFinal:   dataFinal[i],
		}

		alocacoes.Add(aloc)
	}

	var atrasadas = alocacoes.Select(func(index int, value interface{}) bool {
		aloc, _ := value.(Alocacao)
		return aloc.dataFinal < aloc.dataEntrega
	})


	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("Relatório mensal das vendas realizadas em %s", time.Now().Format("2006-01-02")))
	pdf.SetAuthor("Netflix", true)
	pdf.SetSubject("Relatório mensal", true)
	pdf.SetCreationDate(time.Now())

	sbc := chart.BarChart{
		Title:      "Alocações realizadas",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		XAxis:    chart.StyleShow(),
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Bars: []chart.Value{
			{Value: , Label: "Atrasadas"},
			{Value: 5.25, Label: "America"},
			{Value: 4.88, Label: "Brasil"},
			{Value: 4.74, Label: "USA"},
			{Value: 3.22, Label: "CNA"},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err2 := sbc.Render(chart.PNG, buffer)

	if err2 != nil {
		panic(err2)
	}

	image := pdf.RegisterImageOptionsReader("grafico", gofpdf.ImageOptions{ImageType: "png", ReadDpi: true}, io.Reader(buffer))

	pdf.Image("grafico", 0, 10, image.Width()/3, image.Height()/3, false, "", 0, "")

	pdfBuffer := bytes.NewBuffer([]byte{})

	err := pdf.Output(pdfBuffer)

	return pdfBuffer.Bytes(), err
}

func main() {
	fmt.Print("hello")
	/*pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")


	sbc := chart.BarChart{
		Title:      "Netflix Gráfico Exemplo",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		XAxis:    chart.StyleShow(),
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Bars: []chart.Value{
			{Value: 5.25, Label: "America"},
			{Value: 4.88, Label: "Brasil"},
			{Value: 4.74, Label: "USA"},
			{Value: 3.22, Label: "CNA"},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err2 := sbc.Render(chart.PNG, buffer)

	image := pdf.RegisterImageOptionsReader("grafico", gofpdf.ImageOptions{ ImageType: "png", ReadDpi: true}, io.Reader(buffer))

	pdf.Image("grafico", 0, 10, image.Width()/3, image.Height()/3, false, "", 0, "")

	image.Height()

	err := pdf.OutputFileAndClose("hello.pdf")

	fmt.Print(err, err2)*/

}
