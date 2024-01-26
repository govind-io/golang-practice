package price

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"price_calculators/utils"
	"strconv"
)

type taxPrices map[string]float64

type TaxIncludedPricesJob struct {
	TaxRate           float64
	InputPrice        []float64 `json:"-"`
	TaxIncludedPrices taxPrices
}

func WritesJson(data interface{}, path string) (bool, error) {
	file, err := os.Create(fmt.Sprintf("./db/%v", path))
	if err != nil {
		return false, errors.New("unable to write to file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	file.Close()
	if err != nil {
		return false, errors.New("unable to write to file")
	}

	return true, nil

}

func LoadPrices() ([]float64, error) {
	inputPrice := []float64{}

	if file, err := os.Open("./db/inputPrices.txt"); err != nil {
		return inputPrice, err
	} else {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			parsedInputPrice, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				continue
			}
			inputPrice = append(inputPrice, float64(parsedInputPrice))
		}

		err = scanner.Err()

		file.Close()

		if err != nil {
			fmt.Println("Some of the prices were invalid in the input file")
		}

	}

	return inputPrice, nil
}

func NewTaxIncludedPricesJob(taxRate float64) *TaxIncludedPricesJob {
	inputPrice, err := LoadPrices()

	if err != nil {
		return &TaxIncludedPricesJob{}
	}

	taxIncludedPrices := &TaxIncludedPricesJob{
		TaxRate:    taxRate,
		InputPrice: inputPrice,
	}

	taxIncludedPrices.PopulateTaxIncludedPrice()

	WritesJson(taxIncludedPrices, fmt.Sprintf("result %v.json", taxIncludedPrices.TaxRate))

	return taxIncludedPrices
}

func (taxIncludedPrice *TaxIncludedPricesJob) PopulateTaxIncludedPrice() {
	inputPrices := taxIncludedPrice.InputPrice
	taxRatePrices := make(taxPrices, len(inputPrices))
	utils.Map(&inputPrices, func(price float64, _ int) {
		formattedPrice, _ := strconv.ParseFloat((fmt.Sprintf("%.2f", (taxIncludedPrice.TaxRate+1)*price)), 64)
		taxRatePrices[fmt.Sprintf("%.0f", price)] = formattedPrice
	})

	taxIncludedPrice.TaxIncludedPrices = taxRatePrices
}
