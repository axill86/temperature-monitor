# About
This is the simple temperature monitor for raspberry pi.
That uses ic2-compatible temperature sensor BME280 and [Library](github.com/d2r2/go-bsbmp) to work with that.
## How tu build
``make build-pi``
## How to provision
In deployments folder there is a bunch of resources which is provisioned on yandex cloud
* mqtt broker
* function and trigger for metrics

In order to run terraform scripts make sure YC_SERVICE_ACCOUNT_KEY_FILE variable set

``export YC_SERVICE_ACCOUNT_KEY_FILE=/path/to/key/file``
``eport YC_FOLDER_ID=<Your Folder Id>``

## How to build

``make build-pi``

