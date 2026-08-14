# OpenRCS - Open-Source Rail Control Simulator
## Project Background
OpenRCS is a free, open-source rail signalling simulator, that aims to have the following features:
- Interlocking and timetable logic
- Support for multiple signalling systems, including the US, UK, and German systems
- Easy-to-edit timetables, maps, and game data
- Provide an enjoyable and easy cross-platform user experience

**OpenRCS is licensed under GPLv3.**

## Contributors
OpenRCS is developed in Go, Svelte, and Wails. General guiding principles for contributors include:
- Composition over inheritance
- Single source of truth
- Keep it simple, limit abstractions if possible
  
## Running a Dev Environment
This assumes that you have a working knowledge of Git and development workflows
1. Clone the repository to your local computer
2. Ensure you have [Golang](https://go.dev/doc/install), [Node.js](https://nodejs.org/en/download), and [Wails](https://v3.wails.io/quick-start/installation/) installed.
3. Run `wails3 dev`

Note to linux users on webkit2gtk-4.1: use -tags `webkit2_41` in all commands
