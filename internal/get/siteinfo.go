package get

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/kahnwong/swissknife/configs/color"
	"github.com/kahnwong/swissknife/internal/utils"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
	"github.com/rs/zerolog/log"
)

// categoryMapping maps wappalyzer category names to our display categories
var categoryMapping = map[string]string{
	"CMS":                             "CMS",
	"Ecommerce":                       "Ecommerce",
	"Analytics":                       "Analytics",
	"Blogs":                           "CMS",
	"JavaScript frameworks":           "JavaScript Frameworks",
	"Web frameworks":                  "Web Frameworks",
	"Web servers":                     "Web Servers",
	"CDN":                             "CDN",
	"Caching":                         "Caching",
	"Programming languages":           "Programming Languages",
	"Operating systems":               "Operating Systems",
	"Databases":                       "Databases",
	"Message boards":                  "CMS",
	"Payment processors":              "Payment",
	"Security":                        "Security",
	"Tag managers":                    "Analytics",
	"Marketing automation":            "Marketing",
	"Advertising":                     "Marketing",
	"SEO":                             "Marketing",
	"Live chat":                       "Communication",
	"Font scripts":                    "UI",
	"Mobile frameworks":               "Mobile",
	"PaaS":                            "Cloud",
	"IaaS":                            "Cloud",
	"Reverse proxies":                 "Infrastructure",
	"Load balancers":                  "Infrastructure",
	"Web server extensions":           "Web Servers",
	"JavaScript libraries":            "JavaScript Frameworks",
	"UI frameworks":                   "UI",
	"Hosting panels":                  "Hosting",
	"Comment systems":                 "Communication",
	"Widgets":                         "UI",
	"Video players":                   "Media",
	"Maps":                            "Maps",
	"Remote access":                   "Infrastructure",
	"Network devices":                 "Infrastructure",
	"Control systems":                 "Infrastructure",
	"Static site generator":           "Static Site Generator",
	"Development":                     "Development",
	"CI":                              "Development",
	"Page builders":                   "CMS",
	"Wikis":                           "CMS",
	"Documentation":                   "Documentation",
	"Issue trackers":                  "Development",
	"Photo galleries":                 "Media",
	"LMS":                             "Education",
	"Rich text editors":               "UI",
	"Editors":                         "UI",
	"Search engines":                  "Search",
	"Feed readers":                    "Content",
	"DMS":                             "Content",
	"Webmail":                         "Communication",
	"CRM":                             "Business",
	"Accounting":                      "Business",
	"User onboarding":                 "UI",
	"A/B Testing":                     "Marketing",
	"Accessibility":                   "Accessibility",
	"Authentication":                  "Security",
	"Build/Task runners":              "Development",
	"Containers":                      "Infrastructure",
	"Cookie compliance":               "Legal",
	"Cryptominers":                    "Cryptocurrency",
	"Database managers":               "Databases",
	"Date & time pickers":             "UI",
	"DevOps":                          "Development",
	"Error logging":                   "Monitoring",
	"Feature management":              "Development",
	"Geolocation":                     "Location",
	"GraphQL":                         "API",
	"Headless CMS":                    "CMS",
	"Media servers":                   "Media",
	"Miscellaneous":                   "Miscellaneous",
	"Monitoring":                      "Monitoring",
	"Network storage":                 "Storage",
	"Performance":                     "Performance",
	"Personalisation":                 "Marketing",
	"Privacy":                         "Privacy",
	"Proxies":                         "Infrastructure",
	"RUM":                             "Performance",
	"Retargeting":                     "Marketing",
	"SSL/TLS certificate authorities": "Security",
	"Search engine crawlers":          "SEO",
	"Server-side rendering":           "Web Frameworks",
	"Webcams":                         "Media",
}

func GetSiteInfo(args []string) {
	// set URL
	url := utils.SetURL(args)
	fmt.Println(url)

	// fetch site
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to fetch URL: %s", url)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close response body")
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read response body")
	}

	// init wappalyzer
	wappalyzerClient, err := wappalyzer.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize wappalyzer")
	}

	appsInfo := wappalyzerClient.FingerprintWithInfo(resp.Header, data)

	// categorize tech info
	categoryTechs := make(map[string][]string)
	for appName, info := range appsInfo {
		cleanName := strings.Split(appName, ":")[0]

		for _, wapCat := range info.Categories {
			displayCat := categoryMapping[wapCat]
			if displayCat == "" {
				displayCat = wapCat // Use original if no mapping exists
			}

			techs := categoryTechs[displayCat]
			found := false
			for _, t := range techs {
				if t == cleanName {
					found = true
					break
				}
			}
			if !found {
				categoryTechs[displayCat] = append(techs, cleanName)
			}
		}
	}

	if len(categoryTechs) == 0 {
		fmt.Println("No technologies detected")
	}

	categories := make([]string, 0, len(categoryTechs))
	for cat := range categoryTechs {
		categories = append(categories, cat)
	}
	sort.Strings(categories)

	// display technologies
	for _, category := range categories {
		techs := categoryTechs[category]
		sort.Strings(techs)

		fmt.Printf("%s:\n", color.Green(category))
		for _, tech := range techs {
			fmt.Printf("- %s\n", tech)
		}
	}
}
