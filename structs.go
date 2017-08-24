package gitbook

type Book struct {
	BasePath     string
	GithubToken  string
	UserName     string
	Repo         string
	RepoPath     string
	OutputFolder string
	OutputFile   string
	Slack        Slack
}

type Slack struct {
	Message string
	Channel string
	Token   string
}
