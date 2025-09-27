package stereo_data

type RedditMarkdown struct{}

func (r *RedditMarkdown) RedditSuperscript(val string) string {
	return "^(" + val + ")"
}

func NewRedditMarkdown() RedditMarkdown {
	return RedditMarkdown{}
}
