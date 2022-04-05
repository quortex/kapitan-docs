# kapitan-docs

kapitan-docs is a tool to automatically generate documentation for a [Kapitan](https://kapitan.dev) project.

The markdown generation is managed by [go text/template package](https://pkg.go.dev/text/template). kapitan-docs parses metadata from a Kapitan project and generates a number of sub-templates that can be referenced in a template file (by default `README.md.gotmpl`). If no template file is provided, the tool has a default internal template that will generate a reasonably formatted README.

**Currently only class documentation is supported !**

The most useful aspect of this tool is the auto-detection of field descriptions from comments, e.g for a class `./foo/bar.yaml`:

```yml
# -- Foo is a test class, used to display an example.
#
# Class comment support **markdown** in comments.

classes:
  - bar.baz

parameters:
  foo:
    # -- A `number`
    number: 2
    # -- A `string`
    string: "foo"
    # -- An `object`
    object:
      bar: baz
    # -- A `list`
    list:
      - myvalue1
      - myvalue2
      - myvalue3
```

Resulting in a resulting README section like so:

### <a name="foo.bar"></a>foo.bar

Foo is a test class, used to display an example.

Class comment support **markdown** in comments.
#### Uses
* [bar.baz](#bar.baz)

#### Parameters

| Key | Type | Default | Description |
| --- | ---- | ------- | ----------- |
| foo.number | number | <pre>2<br /></pre> | <pre>A number</pre> |
| foo.string | string | <pre>"foo"<br /></pre> | <pre>A string</pre> |
| foo.object | object | <pre>bar: baz<br /></pre> | <pre>An object</pre> |
| foo.list | list | <pre>- myvalue1<br />- myvalue2<br />- myvalue3<br /></pre> | <pre>A list</pre> |

## Usage

To generate documentation in README for a Kapitan project, run `kapitan-docs .` at the root of that project.

```
Usage:
  kapitan-docs [OPTIONS] [Directory]

Application Options:
  -d, --dry-run        Don't render any markdown file, just print in the console.
  -l, --log-level=     Level of logs that should printed, one of (panic, fatal, error, warning, info, debug, trace). (default: error)
  -t, --template-file= gotemplate file path from which documentation will be generated. (default: README.md.gotmpl)

Help Options:
  -h, --help           Show this help message

Arguments:
  Directory:           Kapitan project directory.
```

### Using docker

You can mount a Kapitan project under `/kapitan-docs` within the container.

Then run:

```bash
docker run --rm --volume "$(pwd):/kapitan-docs" -u $(id -u) kapitan-docs:latest
```
