package api

import (
	"encoding/json"
	"fmt"
	"time"
)

type Localized map[string]string

func (l *Localized) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, (*map[string]string)(l))
	if err != nil {
		slice := make([]string, 0)
		if err := json.Unmarshal(data, &slice); err == nil && len(slice) == 0 {
			return nil
		}
	}

	return err
}

type Manga struct {
	Result   string
	Response string
	Data     MangaData
}

type MangaData struct {
	ID         string
	Type       string
	Attributes struct {
		Title                          Localized
		AltTitles                      []Localized
		Description                    Localized
		IsLocked                       bool
		Links                          map[string]string
		OriginalLanguage               string
		LastVolume                     string
		LastChapter                    string
		PublicationDemographic         string
		Status                         string
		Year                           int
		ContentRating                  string
		Tags                           Relationships
		State                          string
		ChapterNumbersResetOnNewVolume bool
		CreatedAt                      time.Time
		UpdatedAt                      time.Time
		Version                        int
	}
	Relationships Relationships
}

type ChapterList struct {
	Result   string
	Response string
	Data     []ChapterData
	Limit    int
	Offset   int
	Total    int
}

type Chapter struct {
	Result   string
	Response string
	Data     ChapterData
}

type ChapterData struct {
	ID         string
	Type       string
	Attributes struct {
		Volume             string
		Chapter            string
		Title              string
		TranslatedLanguage string
		Uploader           string
		ExternalURL        string
		PublishAt          time.Time
		ReadableAt         time.Time
		CreatedAt          time.Time
		UpdatedAt          time.Time
		Pages              int
		Version            int
	}
	Relationships Relationships
}

type CoverList struct {
	Result   string
	Response string
	Data     []CoverData
	Limit    int
	Offset   int
	Total    int
}

type Cover struct {
	Result   string
	Response string
	Data     CoverData
}

type CoverData struct {
	ID         string
	Type       string
	Attributes struct {
		Description string
		Volume      string
		FileName    string
		CreatedAt   string
		UpdatedAt   string
		Version     int
	}
	Relationships Relationships
}

type AuthorList struct {
	Result   string
	Response string
	Data     []AuthorData
	Limit    int
	Offset   int
	Total    int
}

type AuthorData struct {
	ID         string
	Type       string
	Attributes struct {
		Name      string
		ImageUrl  string
		Biography Localized
		CreatedAt time.Time
		UpdatedAt time.Time
		Twitter   string
		Pixiv     string
		MelonBook string
		FanBox    string
		Booth     string
		NicoVideo string
		Skeb      string
		Fantia    string
		Tumblr    string
		Youtube   string
		Weibo     string
		Naver     string
		Website   string
		Version   int
	}
	Relationships Relationships
}

type GroupList struct {
	Result   string
	Response string
	Data     []GroupData
	Limit    int
	Offset   int
	Total    int
}

type Group struct {
	Result   string
	Response string
	Data     GroupData
}

type GroupData struct {
	ID         string
	Type       string
	Attributes struct {
		Name             string
		AltNames         []Localized
		Description      string
		Leader           Relationship
		Members          Relationships
		FocusedLanguages []string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		Website          string
		IRCServer        string
		IRCChannel       string
		Discord          string
		ContactEmail     string
		Twitter          string
		MangaUpdates     string
		Locked           bool
		Verified         bool
		Official         bool
		Version          int
		PublishDelay     int
		Inactive         bool
	}
	Relationships Relationships
}

type IDMappingList struct {
	Result   string
	Response string
	Data     []IDMappingData
	Limit    int
	Offset   int
	Total    int
}

type IDMappingData struct {
	ID         string
	Type       string
	Attributes struct {
		LegacyID int
		NewID    string
		Type     string
	}
	Relationships Relationships
}

type AtHome struct {
	Result  string
	BaseURL string
	Chapter struct {
		Hash      string
		Data      []string
		DataSaver []string
	}
}

type Relationships struct {
	Manga      []string
	Chapter    []string
	Author     []string
	Artist     []string
	Group      []string
	Tag        []string
	User       []string
	CustomList []string
	CoverArt   []string
	Leader     []string
	Member     []string
}

func (rs *Relationships) UnmarshalJSON(data []byte) error {
	parsed := make([]Relationship, 0)
	if err := json.Unmarshal(data, &parsed); err != nil {
		return err
	}

	for _, r := range parsed {
		switch r.Type {
		case "manga":
			rs.Manga = append(rs.Manga, r.ID)
		case "chapter":
			rs.Chapter = append(rs.Chapter, r.ID)
		case "author":
			rs.Author = append(rs.Author, r.ID)
		case "artist":
			rs.Artist = append(rs.Artist, r.ID)
		case "scanlation_group":
			rs.Group = append(rs.Group, r.ID)
		case "tag":
			rs.Tag = append(rs.Tag, r.ID)
		case "user":
			rs.User = append(rs.User, r.ID)
		case "custom_list":
			rs.CustomList = append(rs.CustomList, r.ID)
		case "cover_art":
			rs.CoverArt = append(rs.CoverArt, r.ID)
		case "leader":
			rs.Leader = append(rs.Leader, r.ID)
		case "member":
			rs.Member = append(rs.Member, r.ID)
		default:
			return fmt.Errorf("unsupported relationship: %v", r.Type)
		}
	}

	return nil
}

type Relationship struct {
	ID         string
	Type       string
	Attributes map[string]interface{}
}

type Errors struct {
	Errors []ErrorData
	Result string
}

type ErrorData struct {
	Context string
	Detail  string
	ID      string
	Status  int
	Title   string
}
