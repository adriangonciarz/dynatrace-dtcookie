package notifications

type XMattersServiceClient struct {
	client *ServiceClient
}

func NewXMattersService(baseURL string, token string) *XMattersServiceClient {
	return &XMattersServiceClient{client: NewService(baseURL, token)}
}

func (cs *XMattersServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *XMattersServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *XMattersServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *XMattersServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *XMattersServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *XMattersServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *XMattersServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.XMatters {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
