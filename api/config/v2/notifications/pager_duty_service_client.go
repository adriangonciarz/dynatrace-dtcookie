package notifications

type PagerDutyServiceClient struct {
	client *ServiceClient
}

func NewPagerDutyService(baseURL string, token string) *PagerDutyServiceClient {
	return &PagerDutyServiceClient{client: NewService(baseURL, token)}
}

func (cs *PagerDutyServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *PagerDutyServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *PagerDutyServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *PagerDutyServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *PagerDutyServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *PagerDutyServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *PagerDutyServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.PagerDuty {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
