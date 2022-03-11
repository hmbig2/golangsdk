package nics

import (
	"github.com/chnsz/golangsdk"
)

type CreateOps struct {
	Nics []NicReq `json:"nics" required:"true"`
}

type NicReq struct {
	SubnetId       string   `json:"subnet_id" required:"true"`
	IpAddress      string   `json:"ip_address,omitempty"`
	Ipv6Enable     *bool    `json:"ipv6_enable,omitempty"`
	Ipv6Bandwidth  []IdInfo `json:"ipv6_bandwidth,omitempty"`
	SecurityGroups []IdInfo `json:"security_groups,omitempty"`
}

type IdInfo struct {
	Id string `json:"id" required:"true"`
}

type DeleteOps struct {
	Nics []IdInfo `json:"nics" required:"true"`
}

var RequestOpts = golangsdk.RequestOpts{
	MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
}

func Create(c *golangsdk.ServiceClient, serverId string, opts CreateOps) *golangsdk.ErrResult {
	var r golangsdk.ErrResult
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		r.Err = err
		return &r
	}

	_, r.Err = c.Post(createURL(c, serverId), b, &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	return &r
}

func Delete(c *golangsdk.ServiceClient, serverId string, opts DeleteOps) *golangsdk.ErrResult {
	var r golangsdk.ErrResult
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		r.Err = err
		return &r
	}

	_, r.Err = c.Post(deleteURL(c, serverId), b, &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	return &r
}

func Get(c *golangsdk.ServiceClient, serverId string) (*NicsResp, error) {
	var rst NicsResp
	_, err := c.Get(getURL(c, serverId), &rst, &golangsdk.RequestOpts{
		MoreHeaders: RequestOpts.MoreHeaders,
	})
	return &rst, err
}
