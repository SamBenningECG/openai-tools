package main

import (
	"context"
	"encoding/csv"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Suburban-Street-Trading/openai-tools/openai"
)

type ResultRow struct {
	ParsedIdentifier string
	WhiteBackground  bool
	Color            string
	HasText          bool
	HasWatermark     bool
	ImgUrl           string
}

func main() {

	ctx := context.Background()

	apiKey := readApiKey()

	client := openai.NewClient(apiKey)

	urls := readImgUrls(ctx)

	textPrompt := "Does this product image have a white, plain background? Does it feature text? Does it have a watermark? What is the color of the product? Provide your answer in the following format: \n white_background: [true/false] | color: [insert color here] | has_text: [true/false] | has_watermark: [true/false]"

	var resultRows []ResultRow

	for _, url := range urls {

		res, err := client.LookAtImage(ctx, url, textPrompt)
		if err != nil {
			os.Exit(1)
		}

		resultRows = append(resultRows, createResult(res, url, getDobaIdentifier(url)))

	}

	writeToFile(resultRows)

}

func getDobaIdentifier(url string) string {

	parts := strings.Split(url, "/")

	identifierWithExtension := parts[len(parts)-1]

	return strings.Split(identifierWithExtension, ".")[0]

}

func createResult(aiResponse, imgURL, identifier string) ResultRow {

	parts := strings.Split(aiResponse, "|")

	if len(parts) < 4 {
		return ResultRow{ImgUrl: imgURL}
	}

	// whiteBackgroundText, colorText, hasTextText, hasWaterMarkText := parts[0], parts[1], parts[2], parts[3]

	whiteBackgroundText := parseAiResponseField(parts[0])
	colorText := parseAiResponseField(parts[1])
	hasTextText := parseAiResponseField(parts[2])
	hasWaterMarkText := parseAiResponseField(parts[3])

	whiteBackground, _ := strconv.ParseBool(whiteBackgroundText)
	hasText, _ := strconv.ParseBool(hasTextText)
	hasWaterMark, _ := strconv.ParseBool(hasWaterMarkText)

	return ResultRow{
		ParsedIdentifier: identifier,
		WhiteBackground:  whiteBackground,
		Color:            colorText,
		HasText:          hasText,
		HasWatermark:     hasWaterMark,
		ImgUrl:           imgURL,
	}

}

func parseAiResponseField(input string) string {

	pattern := `:\s*\[?([^\]]+)\]?`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	} else {
		return input
	}

}

func readImgUrls(ctx context.Context) []string {

	var urls []string

	file, err := os.Open("image_urls.csv")
	if err != nil {
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		os.Exit(1)
	}

	records = records[1:]

	for _, rec := range records {
		urls = append(urls, rec[0])
	}

	return urls

}

func writeToFile(results []ResultRow) {

	file, err := os.Create("results.csv")
	if err != nil {
		os.Exit(1)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := [][]string{
		{
			"identifier",
			"color",
			"has_text",
			"has_watermark",
			"has_white_background",
			"img_url",
		},
	}

	for _, row := range results {

		data = append(data, []string{
			row.ParsedIdentifier,
			row.Color,
			strconv.FormatBool(row.HasText),
			strconv.FormatBool(row.HasWatermark),
			strconv.FormatBool(row.WhiteBackground),
			row.ImgUrl,
		})
	}

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			os.Exit(1)
		}
	}
}

func readApiKey() string {

	file, err := os.ReadFile("apikey.txt")
	if err != nil {
		os.Exit(1)
	}

	return string(file)
}
