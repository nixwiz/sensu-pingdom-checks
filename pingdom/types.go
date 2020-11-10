package pingdom

import (
	"fmt"
)

// PingdomResponse represents a general response from the Pingdom API.
type PingdomResponse struct {
	Message string `json:"message"`
}

// PingdomError represents an error response from the Pingdom API.
type PingdomError struct {
	StatusCode int    `json:"statuscode"`
	StatusDesc string `json:"statusdesc"`
	Message    string `json:"errormessage"`
}

type CheckResponse struct {
	Type                     CheckResponseType  `json:"type,omitempty"`
	SendNotificationWhenDown int                `json:"sendnotificationwhendown,omitempty"`
	NotifyAgainEvery         int                `json:"notifyagainevery,omitempty"`
	NotifyWhenBackup         bool               `json:"notifywhenbackup,omitempty"`
	ResponseTimeThreshold    int                `json:"responsetime_threshold,omitempty"`
	CustomMessage            string             `json:"custom_message,omitempty"`
	IntegrationIDs           []int              `json:"integrationids,omitempty"`
	ID                       int                `json:"id"`
	Name                     string             `json:"name"`
	LastErrorTime            int64              `json:"lasterrortime,omitempty"`
	LastTestTime             int64              `json:"lasttesttime,omitempty"`
	LastResponseTime         int64              `json:"lastresponsetime,omitempty"`
	Status                   string             `json:"status,omitempty"`
	Resolution               int                `json:"resolution,omitempty"`
	Hostname                 string             `json:"hostname,omitempty"`
	Created                  int64              `json:"created,omitempty"`
	Tags                     []CheckResponseTag `json:"tags,omitempty"`
	Paused                   bool               `json:"paused,omitempty"`
	ProbeFilters             []string           `json:"probe_filters,omitempty"`
	IPv6                     bool               `json:"ipv6,omitempty"`
	VerifyCertificate        bool               `json:"verify_certificate,omitempty"`
	SSLDownDaysBefore        int                `json:"ssl_down_days_before,omitempty"`
}

type CheckResponseType struct {
	Name              string                          `json:"-"`
	HTTP              *CheckResponseHTTPDetails       `json:"http,omitempty"`
	HTTPCustomDetails *CheckResponseHTTPCustomDetails `json:"httpcustom,omitempty"`
	TCP               *CheckResponseTCPDetails        `json:"tcp,omitempty"`
	UDP               *CheckResponseUDPDetails        `json:"udp,omitempty"`
	DNS               *CheckResponseDNSDetails        `json:"dns,omitempty"`
	SMTP              *CheckResponseSMTPDetails       `json:"smtp,omitempty"`
	POP3              *CheckResponsePOP3Details       `json:"pop3,omitempty"`
	IMAP              *CheckResponseIMAPDetails       `json:"imap,omitempty"`
}

type CheckResponseTag struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Count interface{} `json:"count"`
}

type CheckResponseHTTPDetails struct {
	Username         string            `json:"username,omitempty"`
	Password         string            `json:"password,omitempty"`
	URL              string            `json:"url,omitempty"`
	Encryption       bool              `json:"encryption,omitempty"`
	Port             int               `json:"port,omitempty"`
	ShouldContain    string            `json:"shouldcontain,omitempty"`
	ShouldNotContain string            `json:"shouldnotcontain,omitempty"`
	PostData         string            `json:"postdata,omitempty"`
	RequestHeaders   map[string]string `json:"requestheaders,omitempty"`
}

type CheckResponseHTTPCustomDetails struct {
	URL               string `json:"url,omitempty"`
	Encryption        bool   `json:"encryption,omitempty"`
	Port              int    `json:"port,omitempty"`
	AdditionalURLs    string `json:"additionalurls,omitempty"`
	VerifyCertificate bool   `json:"verify_certificate,omitempty"`
	SSLDownDaysBefore int    `json:"ssl_down_days_before,omitempty"`
}

type CheckResponseTCPDetails struct {
	Port           int    `json:"port,omitempty"`
	StringToSend   string `json:"stringtosend,omitempty"`
	StringToExpect string `json:"stringtoexpect,omitempty"`
}

type CheckResponseUDPDetails struct {
	Port           int    `json:"port,omitempty"`
	StringToSend   string `json:"stringtosend,omitempty"`
	StringToExpect string `json:"stringtoexpect,omitempty"`
}

type CheckResponseDNSDetails struct {
	NameServer string `json:"nameserver,omitempty"`
	ExpectedIP string `json:"expectedip,omitempty"`
}

type CheckResponseSMTPDetails struct {
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	Port           int    `json:"port,omitempty"`
	Encryption     bool   `json:"encryption,omitempty"`
	StringToExpect string `json:"stringtoexpect,omitempty"`
}

type CheckResponsePOP3Details struct {
	Port           int    `json:"port,omitempty"`
	StringToExpect string `json:"stringtoexpect,omitempty"`
}

type CheckResponseIMAPDetails struct {
	Port           int    `json:"port,omitempty"`
	StringToExpect string `json:"stringtoexpect,omitempty"`
}

type CreditResponse struct {
	Credits CreditResponseDetails `json:"credits"`
}

type CreditResponseDetails struct {
	CheckLimit                 int  `json:"checklimit,omitempty"`
	DefaultCheckLimit          int  `json:"defaultchecklimit,omitempty"`
	TransactionCheckLimit      int  `json:"transactionchecklimit,omitempty"`
	AvailableChecks            int  `json:"availablechecks,omitempty"`
	AvailableDefaultChecks     int  `json:"availabledefaultchecks,omitempty"`
	AvailableTransactionChecks int  `json:"availabletransactionchecks,omitempty"`
	UsedDefault                int  `json:"useddefault,omitempty"`
	UsedTransaction            int  `json:"usedtransaction,omitempty"`
	AvailableSMS               int  `json:"availablesms,omitempty"`
	AvailableSMSTests          int  `json:"availablesmstests,omitempty"`
	AutoFillSMS                bool `json:"autofillsms,omitempty"`
	AutoFillSMSAmount          int  `json:"autofillsms_amount,omitempty"`
	AutoFillSMSWhenLeft        int  `json:"autofillsms_when_left,omitempty"`
	MaxSMSOverage              int  `json:"max_sms_overage,omitempty"`
	AvailableRUMSites          int  `json:"availablerumsites,omitempty"`
	UsedRUMSites               int  `json:"usedrumsites,omitempty"`
	MaxRUMFilters              int  `json:"maxrumfilters,omitempty"`
	MaxRUMPageViews            int  `json:"maxrumpageviews,omitempty"`
	MaxAlertingFullUsers       int  `json:"maxalertingfullusers,omitempty"`
	AvailableAlertingFullUsers int  `json:"availablealertingfullusers,omitempty"`
}

// Return string representation of the PingdomError.
func (pe *PingdomError) Error() string {
	return fmt.Sprintf("%d %v: %v", pe.StatusCode, pe.StatusDesc, pe.Message)
}

type ListChecksResponse struct {
	Checks []ListChecksDetails       `json:"checks"`
	Counts ListChecksRespounseCounts `json:"counts"`
}

type ListChecksDetails struct {
	Type             string             `json:"type"`
	ID               int                `json:"id"`
	Name             string             `json:"name"`
	LastErrorTime    int64              `json:"lasterrortime,omitempty"`
	LastTestTime     int64              `json:"lasttesttime,omitempty"`
	LastResponseTime int64              `json:"lastresponsetime,omitempty"`
	Status           string             `json:"status,omitempty"`
	Resolution       int                `json:"resolution,omitempty"`
	Hostname         string             `json:"hostname,omitempty"`
	Created          int64              `json:"created,omitempty"`
	Tags             []CheckResponseTag `json:"tags,omitempty"`
	Paused           bool               `json:"paused,omitempty"`
	ProbeFilters     []string           `json:"probe_filters,omitempty"`
	IPv6             bool               `json:"ipv6,omitempty"`
}

type ListChecksRespounseCounts struct {
	Total    int `json:"total"`
	Limited  int `json:"limited"`
	Filtered int `json:"filtered"`
}

type errorJSONResponse struct {
	Error *PingdomError `json:"error"`
}
