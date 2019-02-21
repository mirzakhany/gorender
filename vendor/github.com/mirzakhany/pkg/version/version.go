package version

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mirzakhany/pkg/stringsutils"
)

// Settings version settings
type Settings struct {
	CopyRightYear string
	LongHash      string
	ShortHash     string
	CommitDate    string
	CommitCount   string
	BuildDate     string
	Version       string
	ServiceName   string
	CompanyName   string
}

var versionSettings Settings

// SetupVersion for setup version string.
func SetupVersion(settings *Settings) {
	versionSettings = *settings
}

func drawLine(w int) {
	for i := 0; i <= w; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

// GetVersion return version
func GetVersion() string {
	return versionSettings.Version
}

// PrintServiceVersion provide print server engine
func PrintServiceVersion() {

	var lines []string
	var linesLen int

	lines = append(lines, fmt.Sprintf(`Version %s, Compiler: %s %s, Copyright (C) %s %s, Inc.`,
		versionSettings.Version,
		runtime.Compiler,
		runtime.Version(),
		versionSettings.CopyRightYear,
		versionSettings.CompanyName))

	lines = append(lines, "Commit Hash: "+versionSettings.LongHash)
	lines = append(lines, "Commit ShortHash:"+versionSettings.ShortHash)
	lines = append(lines, "Commit Count:"+versionSettings.CommitCount)
	lines = append(lines, "Commit Date:"+versionSettings.CommitDate)
	lines = append(lines, "Build Date:"+versionSettings.BuildDate)
	linesLen = stringsutils.MaxLen(lines)

	drawLine(linesLen)
	for _, sl := range lines {
		fmt.Println(sl)
	}
	drawLine(linesLen)
}

// HeaderVersionMiddleware : add version on header.
func HeaderVersionMiddleware() gin.HandlerFunc {
	// Set out header value for each response
	return func(c *gin.Context) {
		k := fmt.Sprintf("X-%s-VERSION", strings.ToUpper(versionSettings.ServiceName))
		c.Header(k, versionSettings.Version)
		c.Next()
	}
}
