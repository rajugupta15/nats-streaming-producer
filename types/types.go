package Types

type J struct {
	Key     string   `json:"key"`
	Events  []Events `json:"events"`
	Context Context  `json:"context"`
}

type Events struct {
	Event      string     `json:"event"`
	Timestamp  int64      `json:"timestamp"`
	Properties Properties `json:"properties"`
}
type Properties struct {
	LocalOrderID       string          `json:"local_order_id"`
	LocalPaymentID     string          `json:"local_payment_id"`
	MarchantAppName    string          `json:"merchant_app_name"`
	MerchantAppVersion string          `json:"merchant_app_version"`
	MerchantAppBuild   int             `json:"merchant_app_build"`
	MerchantOptions    MerchantOptions `json:"merchant_options"`
}
type MerchantOptions struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Currency    string   `json:"currency"`
	Amount      string   `json:"amount"`
	Prefill     Prefill  `json:"prefill"`
	External    External `json:"external"`
	Key         string   `json:"key"`
}
type Prefill struct {
	Contact string `json:"contact"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}
type External struct {
	Wallets []Wallets `json:"wallets"`
}
type Wallets struct {
	Paytm string `json:"paytm"`
}
type Context struct {
	Mode             string  `json:"mode"`
	Device           Device  `json:"device"`
	Sdk              Sdk     `json:"sdk"`
	Network          Network `json:"network"`
	Screen           Screen  `json:"screen"`
	Locale           string  `json:"locale"`
	Timezone         string  `json:"timezone"`
	UserAgent        string  `json:"user_agent"`
	WebviewUserAgent string  `json:"webview_user_agent"`
}
type Device struct {
	ID           string `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Version      string `json:"version"`
}
type Sdk struct {
	Version string `json:"version"`
	Type    string `json:"type"`
}
type Network struct {
	Bluetooth           bool   `json:"bluetooth"`
	Carrier             string `json:"carrier"`
	Cellular            bool   `json:"cellular"`
	CellularNetworkType string `json:"cellular_network_type"`
	Wifi                bool   `json:"wifi"`
}

type Screen struct {
	Density float64 `json:"density"`
	Width   int64   `json:"width"`
	Height  int64   `json:"height"`
}
