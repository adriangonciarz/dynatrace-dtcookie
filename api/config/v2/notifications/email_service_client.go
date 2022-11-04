package notifications

type EmailServiceClient struct {
	client *ServiceClient
}

func NewEmailService(baseURL string, token string) *EmailServiceClient {
	return &EmailServiceClient{client: NewService(baseURL, token)}
}

func (cs *EmailServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *EmailServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *EmailServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *EmailServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *EmailServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *EmailServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *EmailServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.Email {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
