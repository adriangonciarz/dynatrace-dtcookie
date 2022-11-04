package notifications

type VictorOpsServiceClient struct {
	client *ServiceClient
}

func NewVictorOpsService(baseURL string, token string) *VictorOpsServiceClient {
	return &VictorOpsServiceClient{client: NewService(baseURL, token)}
}

func (cs *VictorOpsServiceClient) Create(config *Notification) (string, error) {
	return cs.client.Create(config)
}

func (cs *VictorOpsServiceClient) Update(id string, config *Notification) error {
	return cs.client.Update(id, config)
}

func (cs *VictorOpsServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *VictorOpsServiceClient) Get(id string, target *Notification) error {
	return cs.client.Get(id, target)
}

func (cs *VictorOpsServiceClient) List() ([]string, error) {
	return cs.client.List()
}

func (cs *VictorOpsServiceClient) GET(id string) (interface{}, error) {
	var notification Notification
	if err := cs.Get(id, &notification); err != nil {
		return nil, err
	}
	return &notification, nil
}

func (cs *VictorOpsServiceClient) LIST() ([]string, error) {
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
		if notification.Type == Types.VictorOps {
			filteredList = append(filteredList, id)
		}
	}
	return filteredList, err
}
