#!/bin/bash

until timeout 1 bash -c "9<>/dev/tcp/127.0.0.1/5672"; do
  echo "Trying connection on 127.0.0.1:5672"
  sleep 10
done

while [ "$(rabbitmq-diagnostics -q ping > /dev/null 2>&1 ; echo $?)" -ne 0 ]; do
  echo "Waiting rabbitmq ping to be successfull"
  sleep 5
done

rabbitmqctl add_user "${RABBITMQ_USERNAME}" "${RABBITMQ_PASSWORD}"

rabbitmqctl set_user_tags "${RABBITMQ_USERNAME}" administrator
rabbitmqctl set_permissions -p "/" "${RABBITMQ_USERNAME}" ".*" ".*" ".*"
rabbitmqctl set_permissions -p "/" guest ".*" ".*" ".*"

# setup
rabbitmqadmin declare --vhost="/" queue name="books_queue" auto_delete=false durable=true
rabbitmqadmin declare --vhost="/" exchange name="books_exchange" type=topic auto_delete=false durable=true
rabbitmqadmin declare --vhost="/" binding source="books_exchange" destination="books_queue" routing_key="books.#" destination_type=queue

# messages on queue with higher priority will be consumed first
rabbitmqadmin declare --vhost="/" queue name="recommendations_queue" auto_delete=false durable=true arguments='{"x-max-priority":10}'
rabbitmqadmin declare --vhost="/" exchange name="recommendations_exchange" type=topic auto_delete=false durable=true
rabbitmqadmin declare --vhost="/" binding source="recommendations_exchange" destination="recommendations_queue" routing_key="recommendations.#" destination_type=queue
