package notifications

type AnsibleTowerServiceClient struct {
	client *ServiceClient
}

func NewAnsibleTowerService(baseURL string, token string) *AnsibleTowerServiceClient {
	return &AnsibleTowerServiceClient{client: NewService(baseURL, token)}
}

func (cs *AnsibleTowerServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *AnsibleTowerServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *AnsibleTowerServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *AnsibleTowerServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *AnsibleTowerServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *AnsibleTowerServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *AnsibleTowerServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.AnsibleTower {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
