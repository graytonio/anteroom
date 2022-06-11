package lib

import (
	"fmt"
	"os"
	"regexp"
)

type SrcParse struct {
	Site   string
	User   string
	Name   string
	Ref    string
	Url    string
	Ssh    string
	Subdir string
}

var supported = []string{"github.com"}

func parse_repo(str string) (*SrcParse, *ChamberError) {
	regex := regexp.MustCompile(`^(?:(?:https:\/\/)?(?P<Domain1>[^:/]+\.[^:/]+)\/|git@(?P<Domain2>[^:/]+)[:/]|(?P<Domain3>[^/]+):)?(?P<User>[^/\s]+)\/(?P<Name>[^/\s#]+)(?:(?P<Subdir>(?:\/[^/\s#]+)+))?(?:\/)?(?:#(?P<Ref>.+))?`)
	remove_ext := regexp.MustCompile(`\.git$`)

	matches := regex.FindStringSubmatch(str)
	var parseMap = make(map[string]string)
	for i, param := range regex.SubexpNames() {
		if i > 0 && i <= len(matches) {
			parseMap[param] = matches[i]
		}
	}

	var site string
	if parseMap["Domain1"] != "" {
		site = parseMap["Domain1"]
	} else if parseMap["Domain2"] != "" {
		site = parseMap["Domain2"]
	} else if parseMap["Domain3"] != "" {
		site = parseMap["Domain3"]
	} else {
		site = "github.com"
	}

	if !string_slice_contains(supported, site) {
		return nil, &UnsupportedHostError
	}

	user := parseMap["User"]
	name := remove_ext.ReplaceAllString(parseMap["Name"], "")

	var ref string
	if ref = parseMap["Ref"]; parseMap["Ref"] != "" {
		ref = "HEAD"
	}

	url := fmt.Sprintf("https://%s/%s/%s", site, user, name)
	ssh := fmt.Sprintf("git@%s:%s/%s", site, user, name)
	subdir := parseMap["Subdir"]

	return &SrcParse{
		Site:   site,
		User:   user,
		Name:   name,
		Ref:    ref,
		Url:    url,
		Ssh:    ssh,
		Subdir: subdir,
	}, nil
}

func CloneRepo(src string, template string, path string) {
	repo := fmt.Sprintf("%s/%s", src, template)
	_, err := parse_repo(repo)
	if err != nil {
		LogError(err.Error())
		os.Exit(err.StatusCode)
	}

	fmt.Printf("Cloned %s to %s\n", repo, path)
}
