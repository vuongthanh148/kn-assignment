
# Deploy

The scripts in this directory are meant to be run as steps in Cloud Build.

`deploy-env.yaml` deploys the config, worker, and frontend for a single
environment. The environment is configurable from the Cloud Build trigger
execution page.

`deploy.yaml` deploys to staging, runs e2e tests, and if the tests pass,
deploys to production.
