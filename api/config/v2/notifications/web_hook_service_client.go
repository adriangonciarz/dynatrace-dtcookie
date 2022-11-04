package notifications

type WebHookServiceClient struct {
	client *ServiceClient
}

func NewWebHookService(baseURL string, token string) *WebHookServiceClient {
	return &WebHookServiceClient{client: NewService(baseURL, token)}
}

func (cs *WebHookServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *WebHookServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *WebHookServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *WebHookServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *WebHookServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *WebHookServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *WebHookServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.WebHook {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
