package cheevos

import (
	"context"
)

type OrganizationLogger struct {
	Svc interface {
		AddUserToOrganization(ctx context.Context, req AddUserToOrganizationRequest) (*AddUserToOrganizationResponse, error)
		CreateOrganization(ctx context.Context, req CreateOrganizationRequest) (*CreateOrganizationResponse, error)
	}
	Logger Logger
}

func (ol *OrganizationLogger) AddUserToOrganization(ctx context.Context, req AddUserToOrganizationRequest) (*AddUserToOrganizationResponse, error) {
	ol.Logger.Debug(ctx, "adding user to organization", Fields{"Organization": req.Organization, "User": req.User})

	resp, err := ol.Svc.AddUserToOrganization(ctx, req)
	if err != nil {
		ol.Logger.Error(ctx, "add user to organization failed", err)
		return nil, err
	}
	ol.Logger.Log(ctx, "user added to organization", Fields{
		"Organization": resp.Organization,
		"User":         resp.User,
	})

	return resp, nil
}

func (ol *OrganizationLogger) CreateOrganization(ctx context.Context, req CreateOrganizationRequest) (*CreateOrganizationResponse, error) {
	ol.Logger.Debug(ctx, "creating organization", Fields{"Name": req.Name, "Owner": req.Owner})

	resp, err := ol.Svc.CreateOrganization(ctx, req)
	if err != nil {
		ol.Logger.Error(ctx, "create organization failed", err)
		return nil, err
	}
	ol.Logger.Log(ctx, "organization created", Fields{"Organization": resp.Organization})

	return resp, nil
}
