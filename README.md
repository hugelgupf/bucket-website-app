# Host a static website using Cloud Run and Cloud Storage

`index.html` is the index page, served when a request URL ends in `/`.

# Run Button

Create a service account that has Storage Object Viewer access. (This can be
done in the Web Console UI as well.)

```
export PROJECT_ID=...

# Create service account
gcloud iam service-accounts create cloud-run-app

# Add Storage Object Viewer access to the service account role.
gcloud projects add-iam-policy-binding $PROJECT_ID --member "serviceAccount:cloud-run-app@$PROJECT_ID.iam.gserviceaccount.com" --role "roles/storage.objectViewer"
```

Now hit the run button. It will ask for two env vars:

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)

* `BUCKET`: name of the GCS bucket
* `SERVICE_ACCOUNT`: name of a service account that has Storage Object Viewer
  access. If you created one using the steps above, use
  `cloud-run-app@$PROJECT_ID.iam.gserviceaccount.com`.
