package monitors

type BrowserServiceClient struct {
	client *ServiceClient
}

func NewBrowserService(baseURL string, token string) *BrowserServiceClient {
	return &BrowserServiceClient{client: NewService(baseURL, token)}
}

func (cs *BrowserServiceClient) Create(config *BrowserSyntheticMonitorUpdate) (*string, error) {
	return cs.client.CreateBrowser(config)
}

func (cs *BrowserServiceClient) Update(config *BrowserSyntheticMonitorUpdate) error {
	return cs.client.UpdateBrowser(config)
}

func (cs *BrowserServiceClient) Delete(id string) error {
	return cs.client.Delete(id)
}

func (cs *BrowserServiceClient) Get(id string) (*BrowserSyntheticMonitorUpdate, error) {
	return cs.client.GetBrowser(id)
}

func (cs *BrowserServiceClient) List() (*Monitors, error) {
	return cs.client.ListBrowser()
}

func (cs *BrowserServiceClient) GET(id string) (interface{}, error) {
	return cs.Get(id)
}

func (cs *BrowserServiceClient) LIST() ([]string, error) {
	ids := []string{}
	var err error
	if monitors, err := cs.List(); err == nil {
		for _, monitor := range monitors.Monitors {
			ids = append(ids, monitor.EntityID)
		}
	}
	return ids, err
}
