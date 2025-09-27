package stereo_data_test

import (
	stereo_data "stereo-server/internal/stereo-data"
	"testing"
)

func TestRedditHandler(t *testing.T) {
	redditMarkdown := stereo_data.NewRedditMarkdown()

	expected := "^(test)"

	res := redditMarkdown.RedditSuperscript("test")

	if expected != res {
		t.Errorf("expected superscript to be %v; got %v", expected, res)
	}
}
