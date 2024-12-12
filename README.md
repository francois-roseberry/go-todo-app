# go-todo-app
Simple todo app using Go and HTMX as a practice and personal reference point of go project setup

## Features

2 modes of operation, toggled via clicking on the lock icon:
- while unlocked, tasks can be re-ordered, created, deleted or renamed
  - re-ordering is done via drag-and-drop
  - adding a task is done via the button at the button of the list
  - deleting  task is done via the 'minus' icon at the right of each task
  - renaming a task is done by clicking on a task name, entering a new name, and pressing enter
    - to cancel the renaming, press Esc when editinga task name
- while locked, the list of tasks is frozen, it can only be checked/unchecked

The list of tasks is in-memory only, no persistence has been added since this would just be noise for now; this demo is about HTMX and Templ

## Development setup

Install [Go](https://go.dev/doc/install), [Air](https://github.com/air-verse/air), [Make](https://www.gnu.org/software/make/) (probably already installed by other tools) and [Templ](https://templ.guide/quick-start/installation)

To lint:

`make lint`

To run the unit tests:

`make test`

To run the unit tests and update the snapshots:

`UPDATE_SNAPS=true make test`

To build:

`make build`

To build and start the server in non-development mode:

`make start`

To start the server in development mode (with hot reloading):

`make dev`

Point your browser to `http://localhost:3000`
