const clean = require("esbuild-plugin-clean").clean;

require("esbuild")
  .build({
    entryPoints: ["./src/index.tsx"],
    bundle: true,
    entryNames: "[name].[hash]",
    outdir: "../public/js/app",
    sourcemap: true,
    plugins: [
      clean({
        options: {
          force: true,
        },
        patterns: ["../public/js/app/*"],
      }),
    ],
  })
  .catch(() => process.exit(1));
