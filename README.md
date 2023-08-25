# SSO-X

`ssox` is a TUI (Termnial User Interface) to make AWS SSO's experience nice and smooth.

## Installation

#### On a mac

```bash
brew install omerxx/tools/ssox
```

#### Other platforms

Head on to the releases page and find your OS / Arch. If it's not there, open an issue.

## Usage

```bash
# Start the UI
ssox
```
It'll read your `$HOME/.aws/config` file and list the available configured profiles.
Pick one and it'll pop a browser for authentication.

You can use `/` to filter with a fuzzy search

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
