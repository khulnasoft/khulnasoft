/**
 * This script is used to update the dependencies of all Khulnasoft packages in the monorepo.
 *
 * It can either update all dependencies to the latest version or restore the workspace:* pnpm protocol in all package.json files
 */

import fs from 'fs-extra';
import path from 'node:path';
import { glob } from 'glob';
import { execa } from 'execa';

// Parse CLI arguments
const ALLOWED_REPLACEMENTS = new Set(['workspace:*', 'latest']);
const replacement = process.argv[2];

if (!ALLOWED_REPLACEMENTS.has(replacement)) {
  exit('Usage: node scripts/set-package-dependencies.mjs <workspace:*|latest>');
}

// Remove quoted lines from the nx show projects command
function removeQuotedLines(input) {
  // Split the input into an array of lines
  let lines = input.split('\n');

  // Filter out lines that start with the '>' symbol
  let filteredLines = lines.filter((line) => !line.trim().startsWith('>'));

  // Join the filtered lines back into a single string
  return filteredLines.join('\n');
}

// Get all package names in the monorepo
async function getPackageNames() {
  const { stdout } = await execa`pnpm nx show projects --json`;
  return JSON.parse(removeQuotedLines(stdout));
}

const khulnasoftPackages = new Set(await getPackageNames());

// Update versions of all @khulnasoft dependencies
function updateKhulnasoftDependencies(dependencies) {
  for (const [key] of Object.entries(dependencies || {})) {
    if (key.startsWith('@khulnasoft/') && khulnasoftPackages.has(key)) {
      dependencies[key] = replacement;
    }
  }
}

// Update all dependency fields in a package.json file
function processPackageJson(filePath) {
  const packageJson = fs.readJsonSync(filePath);

  updateKhulnasoftDependencies(packageJson.dependencies);
  updateKhulnasoftDependencies(packageJson.devDependencies);
  updateKhulnasoftDependencies(packageJson.peerDependencies);
  updateKhulnasoftDependencies(packageJson.optionalDependencies);

  fs.writeJsonSync(filePath, packageJson, { spaces: 2 });
  console.log(`Set Khulnasoft packages dependencies to ${replacement} at ${filePath}`);
}

// Find all package.json files in the repo
const files = await glob('packages/**/package.json', { ignore: 'packages/**/node_modules/**' });

// Update all package.json files
files.forEach((file) => processPackageJson(path.resolve(file)));
