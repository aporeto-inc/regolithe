# Model
model:
  rest_name: task
  resource_name: tasks
  entity_name: Task
  package: todo-list
  group: core
  description: Represent a task to do in a listd.
  aliases:
  - tsk
  get:
    description: Retrieve the task with the given ID.
  update:
    description: Updates the task with the given ID.
  delete:
    description: Deletes the task with the given ID.
  extends:
  - '@base'

# Attributes
attributes:
  v1:
  - name: description
    description: The description.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: name
    description: The name.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: the name
    filterable: true
    getter: true
    setter: true
    orderable: true
    validations:
    - $nospace
    - $nocap

  - name: status
    description: The status of the task.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - DONE
    - PROGRESS
    - TODO
    default_value: TODO
    filterable: true
    orderable: true
