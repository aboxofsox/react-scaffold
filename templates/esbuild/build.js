import esbuild from 'esbuild'

esbuild.build({
    entryPoints: ['./src/index.tsx'],
    bundle: true,
    outfile: './dist/app.ts',
    sourcemap: true,
}).catch(err => console.error(err))