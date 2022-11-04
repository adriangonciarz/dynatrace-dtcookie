package monitors

type HttpServiceClient struct {
	client *ServiceClient
}

func NewHTTPService(baseURL string, token string) *HttpServiceClient {
	return &HttpServiceClient{client: NewService(baseURL, token)}
}

func (cs *HttpServiceClient) Create(config *HTTPSyntheticMonitorUpdate) (*string, error) {
	return cs.client.CreateHTTP(config)
}

func (cs *HttpServiceClient) Update(config *HTTPSyntheticMonitorUpdate) error {
	return cs.client.UpdateHTTP(config)
}

func (cs *HttpServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *HttpServiceClient) Get(id string) (*HTTPSyntheticMonitorUpdate, error) {
	return cs.client.GetHTTP(id)
}

func (cs *HttpServiceClient) List() (*Monitors, error) {
	return cs.client.ListHTTP()
}

func (cs *HttpServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *HttpServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if monitors, err := cs.List(); err == nil {
		for _, monitor := range monitors.Monitors {
			ids = append(ids, monitor.EntityID)
		}
	}
	return ids, err
}
