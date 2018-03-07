aws dynamodb create-table \
    --table-name Shipt.Product \
    --attribute-definitions AttributeName=ID,AttributeType=S AttributeName=Category,AttributeType=S \
    --key-schema KeyType=HASH,AttributeName=ID KeyType=RANGE,AttributeName=Category \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1