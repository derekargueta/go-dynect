
package dynect

// QpsReportRequest is used to model the request data for a QPS Report
// API documentation: https://help.dyn.com/create-qps-report-api/
type QpsReportRequest struct {
  // Required. The timestamp at the beginning of the report range
  StartTs       int64       `json:"start_ts"`
  // Required. The timestamp at the end of the report range
  EndTs         int64       `json:"end_ts"`

  // Fields by which to break down the data. Valid values are `hosts`, `rrecs`,
  // and `zones`.
  Breakdown     []string    `json:"breakdown"`

  // Hostnames to include in the report. An individual item can begin with a ‘!’
  // to indicate a value to exclude. Defaults to all record types.
  Hosts         []string    `json:"hosts"`

  // Record types to include in the report. An individual item can begin with a
  // ‘!’ to indicate a value to exclude. Defaults to all record types.
  RRecs         []string    `json:"rrecs"`

  // Zones to include in the report. An empty list defaults to all zones.
  Zones         []string    `json:"zones"`
}


// QpsReportResponse is used for holding the data returned by a call to
// "https://api.dynect.net/REST/QPSReport/".
type QpsReportResponse struct {
  ResponseBlock
  Data struct {
    Csv     string    `json:"csv"`
  }                   `json:"data"`
}
