package catalog

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/go-gota/gota/dataframe"
)

type paifuList struct {
	south3 []paifu
	south4 []paifu
	east3  []paifu
	east4  []paifu
}

type paifu struct {
	Date     string `json:"date"`
	GameType string `json:"game_type"`
	LogID    string `json:"log_id"`
}

type catalog struct {
	inputPath  string
	filesPaths []string
	data       *paifuList
}

const DEST_PATH = "input/catalog"

func NewCatalog(iPath string) *catalog {
	return &catalog{
		inputPath:  iPath,
		filesPaths: make([]string, 0),
		data: &paifuList{
			south3: make([]paifu, 0),
			south4: make([]paifu, 0),
			east3:  make([]paifu, 0),
			east4:  make([]paifu, 0),
		},
	}
}

func (c *catalog) Archive() {
	files, _ := c.GetFiles(c.inputPath)
	c.filesPaths = append(c.filesPaths, files...)
	for _, v := range c.filesPaths {
		c.ReadSingleFile(v)
	}

	east4csv, err := os.OpenFile(path.Join(DEST_PATH, "east4.csv"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer east4csv.Close()

	fmt.Println(len(c.data.south4))
	df := dataframe.LoadStructs(c.data.south4)
	if err := df.WriteCSV(east4csv); err != nil {
		fmt.Println(err)
	}
}

func (c *catalog) GetFiles(inputPath string) ([]string, error) {
	var filesPaths []string
	filepath.Walk(inputPath, func(filePath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		dir := path.Dir(filePath)
		ext := path.Ext(filePath)

		if ext != ".html" {
			return nil
		}

		file := filepath.Join(dir, info.Name())
		filesPaths = append(filesPaths, file)
		return nil
	})

	return filesPaths, nil
}

func (c *catalog) ReadSingleFile(filePath string) {
	filename := path.Base(filePath)
	re := regexp.MustCompile(`^scc(.*)\.html$`)
	fileDate := re.ReplaceAllString(filename, "$1")
	// fileTime, _ := time.Parse(fileDate, "20060102")
	// yearS := string(rune(fileTime.Year()))

	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// ReadLine is a low-level line-reading primitive.
		// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		line := string(bytes)
		lineContentList := strings.Split(line, "|")
		min := strings.TrimSpace(lineContentList[0])
		sec := strings.TrimSpace(lineContentList[1])
		// gameType := strings.TrimSpace(lineContentList[2])

		doc, err := htmlquery.Parse(strings.NewReader(strings.TrimSpace(lineContentList[3])))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		link := htmlquery.SelectAttr(htmlquery.FindOne(doc, "//a/@href"), "href")
		u, _ := url.Parse(link)
		date, _ := time.Parse("20060102 15:04:05", fileDate+" "+min+":"+sec)

		var data = paifu{}
		data.Date = date.Format(time.RFC3339)
		data.LogID = u.Query().Get("log")
		data.GameType = strings.Split(data.LogID, "-")[1]
		switch data.GameType {
		case "00b1":
			c.data.east3 = append(c.data.east3, data)
		case "00f1":
			c.data.east3 = append(c.data.south3, data)

		case "00b9":
			c.data.south3 = append(c.data.south3, data)
		case "00f9":
			c.data.south3 = append(c.data.south3, data)

		case "00a1":
			c.data.east4 = append(c.data.east4, data)
		case "00e1":
			c.data.east4 = append(c.data.east4, data)

		case "00a9":
			c.data.south4 = append(c.data.south4, data)
		case "00e9":
			c.data.south4 = append(c.data.south4, data)
		}
	}
}
