package notifications

type JiraServiceClient struct {
	client *ServiceClient
}

func NewJiraService(baseURL string, token string) *JiraServiceClient {
	return &JiraServiceClient{client: NewService(baseURL, token)}
}

func (cs *JiraServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *JiraServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *JiraServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *JiraServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *JiraServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *JiraServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *JiraServiceClient) LIST() ([]string, error) {
	var ids []string
	var err error
	if ids, err = cs.List(); err != nil {
		return nil, err
	}

	var filteredList []string
	var notification Notification
	for _, id := range ids {
		if err = cs.Get(id, &notification); err != nil {
			return nil, err
		}
		if notification.Type == Types.Jira {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
