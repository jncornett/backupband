# TODO

- Add request models for all request types
- Add handlers for all request types
- Add response model
- Add unit tests and 'integration' tests

## Architecture

    /---------\    /--------------------\    /-------------\
    | Request |--->| backupband.Handler |--->| Request Mux |
    \---------/    \--------------------/    \-------------/


### LaunchRequest

Setup Session.

### SessionEndedRequest

Cleanup Session.

### IntentRequest

    /---------------\    /------------\
    | IntentRequest |--->| Intent Mux |
    \---------------/    \------------/

See `intents.json`.