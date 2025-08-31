# Mapture
Mapture is a CLI tool that checks your config files for you and displays your keymaps with fuzzy search — lightning fast and frustration‑free.

Imagine you’re hunting through your config files, trying to remember “what did I bind fullscreen toggle to again?”. With Mapture, you just ask… and it tells you.

For example, given an aerospace.toml config with keybindings inside, you can run a fuzzy lookup straight from the CLI, and see all your shortcuts in one place.

### 💡 What is Mapture?
A CLI tool that:

Knows where to find all the popular config files
Parses them and extracts only the keymap definitions
Lets you run fuzzy search over your shortcuts like a wizard typing incantations at high speed
Because life’s too short to grep through dotfiles. ✨

---

### 🔎 Example Usage
```bash
$ mapture aerospace list
mod+shift+q -> close_window 
mod+f       -> toggle_fullscreen 
mod+ent     -> new_window
```

```bash
$ mapture search fullscreen
[aerospace] mod+f        -> toggle_fullscreen
[nvim]      <leader>ff   -> telescope find_files
```

### 🏗️ Implementation Plan
Config Parsers Layer
Parsers that understand how different apps define keymaps (e.g. Neovim, Tmux, i3, Aerospace).

Unified Keymap Model
A single internal representation for all those shortcuts, no matter the tool.

Search Engine (Fuzzy)
A fast fuzzy matcher to power search queries.

CLI Layer
Clean, ergonomic CLI commands (built on Cobra or another Go CLI toolkit).

Pretty Output
Colorful, nicely formatted results you can actually read.

---
### 🌟 Nice to Haves
Live Mode:
Run mapture as an interactive TUI (like fzf): type, see instant results.

Cross‑tool Search:
Show the same shortcut across multiple tools — e.g. ctrl-p in both Neovim and Tmux.

#### Export as Markdown

Generate cheatsheets:
```bash
$ mapture export Aerospace
```

Produces a neat Markdown table for your keymaps.

Global Indexing:
Automatically scan ~/.config/ to index shortcuts across all supported apps in one go.

### 🚀 Why Mapture?
Because your muscle memory is both precious and sneaky.
Mapture lets you capture and search your keymaps, instead of rummaging through config jungles.

One command. All your shortcuts. Zero hair‑pulling.

---
### ⚡ Quick Start (coming soon!)
Install and run mapture
Start searching your configs instantly
Profit 🎉
