package blog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func GenerateHtmlIndiv(id, workspace string) error {
	url := fmt.Sprintf("https://note.com/api/v3/notes/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get note by id: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get note by id: %s", resp.Status)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	htmlContent, err := parseHtmlFromByte(body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %v", err)
	}

	formattedHtmlContent, err := formatHtml(htmlContent)
	if err != nil {
		return fmt.Errorf("failed to format HTML: %v", err)
	}

	err = generateFile(id, formattedHtmlContent, workspace)
	if err != nil {
		return fmt.Errorf("failed to generate file: %v", err)
	}

	return nil
}

func parseHtmlFromByte(byteContens []byte) (string, error) {
	var data map[string]map[string]interface{}

	if err := json.Unmarshal(byteContens, &data); err != nil {
		return "", err
	}

	htmlContent := html.UnescapeString(data["data"]["body"].(string))

	formattedContent, err := formatHtml(htmlContent)

	if err != nil {
		return "", err
	}

	return formattedContent, nil
}

func formatHtml(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))

	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := html.Render(&buf, doc); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func generateFile(filename, content, workspace string) error {
	outputFile := filepath.Join(workspace, filename+".html")

	return os.WriteFile(outputFile, []byte(content), 0644)
}
