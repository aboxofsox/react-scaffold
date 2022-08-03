# React Scaffold
A simple CLI tool to get a simple React app started. No need for CRA.

### Usage
| Argument         | Description                                     |
|:-----------------|:------------------------------------------------|
| `rscaf scaffold` | Run the scaffolding tool only.                  |
| `rscaf serve`    | Simple HTTP server for development              |
| `rscaf both`     | Run the scaffolding tool, then the HTTP server  |

### Exepcted Results <small>*excludes node_modules*</small>
```
.
├── dist
├── index.html
├── package-lock.json
├── package.json
├── scripts
│   ├── build.js
│   └── dev.js
├── src
│   ├── components
│   ├── index.tsx
│   └── styles
│       └── styles.css
├── tailwind.config.js
└── tsconfig.json
```

### Installation
Download the binary for your system. How you install it is up to you.
