package useragent

var sortedKnownBrowser = map[string]string{
	"CriOS":           "Chrome",
	"FxiOS":           "Firefox",
	"GSA":             "Google App",
	"Opera":           "Opera",
	"OPR":             "Opera",
	"Edge":            "Edge",
	"Edg":             "Edge",
	"EdgA":            "Edge",
	"EdgiOS":          "Edge",
	"YaBrowser":       "YaBrowser",
	"coc_coc_browser": "Coc Coc",
	"Electron":        "Electron",
	"DuckDuckGo":      "DuckDuckGo",
	"PhantomJS":       "PhantomJS",
	"HeadlessChrome":  "Headless Chrome",
	"Chromium":        "Chromium",
	"Chrome":          "Chrome",
	"Mobile":          "Mobile App",
	"Safari":          "Safari",
}

var sortedKnownEngine = map[string]string{
	"AppleWebKit": "AppleWebKit",
	"Presto":      "Presto",
	"Gecko":       "Gecko",
}

func (p *UserAgent) detectBrowser1(sections []section) {

	//parse engine
	for k, v := range sortedKnownEngine {
		s := getSectionByName(k, sections)
		if s != nil {
			p.browser.Engine = v
			p.browser.EngineVersion = s.version
			break
		}
	}

	for k, v := range sortedKnownBrowser {
		s := getSectionByName(k, sections)
		if s != nil {
			p.browser.Name = v
			p.browser.Version = s.version
			break
		}
	}

}

func getSectionByName(name string, sections []section) *section {
	for _, s := range sections {
		if s.name == name {
			return &s
		}
	}
	return nil
}
