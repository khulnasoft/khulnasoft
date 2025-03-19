const { join, basename } = require('path');
const { copyFileSync, existsSync, mkdirSync } = require('fs');
const Module = require('module');
const { khulnasoftVersion } = require('../utils/versions');
const { getNativeFileCacheLocation } = require('./native-file-cache-location');

// WASI is still experimental and throws a warning when used
// We spawn many many processes so the warning gets printed a lot
// We have a different warning elsewhere to warn people using WASI
const originalEmit = process.emit;
process.emit = function (eventName, eventData) {
  if (
    eventName === `warning` &&
    typeof eventData === `object` &&
    eventData?.name === `ExperimentalWarning` &&
    eventData?.message?.includes(`WASI`)
  ) {
    return false;
  }
  return originalEmit.apply(process, arguments);
};

const khulnasoftPackages = new Set([
  '@khulnasoft/khulnasoft-android-arm64',
  '@khulnasoft/khulnasoft-android-arm-eabi',
  '@khulnasoft/khulnasoft-win32-x64-msvc',
  '@khulnasoft/khulnasoft-win32-ia32-msvc',
  '@khulnasoft/khulnasoft-win32-arm64-msvc',
  '@khulnasoft/khulnasoft-darwin-universal',
  '@khulnasoft/khulnasoft-darwin-x64',
  '@khulnasoft/khulnasoft-darwin-arm64',
  '@khulnasoft/khulnasoft-freebsd-x64',
  '@khulnasoft/khulnasoft-linux-x64-musl',
  '@khulnasoft/khulnasoft-linux-x64-gnu',
  '@khulnasoft/khulnasoft-linux-arm64-musl',
  '@khulnasoft/khulnasoft-linux-arm64-gnu',
  '@khulnasoft/khulnasoft-linux-arm-gnueabihf',
]);

const localNodeFiles = [
  'khulnasoft.android-arm64.node',
  'khulnasoft.android-arm-eabi.node',
  'khulnasoft.win32-x64-msvc.node',
  'khulnasoft.win32-ia32-msvc.node',
  'khulnasoft.win32-arm64-msvc.node',
  'khulnasoft.darwin-universal.node',
  'khulnasoft.darwin-x64.node',
  'khulnasoft.darwin-arm64.node',
  'khulnasoft.freebsd-x64.node',
  'khulnasoft.linux-x64-musl.node',
  'khulnasoft.linux-x64-gnu.node',
  'khulnasoft.linux-arm64-musl.node',
  'khulnasoft.linux-arm64-gnu.node',
  'khulnasoft.linux-arm-gnueabihf.node',
];

const originalLoad = Module._load;

// We override the _load function so that when a native file is required,
// we copy it to a cache directory and require it from there.
// This prevents the file being loaded from node_modules and causing file locking issues.
// Will only be called once because the require cache takes over afterwards.
Module._load = function (request, parent, isMain) {
  const modulePath = request;
  if (khulnasoftPackages.has(modulePath) || localNodeFiles.some((f) => modulePath.endsWith(f))) {
    const nativeLocation = require.resolve(modulePath);
    const fileName = basename(nativeLocation);

    // we copy the file to a workspace-scoped tmp directory and prefix with khulnasoftVersion to avoid stale files being loaded
    const nativeFileCacheLocation = getNativeFileCacheLocation();
    const tmpFile = join(nativeFileCacheLocation, khulnasoftVersion + '-' + fileName);
    if (existsSync(tmpFile)) {
      return originalLoad.apply(this, [tmpFile, parent, isMain]);
    }
    if (!existsSync(nativeFileCacheLocation)) {
      mkdirSync(nativeFileCacheLocation, { recursive: true });
    }
    copyFileSync(nativeLocation, tmpFile);
    return originalLoad.apply(this, [tmpFile, parent, isMain]);
  } else {
    // call the original _load function for everything else
    return originalLoad.apply(this, arguments);
  }
};

const indexModulePath = require.resolve('./native-bindings.js');
delete require.cache[indexModulePath];
const indexModule = require('./native-bindings.js');

module.exports = indexModule;
Module._load = originalLoad;
