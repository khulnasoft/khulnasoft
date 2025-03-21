import { async as glob } from "fast-glob";
import os from "os";
import fs from "fs/promises";
import path from "path";
import { cyan, bold } from "picocolors";
import { Sema } from "async-sema";
import { copy } from "../helpers/copy";
import { install } from "../helpers/install";

import {
  GetTemplateFileArgs,
  InstallTemplateArgs,
  TemplateTypeEnum,
} from "./types";
/**
 * Get the file path for a given file in a template, e.g. "next.config.js".
 */
export const getTemplateFile = ({
  template,
  mode,
  file,
}: GetTemplateFileArgs): string => {
  return path.join(__dirname, template, mode, file);
};

export const SRC_DIR_NAMES = ["app", "pages", "styles"];

/**
 * Install a Next.js internal template to a given `root` directory.
 */
export const installTemplate = async ({
  appName,
  root,
  packageManager,
  isOnline,
  template,
  mode,
  eslint,
  srcDir,
  importAlias,
  secretKey,
  applicationId,
  userId,
}: InstallTemplateArgs) => {
  console.log(bold(`Using ${packageManager}.`));

  /**
   * Copy the template files to the target directory.
   */
  console.log("\nInitializing project with template:", template, "\n");
  const templatePath = path.join(__dirname, template, mode);
  const copySource = ["**"];
  if (!eslint) copySource.push("!eslintrc.json");
  if (!template.includes("react")) {
    copySource.push(
      mode === "ts" ? "tailwind.config.ts" : "!tailwind.config.js",
      "!postcss.config.cjs",
    );
  }

  await copy(copySource, root, {
    parents: true,
    cwd: templatePath,
    rename(name) {
      switch (name) {
        case "gitignore":
        case "eslintrc.json": {
          return `.${name}`;
        }
        /*
         * README.md is ignored by webpack-asset-relocator-loader used by ncc:
         * https://github.com/vercel/webpack-asset-relocator-loader/blob/e9308683d47ff507253e37c9bcbb99474603192b/src/asset-relocator.js#L227
         */
        case "README-template.md": {
          return "README.md";
        }
        default: {
          return name;
        }
      }
    },
  });

  const tsconfigFile = path.join(root, "tsconfig.json");
  await fs.writeFile(
    tsconfigFile,
    (await fs.readFile(tsconfigFile, "utf8"))
      .replace(
        `"@/*": ["./*"]`,
        srcDir ? `"@/*": ["./src/*"]` : `"@/*": ["./*"]`,
      )
      .replace(`"@/*":`, `"${importAlias}":`),
  );

  // update import alias in any files if not using the default
  if (importAlias !== "@/*") {
    const files = await glob("**/*", {
      cwd: root,
      dot: true,
      stats: false,
      /*
       * We don't want to modify compiler options in [ts/js]config.json
       * and none of the files in the .git folder
       */
      ignore: ["tsconfig.json", "jsconfig.json", ".git/**/*"],
    });
    const writeSema = new Sema(8, { capacity: files.length });
    await Promise.all(
      files.map(async (file) => {
        await writeSema.acquire();
        const filePath = path.join(root, file);
        if ((await fs.stat(filePath)).isFile()) {
          await fs.writeFile(
            filePath,
            (await fs.readFile(filePath, "utf8")).replace(
              `@/`,
              `${importAlias.replace(/\*/g, "")}`,
            ),
          );
        }
        writeSema.release();
      }),
    );
  }

  if (srcDir) {
    await fs.mkdir(path.join(root, "src"), { recursive: true });
    await Promise.all(
      SRC_DIR_NAMES.map(async (file) => {
        await fs
          .rename(path.join(root, file), path.join(root, "src", file))
          .catch((err) => {
            if (err.code !== "ENOENT") {
              throw err;
            }
          });
      }),
    );

    const isAppTemplate = template.startsWith("app");

    // Change the `Get started by editing pages/index` / `app/page` to include `src`
    const indexPageFile = path.join(
      "src",
      isAppTemplate ? "app" : "pages",
      `${isAppTemplate ? "page" : "index"}.${mode === "ts" ? "tsx" : "js"}`,
    );

    await fs.writeFile(
      indexPageFile,
      (await fs.readFile(indexPageFile, "utf8")).replace(
        isAppTemplate ? "app/page" : "pages/index",
        isAppTemplate ? "src/app/page" : "src/pages/index",
      ),
    );

    if (template === TemplateTypeEnum.APP_REACT_EMAIL) {
      const tailwindConfigFile = path.join(
        root,
        mode === "ts" ? "tailwind.config.ts" : "tailwind.config.js",
      );
      await fs.writeFile(
        tailwindConfigFile,
        (await fs.readFile(tailwindConfigFile, "utf8")).replace(
          /\.\/(\w+)\/\*\*\/\*\.\{js,ts,jsx,tsx,mdx\}/g,
          "./src/$1/**/*.{js,ts,jsx,tsx,mdx}",
        ),
      );
    }
  }

  /* write .env file */
  const val = Object.entries({
    KHULNASOFT_SECRET_KEY: secretKey,
    NEXT_PUBLIC_KHULNASOFT_APPLICATION_IDENTIFIER: applicationId,
    NEXT_PUBLIC_KHULNASOFT_SUBSCRIBER_ID: userId,
  }).reduce((acc, [key, value]) => {
    return `${acc}${key}=${value}${os.EOL}`;
  }, "");

  await fs.writeFile(path.join(root, ".env.local"), val);

  /* write github action */
  await copy(copySource, `${root}/.github`, {
    parents: true,
    cwd: path.join(__dirname, `./github`),
  });

  /** Copy the version from package.json or override for tests. */
  const version = "14.2.3";

  /** Create a package.json for the new project and write it to disk. */
  const packageJson: any = {
    name: appName,
    version: "0.1.0",
    private: true,
    scripts: {
      dev: `next dev --port=4000`,
      build: "next build",
      start: "next start",
      lint: "next lint",
    },
    /**
     * Default dependencies.
     */
    dependencies: {
      react: "^18",
      "react-dom": "^18",
      next: version,
      "@khulnasoft/framework": "latest",
      "@khulnasoft/nextjs": "^2.5.0",
    },
    devDependencies: {},
  };

  /**
   * TypeScript projects will have type definitions and other devDependencies.
   */
  if (mode === "ts") {
    packageJson.devDependencies = {
      ...packageJson.devDependencies,
      typescript: "^5",
      "@types/node": "^20",
      "@types/react": "^18",
      "@types/react-dom": "^18",
    };
  }

  /* Add Tailwind CSS dependencies. */
  if (template === TemplateTypeEnum.APP_REACT_EMAIL) {
    packageJson.devDependencies = {
      ...packageJson.devDependencies,
      postcss: "^8",
      tailwindcss: "^3.4.1",
    };

    packageJson.dependencies = {
      ...packageJson.dependencies,
      "@react-email/components": "0.0.18",
      "@react-email/tailwind": "0.0.18",
    };

    /* Zod dependencies used in react email example */
    packageJson.dependencies = {
      ...packageJson.dependencies,
      zod: "^3.23.8",
      "zod-to-json-schema": "^3.23.1",
    };
  }

  /* Default ESLint dependencies. */
  if (eslint) {
    packageJson.devDependencies = {
      ...packageJson.devDependencies,
      eslint: "^8",
      "eslint-config-next": version,
    };
  }

  const devDeps = Object.keys(packageJson.devDependencies).length;
  if (!devDeps) delete packageJson.devDependencies;

  await fs.writeFile(
    path.join(root, "package.json"),
    JSON.stringify(packageJson, null, 2) + os.EOL,
  );

  console.log("\nInstalling dependencies:");
  // eslint-disable-next-line guard-for-in
  for (const dependency in packageJson.dependencies)
    console.log(`- ${cyan(dependency)}`);

  if (devDeps) {
    console.log("\nInstalling devDependencies:");
    // eslint-disable-next-line guard-for-in
    for (const dependency in packageJson.devDependencies)
      console.log(`- ${cyan(dependency)}`);
  }

  console.log();

  await install(packageManager, isOnline);
};

export * from "./types";
