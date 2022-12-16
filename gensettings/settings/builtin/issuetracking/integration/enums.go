package integration

type IssueTheme string

var IssueThemes = struct {
	Info     IssueTheme
	Error    IssueTheme
	Resolved IssueTheme
}{
	IssueTheme("INFO"),
	IssueTheme("ERROR"),
	IssueTheme("RESOLVED"),
}

type IssueTrackerSystem string

var IssueTrackerSystems = struct {
	Gitlab        IssueTrackerSystem
	Servicenow    IssueTrackerSystem
	JiraOnPremise IssueTrackerSystem
	JiraCloud     IssueTrackerSystem
	Jira          IssueTrackerSystem
	Github        IssueTrackerSystem
}{
	IssueTrackerSystem("GITLAB"),
	IssueTrackerSystem("SERVICENOW"),
	IssueTrackerSystem("JIRA_ON_PREMISE"),
	IssueTrackerSystem("JIRA_CLOUD"),
	IssueTrackerSystem("JIRA"),
	IssueTrackerSystem("GITHUB"),
}
