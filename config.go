package logger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	
	"github.com/jinzhu/configor"
)

var Config = struct {
	Log struct {
		FilePath      string `default:"log/app"`
		FileExtension string `default:".log"`
		MaxSize       int    `default:"100"` // megabytes
		MaxBackups    int    `default:"7"`
		MaxAge        int    `default:"30"` // days
		Compress      bool   `default:"false"`
	}
}{}

func init() {
	homePath := os.Getenv("GWAKUH_LOGGER_APP_HOME")
	filePath := fmt.Sprintf("%s%c%s%c%s", homePath, os.PathSeparator, "etc", os.PathSeparator, "logger.yml")
	
	err := configor.Load(&Config, filePath)
	if err != nil {
		panic(err)
	}
	
	if !strings.HasPrefix(Config.Log.FilePath, strconv.QuoteRune(os.PathSeparator)) {
		Config.Log.FilePath = fmt.Sprintf("%s%c%s", homePath, os.PathSeparator, Config.Log.FilePath); 
	}
}