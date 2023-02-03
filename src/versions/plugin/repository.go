package plugin

import (
	"bufio"
	"encoding/xml"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"lombok-plugin-action/src/git/github"
	"lombok-plugin-action/src/versions/iu"
	"os"
	"strings"
	"time"
)

type _IdeaPluginName struct {
	Text string `xml:",chardata"`
}
type _IdeaPluginId struct {
	Text string `xml:",chardata"`
}
type _IdeaPluginDescription struct {
	Text string `xml:",chardata"`
}
type _IdeaPluginVersion struct {
	Text string `xml:",chardata"`
}
type _IdeaPluginVendor struct {
	Email string `xml:"email,attr"`
	URL   string `xml:"url,attr"`
}
type _IdeaPluginRating struct {
	Text string `xml:",chardata"`
}
type _IdeaPluginChangeNotes struct {
	Text string `xml:",chardata"`
}
type _IdeaPluginDownloadURL struct {
	Text string `xml:",chardata"`
}
type _IdeaPluginIdeaVersion struct {
	Min        string `xml:"min,attr"`
	Max        string `xml:"max,attr"`
	SinceBuild string `xml:"since-build,attr"`
	UntilBuild string `xml:"until-build,attr"`
}

type _IdeaPlugin struct {
	Downloads   int                    `xml:"downloads,attr"`
	Size        int                    `xml:"size,attr"`
	Date        int64                  `xml:"date,attr"`
	UpdatedDate int64                  `xml:"updatedDate,attr"`
	URL         string                 `xml:"url,attr"`
	Name        _IdeaPluginName        `xml:"name"`
	ID          _IdeaPluginId          `xml:"id"`
	Description _IdeaPluginDescription `xml:"description"`
	Version     _IdeaPluginVersion     `xml:"version"`
	Vendor      _IdeaPluginVendor      `xml:"vendor"`
	Rating      _IdeaPluginRating      `xml:"rating"`
	ChangeNotes _IdeaPluginChangeNotes `xml:"change-notes"`
	DownloadURL _IdeaPluginDownloadURL `xml:"download-url"`
	IdeaVersion _IdeaPluginIdeaVersion `xml:"idea-version"`
}

type _PluginRepositoryFf struct {
	Text string `xml:",chardata"`
}
type _PluginRepositoryCategory struct {
	Name       string        `xml:"name,attr"`
	IdeaPlugin []_IdeaPlugin `xml:"idea-plugin"`
}

type _PluginRepository struct {
	XMLName  xml.Name                  `xml:"plugin-repository"`
	Ff       _PluginRepositoryFf       `xml:"ff"`
	Category _PluginRepositoryCategory `xml:"category"`
}

func CreateRepositoryXml(verTags *priorityqueue.Queue, verInfos *hashmap.Map, sizes *hashmap.Map) (string, error) {
	content := _PluginRepository{
		Ff: _PluginRepositoryFf{Text: "\"Tools Integration\""},
	}
	var categories []_IdeaPlugin

	var item interface{}
	var hasNext bool
	for {
		item, hasNext = verTags.Dequeue()
		if !hasNext {
			break
		}
		verTag := item.(string)
		tmp, _ := verInfos.Get(verTag)
		release := tmp.(iu.IdeaRelease)
		date, _ := time.Parse("2006-01-02", release.Date)
		unix := date.Unix() * 1000
		untilBuild := strings.Split(verTag, ".")[0] + ".*"
		categories = append(categories, _IdeaPlugin{
			Downloads:   0,
			Size:        0,
			Date:        unix,
			UpdatedDate: unix,
			URL:         "",
			Name:        _IdeaPluginName{Text: "Lombok"},
			ID:          _IdeaPluginId{Text: "Lombook Plugin"},
			Description: _IdeaPluginDescription{
				Text: "<![CDATA[<h1>IntelliJ Lombok plugin</h1>\n" +
					"<br/>\n" +
					"<a href=\"https://github.com/mplushnikov/lombok-intellij-plugin\">GitHub</a> |\n" +
					"<a href=\"https://github.com/mplushnikov/lombok-intellij-plugin/issues\">Issues</a> | Donate (\n" +
					"<a href=\"https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=3F9HXD7A2SMCN\">PayPal</a> )\n" +
					"<br/>\n" +
					"<br/>\n" +
					"\n" +
					"<b>A plugin that adds first-class support for Project Lombok</b>\n" +
					"<br/>\n" +
					"<br/>\n" +
					"\n" +
					"<b>Features</b>\n" +
					"<ul>\n" +
					"  <li><a href=\"https://projectlombok.org/features/GetterSetter.html\">@Getter and @Setter</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/experimental/FieldNameConstants\">@FieldNameConstants</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/ToString.html\">@ToString</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/EqualsAndHashCode.html\">@EqualsAndHashCode</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/Constructor.html\">@AllArgsConstructor, @RequiredArgsConstructor and\n" +
					"    @NoArgsConstructor</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/Log.html\">@Log, @Log4j, @Log4j2, @Slf4j, @XSlf4j, @CommonsLog,\n" +
					"    @JBossLog, @Flogger, @CustomLog</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/Data.html\">@Data</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/Builder.html\">@Builder</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/experimental/SuperBuilder\">@SuperBuilder</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/Builder.html#singular\">@Singular</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/Delegate.html\">@Delegate</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/Value.html\">@Value</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/experimental/Accessors.html\">@Accessors</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/experimental/Wither.html\">@Wither</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/With.html\">@With</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/SneakyThrows.html\">@SneakyThrows</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/val.html\">@val</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/var.html\">@var</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/experimental/var.html\">experimental @var</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/experimental/UtilityClass.html\">@UtilityClass</a></li>\n" +
					"  <li><a href=\"https://projectlombok.org/features/configuration.html\">Lombok config system</a></li>\n" +
					"  <li>Code inspections</li>\n" +
					"  <li>Refactoring actions (lombok and delombok)</li>\n" +
					"</ul>\n" +
					"<br/>]]>",
			},
			Version: _IdeaPluginVersion{Text: verTag},
			Vendor: _IdeaPluginVendor{
				URL:   "https://github.com/" + github.REPO + "/release/tag/" + verTag,
				Email: "",
			},
			Rating:      _IdeaPluginRating{Text: "5.0"},
			ChangeNotes: _IdeaPluginChangeNotes{Text: "<![CDATA[]]>"},
			DownloadURL: _IdeaPluginDownloadURL{Text: "https://github.com/" + github.REPO + "/release/download/lombok-" + verTag + ".zip"},
			IdeaVersion: _IdeaPluginIdeaVersion{
				Max:        "n/a",
				Min:        "n/a",
				SinceBuild: verTag,
				UntilBuild: untilBuild,
			},
		})
	}

	content.Category = _PluginRepositoryCategory{
		Name:       "Tools Integration",
		IdeaPlugin: categories,
	}

	marshal, err := xml.Marshal(content)
	if err != nil {
		return "", err
	}

	name := "/tmp/lombok-plugin/plugin-repository"
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}
	write := bufio.NewWriter(file)
	_, err = write.Write(marshal)
	if err != nil {
		return "", err
	}
	err = write.Flush()
	if err != nil {
		return "", err
	}
	file.Close()
	return name, github.CreatePluginRepository(name)
}