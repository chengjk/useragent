package useragent

import "github.com/emirpasic/gods/maps/linkedhashmap"

var sortedKnownBrowser = linkedhashmap.New()
var sortedKnownEngine = linkedhashmap.New()

func init() {
	sortedKnownBrowser.Put("CriOS", "Chrome")
	sortedKnownBrowser.Put("FxiOS", "Firefox")
	sortedKnownBrowser.Put("GSA", "Google App")
	sortedKnownBrowser.Put("Opera", "Opera")
	sortedKnownBrowser.Put("OPR", "Opera")
	sortedKnownBrowser.Put("Edge", "Edge")
	sortedKnownBrowser.Put("Edg", "Edge")
	sortedKnownBrowser.Put("EdgA", "Edge")
	sortedKnownBrowser.Put("EdgiOS", "Edge")
	sortedKnownBrowser.Put("YaBrowser", "YaBrowser")
	sortedKnownBrowser.Put("coc_coc_browser", "Coc Coc")
	sortedKnownBrowser.Put("Electron", "Electron")
	sortedKnownBrowser.Put("DuckDuckGo", "DuckDuckGo")
	sortedKnownBrowser.Put("PhantomJS", "PhantomJS")
	sortedKnownBrowser.Put("HeadlessChrome", "Headless Chrome")
	sortedKnownBrowser.Put("Chromium", "Chromium")
	sortedKnownBrowser.Put("Chrome", "Chrome")
	sortedKnownBrowser.Put("Mobile", "Mobile App")
	sortedKnownBrowser.Put("Safari", "Safari")

	sortedKnownEngine.Put("AppleWebKit", "AppleWebKit")
	sortedKnownEngine.Put("Presto", "Presto")
	sortedKnownEngine.Put("Gecko", "Gecko")

}

func (p *UserAgent) detectBrowser1(sections []section) {

	//parse engine
	for _, key := range sortedKnownEngine.Keys() {
		s := getSectionByName(key.(string), sections)
		if s != nil {
			value, _ := sortedKnownEngine.Get(key)
			p.browser.Engine = value.(string)
			p.browser.EngineVersion = s.version
			break
		}
	}

	for _, key := range sortedKnownBrowser.Keys() {
		s := getSectionByName(key.(string), sections)
		if s != nil {
			value, _ := sortedKnownEngine.Get(key)
			p.browser.Name = value.(string)
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
