# idp-sample-service

A small web service deployed by the [internal developer platform](https://github.com/kushan94/eks-internal-developer-platform).
Pushing to main builds the image, scans and signs it, and records the new tag in `deploy/web.yaml`, which Argo CD deploys. See the platform's `apps/README.md`.
