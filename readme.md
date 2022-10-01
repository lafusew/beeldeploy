# Beeldeploy 

Beeldeploy is a simple tool to allow us to store Cloud Functions configurations in yaml instead of having endless command line arguments.

## How to install 

Currently no release is out yet you can simply build it from source.

Install Golang 1.19 and run `go build .` inside the project.

## How to use

In order to use it you need to have `gcloud cli` installed and to be logged in. You also need to bet setted up on the project you want to deploy to.

In later version you'll be able to choose the project you want to deploy on.

It will change, but for now create a .yaml config file wherever you want. Like the cf.yaml there's in the repo.

And run `beeldeploy -f <function-name> -c <config file path>`

## Supported flags.

**/!\ name need to be set using -f flag on beeldeploy cmd**

Currently all google cloud functions deployement -flags aren't supported you find the list here:

|               Flags                | Supported |
| :--------------------------------: | :-------: |
|               region               |     ✅     |
|       allow-unauthenticated        |     ✅     |
|          docker-registery          |     ✅     |
|          egress-settings           |     ❌     |
|            entry-point             |     ✅     |
|                gen2                |     ✅     |
|            ignore-file             |     ❌     |
|          ingress-settings          |     ❌     |
|               memory               |     ✅     |
|               retry                |     ❌     |
|        run-service-account         |     ❌     |
|              runtime               |     ✅     |
|           security-level           |     ❌     |
| serve-all-traffic-latest-revision  |     ✅     |
|          service-account           |     ✅     |
|               source               |     ✅     |
|            stage-bucket            |     ❌     |
|              timeout               |     ✅     |
|          trigger-location          |     ❌     |
|      trigger-service-account       |     ❌     |
|           update-labels            |     ❌     |
|        build-env-vars-file         |     ❌     |
|        clear-build-env-vars        |     ❌     |
|         set-build-env-vars         |     ❌     |
|       remove-build-env-vars        |     ❌     |
|       update-build-env-vars        |     ❌     |
|         build-worker-pool          |     ❌     |
|      clear-build-worker-pool       |     ❌     |
|      clear-docker-repository       |     ❌     |
|         docker-repository          |     ❌     |
|           clear-env-vars           |     ❌     |
|            set-env-vars            |     ✅     |
|          remove-env-vars           |     ❌     |
|          update-env-vars           |     ❌     |
|           clear-kms-key            |     ❌     |
|              kms-key               |     ❌     |
|            clear-labels            |     ❌     |
|           remove-labels            |     ❌     |
|         clear-max-instance         |     ❌     |
|           max-instances            |     ✅     |
|         clear-min-instance         |     ❌     |
|           min-instances            |     ✅     |
|           clear-secrets            |     ❌     |
|            set-secrets             |     ✅     |
|           remove-secrets           |     ❌     |
|           update-secrets           |     ❌     |
|        clear-vpc-connector         |     ❌     |
|           vpc-connector            |     ✅     |
|           trigger-bucket           |     ❌     |
|            tigger-http             |     ✅     |
|           trigger-topic            |     ❌     |
|           trigger-event            |     ❌     |
|         trigger-ressource          |     ❌     |
|       trigger-event-filters        |     ❌     |
| trigger-event-filters-path-pattern |     ❌     |













    
