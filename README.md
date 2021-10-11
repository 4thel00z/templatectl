# templatectl

![templatectl.png](https://raw.githubusercontent.com/4thel00z/logos/master/templatectl.png)

## Motivation

[boilr](https://github.com/tmrts/boilr) is abandonware, let's do our own thing.

## Usage


### Download a Template
In order to download a template from a github repository, use the following command:

```bash
templatectl template download <github-repo-path> <template-tag>
templatectl template download 4thel00z/service-module sm
```

The downloaded template will be saved to local `templatectl` registry.

### Save a Local Template
In order to save a template from filesystem to the template registry use the following command:

```bash
templatectl template save <template-path> <template-tag>
templatectl template save ~/templates/node-noob node-noob
```

The saved template will be saved to local `templatectl` registry.

### Use a Template
For a templatectl template with the given directory structure:

```tree
.
├── project.json
├── README.md
└── template
    └── package.json
```

And the following `project.json` context file:

```json
{
    "Author": "4thel00z",
    "Year": "2021",
    "TaskRunner": [
        "gulp",
        "make",
        "just"
    ]
}
```

When using the template with the following command:

```bash
templatectl template use <template-tag> <target-dir>
templatectl template use node-noob /home/noob/node/project
```


## License

This project is licensed under the GPL-3 license.
