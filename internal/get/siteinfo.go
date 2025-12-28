package get

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

var (
	categories = map[string]string{
		"1":  "CMS",
		"2":  "Message boards",
		"3":  "Database managers",
		"4":  "Documentation",
		"5":  "Widgets",
		"6":  "Ecommerce",
		"7":  "Photo galleries",
		"8":  "Wikis",
		"9":  "Hosting panels",
		"10": "Analytics",
		"11": "Blogs",
		"12": "JavaScript frameworks",
		"13": "Issue trackers",
		"14": "Video players",
		"15": "Comment systems",
		"16": "Security",
		"17": "Font scripts",
		"18": "Web frameworks",
		"19": "Miscellaneous",
		"20": "Editors",
		"21": "LMS",
		"22": "Web servers",
		"23": "Caching",
		"24": "Rich text editors",
		"25": "JavaScript graphics",
		"26": "Mobile frameworks",
		"27": "Programming languages",
		"28": "Operating systems",
		"29": "Search engines",
		"30": "Webmail",
		"31": "CDN",
		"32": "Marketing automation",
		"33": "Web server extensions",
		"34": "Databases",
		"35": "Maps",
		"36": "Advertising",
		"37": "Network devices",
		"38": "Media servers",
		"39": "Webcams",
		"41": "Payment processors",
		"42": "Tag managers",
		"44": "CI",
		"45": "Control systems",
		"46": "Remote access",
		"47": "Development",
		"48": "Network storage",
		"49": "Feed readers",
		"50": "DMS",
		"51": "Page builders",
		"52": "Live chat",
		"53": "CRM",
		"54": "SEO",
		"55": "Accounting",
		"56": "Cryptominers",
		"57": "Static site generator",
		"58": "User onboarding",
		"59": "JavaScript libraries",
		"60": "Containers",
		"62": "PaaS",
		"63": "IaaS",
		"64": "Reverse proxies",
		"65": "Load balancers",
		"66": "UI frameworks",
		"67": "Cookie compliance",
		"68": "Accessibility",
		"69": "Authentication",
		"70": "SSL/TLS certificate authorities",
		"71": "Affiliate programs",
		"72": "Appointment scheduling",
		"73": "Surveys",
		"74": "A/B Testing",
		"75": "Email",
		"76": "Personalisation",
		"77": "Retargeting",
		"78": "RUM",
		"79": "Geolocation",
		"80": "WordPress themes",
		"81": "Shopify themes",
		"82": "Drupal themes",
		"83": "Browser fingerprinting",
		"84": "Loyalty & rewards",
		"85": "Feature management",
		"86": "Segmentation",
		"87": "WordPress plugins",
		"88": "Hosting",
		"89": "Translation",
		"90": "Reviews",
		"91": "Buy now pay later",
		"92": "Performance",
		"93": "Reservations & delivery",
		"94": "Referral marketing",
		"95": "Digital asset management",
		"96": "Content curation",
		"97": "Customer data platform",
		"98": "Cart abandonment",
		"99": "Shipping carriers",
		"100": "Shopify apps",
		"101": "Recruitment & staffing",
		"102": "Returns",
		"103": "Livestreaming",
		"104": "Ticket booking",
		"105": "Augmented reality",
		"106": "Cross border ecommerce",
		"107": "Fulfilment",
		"108": "Ecommerce frontends",
		"109": "Domain parking",
		"110": "Form builders",
		"111": "Fundraising & donations",
	}
	technologies = map[string][]string{
		"Acquia Cloud Platform": {"88", "62"},
		"Amazon EC2":            {"63"},
		"Apache":                {"22"},
		"Cloudflare":            {"31"},
		"Drupal":                {"1", "80", "82"},
		"PHP":                   {"27"},
		"Percona":               {"34"},
		"React":                 {"12"},
		"Varnish":               {"23"},
		"DigitalOcean":          {"63"},
		"CloudFront":            {"31"},
	}
)

func GetSiteInfo(url string) (string, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get url: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body: %w", err)
	}

	wappalyzerClient, err := wappalyzer.New()
	if err != nil {
		return "", fmt.Errorf("failed to create wappalyzer client: %w", err)
	}

	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	if len(fingerprints) == 0 {
		return "No technologies found", nil
	}

	categorized := make(map[string][]string)
	for tech := range fingerprints {
		catIDs, ok := technologies[tech]
		if !ok {
			continue
		}

		for _, catID := range catIDs {
			catName, ok := categories[catID]
			if !ok {
				continue
			}
			categorized[catName] = append(categorized[catName], tech)
		}
	}

	// Sort categories for consistent output
	sortedCategories := make([]string, 0, len(categorized))
	for catName := range categorized {
		sortedCategories = append(sortedCategories, catName)
	}
	sort.Strings(sortedCategories)

	var builder strings.Builder
	for _, catName := range sortedCategories {
		builder.WriteString(fmt.Sprintf("%s:\n", catName))
		for _, tech := range categorized[catName] {
			builder.WriteString(fmt.Sprintf("- %s\n", tech))
		}
		builder.WriteString("\n")
	}

	return strings.TrimSpace(builder.String()), nil
}
