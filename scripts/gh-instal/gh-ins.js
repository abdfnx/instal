const fs = require("fs");
const path = require("path");
const rm = require("rimraf");
const mkdirp = require("mkdirp");
const sh = require("shelljs");

const VERSION_CMD = sh.exec("git describe --abbrev=0 --tags");
const VERSION_DATE_CMD = sh.exec("go run ./scripts/date.go");

const VERSION = VERSION_CMD.replace("\n", "").replace("\r", "");
const VERSION_DATE = VERSION_DATE_CMD.replace("\n", "").replace("\r", "");

const ROOT = __dirname;
const TEMPLATES = path.join(ROOT, "templates");

async function updateInstalExtension(ghInstalDir) {
  const templatePath = path.join(TEMPLATES, "gh-instal");
  const template = fs.readFileSync(templatePath).toString("utf-8");

  const templateReplaced = template
    .replace("CLI_VERSION", VERSION)
    .replace("CLI_VERSION_DATE", VERSION_DATE);

  fs.writeFileSync(path.join(ghInstalDir, "gh-instal"), templateReplaced);
}

async function updateExtension() {
  const tmp = path.join(__dirname, "tmp");
  const extensionDir = path.join(tmp, "gh-instal");

  mkdirp.sync(tmp);
  rm.sync(extensionDir);

  console.log(`cloning https://github.com/abdfnx/gh-instal to ${extensionDir}`);

  sh.exec(`git clone https://github.com/abdfnx/gh-instal.git ${extensionDir}`)

  console.log(`done cloning abdfnx/gh-instal to ${extensionDir}`);

  console.log("updating local git...");

  await updateInstalExtension(extensionDir);
}

updateExtension().catch((err) => {
  console.error(`error running scripts/gh-instal/gh-ins.js`, err);
  process.exit(1);
});
