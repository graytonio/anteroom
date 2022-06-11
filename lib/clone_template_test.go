package lib

import (
	"reflect"
	"testing"
)

var base_want = &SrcParse{
	Site:   "github.com",
	User:   "graytonio",
	Name:   "anteroom-templates",
	Ref:    "",
	Url:    "https://github.com/graytonio/anteroom-templates",
	Ssh:    "git@github.com:graytonio/anteroom-templates",
	Subdir: "",
}

var subdir_want = &SrcParse{
	Site:   "github.com",
	User:   "graytonio",
	Name:   "anteroom-templates",
	Ref:    "",
	Url:    "https://github.com/graytonio/anteroom-templates",
	Ssh:    "git@github.com:graytonio/anteroom-templates",
	Subdir: "/solid-ts",
}

func TestParseRepoGithubBase(t *testing.T) {
	want := base_want

	res, err := parse_repo("graytonio/anteroom-templates")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !reflect.DeepEqual(want, res) {
		t.Errorf("Incorrect Parse\n\tWanted:\n%v\nGot\n%v", want, res)
	}
}

func TestParseRepoGithubSubdir(t *testing.T) {
	want := subdir_want

	res, err := parse_repo("graytonio/anteroom-templates/solid-ts")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !reflect.DeepEqual(want, res) {
		t.Errorf("Incorrect Parse\n\tWanted:\n%v\nGot\n%v", want, res)
	}
}

func TestParseRepoGithubURLBase(t *testing.T) {
	want := base_want

	res, err := parse_repo("https://github.com/graytonio/anteroom-templates")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !reflect.DeepEqual(want, res) {
		t.Errorf("Incorrect Parse\n\tWanted:\n%v\nGot\n%v", want, res)
	}
}

func TestParseRepoGithubURLSubdir(t *testing.T) {
	want := subdir_want

	res, err := parse_repo("https://github.com/graytonio/anteroom-templates/solid-ts")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !reflect.DeepEqual(want, res) {
		t.Errorf("Incorrect Parse\n\tWanted:\n%v\nGot\n%v", want, res)
	}
}

func TestParseRepoGithubGitBase(t *testing.T) {
	want := base_want

	res, err := parse_repo("git@github.com:graytonio/anteroom-templates")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !reflect.DeepEqual(want, res) {
		t.Errorf("Incorrect Parse\n\tWanted:\n%v\nGot\n%v", want, res)
	}
}

func TestParseRepoGithubGitSubdir(t *testing.T) {
	want := subdir_want

	res, err := parse_repo("git@github.com:graytonio/anteroom-templates/solid-ts")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if !reflect.DeepEqual(want, res) {
		t.Errorf("Incorrect Parse\n\tWanted:\n%v\nGot\n%v", want, res)
	}
}
