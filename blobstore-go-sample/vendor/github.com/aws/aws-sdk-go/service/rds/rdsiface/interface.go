// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package rdsiface provides an interface for the Amazon Relational Database Service.
package rdsiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/rds"
)

// RDSAPI is the interface type for rds.RDS.
type RDSAPI interface {
	AddSourceIdentifierToSubscriptionRequest(*rds.AddSourceIdentifierToSubscriptionInput) (*request.Request, *rds.AddSourceIdentifierToSubscriptionOutput)

	AddSourceIdentifierToSubscription(*rds.AddSourceIdentifierToSubscriptionInput) (*rds.AddSourceIdentifierToSubscriptionOutput, error)

	AddTagsToResourceRequest(*rds.AddTagsToResourceInput) (*request.Request, *rds.AddTagsToResourceOutput)

	AddTagsToResource(*rds.AddTagsToResourceInput) (*rds.AddTagsToResourceOutput, error)

	ApplyPendingMaintenanceActionRequest(*rds.ApplyPendingMaintenanceActionInput) (*request.Request, *rds.ApplyPendingMaintenanceActionOutput)

	ApplyPendingMaintenanceAction(*rds.ApplyPendingMaintenanceActionInput) (*rds.ApplyPendingMaintenanceActionOutput, error)

	AuthorizeDBSecurityGroupIngressRequest(*rds.AuthorizeDBSecurityGroupIngressInput) (*request.Request, *rds.AuthorizeDBSecurityGroupIngressOutput)

	AuthorizeDBSecurityGroupIngress(*rds.AuthorizeDBSecurityGroupIngressInput) (*rds.AuthorizeDBSecurityGroupIngressOutput, error)

	CopyDBClusterParameterGroupRequest(*rds.CopyDBClusterParameterGroupInput) (*request.Request, *rds.CopyDBClusterParameterGroupOutput)

	CopyDBClusterParameterGroup(*rds.CopyDBClusterParameterGroupInput) (*rds.CopyDBClusterParameterGroupOutput, error)

	CopyDBClusterSnapshotRequest(*rds.CopyDBClusterSnapshotInput) (*request.Request, *rds.CopyDBClusterSnapshotOutput)

	CopyDBClusterSnapshot(*rds.CopyDBClusterSnapshotInput) (*rds.CopyDBClusterSnapshotOutput, error)

	CopyDBParameterGroupRequest(*rds.CopyDBParameterGroupInput) (*request.Request, *rds.CopyDBParameterGroupOutput)

	CopyDBParameterGroup(*rds.CopyDBParameterGroupInput) (*rds.CopyDBParameterGroupOutput, error)

	CopyDBSnapshotRequest(*rds.CopyDBSnapshotInput) (*request.Request, *rds.CopyDBSnapshotOutput)

	CopyDBSnapshot(*rds.CopyDBSnapshotInput) (*rds.CopyDBSnapshotOutput, error)

	CopyOptionGroupRequest(*rds.CopyOptionGroupInput) (*request.Request, *rds.CopyOptionGroupOutput)

	CopyOptionGroup(*rds.CopyOptionGroupInput) (*rds.CopyOptionGroupOutput, error)

	CreateDBClusterRequest(*rds.CreateDBClusterInput) (*request.Request, *rds.CreateDBClusterOutput)

	CreateDBCluster(*rds.CreateDBClusterInput) (*rds.CreateDBClusterOutput, error)

	CreateDBClusterParameterGroupRequest(*rds.CreateDBClusterParameterGroupInput) (*request.Request, *rds.CreateDBClusterParameterGroupOutput)

	CreateDBClusterParameterGroup(*rds.CreateDBClusterParameterGroupInput) (*rds.CreateDBClusterParameterGroupOutput, error)

	CreateDBClusterSnapshotRequest(*rds.CreateDBClusterSnapshotInput) (*request.Request, *rds.CreateDBClusterSnapshotOutput)

	CreateDBClusterSnapshot(*rds.CreateDBClusterSnapshotInput) (*rds.CreateDBClusterSnapshotOutput, error)

	CreateDBInstanceRequest(*rds.CreateDBInstanceInput) (*request.Request, *rds.CreateDBInstanceOutput)

	CreateDBInstance(*rds.CreateDBInstanceInput) (*rds.CreateDBInstanceOutput, error)

	CreateDBInstanceReadReplicaRequest(*rds.CreateDBInstanceReadReplicaInput) (*request.Request, *rds.CreateDBInstanceReadReplicaOutput)

	CreateDBInstanceReadReplica(*rds.CreateDBInstanceReadReplicaInput) (*rds.CreateDBInstanceReadReplicaOutput, error)

	CreateDBParameterGroupRequest(*rds.CreateDBParameterGroupInput) (*request.Request, *rds.CreateDBParameterGroupOutput)

	CreateDBParameterGroup(*rds.CreateDBParameterGroupInput) (*rds.CreateDBParameterGroupOutput, error)

	CreateDBSecurityGroupRequest(*rds.CreateDBSecurityGroupInput) (*request.Request, *rds.CreateDBSecurityGroupOutput)

	CreateDBSecurityGroup(*rds.CreateDBSecurityGroupInput) (*rds.CreateDBSecurityGroupOutput, error)

	CreateDBSnapshotRequest(*rds.CreateDBSnapshotInput) (*request.Request, *rds.CreateDBSnapshotOutput)

	CreateDBSnapshot(*rds.CreateDBSnapshotInput) (*rds.CreateDBSnapshotOutput, error)

	CreateDBSubnetGroupRequest(*rds.CreateDBSubnetGroupInput) (*request.Request, *rds.CreateDBSubnetGroupOutput)

	CreateDBSubnetGroup(*rds.CreateDBSubnetGroupInput) (*rds.CreateDBSubnetGroupOutput, error)

	CreateEventSubscriptionRequest(*rds.CreateEventSubscriptionInput) (*request.Request, *rds.CreateEventSubscriptionOutput)

	CreateEventSubscription(*rds.CreateEventSubscriptionInput) (*rds.CreateEventSubscriptionOutput, error)

	CreateOptionGroupRequest(*rds.CreateOptionGroupInput) (*request.Request, *rds.CreateOptionGroupOutput)

	CreateOptionGroup(*rds.CreateOptionGroupInput) (*rds.CreateOptionGroupOutput, error)

	DeleteDBClusterRequest(*rds.DeleteDBClusterInput) (*request.Request, *rds.DeleteDBClusterOutput)

	DeleteDBCluster(*rds.DeleteDBClusterInput) (*rds.DeleteDBClusterOutput, error)

	DeleteDBClusterParameterGroupRequest(*rds.DeleteDBClusterParameterGroupInput) (*request.Request, *rds.DeleteDBClusterParameterGroupOutput)

	DeleteDBClusterParameterGroup(*rds.DeleteDBClusterParameterGroupInput) (*rds.DeleteDBClusterParameterGroupOutput, error)

	DeleteDBClusterSnapshotRequest(*rds.DeleteDBClusterSnapshotInput) (*request.Request, *rds.DeleteDBClusterSnapshotOutput)

	DeleteDBClusterSnapshot(*rds.DeleteDBClusterSnapshotInput) (*rds.DeleteDBClusterSnapshotOutput, error)

	DeleteDBInstanceRequest(*rds.DeleteDBInstanceInput) (*request.Request, *rds.DeleteDBInstanceOutput)

	DeleteDBInstance(*rds.DeleteDBInstanceInput) (*rds.DeleteDBInstanceOutput, error)

	DeleteDBParameterGroupRequest(*rds.DeleteDBParameterGroupInput) (*request.Request, *rds.DeleteDBParameterGroupOutput)

	DeleteDBParameterGroup(*rds.DeleteDBParameterGroupInput) (*rds.DeleteDBParameterGroupOutput, error)

	DeleteDBSecurityGroupRequest(*rds.DeleteDBSecurityGroupInput) (*request.Request, *rds.DeleteDBSecurityGroupOutput)

	DeleteDBSecurityGroup(*rds.DeleteDBSecurityGroupInput) (*rds.DeleteDBSecurityGroupOutput, error)

	DeleteDBSnapshotRequest(*rds.DeleteDBSnapshotInput) (*request.Request, *rds.DeleteDBSnapshotOutput)

	DeleteDBSnapshot(*rds.DeleteDBSnapshotInput) (*rds.DeleteDBSnapshotOutput, error)

	DeleteDBSubnetGroupRequest(*rds.DeleteDBSubnetGroupInput) (*request.Request, *rds.DeleteDBSubnetGroupOutput)

	DeleteDBSubnetGroup(*rds.DeleteDBSubnetGroupInput) (*rds.DeleteDBSubnetGroupOutput, error)

	DeleteEventSubscriptionRequest(*rds.DeleteEventSubscriptionInput) (*request.Request, *rds.DeleteEventSubscriptionOutput)

	DeleteEventSubscription(*rds.DeleteEventSubscriptionInput) (*rds.DeleteEventSubscriptionOutput, error)

	DeleteOptionGroupRequest(*rds.DeleteOptionGroupInput) (*request.Request, *rds.DeleteOptionGroupOutput)

	DeleteOptionGroup(*rds.DeleteOptionGroupInput) (*rds.DeleteOptionGroupOutput, error)

	DescribeAccountAttributesRequest(*rds.DescribeAccountAttributesInput) (*request.Request, *rds.DescribeAccountAttributesOutput)

	DescribeAccountAttributes(*rds.DescribeAccountAttributesInput) (*rds.DescribeAccountAttributesOutput, error)

	DescribeCertificatesRequest(*rds.DescribeCertificatesInput) (*request.Request, *rds.DescribeCertificatesOutput)

	DescribeCertificates(*rds.DescribeCertificatesInput) (*rds.DescribeCertificatesOutput, error)

	DescribeDBClusterParameterGroupsRequest(*rds.DescribeDBClusterParameterGroupsInput) (*request.Request, *rds.DescribeDBClusterParameterGroupsOutput)

	DescribeDBClusterParameterGroups(*rds.DescribeDBClusterParameterGroupsInput) (*rds.DescribeDBClusterParameterGroupsOutput, error)

	DescribeDBClusterParametersRequest(*rds.DescribeDBClusterParametersInput) (*request.Request, *rds.DescribeDBClusterParametersOutput)

	DescribeDBClusterParameters(*rds.DescribeDBClusterParametersInput) (*rds.DescribeDBClusterParametersOutput, error)

	DescribeDBClusterSnapshotAttributesRequest(*rds.DescribeDBClusterSnapshotAttributesInput) (*request.Request, *rds.DescribeDBClusterSnapshotAttributesOutput)

	DescribeDBClusterSnapshotAttributes(*rds.DescribeDBClusterSnapshotAttributesInput) (*rds.DescribeDBClusterSnapshotAttributesOutput, error)

	DescribeDBClusterSnapshotsRequest(*rds.DescribeDBClusterSnapshotsInput) (*request.Request, *rds.DescribeDBClusterSnapshotsOutput)

	DescribeDBClusterSnapshots(*rds.DescribeDBClusterSnapshotsInput) (*rds.DescribeDBClusterSnapshotsOutput, error)

	DescribeDBClustersRequest(*rds.DescribeDBClustersInput) (*request.Request, *rds.DescribeDBClustersOutput)

	DescribeDBClusters(*rds.DescribeDBClustersInput) (*rds.DescribeDBClustersOutput, error)

	DescribeDBEngineVersionsRequest(*rds.DescribeDBEngineVersionsInput) (*request.Request, *rds.DescribeDBEngineVersionsOutput)

	DescribeDBEngineVersions(*rds.DescribeDBEngineVersionsInput) (*rds.DescribeDBEngineVersionsOutput, error)

	DescribeDBEngineVersionsPages(*rds.DescribeDBEngineVersionsInput, func(*rds.DescribeDBEngineVersionsOutput, bool) bool) error

	DescribeDBInstancesRequest(*rds.DescribeDBInstancesInput) (*request.Request, *rds.DescribeDBInstancesOutput)

	DescribeDBInstances(*rds.DescribeDBInstancesInput) (*rds.DescribeDBInstancesOutput, error)

	DescribeDBInstancesPages(*rds.DescribeDBInstancesInput, func(*rds.DescribeDBInstancesOutput, bool) bool) error

	DescribeDBLogFilesRequest(*rds.DescribeDBLogFilesInput) (*request.Request, *rds.DescribeDBLogFilesOutput)

	DescribeDBLogFiles(*rds.DescribeDBLogFilesInput) (*rds.DescribeDBLogFilesOutput, error)

	DescribeDBLogFilesPages(*rds.DescribeDBLogFilesInput, func(*rds.DescribeDBLogFilesOutput, bool) bool) error

	DescribeDBParameterGroupsRequest(*rds.DescribeDBParameterGroupsInput) (*request.Request, *rds.DescribeDBParameterGroupsOutput)

	DescribeDBParameterGroups(*rds.DescribeDBParameterGroupsInput) (*rds.DescribeDBParameterGroupsOutput, error)

	DescribeDBParameterGroupsPages(*rds.DescribeDBParameterGroupsInput, func(*rds.DescribeDBParameterGroupsOutput, bool) bool) error

	DescribeDBParametersRequest(*rds.DescribeDBParametersInput) (*request.Request, *rds.DescribeDBParametersOutput)

	DescribeDBParameters(*rds.DescribeDBParametersInput) (*rds.DescribeDBParametersOutput, error)

	DescribeDBParametersPages(*rds.DescribeDBParametersInput, func(*rds.DescribeDBParametersOutput, bool) bool) error

	DescribeDBSecurityGroupsRequest(*rds.DescribeDBSecurityGroupsInput) (*request.Request, *rds.DescribeDBSecurityGroupsOutput)

	DescribeDBSecurityGroups(*rds.DescribeDBSecurityGroupsInput) (*rds.DescribeDBSecurityGroupsOutput, error)

	DescribeDBSecurityGroupsPages(*rds.DescribeDBSecurityGroupsInput, func(*rds.DescribeDBSecurityGroupsOutput, bool) bool) error

	DescribeDBSnapshotAttributesRequest(*rds.DescribeDBSnapshotAttributesInput) (*request.Request, *rds.DescribeDBSnapshotAttributesOutput)

	DescribeDBSnapshotAttributes(*rds.DescribeDBSnapshotAttributesInput) (*rds.DescribeDBSnapshotAttributesOutput, error)

	DescribeDBSnapshotsRequest(*rds.DescribeDBSnapshotsInput) (*request.Request, *rds.DescribeDBSnapshotsOutput)

	DescribeDBSnapshots(*rds.DescribeDBSnapshotsInput) (*rds.DescribeDBSnapshotsOutput, error)

	DescribeDBSnapshotsPages(*rds.DescribeDBSnapshotsInput, func(*rds.DescribeDBSnapshotsOutput, bool) bool) error

	DescribeDBSubnetGroupsRequest(*rds.DescribeDBSubnetGroupsInput) (*request.Request, *rds.DescribeDBSubnetGroupsOutput)

	DescribeDBSubnetGroups(*rds.DescribeDBSubnetGroupsInput) (*rds.DescribeDBSubnetGroupsOutput, error)

	DescribeDBSubnetGroupsPages(*rds.DescribeDBSubnetGroupsInput, func(*rds.DescribeDBSubnetGroupsOutput, bool) bool) error

	DescribeEngineDefaultClusterParametersRequest(*rds.DescribeEngineDefaultClusterParametersInput) (*request.Request, *rds.DescribeEngineDefaultClusterParametersOutput)

	DescribeEngineDefaultClusterParameters(*rds.DescribeEngineDefaultClusterParametersInput) (*rds.DescribeEngineDefaultClusterParametersOutput, error)

	DescribeEngineDefaultParametersRequest(*rds.DescribeEngineDefaultParametersInput) (*request.Request, *rds.DescribeEngineDefaultParametersOutput)

	DescribeEngineDefaultParameters(*rds.DescribeEngineDefaultParametersInput) (*rds.DescribeEngineDefaultParametersOutput, error)

	DescribeEngineDefaultParametersPages(*rds.DescribeEngineDefaultParametersInput, func(*rds.DescribeEngineDefaultParametersOutput, bool) bool) error

	DescribeEventCategoriesRequest(*rds.DescribeEventCategoriesInput) (*request.Request, *rds.DescribeEventCategoriesOutput)

	DescribeEventCategories(*rds.DescribeEventCategoriesInput) (*rds.DescribeEventCategoriesOutput, error)

	DescribeEventSubscriptionsRequest(*rds.DescribeEventSubscriptionsInput) (*request.Request, *rds.DescribeEventSubscriptionsOutput)

	DescribeEventSubscriptions(*rds.DescribeEventSubscriptionsInput) (*rds.DescribeEventSubscriptionsOutput, error)

	DescribeEventSubscriptionsPages(*rds.DescribeEventSubscriptionsInput, func(*rds.DescribeEventSubscriptionsOutput, bool) bool) error

	DescribeEventsRequest(*rds.DescribeEventsInput) (*request.Request, *rds.DescribeEventsOutput)

	DescribeEvents(*rds.DescribeEventsInput) (*rds.DescribeEventsOutput, error)

	DescribeEventsPages(*rds.DescribeEventsInput, func(*rds.DescribeEventsOutput, bool) bool) error

	DescribeOptionGroupOptionsRequest(*rds.DescribeOptionGroupOptionsInput) (*request.Request, *rds.DescribeOptionGroupOptionsOutput)

	DescribeOptionGroupOptions(*rds.DescribeOptionGroupOptionsInput) (*rds.DescribeOptionGroupOptionsOutput, error)

	DescribeOptionGroupOptionsPages(*rds.DescribeOptionGroupOptionsInput, func(*rds.DescribeOptionGroupOptionsOutput, bool) bool) error

	DescribeOptionGroupsRequest(*rds.DescribeOptionGroupsInput) (*request.Request, *rds.DescribeOptionGroupsOutput)

	DescribeOptionGroups(*rds.DescribeOptionGroupsInput) (*rds.DescribeOptionGroupsOutput, error)

	DescribeOptionGroupsPages(*rds.DescribeOptionGroupsInput, func(*rds.DescribeOptionGroupsOutput, bool) bool) error

	DescribeOrderableDBInstanceOptionsRequest(*rds.DescribeOrderableDBInstanceOptionsInput) (*request.Request, *rds.DescribeOrderableDBInstanceOptionsOutput)

	DescribeOrderableDBInstanceOptions(*rds.DescribeOrderableDBInstanceOptionsInput) (*rds.DescribeOrderableDBInstanceOptionsOutput, error)

	DescribeOrderableDBInstanceOptionsPages(*rds.DescribeOrderableDBInstanceOptionsInput, func(*rds.DescribeOrderableDBInstanceOptionsOutput, bool) bool) error

	DescribePendingMaintenanceActionsRequest(*rds.DescribePendingMaintenanceActionsInput) (*request.Request, *rds.DescribePendingMaintenanceActionsOutput)

	DescribePendingMaintenanceActions(*rds.DescribePendingMaintenanceActionsInput) (*rds.DescribePendingMaintenanceActionsOutput, error)

	DescribeReservedDBInstancesRequest(*rds.DescribeReservedDBInstancesInput) (*request.Request, *rds.DescribeReservedDBInstancesOutput)

	DescribeReservedDBInstances(*rds.DescribeReservedDBInstancesInput) (*rds.DescribeReservedDBInstancesOutput, error)

	DescribeReservedDBInstancesPages(*rds.DescribeReservedDBInstancesInput, func(*rds.DescribeReservedDBInstancesOutput, bool) bool) error

	DescribeReservedDBInstancesOfferingsRequest(*rds.DescribeReservedDBInstancesOfferingsInput) (*request.Request, *rds.DescribeReservedDBInstancesOfferingsOutput)

	DescribeReservedDBInstancesOfferings(*rds.DescribeReservedDBInstancesOfferingsInput) (*rds.DescribeReservedDBInstancesOfferingsOutput, error)

	DescribeReservedDBInstancesOfferingsPages(*rds.DescribeReservedDBInstancesOfferingsInput, func(*rds.DescribeReservedDBInstancesOfferingsOutput, bool) bool) error

	DownloadDBLogFilePortionRequest(*rds.DownloadDBLogFilePortionInput) (*request.Request, *rds.DownloadDBLogFilePortionOutput)

	DownloadDBLogFilePortion(*rds.DownloadDBLogFilePortionInput) (*rds.DownloadDBLogFilePortionOutput, error)

	DownloadDBLogFilePortionPages(*rds.DownloadDBLogFilePortionInput, func(*rds.DownloadDBLogFilePortionOutput, bool) bool) error

	FailoverDBClusterRequest(*rds.FailoverDBClusterInput) (*request.Request, *rds.FailoverDBClusterOutput)

	FailoverDBCluster(*rds.FailoverDBClusterInput) (*rds.FailoverDBClusterOutput, error)

	ListTagsForResourceRequest(*rds.ListTagsForResourceInput) (*request.Request, *rds.ListTagsForResourceOutput)

	ListTagsForResource(*rds.ListTagsForResourceInput) (*rds.ListTagsForResourceOutput, error)

	ModifyDBClusterRequest(*rds.ModifyDBClusterInput) (*request.Request, *rds.ModifyDBClusterOutput)

	ModifyDBCluster(*rds.ModifyDBClusterInput) (*rds.ModifyDBClusterOutput, error)

	ModifyDBClusterParameterGroupRequest(*rds.ModifyDBClusterParameterGroupInput) (*request.Request, *rds.DBClusterParameterGroupNameMessage)

	ModifyDBClusterParameterGroup(*rds.ModifyDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error)

	ModifyDBClusterSnapshotAttributeRequest(*rds.ModifyDBClusterSnapshotAttributeInput) (*request.Request, *rds.ModifyDBClusterSnapshotAttributeOutput)

	ModifyDBClusterSnapshotAttribute(*rds.ModifyDBClusterSnapshotAttributeInput) (*rds.ModifyDBClusterSnapshotAttributeOutput, error)

	ModifyDBInstanceRequest(*rds.ModifyDBInstanceInput) (*request.Request, *rds.ModifyDBInstanceOutput)

	ModifyDBInstance(*rds.ModifyDBInstanceInput) (*rds.ModifyDBInstanceOutput, error)

	ModifyDBParameterGroupRequest(*rds.ModifyDBParameterGroupInput) (*request.Request, *rds.DBParameterGroupNameMessage)

	ModifyDBParameterGroup(*rds.ModifyDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error)

	ModifyDBSnapshotAttributeRequest(*rds.ModifyDBSnapshotAttributeInput) (*request.Request, *rds.ModifyDBSnapshotAttributeOutput)

	ModifyDBSnapshotAttribute(*rds.ModifyDBSnapshotAttributeInput) (*rds.ModifyDBSnapshotAttributeOutput, error)

	ModifyDBSubnetGroupRequest(*rds.ModifyDBSubnetGroupInput) (*request.Request, *rds.ModifyDBSubnetGroupOutput)

	ModifyDBSubnetGroup(*rds.ModifyDBSubnetGroupInput) (*rds.ModifyDBSubnetGroupOutput, error)

	ModifyEventSubscriptionRequest(*rds.ModifyEventSubscriptionInput) (*request.Request, *rds.ModifyEventSubscriptionOutput)

	ModifyEventSubscription(*rds.ModifyEventSubscriptionInput) (*rds.ModifyEventSubscriptionOutput, error)

	ModifyOptionGroupRequest(*rds.ModifyOptionGroupInput) (*request.Request, *rds.ModifyOptionGroupOutput)

	ModifyOptionGroup(*rds.ModifyOptionGroupInput) (*rds.ModifyOptionGroupOutput, error)

	PromoteReadReplicaRequest(*rds.PromoteReadReplicaInput) (*request.Request, *rds.PromoteReadReplicaOutput)

	PromoteReadReplica(*rds.PromoteReadReplicaInput) (*rds.PromoteReadReplicaOutput, error)

	PromoteReadReplicaDBClusterRequest(*rds.PromoteReadReplicaDBClusterInput) (*request.Request, *rds.PromoteReadReplicaDBClusterOutput)

	PromoteReadReplicaDBCluster(*rds.PromoteReadReplicaDBClusterInput) (*rds.PromoteReadReplicaDBClusterOutput, error)

	PurchaseReservedDBInstancesOfferingRequest(*rds.PurchaseReservedDBInstancesOfferingInput) (*request.Request, *rds.PurchaseReservedDBInstancesOfferingOutput)

	PurchaseReservedDBInstancesOffering(*rds.PurchaseReservedDBInstancesOfferingInput) (*rds.PurchaseReservedDBInstancesOfferingOutput, error)

	RebootDBInstanceRequest(*rds.RebootDBInstanceInput) (*request.Request, *rds.RebootDBInstanceOutput)

	RebootDBInstance(*rds.RebootDBInstanceInput) (*rds.RebootDBInstanceOutput, error)

	RemoveSourceIdentifierFromSubscriptionRequest(*rds.RemoveSourceIdentifierFromSubscriptionInput) (*request.Request, *rds.RemoveSourceIdentifierFromSubscriptionOutput)

	RemoveSourceIdentifierFromSubscription(*rds.RemoveSourceIdentifierFromSubscriptionInput) (*rds.RemoveSourceIdentifierFromSubscriptionOutput, error)

	RemoveTagsFromResourceRequest(*rds.RemoveTagsFromResourceInput) (*request.Request, *rds.RemoveTagsFromResourceOutput)

	RemoveTagsFromResource(*rds.RemoveTagsFromResourceInput) (*rds.RemoveTagsFromResourceOutput, error)

	ResetDBClusterParameterGroupRequest(*rds.ResetDBClusterParameterGroupInput) (*request.Request, *rds.DBClusterParameterGroupNameMessage)

	ResetDBClusterParameterGroup(*rds.ResetDBClusterParameterGroupInput) (*rds.DBClusterParameterGroupNameMessage, error)

	ResetDBParameterGroupRequest(*rds.ResetDBParameterGroupInput) (*request.Request, *rds.DBParameterGroupNameMessage)

	ResetDBParameterGroup(*rds.ResetDBParameterGroupInput) (*rds.DBParameterGroupNameMessage, error)

	RestoreDBClusterFromS3Request(*rds.RestoreDBClusterFromS3Input) (*request.Request, *rds.RestoreDBClusterFromS3Output)

	RestoreDBClusterFromS3(*rds.RestoreDBClusterFromS3Input) (*rds.RestoreDBClusterFromS3Output, error)

	RestoreDBClusterFromSnapshotRequest(*rds.RestoreDBClusterFromSnapshotInput) (*request.Request, *rds.RestoreDBClusterFromSnapshotOutput)

	RestoreDBClusterFromSnapshot(*rds.RestoreDBClusterFromSnapshotInput) (*rds.RestoreDBClusterFromSnapshotOutput, error)

	RestoreDBClusterToPointInTimeRequest(*rds.RestoreDBClusterToPointInTimeInput) (*request.Request, *rds.RestoreDBClusterToPointInTimeOutput)

	RestoreDBClusterToPointInTime(*rds.RestoreDBClusterToPointInTimeInput) (*rds.RestoreDBClusterToPointInTimeOutput, error)

	RestoreDBInstanceFromDBSnapshotRequest(*rds.RestoreDBInstanceFromDBSnapshotInput) (*request.Request, *rds.RestoreDBInstanceFromDBSnapshotOutput)

	RestoreDBInstanceFromDBSnapshot(*rds.RestoreDBInstanceFromDBSnapshotInput) (*rds.RestoreDBInstanceFromDBSnapshotOutput, error)

	RestoreDBInstanceToPointInTimeRequest(*rds.RestoreDBInstanceToPointInTimeInput) (*request.Request, *rds.RestoreDBInstanceToPointInTimeOutput)

	RestoreDBInstanceToPointInTime(*rds.RestoreDBInstanceToPointInTimeInput) (*rds.RestoreDBInstanceToPointInTimeOutput, error)

	RevokeDBSecurityGroupIngressRequest(*rds.RevokeDBSecurityGroupIngressInput) (*request.Request, *rds.RevokeDBSecurityGroupIngressOutput)

	RevokeDBSecurityGroupIngress(*rds.RevokeDBSecurityGroupIngressInput) (*rds.RevokeDBSecurityGroupIngressOutput, error)
}

var _ RDSAPI = (*rds.RDS)(nil)
