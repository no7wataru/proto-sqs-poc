import os
import sys
import base64
import boto3

sys.path.append(os.path.normpath(os.path.join(
    os.path.dirname(os.path.abspath(__file__)), '..')))
from proto.message_pb2 import Message


endpoint = 'http://localhost:9324'
region = 'ap-northeast-1'
queue_name = 'proto-test-mq'
access_key = 'x'
secret_key = 'x'

client = boto3.resource('sqs',
                        endpoint_url=endpoint,
                        region_name=region,
                        aws_secret_access_key=access_key,
                        aws_access_key_id=secret_key,
                        use_ssl=False)
queue = client.get_queue_by_name(QueueName=queue_name)

messages = queue.receive_messages(
    AttributeNames=[
        'All'
    ],
    MaxNumberOfMessages=1,
    MessageAttributeNames=[
        'All'
    ],
    VisibilityTimeout=0,
    WaitTimeSeconds=0
)

for message in messages:
    data = base64.b64decode(message.body)
    msg = Message()
    msg.ParseFromString(data)
    print(msg.title)
    for sentence in msg.sentences:
        print(sentence.id)
        print(sentence.text)
    message.delete()
