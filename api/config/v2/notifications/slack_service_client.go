package notifications

type SlackServiceClient struct {
	client *ServiceClient
}

func NewSlackService(baseURL string, token string) *SlackServiceClient {
	return &SlackServiceClient{client: NewService(baseURL, token)}
}

func (cs *SlackServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *SlackServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *SlackServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *SlackServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *SlackServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *SlackServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *SlackServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.Slack {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
