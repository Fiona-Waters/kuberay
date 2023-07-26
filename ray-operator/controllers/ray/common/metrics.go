package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// Define all the prometheus counters for all clusters
var (
	clustersCreatedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ray_operator_clusters_created_total",
			Help: "Counts number of clusters created",
		},
		[]string{"namespace"},
	)
	clustersDeletedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ray_operator_clusters_deleted_total",
			Help: "Counts number of clusters deleted",
		},
		[]string{"namespace"},
	)
	clustersSuccessfulCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ray_operator_clusters_successful_total",
			Help: "Counts number of clusters successful",
		},
		[]string{"namespace"},
	)
	clustersFailedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ray_operator_clusters_failed_total",
			Help: "Counts number of clusters failed",
		},
		[]string{"namespace"},
	)
	jobCreatedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ray_job_created_total",
			Help: "Counts number of jobs created",
		},
		[]string{"namespace"},
	) 
	jobsStoppedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ray_jobs_stopped_total",
			Help: "Counts number of jobs stopped",
		},
		[]string{"namespace"},
	)
	jobsFailedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ray_jobs_failed_total",
			Help: "Counts number of failed jobs",
		},
		[]string{"namespace"},
	)
	successfulJobsCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "successful_ray_jobs",
			Help: "Counts number of jobs that have succeeded",
		},
		[]string{
			"namespace",
		},
	)
)

// TODO add a duration histogram, with duration between cluster creation and completion?

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(clustersCreatedCount,
		clustersDeletedCount,
		clustersSuccessfulCount,
		clustersFailedCount, 
		jobCreatedCount,
		jobsFailedCount,
		jobsStoppedCount,
		successfulJobsCount,
	)
}

func CreatedClustersCounterInc(namespace string) {
	clustersCreatedCount.WithLabelValues(namespace).Inc()
}

// TODO: We don't handle the delete events in new reconciler mode, how to emit deletion metrics?
func DeletedClustersCounterInc(namespace string) {
	clustersDeletedCount.WithLabelValues(namespace).Inc()
}

func SuccessfulClustersCounterInc(namespace string) {
	clustersSuccessfulCount.WithLabelValues(namespace).Inc()
}

func FailedClustersCounterInc(namespace string) {
	clustersFailedCount.WithLabelValues(namespace).Inc()
}

func CreatedJobsCounterInc(namespace string) {
	jobCreatedCount.WithLabelValues(namespace).Inc()
}

func StoppedJobsCounterInc(namespace string) {
	jobsStoppedCount.WithLabelValues(namespace).Inc()
}
// failed to create
func FailedJobsCounterInc(namespace string) {
	jobsFailedCount.WithLabelValues(namespace).Inc()
}

// succeeded
func SuccessfulJobsCounterInc(namespace string) {
	successfulJobsCount.WithLabelValues(namespace).Inc()
}
