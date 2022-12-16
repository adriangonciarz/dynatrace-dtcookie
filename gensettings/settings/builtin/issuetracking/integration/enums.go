package integration

type IssueTheme string

var IssueThemes = struct {
	Error    IssueTheme
	Info     IssueTheme
	Resolved IssueTheme
}{
	IssueTheme("ERROR"),
	IssueTheme("INFO"),
	IssueTheme("RESOLVED"),
}

type IssueTrackerSystem string

var IssueTrackerSystems = struct {
	Github        IssueTrackerSystem
	Gitlab        IssueTrackerSystem
	Jira          IssueTrackerSystem
	JiraCloud     IssueTrackerSystem
	JiraOnPremise IssueTrackerSystem
	Servicenow    IssueTrackerSystem
}{
	IssueTrackerSystem("GITHUB"),
	IssueTrackerSystem("GITLAB"),
	IssueTrackerSystem("JIRA"),
	IssueTrackerSystem("JIRA_CLOUD"),
	IssueTrackerSystem("JIRA_ON_PREMISE"),
	IssueTrackerSystem("SERVICENOW"),
}
