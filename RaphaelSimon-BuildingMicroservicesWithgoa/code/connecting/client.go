package acl

import "public_service/clients/acl/client" // HL

type Client struct {
	*client.Client // HL
}

func (c *Client) ListACLs(resource *Resource, user *User) ([]*ACL, error) {
	resourceHref := BuildResourceHref(resource)
	userHref := BuildUserHref(user)
	acls, err := c.Client.ListACLs(resourceHref, userHref) // HL
	if err != nil {
		return nil, ErrBadGateway(err)
	}
	res := make([]*ACL, len(acls))
	for i, acl := range acls {
		res[i] = aclFromResponse(acl)
	}
	return res, nil
}
