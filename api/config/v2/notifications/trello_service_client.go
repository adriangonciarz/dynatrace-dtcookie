package notifications

type TrelloServiceClient struct {
	client *ServiceClient
}

func NewTrelloService(baseURL string, token string) *TrelloServiceClient {
	return &TrelloServiceClient{client: NewService(baseURL, token)}
}

func (cs *TrelloServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *TrelloServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *TrelloServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *TrelloServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *TrelloServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *TrelloServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *TrelloServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.Trello {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
