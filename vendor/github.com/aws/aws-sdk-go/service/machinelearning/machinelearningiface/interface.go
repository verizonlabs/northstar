// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package machinelearningiface provides an interface for the Amazon Machine Learning.
package machinelearningiface

import (
	"github.com/aws/aws-sdk-go/service/machinelearning"
)

// MachineLearningAPI is the interface type for machinelearning.MachineLearning.
type MachineLearningAPI interface {
	CreateBatchPrediction(*machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error)

	CreateDataSourceFromRDS(*machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error)

	CreateDataSourceFromRedshift(*machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)

	CreateDataSourceFromS3(*machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error)

	CreateEvaluation(*machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error)

	CreateMLModel(*machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error)

	CreateRealtimeEndpoint(*machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error)

	DeleteBatchPrediction(*machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error)

	DeleteDataSource(*machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error)

	DeleteEvaluation(*machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error)

	DeleteMLModel(*machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error)

	DeleteRealtimeEndpoint(*machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error)

	DescribeBatchPredictions(*machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error)

	DescribeDataSources(*machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error)

	DescribeEvaluations(*machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error)

	DescribeMLModels(*machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error)

	GetBatchPrediction(*machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error)

	GetDataSource(*machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error)

	GetEvaluation(*machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error)

	GetMLModel(*machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error)

	Predict(*machinelearning.PredictInput) (*machinelearning.PredictOutput, error)

	UpdateBatchPrediction(*machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error)

	UpdateDataSource(*machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error)

	UpdateEvaluation(*machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error)

	UpdateMLModel(*machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error)
}
