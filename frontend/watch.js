const esbuild = require("esbuild");
const clean = require("esbuild-plugin-clean").clean;

async function watch() {
  let ctx = await esbuild.context({
    entryPoints: ["./src/index.tsx"],
    minify: false,
    entryNames: "[name].[hash]",
    outdir: "../public/js/app",
    bundle: true,
    plugins: [
      clean({
        options: {
          force: true,
        },
        patterns: ["../public/js/app/*"],
      }),
    ],
  });
  await ctx.watch();
  console.log("Watching...");
}
watch();
