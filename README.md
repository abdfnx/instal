<p align="center">
  <img src="https://raw.githubusercontent.com/abdfnx/instal/main/.github/assets/logo.svg" height="120px" />
</p>

![preview](https://user-images.githubusercontent.com/64256993/154235804-d29d7b10-c7d2-4732-8c06-63ee994865eb.gif)

> 🛰️ Install any binary app from a script URL.

## Installation

### Using script

* Shell

```
curl -fsSL https://bit.ly/instal-cli | bash
```

* PowerShell

```
iwr -useb https://bit.ly/instal-win | iex
```

**then restart your powershell**

### Homebrew

```bash
brew install abdfnx/tap/instal
```

### GitHub CLI

```bash
gh extension install abdfnx/gh-instal
```

### Via Docker

```bash
docker run -it instalcli/instal
```

## Usage

* Open Instal UI

```bash
instal
```

* Install binary app from script URL and run it

```
instal <SCRIPT_URL>
```

## Flags

```
      --help           Help for instal
  -H, --hidden         hide the output
  -s, --shell string   shell to use (Default: bash | powershell)
```

## Examples

```bash
instal https://get.docker.com

instal https://https://getmic.ro --shell sh
```

### License

instal is licensed under the terms of [MIT](https://github.com/abdfnx/instal/blob/main/LICENSE) license.
