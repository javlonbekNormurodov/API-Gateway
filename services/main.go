package services

import (
	"fmt"

	"google.golang.org/grpc"

	"bitbucket.org/udevs/example_api_gateway/config"
	"bitbucket.org/udevs/example_api_gateway/genproto/company_service"
)

type ServiceManager interface {
	CompanyService() company_service.CompanyServiceClient
}

type grpcClients struct {
	companyService company_service.CompanyServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connCompanyService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CompanyServiceHost, conf.CompanyServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		companyService: company_service.NewCompanyServiceClient(connCompanyService),
	}, nil
}

func (g *grpcClients) CompanyService() company_service.CompanyServiceClient {
	return g.companyService
}
