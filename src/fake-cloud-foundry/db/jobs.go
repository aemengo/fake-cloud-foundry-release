package db

import (
	"github.com/satori/go.uuid"
	"time"
)

type Job struct {
	Guid      string
	Status    string
	CreatedAt time.Time
}

func (db *DB) CreateJob() Job {
	job := Job{
		Guid:      uuid.NewV4().String(),
		Status:    "queued",
		CreatedAt: time.Now().UTC(),
	}

	db.jobs = append(db.jobs, job)
	return job
}

func (db *DB) GetJobByGuid(guid string) (Job, bool) {
	for _, job := range db.jobs {
		if job.Guid == guid {
			return job, true
		}
	}

	return Job{}, false
}

func (db *DB) SaveJob(job Job) {
	var (
		i = -1
	)

	for index, j := range db.jobs {
		if j.Guid == job.Guid {
			i = index
		}
	}

	if i > -1 {
		db.jobs[i] = job
	}
}