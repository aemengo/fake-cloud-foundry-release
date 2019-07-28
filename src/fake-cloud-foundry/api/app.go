package api

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

type putAppForm struct {
	State string `json:"state"`
}

func (a *API) PutApp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	var f putAppForm
	json.NewDecoder(r.Body).Decode(&f)

	guid := ps.ByName("guid")
	app, ok := a.database.GetAppByGuid(guid)
	if !ok {
		//TODO something
		return
	}

	app.State = f.State
	if app.State == "STARTED" {
		go a.startApp(app)
	}

	resource := Resource{
		Metadata: Metadata{
			Guid:      app.Guid,
			URL:       fmt.Sprintf("/v2/apps/%s", app.Guid),
			CreatedAt: app.CreatedAt.Format(time.RFC3339),
			UpdatedAt: app.CreatedAt.Format(time.RFC3339),
		},
		Entity: presentApp(app),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

func (a *API) GetApp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guid := ps.ByName("guid")

	app, ok := a.database.GetAppByGuid(guid)
	if !ok {
		//TODO something
		return
	}

	resource := Resource{
		Metadata: Metadata{
			Guid:      app.Guid,
			URL:       fmt.Sprintf("/v2/apps/%s", app.Guid),
			CreatedAt: app.CreatedAt.Format(time.RFC3339),
			UpdatedAt: app.CreatedAt.Format(time.RFC3339),
		},
		Entity: presentApp(app),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource)
}

func (a *API) startApp(app db.App) {
	time.Sleep(6 * time.Second)

	app.DetectedBuildpack = strPtr("some-buildpack")
	app.DetectedBuildpackGuid = strPtr(uuid.NewV4().String())
	app.DetectedStartCommand = "some-start-command"
	app.PackageState = "STAGED"

	a.database.SaveApp(app)
}

func strPtr(v string) *string {
	return &v
}