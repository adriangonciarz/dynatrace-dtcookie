package notifications

type ServiceNowServiceClient struct {
	client *ServiceClient
}

func NewServiceNowService(baseURL string, token string) *ServiceNowServiceClient {
	return &ServiceNowServiceClient{client: NewService(baseURL, token)}
}

func (cs *ServiceNowServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *ServiceNowServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *ServiceNowServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *ServiceNowServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *ServiceNowServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *ServiceNowServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *ServiceNowServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.ServiceNow {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
