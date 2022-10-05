# Beeldeploy 

Beeldeploy is a simple tool build at Beeldi.com to allow us to store Cloud Functions configurations in yaml instead of having endless command line arguments.

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

|               Flags                | Supported |             YAML Key              |
| :--------------------------------: | :-------: | :-------------------------------: |
|       allow-unauthenticated        |     ✅     |       allow-unauthenticated       |
|          docker-registery          |     ✅     |         docker-registery          |
|          egress-settings           |     ✅     |              egress               |
|            entry-point             |     ✅     |            entry-point            |
|                gen2                |     ✅     |               gen2                |
|            ignore-file             |     ✅     |            ignore-file            |
|          ingress-settings          |     ✅     |              ingress              |
|               memory               |     ✅     |              memory               |
|               retry                |     ✅     |               retry               |
|        run-service-account         |     ✅     |        run-service-account        |
|              runtime               |     ✅     |              runtime              |
|           security-level           |     ❌     |
| serve-all-traffic-latest-revision  |     ✅     | serve-all-traffic-latest-revision |
|          service-account           |     ✅     |          service-account          |
|               region               |     ✅     |              region               |
|               source               |     ✅     |                src                |
|            stage-bucket            |     ❌     |
|              timeout               |     ✅     |              timeout              |
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
|            set-env-vars            |     ✅     |               envs                |
|          remove-env-vars           |     ❌     |
|          update-env-vars           |     ❌     |
|           clear-kms-key            |     ❌     |
|              kms-key               |     ❌     |
|            clear-labels            |     ❌     |
|           remove-labels            |     ❌     |
|         clear-max-instance         |     ❌     |
|           max-instances            |     ✅     |           max-instances           |
|         clear-min-instance         |     ❌     |
|           min-instances            |     ✅     |           min-instances           |
|           clear-secrets            |     ❌     |
|            set-secrets             |     ✅     |              secrets              |
|           remove-secrets           |     ❌     |
|           update-secrets           |     ❌     |
|        clear-vpc-connector         |     ❌     |
|           vpc-connector            |     ✅     |                vpc                |
|           trigger-bucket           |     ❌     |
|            tigger-http             |     ✅     |           trigger-http            |
|           trigger-topic            |     ❌     |
|           trigger-event            |     ❌     |
|         trigger-ressource          |     ❌     |
|       trigger-event-filters        |     ❌     |
| trigger-event-filters-path-pattern |     ❌     |













    