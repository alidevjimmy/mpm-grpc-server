package server

import (
	"context"
	"esme_team/src/model"
	"esme_team/src/pbs/reportpb"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReportServer struct {
	reportpb.UnimplementedReportServiceServer
	Collection *mongo.Collection
}

func (r *ReportServer) CreateReport(ctx context.Context, req *reportpb.SensorReportRequest) (*reportpb.SensorReportResponse, error) {
	report := req.GetReport()
	var s int
	if report.GetStatus() == reportpb.Report_Default {
		s = int(reportpb.Report_LOW)
	}
	model := model.ReportModel{
		FaultID:  int(report.GetFaultId()),
		SensorID: int(report.GetSensorId()),
		Status:   s,
		Tags: report.GetTags(),
	}

	res , err := r.Collection.InsertOne(context.Background() , model)
	if err != nil {
		return nil , status.Errorf(
			codes.Internal,
			fmt.Sprintf("error while store report data into database: %v" , err),
		)
	}
	rid , ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while casting stored data: %v", err),
		)
	}
	return &reportpb.SensorReportResponse{
		Report: &reportpb.Report{
			ID: rid.Hex(),
			FaultId: report.GetFaultId(),
			SensorId: report.GetSensorId(),
			Status: report.GetStatus(),
			Tags: report.GetTags(),
		},
	}, nil
}

func (r *ReportServer) GetReport(ctx context.Context, req *reportpb.GetReportRequest) (*reportpb.GetReportResponse, error) {
	return nil, nil
}

func (r *ReportServer) GetUnCompletedReports(req *reportpb.GetUnCompletedReportsRequest, stream reportpb.ReportService_GetUnCompletedReportsServer) error {
	return nil
}

func (r *ReportServer) GetReportLogs(req *reportpb.GetReportLogRequest, stream reportpb.ReportService_GetReportLogsServer) error {
	return nil
}
