package models

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Status represents Job Satus
type Status string

const (
	// Preparing job is preparing Fastq files
	Preparing Status = "preparing"
	// Ready job is ready to be picked up by worker
	Ready Status = "ready"
	// Running job is now being processed by worker
	Running Status = "running"
	// Completed job is already finished without errors
	Completed Status = "completed"
	// Errored job is already finished WITH ERRORs
	Errored Status = "errored"
)

// Jobs provides Collection of "jobs"
func Jobs(session *mgo.Session) *mgo.Collection {
	return session.DB(os.Getenv(dbname)).C("jobs")
}

// Job represents job
type Job struct {
	ID         bson.ObjectId `json:"_id"         bson:"_id"`
	Resource   Resource      `json:"resource"    bson:"resource"`
	Status     Status        `json:"status"      bson:"status"`
	CreatedAt  time.Time     `json:"created_at"  bson:"created_at"`
	StartedAt  *time.Time    `json:"started_at,omitempty"  bson:"started_at,omitempty"`
	FinishedAt *time.Time    `json:"finished_at,omitempty" bson:"finished_at,omitempty"`
	Errors     []string      `json:"errors"      bson:"errors"`
}

// Resource represents job resource
type Resource struct {
	URL       string                 `json:"url"       bson:"url"`
	Reference string                 `json:"reference" bson:"reference"`
	Reads     []string               `json:"reads"     bson:"reads"`
	Extra     map[string]interface{} `json:"extra,omitempty" bson:"extra,omitempty"`
}

// NewJob ...
func NewJob() *Job {
	job := new(Job)
	job.ID = bson.NewObjectId()
	job.CreatedAt = time.Now()
	job.Resource = Resource{
		Reads: []string{},
		Extra: map[string]interface{}{},
	}
	job.Errors = []string{}
	return job
}
