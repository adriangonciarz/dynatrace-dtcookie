package notifications_test

import (
	"os"
	"sort"
	"testing"

	"github.com/dtcookie/assert"
	"github.com/dtcookie/dynatrace/api/config/v2/notifications"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/opt"
)

type config struct {
	BaseURL  string // https://######.live.dynatrace.com/api/config/v2
	APIToken string // ########
}

func Config(t *testing.T) config {
	if os.Getenv("DYNATRACE_DEBUG") == "true" {
		rest.Verbose = true
	}
	cfg := config{
		BaseURL:  os.Getenv("DYNATRACE_ENV_URL"),
		APIToken: os.Getenv("DYNATRACE_API_TOKEN"),
	}
	if len(cfg.BaseURL) == 0 {
		t.Skipf("You need to specify the environment variable `DYNATRACE_ENV_URL` to run this test.")
	}
	if len(cfg.APIToken) == 0 {
		t.Skipf("You need to specify the environment variable `DYNATRACE_API_TOKEN` to run this test.")
	}
	return cfg
}

func TestAnsibleTower(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "ansible_tower_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.AnsibleTower,
		AnsibleTower: &notifications.AnsibleTower{
			Insecure:       true,
			JobTemplateURL: "https://localhost/#/templates/job_template/666",
			Username:       "foo",
			Password:       "bar",
			CustomMessage:  "some-custom-message",
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)
	config2.AnsibleTower.Password = config.AnsibleTower.Password

	assert.Equals(config, config2)
	client.Delete(id)
}

func TestEmail(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "email_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.Email,
		Email: &notifications.Email{
			Subject:              "EMAIL-SUBJECT",
			Recipients:           []string{"she@home.com", "me@home.com", "you@home.com"},
			CCRecipients:         []string{"she@home.org", "me@home.org", "you@home.org"},
			BCCRecipients:        []string{"she@home.gov", "me@home.gov", "you@home.gov"},
			NotifyClosedProblems: true,
			Body:                 "{ProblemDetailsHTML}",
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	// receipients, cc and bcc are sets - unordered slices
	// the REST API may decide to reorder them
	sort.Strings(config.Email.Recipients)
	sort.Strings(config.Email.CCRecipients)
	sort.Strings(config.Email.BCCRecipients)
	sort.Strings(config2.Email.Recipients)
	sort.Strings(config2.Email.CCRecipients)
	sort.Strings(config2.Email.BCCRecipients)

	assert.Equals(config, config2)
	client.Delete(id)
}

func TestJira(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "jira_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.Jira,
		Jira: &notifications.Jira{
			URL:         "https://localhost:9999/jira",
			Username:    "jira-user-name",
			APIToken:    "jira-api-token",
			ProjectKey:  "jira-project-key",
			IssueType:   "jira-issue-type",
			Summary:     "jira-summary",
			Description: "jira-description",
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config2.Jira.APIToken = config.Jira.APIToken
	assert.Equals(config, config2)
	client.Delete(id)
}

func TestOpsGenie(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "ops_genie_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.OpsGenie,
		OpsGenie: &notifications.OpsGenie{
			Domain:  "ops-genie-domain",
			Message: "ops-genie-message",
			APIKey:  opt.NewString("ops-genie-api-key"),
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config2.OpsGenie.APIKey = config.OpsGenie.APIKey
	assert.Equals(config, config2)
	client.Delete(id)
}

func TestPagerDuty(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "pager_duty_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.PagerDuty,
		PagerDuty: &notifications.PagerDuty{
			Account:     "pager-duty-account",
			ServiceName: "pager-duty-service",
			APIKey:      "pager-duty-api-key",
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config2.PagerDuty.APIKey = config.PagerDuty.APIKey
	assert.Equals(config, config2)
	client.Delete(id)
}

func TestServiceNow(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "service_now_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.ServiceNow,
		ServiceNow: &notifications.ServiceNow{
			InstanceName:  opt.NewString("service-now-instance-name"),
			URL:           nil,
			Username:      "service-now-username",
			Password:      "service-now-password",
			Message:       "service-now-message",
			SendIncidents: true,
			SendEvents:    true,
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config2.ServiceNow.Password = config.ServiceNow.Password
	assert.Equals(config, config2)
	client.Delete(id)
}

func TestSlack(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "slack_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.Slack,
		Slack: &notifications.Slack{
			URL:     "https://slack.home.com",
			Channel: "slack-channel",
			Message: "slack-message",
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config2.Slack.URL = config.Slack.URL
	assert.Equals(config, config2)
	client.Delete(id)
}

func TestTrello(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "ops_genie_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.Trello,
		Trello: &notifications.Trello{
			ApplicationKey:     "trello-application-key",
			BoardID:            "trello-board-id",
			ListID:             "trello-list-id",
			ResolvedListID:     "trello-resolved-list-id",
			Text:               "trello-text",
			Description:        "trello-description",
			AuthorizationToken: "trello-authorization-token",
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config2.Trello.AuthorizationToken = config.Trello.AuthorizationToken
	assert.Equals(config, config2)
	client.Delete(id)
}

func TestVictorOps(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "victor_ops_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.VictorOps,
		VictorOps: &notifications.VictorOps{
			APIKey:     "victor-ops-api-key",
			RoutingKey: "victor-ops-routing-key",
			Message:    "victor-ops-message",
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config2.VictorOps.APIKey = config.VictorOps.APIKey
	assert.Equals(config, config2)
	client.Delete(id)
}

func TestWebHook(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "web_hook_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.WebHook,
		WebHook: &notifications.WebHook{
			URL:                      "https://webhook.site/40bf4d43-1a50-4ebd-913d-bf50ce7c3a1e",
			Insecure:                 true,
			NotifyEventMergesEnabled: true,
			NotifyClosedProblems:     true,
			Payload:                  "web-hook-payload",
			Headers: notifications.HTTPHeaders{
				{
					Name:   "http-header-name-01",
					Secret: false,
					Value:  opt.NewString("http-header-value-01"),
				},
				{
					Name:        "http-header-name-02",
					Secret:      true,
					SecretValue: opt.NewString("http-header-value-02"),
				},
			},
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config.WebHook.Headers.Sort()

	for idx, header := range config2.WebHook.Headers {
		header.SecretValue = config.WebHook.Headers[idx].SecretValue
	}
	assert.Equals(config, config2)
	client.Delete(id)
}

func TestXMatters(t *testing.T) {
	cfg := Config(t)

	client := notifications.NewService(cfg.BaseURL, cfg.APIToken)
	config := notifications.Notification{
		Enabled:   false,
		Name:      "x_matters_test",
		ProfileID: "vu9U3hXa3q0AAAABABhidWlsdGluOmFsZXJ0aW5nLnByb2ZpbGUABnRlbmFudAAGdGVuYW50ACRmNzVlNjhlZi1hY2E3LTNhMDctOWMyMS05NGViMDBlY2ZjNTa-71TeFdrerQ",
		Type:      notifications.Types.XMatters,
		XMatters: &notifications.XMatters{
			URL:      "https://webhook.site/40bf4d43-1a50-4ebd-913d-bf50ce7c3a1e",
			Insecure: true,
			Payload:  "x-matters-payload",
			Headers: notifications.HTTPHeaders{
				{
					Name:   "http-header-name-01",
					Secret: false,
					Value:  opt.NewString("http-header-value-01"),
				},
				{
					Name:        "http-header-name-02",
					Secret:      true,
					SecretValue: opt.NewString("http-header-value-02"),
				},
			},
		},
	}
	id, err := client.Create(&config)
	if err != nil {
		t.Error(err)
	}
	var config2 notifications.Notification
	if err := client.Get(id, &config2); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)

	config.XMatters.Headers.Sort()

	for idx, header := range config2.XMatters.Headers {
		header.SecretValue = config.XMatters.Headers[idx].SecretValue
	}
	assert.Equals(config, config2)
	client.Delete(id)
}
