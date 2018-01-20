# BackupBand

TODO description

## Development

### Building

In the package root, run

    make

This will build the handler binary at `bin/main`.

### Deploying

First, make sure you have these prerequisites:

- serverless (`npm -g install serverless`)
- AWS CLI (`pip install --user awscli`)

Then, make sure the AWS CLI is configured with the appropriate credentials.
Run:

    aws configure

And follow the prompts.

Finally, run:

    serverless deploy

### Invoking the Lambda function

    serverless invoke -f BackupBandAlexaHandler -p request.json -l