#!/bin/sh

gcloud functions deploy $FUNCTION_NAME \
  --project=$PROJECT \
  --gen2 \
  --runtime=go121 \
  --region=$REGION \
  --source=./ \
  --entry-point=$ENTRY_POINT \
  --trigger-http \
  --set-env-vars=ESA_ACCESS_TOKEN=$ESA_ACCESS_TOKEN \
  --max-instances=$MAX_INSTANCES \
  --concurrency=$CONCURRENCY \
  --timeout=$TIMEOUT
