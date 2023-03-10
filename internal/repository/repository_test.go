package repository

import (
	"testing"

	"github.com/loghinalexandru/swears/internal/models"
)

func TestLang(t *testing.T) {
	t.Parallel()

	testLang := "en"
	mock := fileDB{
		lang: testLang,
	}

	res := mock.Lang()

	if testLang != res {
		t.Fatal("Wrong language")
	}
}

func TestGet_SingleValue(t *testing.T) {
	t.Parallel()

	mock := fileDB{
		data: []models.Record{
			{
				Value: "test",
			},
		},
	}

	res, err := mock.Get()

	if err == nil && res.Value != "test" {
		t.Fatal("Wrong object returned")
	}
}

func TestGet_EmptyData(t *testing.T) {
	t.Parallel()

	mock := fileDB{
		data: []models.Record{},
	}

	res, err := mock.Get()

	if err != nil && res.Value != "" {
		t.Fatal("Wrong object returned")
	}
}

func TestLoad(t *testing.T) {
	t.Parallel()

	mock := fileDB{
		lang: "test",
	}

	mock.load("testdata/test.txt")

	if len(mock.data) != 3 {
		t.Fatal("Wrong parsing of data file")
	}

	if mock.data[0].Value != "Test1" ||
		mock.data[1].Value != "Test 2" ||
		mock.data[2].Value != " Test 3" {
		t.Fatal("Invalid parsing!")
	}
}
