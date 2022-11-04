package notifications

type OpsgenieServiceClient struct {
	client *ServiceClient
}

func NewOpsgenieService(baseURL string, token string) *OpsgenieServiceClient {
	return &OpsgenieServiceClient{client: NewService(baseURL, token)}
}

func (cs *OpsgenieServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *OpsgenieServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *OpsgenieServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *OpsgenieServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *OpsgenieServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *OpsgenieServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *OpsgenieServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.OpsGenie {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
