package api

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"time"
)

//{
//  "entity": {
//    "guid": "6a3ce0e7-7fdd-4a63-a5cf-783ff76f1107",
//    "status": "queued"
//  },
//  "metadata": {
//    "created_at": "2019-07-22T20:46:48Z",
//    "guid": "6a3ce0e7-7fdd-4a63-a5cf-783ff76f1107",
//    "url": "/v2/jobs/6a3ce0e7-7fdd-4a63-a5cf-783ff76f1107"
//  }
//}

func (a *API) PutBits(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	job := a.database.CreateJob()

	r.ParseMultipartForm(32 << 20)
	go a.saveBits(r, job)

	resource := Resource{
		Metadata: Metadata{
			Guid:      job.Guid,
			URL:       fmt.Sprintf("/v2/jobs/%s", job.Guid),
			CreatedAt: job.CreatedAt.Format(time.RFC3339),
		},
		Entity: presentJob(job),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

func (a *API) saveBits(r *http.Request, job db.Job) {
	file, handler, err := r.FormFile("application")
	if err != nil {
		//TODO make better
		fmt.Println("Checkpoint 1", err)
		job.Status = "failed"
		a.database.SaveJob(job)
		return
	}
	defer file.Close()

	f, err := os.Create("/tmp/"+handler.Filename)
	if err != nil {
		//TODO make better
		fmt.Println("Checkpoint 2", err)
		job.Status = "failed"
		a.database.SaveJob(job)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	job.Status = "finished"
	a.database.SaveJob(job)
}

func presentJob(job db.Job) map[string]interface{} {
	return map[string]interface{}{
		"guid":   job.Guid,
		"status": job.Status,
	}
}
